package fileFolder

import (
	"net/http"

	"cloud-disk/common/result"

	"cloud-disk/app/disk/cmd/api/internal/logic/fileFolder"
	"cloud-disk/app/disk/cmd/api/internal/svc"
	"cloud-disk/app/disk/cmd/api/internal/types"
)

func MoveHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.MovedReq
		if err := httpx.Parse(r, &req); err != nil {
			result.ParamErrorResult(r, w, err)
			return
		}

		l := fileFolder.NewMoveLogic(r.Context(), svcCtx)
		resp, err := l.Move(&req)
		result.HttpResult(r, w, resp, err)
	}
}
