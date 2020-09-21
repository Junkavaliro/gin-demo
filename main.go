package main

import (
	"github.com/gin-gonic/gin"
	"github.com/go-programming-tour-book/blog-service/global"
	"github.com/go-programming-tour-book/blog-service/internal/model"
	"github.com/go-programming-tour-book/blog-service/internal/routers"
	"github.com/go-programming-tour-book/blog-service/pkg/logger"
	"github.com/go-programming-tour-book/blog-service/pkg/setting"
	"gopkg.in/natefinch/lumberjack.v2"
	"log"
	"net/http"
	"time"
)

func init() {
	err := setupSetting()
	if err != nil {
		log.Fatalf("init setupSetting err: %v", err)
	}

	err = setupDBEngine()
	if err != nil {
		log.Fatalf("init setupDBEngine err: %v", err)
	}
	_ = setupLogger()
}

func main() {
	gin.SetMode(global.ServerSetting.RunMode)
	router := routers.NewRouter()
	s := &http.Server{
		Addr: ":" + global.ServerSetting.HttpPort,
		Handler: router,
		ReadTimeout: global.ServerSetting.ReadTimeout,
		WriteTimeout: global.ServerSetting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}
	global.Logger.Debugf("啓動項目：%s", "blog_service")
	_ = s.ListenAndServe()
}

func setupSetting() error {
	setting, err := setting.NewSetting()
	if err != nil {
		return err
	}
	err = setting.ReadSection("Server", &global.ServerSetting)
	if err != nil {
		return err
	}
	log.Printf("&global.ServerSetting: %v\n", global.ServerSetting)
	err = setting.ReadSection("App", &global.AppSetting)
	if err != nil {
		return err
	}
	log.Printf("&global.AppSetting: %v\n", global.AppSetting)
	err = setting.ReadSection("Database", &global.DatabaseSetting)
	if err != nil {
		return err
	}
	log.Printf("&global.DatabaseSetting: %v\n", global.DatabaseSetting)
	global.ServerSetting.ReadTimeout *= time.Second
	global.ServerSetting.WriteTimeout *= time.Second
	return nil
}

func setupDBEngine() error {
	var err error
	global.DBEngine, err = model.NewDBEngine(global.DatabaseSetting)
	if err != nil {
		return err
	}
	return nil
}

func setupLogger() error {
	global.Logger = logger.NewLogger(&lumberjack.Logger{
		Filename: global.AppSetting.LogSavePath + "/" +
			global.AppSetting.LogFileName + global.AppSetting.LogFileExt,
			MaxSize: 600,
			MaxAge: 10,
			LocalTime: true,
	}, "", log.LstdFlags).WithCaller(2)
	return nil
}