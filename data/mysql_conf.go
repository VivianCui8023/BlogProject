package data

import (
	"BeeProject/models"
	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/config"
	"github.com/beego/beego/v2/core/logs"
	_ "github.com/go-sql-driver/mysql"
)

func init() {

	orm.RegisterModel(new(models.User))

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
	logoInfo := userName + ":" + pwd + "@tcp(" + host + ":" + port + ")/" + dbName + "?charset=utf8"
	err = orm.RegisterDataBase("default", driverName, logoInfo)
	if err != nil {
		logs.Error("连接数据库错误：" + err.Error())
	}
	logs.Info("连接数据库成功")
}

func StartDB() orm.Ormer {
	//设置自动同步表信息
	orm.RunSyncdb("default", false, true)
	//开发环境下设置输出sql语句
	orm.Debug = true
	o := orm.NewOrm()
	return o
}
