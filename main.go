package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/leetpy/daisy/conf"
	"github.com/leetpy/daisy/logger"
	"github.com/leetpy/daisy/router"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

var (
	cfg = pflag.StringP("config", "c", "", "apiserver config file path.")
)

func main() {
	pflag.Parse()

	// init config
	if err := conf.Init(*cfg); err != nil {
		panic(err)
	}

	// init logger
	logger.InitLogger()

	gin.SetMode(viper.GetString("runmode"))
	r := gin.New()
	router.Load(r)
	zap.L().Info(http.ListenAndServe(viper.GetString("addr"), r).Error())
}
