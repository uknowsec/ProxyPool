package ProxyApi

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type ProxyServer struct {
	Server   string `json:"sever"`
	Port     int    `json:"port"`
	User     string `json:"user"`
	Password string `json:"pw"`
}

type Response struct {
	Count  string        `json:"count"`
	Status string        `json:"status"`
	Expire string        `json:"expire"`
	List   []ProxyServer `json:"list"`
}

func GetProxyServer(ProxyApi string) (ProxyServer, error) {
	var serverInfo ProxyServer

	// 发起HTTP GET请求
	resp, err := http.Get(ProxyApi)
	if err != nil {
		return serverInfo, err
	}
	defer resp.Body.Close()

	// 读取响应内容
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return serverInfo, err
	}

	// 解析JSON响应
	var response Response
	err = json.Unmarshal(body, &response)
	if err != nil {
		return serverInfo, err
	}

	// 提取服务器信息
	if len(response.List) > 0 {
		serverInfo = response.List[0]
	} else {
		log.Println("No server information found in the response.")
	}

	return serverInfo, nil
}
