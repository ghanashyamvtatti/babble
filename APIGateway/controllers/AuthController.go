package controllers

import (
	"ds-project/APIGateway/common"
	"ds-project/APIGateway/dtos"
	"ds-project/common/proto/auth"
	"ds-project/common/proto/models"
	"ds-project/common/proto/users"
	"ds-project/common/utilities"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SignIn(clients *common.ServiceClients) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var login dtos.UserLogin

		if err := ctx.ShouldBind(&login); err != nil {
			ctx.JSON(http.StatusBadRequest, dtos.Response{
				Status:  false,
				Message: err.Error(),
			})
			return
		}

		response, err := clients.AuthClient.Login(ctx, &auth.LoginRequest{
			Username: login.Username,
			Password: login.Password,
		})

		if err == nil && response.Ok {
			// Get user
			userResponse, err := clients.UserClient.GetUser(ctx, &users.GetUserRequest{Username: login.Username})
			if err != nil {
				ctx.JSON(http.StatusInternalServerError, dtos.Response{
					Status:  false,
					Message: "Something went wrong while trying to fetch the user",
					Data:    nil,
				})
				return
			}
			tokenResponse, err := clients.AuthClient.GenerateAccessToken(ctx, &auth.GenerateTokenRequest{Username: login.Username})
			if err != nil {
				ctx.JSON(http.StatusInternalServerError, dtos.Response{
					Status:  false,
					Message: "Something went wrong while trying to generate the access token",
					Data:    nil,
				})
				return
			}
			ctx.JSON(http.StatusOK, dtos.Response{
				Status:  true,
				Message: "Successfully logged in",
				Data: gin.H{
					"user": dtos.User{
						Username: userResponse.Username,
						FullName: userResponse.User.FullName,
					},
					"token": tokenResponse.Token,
				},
			})
		} else {
			ctx.JSON(http.StatusInternalServerError, dtos.Response{
				Status:  false,
				Message: "Authentication failed",
				Data:    nil,
			})
		}
	}
}

func SignUp(clients *common.ServiceClients) gin.HandlerFunc {
	return func(ctx *gin.Context) {

		var registerUser dtos.UserRegistration
		if err := ctx.ShouldBind(&registerUser); err != nil {
			ctx.JSON(http.StatusBadRequest, dtos.Response{
				Status:  false,
				Message: err.Error(),
			})
			return
		}

		resp, err := clients.UserClient.CheckUserNameExists(ctx, &users.GetUserRequest{Username: registerUser.Username})
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, dtos.Response{
				Status:  false,
				Message: "Something went wrong while trying to check if the user already exists",
				Data:    nil,
			})
			return
		}

		if !resp.Ok {
			password, err := utilities.HashPassword(registerUser.Password)
			if err != nil {
				ctx.JSON(http.StatusInternalServerError, dtos.Response{
					Status:  false,
					Message: "Unable to create user",
					Data:    nil,
				})
				return
			}
			value := models.User{
				FullName: registerUser.FullName,
				Password: password,
			}
			_, er := clients.UserClient.CreateUser(ctx, &users.CreateUserRequest{Username: registerUser.Username, User: &value})
			if er != nil {
				ctx.JSON(http.StatusInternalServerError, dtos.Response{
					Status:  false,
					Message: "Unable to create user",
					Data:    nil,
				})
				return
			}
		}

		// Get user
		userResponse, err := clients.UserClient.GetUser(ctx, &users.GetUserRequest{Username: registerUser.Username})
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, dtos.Response{
				Status:  false,
				Message: "Something went wrong while trying to fetch the user",
				Data:    nil,
			})
			return
		}

		tokenResponse, err := clients.AuthClient.GenerateAccessToken(ctx, &auth.GenerateTokenRequest{Username: registerUser.Username})
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, dtos.Response{
				Status:  false,
				Message: "Something went wrong while trying to generate the access token",
				Data:    nil,
			})
			return
		}
		ctx.JSON(http.StatusOK, dtos.Response{
			Status:  true,
			Message: "Successfully registered user",
			Data: gin.H{
				"user": dtos.User{
					Username: userResponse.Username,
					FullName: userResponse.User.FullName,
				},
				"token": tokenResponse.Token,
			},
		})
	}
}

func SignOut(clients *common.ServiceClients) gin.HandlerFunc {
	return func(context *gin.Context) {
		_, err := clients.AuthClient.Logout(context, &auth.LogoutRequest{Username: context.Param("username")})
		if err != nil {
			context.JSON(http.StatusInternalServerError, dtos.Response{
				Status:  false,
				Message: "Unable to sign the user out",
				Data:    nil,
			})
			return
		}
		context.JSON(200, dtos.Response{
			Status:  true,
			Message: "Successfully signed out the user",
			Data:    nil,
		})
	}
}

func Authenticate(clients *common.ServiceClients) gin.HandlerFunc {
	return func(context *gin.Context) {
		username := context.Param("username")
		resp, err := clients.UserClient.CheckUserNameExists(context, &users.GetUserRequest{Username: username})

		if err != nil || !(resp.Ok) {
			context.JSON(500, dtos.Response{
				Status:  false,
				Message: "Invalid username",
				Data:    nil,
			})
			context.Abort()
		}
		validityResp, err := clients.AuthClient.CheckAccessTokenValid(context, &auth.TokenValidityRequest{Username: username, Token: context.GetHeader("token")})
		if err != nil || !validityResp.Ok {
			context.JSON(http.StatusUnauthorized, dtos.Response{
				Status:  false,
				Message: "Missing or invalid access token",
				Data:    nil,
			})
			context.Abort()
		}
	}
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With, token, redirect")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
