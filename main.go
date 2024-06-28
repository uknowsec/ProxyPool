package main

import (
	"ProxyPool/Config"
	"ProxyPool/Gost"
	"ProxyPool/ProxyApi"
	"log"
	"time"
)

func run(config Config.Config) error {
	// 获取代理服务器信息
	proxyserv, err := ProxyApi.GetProxyServer(config.ProxyApi.URL)
	if err != nil {
		log.Printf("error fetching proxy server: %v", err)
		return err
	}

	// 更新 JSON 数据
	jsondata := Gost.UpdateJson(proxyserv)

	// 更新 Gost 的设置
	err = Gost.UpdateGostHop(config.Gost.ApiURL, jsondata)
	if err != nil {
		log.Printf("error updating Gost: %v", err)
		return err
	}
	return err
}

func main() {
	// 加载配置文件
	config, err := Config.LoadConfig("config.yaml")
	if err != nil {
		log.Fatalf("error loading config: %v", err)
	}

	err = Gost.RunGostCommand(config.Gost.Socks5User, config.Gost.Socks5Pass, config.Gost.Socks5Port, config.Gost.ApiURL)
	if err != nil {
		log.Fatalf("Failed to run gost command: %v", err)
	}

	// 立即执行一次 run
	if err := run(*config); err != nil {
		log.Printf("error running task: %v", err)
	}

	// 设置定时器
	tickerTime := time.Duration(config.Interval) * time.Minute
	ticker := time.NewTicker(tickerTime)
	defer ticker.Stop() // 确保在 main 函数退出前停止定时器

	// 主循环
	for {
		<-ticker.C // 等待定时器触发
		if err := run(*config); err != nil {
			log.Printf("error running task: %v", err)
		}

	}
}
