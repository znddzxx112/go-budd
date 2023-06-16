package main

import (
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli"
	"github.com/znddzxx112/go-budd/pkg/utils"
	"github.com/znddzxx112/go-budd/server"
	"os"
	"os/signal"
	"runtime"
	"syscall"
	"time"
)

// @title budd
// @version latest
// @description budd
// @author znddzxx112@163.com
// @schemes {{scheme}}

const COMMON_VERSION = "v1.1"
const COMMON_NAME = "budd"

func main() {

	runtime.GOMAXPROCS(runtime.NumCPU())

	utils.SetLogLevel()
	utils.SetDefaultLocation()

	app := cli.NewApp()
	app.Name = COMMON_NAME
	app.Version = COMMON_VERSION
	configPathFlag := cli.StringFlag{
		Name:  "conf, c",
		Usage: "config file path",
		Value: "./conf/config.yaml",
	}
	portFlag := cli.StringFlag{
		Name:  "port, p",
		Usage: "port",
		Value: "7890",
	}
	app.Flags = []cli.Flag{
		configPathFlag,
		portFlag,
	}
	app.Action = Start
	err := app.Run(os.Args)
	if err != nil {
		logrus.Fatal(err)
	}
}

func Start(ctx *cli.Context) error {
	sigCh := make(chan os.Signal)
	defer close(sigCh)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGKILL, syscall.SIGTERM)

	conf := ctx.String("conf")
	port := ctx.String("port")

	// 1、打印配置文件和端口信息
	logrus.Infof("conf is %s", conf)
	logrus.Infof("port is %s", port)

	// 2、打印启动信息
	utils.WelcomeLog(COMMON_NAME, COMMON_VERSION)

	// 3、服务器启动
	defaultServer := server.NewServer(COMMON_NAME)
	err := defaultServer.Run(port, conf)
	if err != nil {
		return err
	}

	// 4、监听信号，服务器退出
	go func() {
		select {
		case <-sigCh:
			if err := defaultServer.Close(); err != nil {
				logrus.Println(err.Error())
			}
			time.Sleep(time.Second)
			utils.ExitLog(COMMON_NAME, COMMON_VERSION)
			os.Exit(0)
		}
	}()

	return nil
}
