package models

import (
	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/config"
	"github.com/beego/beego/v2/core/logs"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	driverName, err := config.String("driverName")
	err = orm.RegisterDriver(driverName, orm.DRMySQL)
	if err != nil {
		logs.Error(err.Error())
	}
	userName, _ := config.String("mysqlUser")
	pwd, _ := config.String("mysqlPwd")
	host, _ := config.String("host")
	port, _ := config.String("port")
	dbName, _ := config.String("dbName")
	//root:123456@tcp(127.0.0.1:3306)/beego?charset=utf8"
	logoInfo := userName + ":" + pwd + "@tcp(" + host + ":" + port + ")/" + dbName + "?charset = utf8"
	err = orm.RegisterDataBase("default", driverName, logoInfo)
	if err != nil {
		logs.Error("连接数据库错误：" + err.Error())
	}
	logs.Info("连接数据库成功")
}
