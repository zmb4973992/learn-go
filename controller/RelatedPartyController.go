package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"learn-go/model"
	"learn-go/util"

	"net/http"
)

func FindProject(context *gin.Context) {
	var a model.RelatedParty
	util.DB.First(&a)
	fmt.Println(a)
	context.JSON(http.StatusOK, gin.H{"data": a})
}
