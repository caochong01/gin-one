package config

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var Db *gorm.DB
var err error

func init() {
	// TODO 数据库练习
	// 参考 https://github.com/go-sql-driver/mysql#dsn-data-source-name 获取详情
	//dsn := "root:caochong0521@tcp(127.0.0.1:3306)/gogogo?charset=utf8mb4&parseTime=True&loc=Local"
	//_, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	Db, err = gorm.Open(mysql.New(mysql.Config{
		DSN: "root:caochong0521@tcp(127.0.0.1:3306)" +
			"/gogogo?charset=utf8mb4&parseTime=True&loc=Local", // DSN data source name
		DefaultStringSize:         1024,  // string 类型字段的默认长度
		DisableDatetimePrecision:  true,  // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,  // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,  // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false, // 根据当前 MySQL 版本自动配置
	}), &gorm.Config{NamingStrategy: schema.NamingStrategy{
		TablePrefix:   "",   // 表名前缀`t_`，则`User` 的表名应该是 `t_users`
		SingularTable: true, // 使用单数表名，启用该选项，此时，`User` 的表名应该是 `user`
	}})
	if err != nil {
		fmt.Println("数据库连接异常错误：")
		panic(err)
	}
}
