package database

import (
	"database/sql"
	"errors"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
	"valtec/pkg/config"
)

var gormDb *gorm.DB

func Count(model any) int64 {
	var res int64
	db := gormDb.Model(model).Count(&res)
	if db.Error != nil && !errors.Is(db.Error, gorm.ErrRecordNotFound) {
		panic("database error: " + db.Error.Error())
	}
	return res
}

func SelectOne(where, result any, colums ...string) *gorm.DB {
	db := gormDb.Select(colums).Where(where).Take(result)
	if db.Error != nil && !errors.Is(db.Error, gorm.ErrRecordNotFound) {
		panic("database error: " + db.Error.Error())
	}
	return db
}

func SelectFirst(where, result any, colums ...string) *gorm.DB {
	db := gormDb.Select(colums).Where(where).First(result)
	if db.Error != nil && !errors.Is(db.Error, gorm.ErrRecordNotFound) {
		panic("database error: " + db.Error.Error())
	}
	return db
}

func SelectLast(where, result any, colums ...string) *gorm.DB {
	db := gormDb.Select(colums).Where(where).Last(result)
	if db.Error != nil && !errors.Is(db.Error, gorm.ErrRecordNotFound) {
		panic("database error: " + db.Error.Error())
	}
	return db
}

func Select(where any, result any, colums ...string) *gorm.DB {
	db := gormDb.Select(colums).Where(where).Find(result)
	if db.Error != nil && !errors.Is(db.Error, gorm.ErrRecordNotFound) {
		panic("database error: " + db.Error.Error())
	}
	return db
}

func AddOne(a any) *gorm.DB {
	res := gormDb.Create(a)
	if res.Error != nil {
		panic("database error: " + res.Error.Error())
	}
	return res
}

//
//func (d dataBase[T]) InsertAll(a []*T) {
//	res := d.gormDb.C(a)
//	if res.Error != nil {
//		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
//			return nil
//		} else {
//			panic("database error: " + res.Error.Error())
//		}
//	}
//	return a
//}

func init() {
	username := config.GetStringSevere("db", "username")
	password := config.GetStringSevere("db", "password")
	host := config.GetStringSevere("db", "host")
	port := config.GetStringSevere("db", "port")
	database := config.GetStringSevere("db", "database")
	maxIdleConns := config.GetIntSevere("db", "maxIdleConns")
	maxOpenConns := config.GetIntSevere("db", "maxOpenConns")
	connMaxLifetime := config.GetIntSevere("db", "connMaxLifetime")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", username, password, host, port, database)
	dbConn, err := sql.Open("mysql", dsn)
	dbConn.SetMaxIdleConns(maxIdleConns)
	dbConn.SetMaxOpenConns(maxOpenConns)
	dbConn.SetConnMaxLifetime(time.Duration(connMaxLifetime) * time.Second)
	if err != nil {
		panic("database init error: " + err.Error())
	}
	gormDb, err = gorm.Open(mysql.New(mysql.Config{
		Conn:                      dbConn,
		DefaultStringSize:         256,   // string 类型字段的默认长度
		DisableDatetimePrecision:  true,  // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,  // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,  // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false, // 根据当前 MySQL 版本自动配置
	}), &gorm.Config{})
}
