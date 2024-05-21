package handlers

import (
	"errors"
	c "inventory/controller"
	m "inventory/model"
	u "inventory/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// Login : to login into system using email & password
func Login(ctx *gin.Context) {
	var (
		reqData m.Login
		err     error
		token   string
	)

	if err := ctx.ShouldBindJSON(&reqData); err != nil {
		var ve validator.ValidationErrors
		if errors.As(err, &ve) {
			out := make([]m.ErrorMsg, len(ve))
			for i, fe := range ve {
				out[i] = m.ErrorMsg{Field: fe.Field(), Message: u.GetErrorMsg(fe)}
			}
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"errors": out})
		}

		return
	}

	if reqData.Email == "" {
		out := make([]m.ErrorMsg, 1)
		out[0] = m.ErrorMsg{Field: "email", Message: "Email is required"}
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"errors": out})
		return
	}

	token, err = c.Login(reqData.Email, reqData.Password)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "email or password is incorrect."})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}
