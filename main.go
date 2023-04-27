package main

import (
	"github.com/GIT_USER_ID/GIT_REPO_ID/handlers"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

    //todo: handle the error!
	c, _ := handlers.NewContainer()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())


	// GetNftId - Get NFT info
	e.GET("/api/v1/nft/:class_id/:nft_id", c.GetNftId)

	// GetNfts - Get Listed NFTs
	e.GET("/api/v1/nft", c.GetNfts)


	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}