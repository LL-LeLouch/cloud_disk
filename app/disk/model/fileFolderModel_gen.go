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
	fileFolderFieldNames          = builder.RawFieldNames(&FileFolder{})
	fileFolderRows                = strings.Join(fileFolderFieldNames, ",")
	fileFolderRowsExpectAutoSet   = strings.Join(stringx.Remove(fileFolderFieldNames, "`id`", "`create_time`", "`update_time`"), ",")
	fileFolderRowsWithPlaceHolder = strings.Join(stringx.Remove(fileFolderFieldNames, "`id`", "`create_time`", "`update_time`"), "=?,") + "=?"

	cacheCloudDiskFileFolderIdPrefix = "cache:cloudDisk:fileFolder:id:"
)

type (
	fileFolderModel interface {
		Insert(ctx context.Context, session sqlx.Session, data *FileFolder) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*FileFolder, error)
		Update(ctx context.Context, session sqlx.Session, data *FileFolder) (sql.Result, error)
		UpdateWithVersion(ctx context.Context, session sqlx.Session, data *FileFolder) error
		Delete(ctx context.Context, session sqlx.Session, id int64) error
	}

	defaultFileFolderModel struct {
		sqlc.CachedConn
		table string
	}

	FileFolder struct {
		Id             int64          `db:"id"`               // 文件夹ID
		FileFolderName sql.NullString `db:"file_folder_name"` // 文件夹名称
		ParentFolderId int64          `db:"parent_folder_id"` // 父文件夹ID
		FileStoreId    sql.NullInt64  `db:"file_store_id"`    // 所属文件仓库ID
		Time           sql.NullString `db:"time"`             // 创建时间
		Version        int64          `db:"version"`          // 乐观锁版本号
		DeleteTime     time.Time      `db:"delete_time"`
		DelState       int64          `db:"del_state"`
	}
)

func newFileFolderModel(conn sqlx.SqlConn, c cache.CacheConf) *defaultFileFolderModel {
	return &defaultFileFolderModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      "`file_folder`",
	}
}

func (m *defaultFileFolderModel) Insert(ctx context.Context, session sqlx.Session, data *FileFolder) (sql.Result, error) {
	data.DeleteTime = time.Unix(0, 0)
	cloudDiskFileFolderIdKey := fmt.Sprintf("%s%v", cacheCloudDiskFileFolderIdPrefix, data.Id)
	return m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?)", m.table, fileFolderRowsExpectAutoSet)
		if session != nil {
			return session.ExecCtx(ctx, query, data.FileFolderName, data.ParentFolderId, data.FileStoreId, data.Time, data.Version, data.DeleteTime, data.DelState)
		}
		return conn.ExecCtx(ctx, query, data.FileFolderName, data.ParentFolderId, data.FileStoreId, data.Time, data.Version, data.DeleteTime, data.DelState)
	}, cloudDiskFileFolderIdKey)
}

func (m *defaultFileFolderModel) FindOne(ctx context.Context, id int64) (*FileFolder, error) {
	cloudDiskFileFolderIdKey := fmt.Sprintf("%s%v", cacheCloudDiskFileFolderIdPrefix, id)
	var resp FileFolder
	err := m.QueryRowCtx(ctx, &resp, cloudDiskFileFolderIdKey, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? and del_state = ? limit 1", fileFolderRows, m.table)
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

func (m *defaultFileFolderModel) Update(ctx context.Context, session sqlx.Session, data *FileFolder) (sql.Result, error) {
	cloudDiskFileFolderIdKey := fmt.Sprintf("%s%v", cacheCloudDiskFileFolderIdPrefix, data.Id)
	return m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, fileFolderRowsWithPlaceHolder)
		if session != nil {
			return session.ExecCtx(ctx, query, data.FileFolderName, data.ParentFolderId, data.FileStoreId, data.Time, data.Version, data.DeleteTime, data.DelState, data.Id)
		}
		return conn.ExecCtx(ctx, query, data.FileFolderName, data.ParentFolderId, data.FileStoreId, data.Time, data.Version, data.DeleteTime, data.DelState, data.Id)
	}, cloudDiskFileFolderIdKey)
}

func (m *defaultFileFolderModel) UpdateWithVersion(ctx context.Context, session sqlx.Session, data *FileFolder) error {

	oldVersion := data.Version
	data.Version += 1

	var sqlResult sql.Result
	var err error

	cloudDiskFileFolderIdKey := fmt.Sprintf("%s%v", cacheCloudDiskFileFolderIdPrefix, data.Id)
	sqlResult, err = m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ? and version = ? ", m.table, fileFolderRowsWithPlaceHolder)
		if session != nil {
			return session.ExecCtx(ctx, query, data.FileFolderName, data.ParentFolderId, data.FileStoreId, data.Time, data.Version, data.DeleteTime, data.DelState, data.Id, oldVersion)
		}
		return conn.ExecCtx(ctx, query, data.FileFolderName, data.ParentFolderId, data.FileStoreId, data.Time, data.Version, data.DeleteTime, data.DelState, data.Id, oldVersion)
	}, cloudDiskFileFolderIdKey)
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

func (m *defaultFileFolderModel) Delete(ctx context.Context, session sqlx.Session, id int64) error {
	cloudDiskFileFolderIdKey := fmt.Sprintf("%s%v", cacheCloudDiskFileFolderIdPrefix, id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		if session != nil {
			return session.ExecCtx(ctx, query, id)
		}
		return conn.ExecCtx(ctx, query, id)
	}, cloudDiskFileFolderIdKey)
	return err
}

func (m *defaultFileFolderModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", cacheCloudDiskFileFolderIdPrefix, primary)
}
func (m *defaultFileFolderModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary interface{}) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? and del_state = ? limit 1", fileFolderRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary, globalkey.DelStateNo)
}

func (m *defaultFileFolderModel) tableName() string {
	return m.table
}
