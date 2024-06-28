package Gost

import (
	"ProxyPool/ProxyApi"
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"
)

type Connector struct {
	Type string `json:"type"`
	Auth struct {
		Username string `json:"username"`
		Password string `json:"password"`
	} `json:"auth"`
}

type Dialer struct {
	Type string `json:"type"`
	TLS  struct {
		ServerName string `json:"serverName"`
	} `json:"tls"`
}

type Node struct {
	Name      string    `json:"name"`
	Addr      string    `json:"addr"`
	Connector Connector `json:"connector"`
	Dialer    Dialer    `json:"dialer"`
}

type Hop struct {
	Name  string `json:"name"`
	Nodes []Node `json:"nodes"`
}

type Chain struct {
	Name string `json:"name"`
	Hops []Hop  `json:"hops"`
}

func UpdateJson(server ProxyApi.ProxyServer) []byte {
	log.Printf("Updating ProxyServer information:")
	log.Printf("Server: %s", server.Server)
	log.Printf("Port: %d", server.Port)
	log.Printf("User: %s", server.User)
	log.Printf("Password: %s", server.Password)
	chain := Chain{
		Name: "chain-0",
		Hops: []Hop{
			{
				Name: "hop-0",
				Nodes: []Node{
					{
						Name: "node-0",
						Addr: fmt.Sprintf("%s:%d", server.Server, server.Port),
						Connector: Connector{
							Type: "socks5",
							Auth: struct {
								Username string `json:"username"`
								Password string `json:"password"`
							}{
								Username: server.User,
								Password: server.Password,
							},
						},
						Dialer: Dialer{
							Type: "tcp",
							TLS: struct {
								ServerName string `json:"serverName"`
							}{
								ServerName: server.Server,
							},
						},
					},
				},
			},
		},
	}

	// Convert struct to JSON
	jsonData, err := json.Marshal(chain)
	if err != nil {
		log.Fatalf("Failed to marshal JSON: %v", err)
	}

	// Print JSON
	log.Printf(string(jsonData))

	return jsonData

}

func UpdateGostHop(url string, jsonData []byte) error {
	url = fmt.Sprintf("http://%s/config/chains/chain-0", url)
	req, err := http.NewRequest("PUT", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return fmt.Errorf("error creating PUT request: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("error sending PUT request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected response status: %s", resp.Status)
	}

	log.Printf("PUT request successful")
	return nil
}

func RunGostCommand(user, pass string, localPort int, apiUrl string) error {
	// 打开文件准备保存输出
	logFile, err := os.OpenFile("gost.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer logFile.Close()

	// 构建命令参数
	cmdArgs := []string{
		"-L", fmt.Sprintf("socks5://%s:%s@:%d", user, pass, localPort),
		"-F", "socks5://user:pass@180.118.241.150:34253",
		"-api", fmt.Sprintf("%s", apiUrl),
	}

	// 构建并执行命令
	cmd := exec.Command("gost", cmdArgs...)

	// 将命令的标准输出和标准错误连接到文件
	cmd.Stdout = logFile
	cmd.Stderr = logFile

	// 启动命令
	err = cmd.Start()
	if err != nil {
		return err
	}

	// 使用 goroutine 等待命令完成
	go func() {
		err := cmd.Wait()
		if err != nil {
			log.Printf("Command finished with error: %v", err)
		}
	}()

	return nil
}
