package handler

import (
	"fmt"
	"github.com/RioTsukiji/MusicShareBackend/internal/config"
	"github.com/RioTsukiji/MusicShareBackend/usecase"
	"github.com/gin-gonic/gin"
	"net/http"
)

type userSignupRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserHandler interface {
	HandleUserGet(c *gin.Context)
	HandleUserSignup(c *gin.Context)
}

type userHandler struct {
	userUseCase usecase.UserUseCase
}

func NewUserHandler(uu usecase.UserUseCase) UserHandler {
	return &userHandler{
		userUseCase: uu,
	}
}

func (uh userHandler) HandleUserGet(c *gin.Context) {
	username := c.Query("username")
	user, err := uh.userUseCase.GetByUserName(config.Db, username)
	if err != nil {
		fmt.Println(fmt.Sprintf("error: %v", err))
		return
	}
	c.JSON(http.StatusOK, user)
}

func (uh userHandler) HandleUserSignup(c *gin.Context) {
	var requestBody userSignupRequest
	if err := c.BindJSON(&requestBody); err != nil {
		return
	}
	if err := uh.userUseCase.InsertUser(config.Db, requestBody.Username, requestBody.Password); err != nil {
		return
	}

	c.JSON(http.StatusOK, "OK")
}
