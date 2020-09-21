package model

import (
	"fmt"
	"github.com/go-programming-tour-book/blog-service/global"
	"github.com/go-programming-tour-book/blog-service/pkg/setting"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

type Model struct {
	ID uint32 `gorm:"primary_key" json:"id"`
	CreatedBy string `json:"created_by"`
	ModifiedBy string `json:"modified_by"`
	CreatedOn uint32 `json:"created_on"`
	ModifiedOn uint32 `json:"modified_on"`
	DeletedOn uint32 `json:"deleted_on"`
	IsDel uint8 `json:"is_del"`
}

func NewDBEngine(s *setting.DatabaseSettingS) (*gorm.DB, error) {
	db, err := gorm.Open(s.DBType,
		fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=%s&parseTime=%t&loc=Local",
			s.Username,
			s.Password,
			s.Host,
			s.DBName,
			s.Charset,
			s.ParseTime))
	if err != nil {
		return nil, err
	}
	if global.ServerSetting.RunMode == "debug" {
		db.LogMode(true)
	}
	db.SingularTable(true)
	db.DB().SetMaxIdleConns(s.MaxIdleConns)
	db.DB().SetMaxOpenConns(s.MaxOpenConns)
	return db, nil
}