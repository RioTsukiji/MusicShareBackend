package handler

import (
	"fmt"
	"github.com/RioTsukiji/MusicShareBackend/internal/config"
	"github.com/RioTsukiji/MusicShareBackend/usecase"
	"github.com/gin-gonic/gin"
	"net/http"
)

type musicRegisterRequest struct {
	Title  string `json:"title"`
	Artist string `json:"artist"`
	Link   string `json:"link"`
	UserID int    `json:"user_id"`
}

type MusicHandler interface {
	HandleMusicGet(c *gin.Context)
	HandleMusicRecord(c *gin.Context)
}

type musicHandler struct {
	musicUseCase usecase.MusicUseCase
}

func NewMusicHandler(mu usecase.MusicUseCase) MusicHandler {
	return &musicHandler{
		musicUseCase: mu,
	}
}

func (mh musicHandler) HandleMusicGet(c *gin.Context) {
	songs, err := mh.musicUseCase.GetAllMusic(config.Db)
	if err != nil {
		fmt.Println(fmt.Sprintf("error: %v", err))
		return
	}
	c.JSON(http.StatusOK, songs)
}

func (mh musicHandler) HandleMusicRecord(c *gin.Context) {
	//リクエストボディを取得
	var requestBody musicRegisterRequest
	if err := c.BindJSON(&requestBody); err != nil {
		return
	}
	if err := mh.musicUseCase.InsertMusic(config.Db, requestBody.Title, requestBody.Artist, requestBody.Link, requestBody.UserID); err != nil {
		return
	}
	c.JSON(http.StatusOK, "OK")
}
