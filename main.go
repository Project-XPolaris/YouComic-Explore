package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"net/http"
	"web-desktop/setup"
)

func main() {
	viper.SetConfigName("config")
	viper.SetConfigType("json")
	viper.AddConfigPath("./config")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	firstStart := viper.GetBool("first")
	portString := fmt.Sprintf(":%s",viper.GetString("port"))
	if firstStart {
		router := gin.Default()
		router.Static("/assets", "./setup/assets")
		router.LoadHTMLGlob("setup/templates/*")
		//router.LoadHTMLFiles("templates/template1.html", "templates/template2.html")
		router.GET("/", setup.HomeHandler)
		router.POST("/submit",setup.SubmitConfigHandler)
		router.GET("/error",setup.ErrorHandler)
		router.GET("/success",setup.SuccessHandler)
		fmt.Println(fmt.Sprintf("\n\n\n\n服务启动成功，请在地址栏输入 http://localhost%s 访问服务\n\n\n\n",portString))
		router.Run(portString)
	}

	http.Handle("/", http.FileServer(http.Dir("./static")))
	fmt.Println(fmt.Sprintf("\n\n\n\n服务启动成功，请在地址栏输入 http://localhost%s 访问服务\n\n\n\n",portString))
	http.ListenAndServe(portString, nil)
}
