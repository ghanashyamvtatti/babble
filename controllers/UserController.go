package controllers
//
//import (
//	userproto "ds-project-bak/Common/Proto/Users"
//	"ds-project/common/proto/dsl"
//	"ds-project/config"
//	"ds-project/common/proto"
//	"ds-project/dtos"
//	"ds-project/services"
//	"github.com/gin-gonic/gin"
//	"net/http"
//)
//
//func GetUsers(appConfig *config.ApplicationConfig) gin.HandlerFunc {
//	return func(ctx *gin.Context) {
//		/*userReq:= &userproto.REQ {Nothing:"nothing"}
//        if userResponse, userErr:= client.GetUsers(ctx, userReq); userErr == nil {
//
//            ctx.JSON(http.StatusOK, dtos.Response{
//			Status:  true,
//			Message: "Successfully fetched the data",
//			Data: gin.H{
//				"result": userResponse,
//			},
//			})
//
//        } else {
//            ctx.JSON(http.StatusInternalServerError, gin.H {
//                "error": userErr.Error(),
//            })
//        }*/
//
//		// ctx.JSON(http.StatusOK, dtos.Response{
//		// 	Status:  true,
//		// 	Message: "Successfully fetched the data",
//		// 	Data: gin.H{
//		// 		"result": services.GetAllUsers(appConfig),
//		// 	},
//		// })
//	}
//}
//
//func GetUserDetails(appConfig *config.ApplicationConfig) gin.HandlerFunc {
//	return func(context *gin.Context) {
//		username := context.Param("username")
//		user := services.GetUserByUsername(appConfig, username)
//		context.JSON(http.StatusOK, dtos.Response{
//			Status:  true,
//			Message: "Successfully fetched the data",
//			Data: gin.H{
//				"user": user,
//			},
//		})
//	}
//}
