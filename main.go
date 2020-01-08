package main

import (
	"fmt"
	"io"
	"log"
	"os"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/renjiniravath/pokemon-unravelled/config"
	"github.com/renjiniravath/pokemon-unravelled/core/logger"
	"github.com/renjiniravath/pokemon-unravelled/routes"
	"github.com/renjiniravath/pokemon-unravelled/services"
)

func main() {
	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization},
	}))
	routes.Set(e)
	config.Load()
	services.Load()

	filepath := config.Current.LogFile
	fmt.Println("filepath ", filepath)

	f, err := os.OpenFile(filepath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		log.Fatalln("File not opened", err)
	}
	writer := io.MultiWriter(os.Stdout, f)
	logger.LogData(writer)

	e.Logger.Fatal(e.Start(":8080"))
}
