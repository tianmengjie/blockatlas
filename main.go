package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"trustwallet.com/blockatlas/platform"
	"os"
)

func main() {
	// Load config
	viper.AddConfigPath(".")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	if err := viper.ReadInConfig(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	// Start server
	gin.SetMode(viper.GetString("gin.mode"))
	router := gin.Default()
	platform.Register(router)
	router.Run(":8080")
}