package psql

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/nizonglonggit/userapisrv/pkg/config"
	"github.com/nizonglonggit/userapisrv/pkg/define"
)

var (
	ExampleDB *gorm.DB
)

// 创建 example 连接
func NewExampleDBConn(tomlConfig *config.Config) {
	// 读取配置
	var err error
	dbConfig, ok := tomlConfig.DBServerConf(define.DBExample)
	if !ok {
		panic(fmt.Sprintf("Postgres: %v no set.", define.DBExample))
	}

	ExampleDB, err = gorm.Open("postgres", dbConfig.ConnectString())
	if err != nil {
		panic(fmt.Sprintf("gorm.Open: err:%v", err))
	}
	// 设置最大链接数
	ExampleDB.DB().SetMaxOpenConns(10)
}
