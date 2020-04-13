package initdbase

import (
	//"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql" // 注意此行必须加
)

// 初始化连接数据库
func init()  {

	/*
	//数据库ip
	dbHost := beego.AppConfig.String("mysql::db_host")
	//数据库端口
	dbPort := beego.AppConfig.String("mysql::db_port")
	//数据库名称
	dbDatabase := beego.AppConfig.String("mysql::db_database")
	//数据库用户名
	dbUsername := beego.AppConfig.String("mysql::db_username")
	//数据库密码
	dbPasword := beego.AppConfig.String("mysql::db_password")
	//数据库字符集
	dbCharset := beego.AppConfig.String("mysql::db_charset")
	*/

	//初始化数据库
	_ = orm.RegisterDriver("mysql", orm.DRMySQL)

	// _ = orm.RegisterDataBase("default", "mysql", dbUsername+":"+dbPasword+"@tcp("+dbHost+":"+dbPort+")/"+dbDatabase+"?charset="+dbCharset)

	_ = orm.RegisterDataBase("default", "mysql", "root:1024@tcp(127.0.0.1:3306)/hilin-name?charset=utf8", 30)

	//设置数据库最大空闲连接
	orm.SetMaxIdleConns("default", 30)
	//设置数据库最大连接数
	orm.SetMaxOpenConns("default", 30)
}