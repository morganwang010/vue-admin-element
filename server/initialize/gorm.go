package initialize

import (
	"fmt"
	"time"
	"vue-admin-element/global"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

// 初始化MySQl数据库
func Pgsql() {
	m := global.Config.Pgsql
	// s := "%s:%s@tcp(%s:%v)/%s?charset=utf8&parseTime=True&loc=Local"

	// var dsn = fmt.Sprintf(s, m.Username, m.Password, m.Host, m.Port, m.Dbname)
	// fmt.Printf("********" + dsn)
	dsn := "host=localhost port=5432 user=postgres password=myS3curepass dbname=cms"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{SingularTable: true},
	})
	if err != nil {
		fmt.Printf("1111pgsql error: %s", err)
		return
	}
	sqlDb, err := db.DB()
	if err != nil {
		fmt.Printf("pgsql error: %s", err)
	}
	sqlDb.SetMaxIdleConns(m.MaxIdleConns)
	sqlDb.SetMaxOpenConns(m.MaxOpenConns)
	sqlDb.SetConnMaxLifetime(time.Duration(m.ConnMaxLifetime))
	global.Db = db
}
