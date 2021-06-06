package controller

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/jmoiron/sqlx"
)

type ShopController struct {
	beego.Controller
}

type ShopConfig struct {
	RedisAddr string
	RedisPort string

	DbDriver   string
	DbUser     string
	DbPwd      string
	DbHost     string
	DbPort     string
	DbDatabase string
	DbProtocol string
}

type Product struct {
	Id          int `json:"id"`
	Name        int `json:"name"`
	Cid         int `json:"cid"`
	PhotoLittle int `json:"photo_little"`
}

var (
	shopConfig = &ShopConfig{}
	Db         *sqlx.DB
	err        error
	product    []Product
	// 限制显示记录
	limit int
	// 执行结果
	executeResult map[string]interface{}
)

func initAppConfig() *ShopConfig {
	redisAddr := beego.AppConfig.String("redis_addr")
	redisPort := beego.AppConfig.String("redis_port")
	dbDriver := beego.AppConfig.String("db_driver")
	dbUser := beego.AppConfig.String("db_user")
	dbPwd := beego.AppConfig.String("db_pwd")
	dbHost := beego.AppConfig.String("db_host")
	dbPort := beego.AppConfig.String("db_port")
	dbDatabase := beego.AppConfig.String("db_database")
	dbProtocol := beego.AppConfig.String("db_protocol")

	shopConfig.RedisAddr = redisAddr
	shopConfig.RedisPort = redisPort
	shopConfig.DbDriver = dbDriver
	shopConfig.DbUser = dbUser
	shopConfig.DbPwd = dbPwd
	shopConfig.DbHost = dbHost
	shopConfig.DbPort = dbPort
	shopConfig.DbDatabase = dbDatabase
	shopConfig.DbProtocol = dbProtocol

	return shopConfig
}

func init() {
	initAppConfig()
	sqlPrameter := fmt.Sprintf("%s:%s@%s(%s:%d)/%s", shopConfig.DbUser, shopConfig.DbPwd, shopConfig.DbProtocol, shopConfig.DbHost, shopConfig.DbPort, shopConfig.DbDatabase)
	if Db, err = sqlx.Open(shopConfig.DbDriver, sqlPrameter); err != nil {
		logs.Debug("数据库连接失败，失败原因：", err.Error())
		return
	}
}

// 通过sqlx操作mysql,crud
func (this *ShopController) GetProduct() {
	executeResult = make(map[string]interface{})
	executeResult["status"] = 1
	executeResult["msg"] = "success"
	executeResult["data"] = []string{"hello", "world"}

	this.Data["json"] = executeResult
	this.ServeJSON()
	return

	if limit, err = this.GetInt("limit"); err != nil {
		logs.Debug("this.GetInt err:", err.Error())
		return
	}
	Db.Select(&product, "select id,name ,cid,photo_little from product limit ?", limit)
}

func (this *ShopController) TestBeego() {
	var names []string = []string{"minishop", "lonzee", "golang"}
	this.Data["json"] = names

	this.ServeJSON()
}
