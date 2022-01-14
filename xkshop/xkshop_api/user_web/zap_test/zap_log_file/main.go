package main

import (
	"fmt"
	"go.uber.org/zap"
	"time"
)

func NewLogger() (*zap.Logger, error) {
	cfg := zap.NewProductionConfig()
	cfg.OutputPaths = []string{ //切片
		"./xk.log",
		//"stderr", //让日志在控制台输出
		"stdout", //让日志在控制台输出
	}
	return cfg.Build()
}

func main() {
	logger, err := NewLogger()
	if err != nil {
		panic(err)
	}
	defer logger.Sync() // flushes buffer, if any

	url := "https://imooc.com"
	sugar := logger.Sugar()
	sugar.Infow("failed to fetch URL",
		// Structured context as loosely typed key-value pairs.
		"url", url,
		"attempt", 3,
		"backoff", time.Second,
	)
	sugar.Infof("Failed to fetch URL: %s", url)
	fmt.Println("==============================")

	logger.Info("kekekekkek")
	fmt.Println("==============================")

	logger.Info("kekekekkek",
		zap.String("url", url),
		zap.Int("nums", 333),
	)
}
