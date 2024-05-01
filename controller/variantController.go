package controller

import (
	"finaProjectGolang/database"
	"finaProjectGolang/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllVariant(ctx *gin.Context) {
	db := database.GetDB()

	results := []model.Variant{}

	err := db.Debug().Find(&results).Error
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad request",
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": results,
	})
}

func GetVariantByID(ctx *gin.Context) {
	db := database.GetDB()
	id := ctx.Param("variantID")

	var variant model.Variant

	err := db.Debug().Where("id = ?", id).First(&variant).Error
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad request",
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": variant,
	})
}

func CreateVariant(ctx *gin.Context) {

}
