package model

import (
	"time"

	"gorm.io/gorm"
)

type PublicRepository struct {
	Id                 int
	Identity           string
	UserIdentity       string
	ParentId           int64
	RepositoryIdentity string
	Ext                string
	Name               string
	CreatedAt          time.Time      `gorm:"created"`
	UpdatedAt          time.Time      `gorm:"updated"`
	DeletedAt          gorm.DeletedAt `gorm:"deleted"`
}

func (table PublicRepository) TableName() string {
	return "public_repository"
}
