package logic

import (
	"cloud-disk/app/user/cmd/api/internal/svc"
	"cloud-disk/app/user/cmd/api/internal/types"
	"cloud-disk/app/user/cmd/rpc/user"
	userModel "cloud-disk/app/user/model"
	"cloud-disk/common/xerr"
	"context"
	"fmt"
	"github.com/pkg/errors"

	wechat "github.com/silenceper/wechat/v2"
	"github.com/silenceper/wechat/v2/cache"
	miniConfig "github.com/silenceper/wechat/v2/miniprogram/config"
	"github.com/zeromicro/go-zero/core/logx"
)

// ErrWxMiniAuthFailError error
var ErrWxMiniAuthFailError = xerr.NewErrMsg("wechat mini auth fail")

type WxAuthLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewWxAuthLogic(ctx context.Context, svcCtx *svc.ServiceContext) *WxAuthLogic {
	return &WxAuthLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// Wechat auth
func (l *WxAuthLogic) WxAuth(req *types.WXMiniAuthReq) (*types.WXMiniAuthResp, error) {
	//1、Wechat-Mini
	miniprogram := wechat.NewWechat().GetMiniProgram(&miniConfig.Config{
		AppID:     l.svcCtx.Config.WxMiniConf.AppId,
		AppSecret: l.svcCtx.Config.WxMiniConf.Secret,
		Cache:     cache.NewMemory(),
	})
	authResult, err := miniprogram.GetAuth().Code2Session(req.Code)
	if err != nil || authResult.ErrCode != 0 || authResult.OpenID == "" {
		return nil, errors.Wrapf(ErrWxMiniAuthFailError, "发起授权请求失败 err : %v , code : %s  , authResult : %+v", err, req.Code, authResult)
	}
	//2、解析 WeChat-Mini 返回的信息
	userData, err := miniprogram.GetEncryptor().Decrypt(authResult.SessionKey, req.EncryptedData, req.IV)
	if err != nil {
		return nil, errors.Wrapf(ErrWxMiniAuthFailError, "解析数据失败 req : %+v , err: %v , authResult:%+v ", req, err, authResult)
	}

	//3、绑定用户或者创建用户
	var userId int64
	rpcRsp, err := l.svcCtx.UserRpc.GetUserAuthByAuthKey(l.ctx, &user.GetUserAuthByAuthKeyReq{
		AuthType: userModel.UserAuthTypeSmallWX,
		AuthKey:  authResult.OpenID,
	})
	if err != nil {
		return nil, errors.Wrapf(ErrWxMiniAuthFailError, "rpc call userAuthByAuthKey err : %v , authResult : %+v", err, authResult)
	}
	if rpcRsp.UserAuth == nil || rpcRsp.UserAuth.Id == 0 {
		//bind user.

		//Wechat-Mini Decrypted data
		Phone := userData.PhoneNumber
		nickName := fmt.Sprintf("cloud-disk%s", Phone[7:]) //防止昵称有重复无法注册
		registerRsp, err := l.svcCtx.UserRpc.Register(l.ctx, &user.RegisterReq{
			AuthKey:  authResult.OpenID,
			AuthType: userModel.UserAuthTypeSmallWX,
			Phone:    Phone,
			Nickname: nickName,
		})
		if err != nil {
			return nil, errors.Wrapf(ErrWxMiniAuthFailError, "UserRpc.Register err :%v, authResult : %+v", err, authResult)
		}

		return &types.WXMiniAuthResp{
			AccessToken:  registerRsp.AccessToken,
			AccessExpire: registerRsp.AccessExpire,
			RefreshAfter: registerRsp.RefreshAfter,
		}, nil

	} else {
		userId = rpcRsp.UserAuth.UserId
		tokenResp, err := l.svcCtx.UserRpc.GenerateToken(l.ctx, &user.GenerateTokenReq{
			UserId: userId,
		})
		if err != nil {
			return nil, errors.Wrapf(ErrWxMiniAuthFailError, "userRpc.GenerateToken err :%v, userId : %d", err, userId)
		}
		return &types.WXMiniAuthResp{
			AccessToken:  tokenResp.AccessToken,
			AccessExpire: tokenResp.AccessExpire,
			RefreshAfter: tokenResp.RefreshAfter,
		}, nil
	}
}
