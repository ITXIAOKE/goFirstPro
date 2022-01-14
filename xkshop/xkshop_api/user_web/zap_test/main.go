package main

import (
	"fmt"
	"go.uber.org/zap"
	"time"
)

func main() {
	logger, _ := zap.NewProduction() //生产环境
	//zap.NewDevelopment()//开发环境

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
