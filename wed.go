package wed

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/qiniu/qmgo"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

var Mongodb *qmgo.Database
var config AllConfig

func init() {
	const configFile = "./config.yaml"
	if _, err := os.Stat(configFile); err != nil {
		downConfig()
		fmt.Println("miss config.yaml file")
		//os.Exit(1)
	}
	viper := viper.New()
	viper.SetConfigFile(configFile)
	viper.ReadInConfig()
	viper.MergeInConfig()
	viper.Unmarshal(&config)
}
func downConfig() {
	url := "https://github.com/smart817/wed/config.yaml"
	res, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	f, err := os.Create("config.yaml")
	if err != nil {
		panic(err)
	}
	io.Copy(f, res.Body)
}

func Run(f func(r *gin.Engine)) {
	if config.App.Debug {
		gin.SetMode("debug")
	} else {
		gin.SetMode("release")
	}
	r := gin.New()
	r.SetTrustedProxies(nil)

	f(r)
	if config.App.Port == "" {
		config.App.Port = "8080"
	}
	r.Run(fmt.Sprintf(":%v", config.App.Port))
}
func Mongo() {
	ctx := context.Background()
	client, err := qmgo.NewClient(ctx, &qmgo.Config{Uri: config.Mongodb.Uri})
	if err != nil {
		fmt.Println("mongodb err")
		os.Exit(1)
	}
	fmt.Println("mongodb 连接成功")
	db := client.Database(config.Mongodb.DB)
	Mongodb = db
	cli := db.Collection("user")
	cli.EnsureIndexes(ctx, []string{"phoneNo"}, []string{}) //唯一索引
}

var MysqlDb *gorm.DB
var MysqlError error

func Mysql() {
	//dns := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", config.Mysql.UserName, config.Mysql.Password, config.Mysql.Addr, config.Mysql.DB)
	MysqlDb, MysqlError = gorm.Open(mysql.New(mysql.Config{
		DSN:               config.Mysql.Uri,
		DefaultStringSize: 256, // string 类型字段的默认长度
	}), &gorm.Config{
		// gorm日志模式：silent
		Logger: logger.Default.LogMode(logger.Silent),
		// 外键约束
		DisableForeignKeyConstraintWhenMigrating: true,
		// 禁用默认事务（提高运行速度）
		//SkipDefaultTransaction: true,
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "gee_", // 表名前缀，`User` 的表名应该是 `t_users`
			SingularTable: true,   // 使用单数表名，启用该选项，此时，`User` 的表名应该是 `gee_user`
		},
	})

	if MysqlError != nil {
		fmt.Println("Myslq连接失败：", MysqlError)
		os.Exit(1)
	}

	// 迁移数据表，在没有数据表结构变更时候，建议注释不执行
	//_ = db.AutoMigrate(&User{}, &Article{}, &Category{}, Profile{}, Comment{})

	sqlDB, _ := MysqlDb.DB()
	// SetMaxIdleCons 设置连接池中的最大闲置连接数。
	sqlDB.SetMaxIdleConns(10)

	// SetMaxOpenCons 设置数据库的最大连接数量。
	sqlDB.SetMaxOpenConns(100)

	// SetConnMaxLifetiment 设置连接的最大可复用时间。
	sqlDB.SetConnMaxLifetime(10 * time.Second)
}
