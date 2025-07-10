package middlewares

import (
	"Event_Management/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Authenticate(context *gin.Context) {
	token := context.Request.Header.Get("Authorization")

	// verify the token is exist or not
	if token == "" {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "unauthorized"})
		return
	}

	// verify the jwt

	err := utils.VerifyToken(token)

	if err != nil {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "unauthorized"})
		return
	}

	// using the middleware we can set some context values ....
	//context.Set("userID", userId)

	// get the value from context
	//context.Get('userId')
	//context.GetInt64('userId')

	// jwt verified
	context.Next()

}

func GetUserIdFromToken(context *gin.Context) {
	token := context.Request.Header.Get("Authorization")

	userId, err := utils.GetUserIdFromToken(token)

	if err != nil {
		context.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	context.Set("user_id", userId)
	context.Next()
}
