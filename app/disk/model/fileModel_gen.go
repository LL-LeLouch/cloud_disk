// Code generated by goctl. DO NOT EDIT!

package model

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"cloud-disk/common/globalkey"
	"cloud-disk/common/xerr"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	fileFieldNames          = builder.RawFieldNames(&File{})
	fileRows                = strings.Join(fileFieldNames, ",")
	fileRowsExpectAutoSet   = strings.Join(stringx.Remove(fileFieldNames, "`id`", "`create_time`", "`update_time`"), ",")
	fileRowsWithPlaceHolder = strings.Join(stringx.Remove(fileFieldNames, "`id`", "`create_time`", "`update_time`"), "=?,") + "=?"

	cacheCloudDiskFileIdPrefix = "cache:cloudDisk:file:id:"
)

type (
	fileModel interface {
		Insert(ctx context.Context, session sqlx.Session, data *File) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*File, error)
		Update(ctx context.Context, session sqlx.Session, data *File) (sql.Result, error)
		UpdateWithVersion(ctx context.Context, session sqlx.Session, data *File) error
		Delete(ctx context.Context, session sqlx.Session, id int64) error
	}

	defaultFileModel struct {
		sqlc.CachedConn
		table string
	}

	File struct {
		Id          int64     `db:"id"`           // 文件ID
		FileName    string    `db:"file_name"`    // 文件名
		FileHash    string    `db:"file_hash"`    // 文件哈希值
		UserId      int64     `db:"user_id"`      // 用户id
		StoreId     int64     `db:"store_id"`     // 文件仓库ID
		FilePath    string    `db:"file_path"`    // 文件存储路径
		Size        int64     `db:"size"`         // 文件大小
		Postfix     string    `db:"postfix"`      // 文件后缀
		DownloadNum int64     `db:"download_num"` // 下载次数
		UpdateTime  time.Time `db:"update_time"`  // 上传时间
		CreateTime  time.Time `db:"create_time"`
		DelState    int64     `db:"del_state"`
		DeleteTime  time.Time `db:"delete_time"`
		Version     int64     `db:"version"` // 乐观锁版本号
	}
)

func newFileModel(conn sqlx.SqlConn, c cache.CacheConf) *defaultFileModel {
	return &defaultFileModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      "`file`",
	}
}

func (m *defaultFileModel) Insert(ctx context.Context, session sqlx.Session, data *File) (sql.Result, error) {
	data.DeleteTime = time.Unix(0, 0)
	cloudDiskFileIdKey := fmt.Sprintf("%s%v", cacheCloudDiskFileIdPrefix, data.Id)
	return m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", m.table, fileRowsExpectAutoSet)
		if session != nil {
			return session.ExecCtx(ctx, query, data.FileName, data.FileHash, data.UserId, data.StoreId, data.FilePath, data.Size, data.Postfix, data.DownloadNum, data.DelState, data.DeleteTime, data.Version)
		}
		return conn.ExecCtx(ctx, query, data.FileName, data.FileHash, data.UserId, data.StoreId, data.FilePath, data.Size, data.Postfix, data.DownloadNum, data.DelState, data.DeleteTime, data.Version)
	}, cloudDiskFileIdKey)
}

func (m *defaultFileModel) FindOne(ctx context.Context, id int64) (*File, error) {
	cloudDiskFileIdKey := fmt.Sprintf("%s%v", cacheCloudDiskFileIdPrefix, id)
	var resp File
	err := m.QueryRowCtx(ctx, &resp, cloudDiskFileIdKey, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? and del_state = ? limit 1", fileRows, m.table)
		return conn.QueryRowCtx(ctx, v, query, id, globalkey.DelStateNo)
	})
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultFileModel) Update(ctx context.Context, session sqlx.Session, data *File) (sql.Result, error) {
	cloudDiskFileIdKey := fmt.Sprintf("%s%v", cacheCloudDiskFileIdPrefix, data.Id)
	return m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, fileRowsWithPlaceHolder)
		if session != nil {
			return session.ExecCtx(ctx, query, data.FileName, data.FileHash, data.UserId, data.StoreId, data.FilePath, data.Size, data.Postfix, data.DownloadNum, data.DelState, data.DeleteTime, data.Version, data.Id)
		}
		return conn.ExecCtx(ctx, query, data.FileName, data.FileHash, data.UserId, data.StoreId, data.FilePath, data.Size, data.Postfix, data.DownloadNum, data.DelState, data.DeleteTime, data.Version, data.Id)
	}, cloudDiskFileIdKey)
}

func (m *defaultFileModel) UpdateWithVersion(ctx context.Context, session sqlx.Session, data *File) error {

	oldVersion := data.Version
	data.Version += 1

	var sqlResult sql.Result
	var err error

	cloudDiskFileIdKey := fmt.Sprintf("%s%v", cacheCloudDiskFileIdPrefix, data.Id)
	sqlResult, err = m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ? and version = ? ", m.table, fileRowsWithPlaceHolder)
		if session != nil {
			return session.ExecCtx(ctx, query, data.FileName, data.FileHash, data.UserId, data.StoreId, data.FilePath, data.Size, data.Postfix, data.DownloadNum, data.DelState, data.DeleteTime, data.Version, data.Id, oldVersion)
		}
		return conn.ExecCtx(ctx, query, data.FileName, data.FileHash, data.UserId, data.StoreId, data.FilePath, data.Size, data.Postfix, data.DownloadNum, data.DelState, data.DeleteTime, data.Version, data.Id, oldVersion)
	}, cloudDiskFileIdKey)
	if err != nil {
		return err
	}
	updateCount, err := sqlResult.RowsAffected()
	if err != nil {
		return err
	}
	if updateCount == 0 {
		return xerr.NewErrCode(xerr.DB_UPDATE_AFFECTED_ZERO_ERROR)
	}

	return nil
}

func (m *defaultFileModel) Delete(ctx context.Context, session sqlx.Session, id int64) error {
	cloudDiskFileIdKey := fmt.Sprintf("%s%v", cacheCloudDiskFileIdPrefix, id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		if session != nil {
			return session.ExecCtx(ctx, query, id)
		}
		return conn.ExecCtx(ctx, query, id)
	}, cloudDiskFileIdKey)
	return err
}

func (m *defaultFileModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", cacheCloudDiskFileIdPrefix, primary)
}
func (m *defaultFileModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary interface{}) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? and del_state = ? limit 1", fileRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary, globalkey.DelStateNo)
}

func (m *defaultFileModel) tableName() string {
	return m.table
}
