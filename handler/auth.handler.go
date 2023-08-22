package handler

import (
	DTO "api_book/dto"
	"api_book/helpers"
	"api_book/middleware"
	"api_book/models"
	"api_book/repository"
	"fmt"
	"net/http"

	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type AuthHandler struct {
	AuthRepository repository.AuthRepository
}

func (auth *AuthHandler) Register(ctx *gin.Context) {
	request := DTO.Register{}
	defer ctx.Request.Body.Close()

	if err := ctx.Bind(&request); err != nil {
		ctx.JSON(http.StatusOK, helpers.ErrorResponse{
			StatusCode: -1,
			Message:    err.Error(),
		})
		return
	}

	if isEmpty := request.IsEmpty(); isEmpty {
		ctx.JSON(http.StatusOK, helpers.ErrorResponse{
			StatusCode: -1,
			Message:    "Please complete all information",
		})
		return
	}

	if _, err := govalidator.ValidateStruct(request); err != nil {
		ctx.JSON(http.StatusBadRequest, helpers.ErrorResponse{
			StatusCode: -1,
			Message:    err.Error(),
		})
		return
	}

	pwdCode, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	helpers.Fatal(err)

	request.Password = string(pwdCode[:])
	user := models.User{
		Name:     request.FullName,
		Email:    request.Email,
		Password: request.Password,
		Role:     "guest",
	}

	result, err := auth.AuthRepository.Register(ctx, user)
	helpers.Fatal(err)
	fmt.Println("result: ", result)

	token, err := middleware.GenToken(user)
	helpers.Fatal(err)

	user.AccessToken = token
	ctx.JSON(http.StatusOK, helpers.Response{
		StatusCode: 1,
		Message:    "",
		Data:       user,
	})
}

func (auth *AuthHandler) Login(ctx *gin.Context) {}

func (auth *AuthHandler) ForgotPassword(ctx *gin.Context) {}
