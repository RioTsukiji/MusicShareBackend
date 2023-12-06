package main

import (
	"github.com/RioTsukiji/MusicShareBackend/infrastructure/persistence"
	"github.com/RioTsukiji/MusicShareBackend/interfaces/handler"
	"github.com/RioTsukiji/MusicShareBackend/usecase"
	"github.com/gin-gonic/gin"
	"os"
)

func main() {
	userPersistence := persistence.NewUserPersistence()
	userUseCase := usecase.NewUserUseCase(userPersistence)
	userHandler := handler.NewUserHandler(userUseCase)

	musicPersistence := persistence.NewMusicPersistence()
	musicUseCase := usecase.NewMusicUseCase(musicPersistence)
	musicHandler := handler.NewMusicHandler(musicUseCase)

	engine := gin.Default()
	// Todo: ログイン機能を実装する
	engine.POST("/login", userHandler.HandleUserLogin)
	engine.POST("/signup", userHandler.HandleUserSignup)

	engine.GET("/", musicHandler.HandleMusicGet)
	engine.POST("/share", musicHandler.HandleMusicRecord)

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000" // デフォルトポート
	}

	engine.Run(":" + port)
}
