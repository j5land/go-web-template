package requests

import (
	"go-web-template/pkg/logger"
	"time"
)

var (
	httpClient *Session
)

func init() {
	//init httpClient
	httpClient = NewSession()
	httpClient.Client.SetTimeout(5 * time.Second)
	logger.Info("init httpClient success.")
}