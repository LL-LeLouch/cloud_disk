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
	folderFieldNames          = builder.RawFieldNames(&Folder{})
	folderRows                = strings.Join(folderFieldNames, ",")
	folderRowsExpectAutoSet   = strings.Join(stringx.Remove(folderFieldNames, "`id`", "`create_time`", "`update_time`"), ",")
	folderRowsWithPlaceHolder = strings.Join(stringx.Remove(folderFieldNames, "`id`", "`create_time`", "`update_time`"), "=?,") + "=?"

	cacheCloudDiskFolderIdPrefix = "cache:cloudDisk:folder:id:"
)

type (
	folderModel interface {
		Insert(ctx context.Context, session sqlx.Session, data *Folder) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*Folder, error)
		Update(ctx context.Context, session sqlx.Session, data *Folder) (sql.Result, error)
		UpdateWithVersion(ctx context.Context, session sqlx.Session, data *Folder) error
		Delete(ctx context.Context, session sqlx.Session, id int64) error
	}

	defaultFolderModel struct {
		sqlc.CachedConn
		table string
	}

	Folder struct {
		Id         int64     `db:"id"`          // 文件夹ID
		FolderName string    `db:"folder_name"` // 文件夹名称
		UserId     int64     `db:"user_id"`     // 用户id
		StoreId    int64     `db:"store_id"`    // 所属文件仓库ID
		FolderPath string    `db:"folder_path"` // 文件夹路径
		CreateTime time.Time `db:"create_time"` // 创建时间
		UpdateTime time.Time `db:"update_time"`
		DelState   int64     `db:"del_state"`
		DeleteTime time.Time `db:"delete_time"`
		Version    int64     `db:"version"` // 乐观锁版本号
	}
)

func newFolderModel(conn sqlx.SqlConn, c cache.CacheConf) *defaultFolderModel {
	return &defaultFolderModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      "`folder`",
	}
}

func (m *defaultFolderModel) Insert(ctx context.Context, session sqlx.Session, data *Folder) (sql.Result, error) {
	data.DeleteTime = time.Unix(0, 0)
	cloudDiskFolderIdKey := fmt.Sprintf("%s%v", cacheCloudDiskFolderIdPrefix, data.Id)
	return m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?)", m.table, folderRowsExpectAutoSet)
		if session != nil {
			return session.ExecCtx(ctx, query, data.FolderName, data.UserId, data.StoreId, data.FolderPath, data.DelState, data.DeleteTime, data.Version)
		}
		return conn.ExecCtx(ctx, query, data.FolderName, data.UserId, data.StoreId, data.FolderPath, data.DelState, data.DeleteTime, data.Version)
	}, cloudDiskFolderIdKey)
}

func (m *defaultFolderModel) FindOne(ctx context.Context, id int64) (*Folder, error) {
	cloudDiskFolderIdKey := fmt.Sprintf("%s%v", cacheCloudDiskFolderIdPrefix, id)
	var resp Folder
	err := m.QueryRowCtx(ctx, &resp, cloudDiskFolderIdKey, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? and del_state = ? limit 1", folderRows, m.table)
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

func (m *defaultFolderModel) Update(ctx context.Context, session sqlx.Session, data *Folder) (sql.Result, error) {
	cloudDiskFolderIdKey := fmt.Sprintf("%s%v", cacheCloudDiskFolderIdPrefix, data.Id)
	return m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, folderRowsWithPlaceHolder)
		if session != nil {
			return session.ExecCtx(ctx, query, data.FolderName, data.UserId, data.StoreId, data.FolderPath, data.DelState, data.DeleteTime, data.Version, data.Id)
		}
		return conn.ExecCtx(ctx, query, data.FolderName, data.UserId, data.StoreId, data.FolderPath, data.DelState, data.DeleteTime, data.Version, data.Id)
	}, cloudDiskFolderIdKey)
}

func (m *defaultFolderModel) UpdateWithVersion(ctx context.Context, session sqlx.Session, data *Folder) error {

	oldVersion := data.Version
	data.Version += 1

	var sqlResult sql.Result
	var err error

	cloudDiskFolderIdKey := fmt.Sprintf("%s%v", cacheCloudDiskFolderIdPrefix, data.Id)
	sqlResult, err = m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ? and version = ? ", m.table, folderRowsWithPlaceHolder)
		if session != nil {
			return session.ExecCtx(ctx, query, data.FolderName, data.UserId, data.StoreId, data.FolderPath, data.DelState, data.DeleteTime, data.Version, data.Id, oldVersion)
		}
		return conn.ExecCtx(ctx, query, data.FolderName, data.UserId, data.StoreId, data.FolderPath, data.DelState, data.DeleteTime, data.Version, data.Id, oldVersion)
	}, cloudDiskFolderIdKey)
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

func (m *defaultFolderModel) Delete(ctx context.Context, session sqlx.Session, id int64) error {
	cloudDiskFolderIdKey := fmt.Sprintf("%s%v", cacheCloudDiskFolderIdPrefix, id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		if session != nil {
			return session.ExecCtx(ctx, query, id)
		}
		return conn.ExecCtx(ctx, query, id)
	}, cloudDiskFolderIdKey)
	return err
}

func (m *defaultFolderModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", cacheCloudDiskFolderIdPrefix, primary)
}
func (m *defaultFolderModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary interface{}) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? and del_state = ? limit 1", folderRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary, globalkey.DelStateNo)
}

func (m *defaultFolderModel) tableName() string {
	return m.table
}
