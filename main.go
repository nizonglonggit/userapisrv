package main

import (
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/nizonglonggit/userapisrv/pkg/config"
	"github.com/nizonglonggit/userapisrv/pkg/model/psql"
	"github.com/nizonglonggit/userapisrv/pkg/router"
	"log"
	"os"
	"os/signal"
	"runtime"
	"runtime/debug"
	"syscall"
)

var (
	tomlFilePath string
	mode         string
)

// init 初始化配置
func init() {
	flag.StringVar(&tomlFilePath, "config", "docs/local.toml", "服务配置文件")
	flag.StringVar(&mode, "mode", gin.DebugMode, "模型-debug还是release还是test")
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func main() {
	defer func() {
		// 停止服务
		log.Fatal("stop server!")
		if err := recover(); err != nil {
			log.Fatal(fmt.Sprintf("error: %s Stack: %s", err, string(debug.Stack())))
			os.Exit(1)
		}
	}()

	flag.Parse()

	// 解析配置文件
	cfg, err := config.UnmarshalConfig(tomlFilePath)
	if err != nil {
		panic(err)
	}

	// 初始化路由
	engine := router.InitEngine(mode)

	// 初始化数据库连接
	initDB(cfg)

	// 启动web服务
	fmt.Println(fmt.Sprintf("userapisrv server run %s", cfg.GetListenAddr()))
	go engine.Run(cfg.GetListenAddr())

	// 阻塞主进程 等待退出信号
	signalChan := make(chan os.Signal, 0)
	signal.Notify(signalChan, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT, syscall.SIGSTOP, syscall.SIGKILL)
	<-signalChan
}

// 初始化数据库连接
func initDB(cfg *config.Config) {
	// db
	psql.NewExampleDBConn(cfg)

}
