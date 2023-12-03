package main

import (
	"github.com/RioTsukiji/MusicShareBackend/infrastructure/persistence"
	"github.com/RioTsukiji/MusicShareBackend/interfaces/handler"
	"github.com/RioTsukiji/MusicShareBackend/usecase"
	"github.com/gin-gonic/gin"
)

func main() {
	userPersistence := persistence.NewUserPersistence()
	userUseCase := usecase.NewUserUseCase(userPersistence)
	userHandler := handler.NewUserHandler(userUseCase)

	musicPersistence := persistence.NewMusicPersistence()
	musicUseCase := usecase.NewMusicUseCase(musicPersistence)
	musicHandler := handler.NewMusicHandler(musicUseCase)

	engine := gin.Default()
	engine.GET("/login", userHandler.HandleUserGet)
	engine.POST("/signup", userHandler.HandleUserSignup)

	engine.GET("/", musicHandler.HandleMusicGet)
	engine.POST("/share", musicHandler.HandleMusicRecord)

	engine.Run(":3000")
}
