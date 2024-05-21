package handlers

import (
	"errors"
	c "inventory/controller"
	m "inventory/model"
	u "inventory/utils"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// GetItems : to get items data based on parameters
func GetItems(ctx *gin.Context) {
	var (
		sku, name, description, status, orderBy string
		retData                                 []*m.Item
		err                                     error
	)

	sku = ctx.Query("sku")
	name = ctx.Query("name")
	description = ctx.Query("description")
	status = ctx.Query("status")
	orderBy = ctx.Query("orderby")

	retData, err = c.GetItems(sku, name, description, status, orderBy)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"errors": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": retData,
	})
}

// AddItem : to add new item
func AddItem(ctx *gin.Context) {
	var (
		reqData *m.Item
		err     error
	)

	err = ctx.ShouldBindJSON(&reqData)
	if err != nil {
		var ve validator.ValidationErrors
		if errors.As(err, &ve) {
			out := make([]m.ErrorMsg, len(ve))
			for i, fe := range ve {
				out[i] = m.ErrorMsg{Field: strings.ToLower(fe.Field()), Message: u.GetErrorMsg(fe)}
			}
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"errors": out})
		}

		return
	}

	err = c.AddItem(reqData)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"errors": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "Success to add item",
	})
}

// ModifyItem : to change data of existing item
func ModifyItem(ctx *gin.Context) {
	var (
		reqData *m.Item
		err     error
		id      string
	)

	err = ctx.ShouldBindJSON(&reqData)
	if err != nil {
		var ve validator.ValidationErrors
		if errors.As(err, &ve) {
			out := make([]m.ErrorMsg, len(ve))
			for i, fe := range ve {
				out[i] = m.ErrorMsg{Field: strings.ToLower(fe.Field()), Message: u.GetErrorMsg(fe)}
			}
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"errors": out})
		}

		return
	}

	id = ctx.Param("id")

	err = c.ModifyItem(id, reqData)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"errors": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "Success to update item",
	})
}

// RemoveItem : to delete existing item. Item deleted will be set it's status flag into 2 and had value in deleted_at field
func RemoveItem(ctx *gin.Context) {
	var (
		id  string
		err error
	)

	id = ctx.Param("id")

	err = c.RemoveItem(id)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"errors": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "Success to remove item",
	})
}
