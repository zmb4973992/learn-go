package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"learn-go/model"
	"learn-go/util"
	"net/http"
)

func FindRelatedParty(context *gin.Context) {
	var relatedParty model.RelatedParty
	util.DB.First(&relatedParty)
	fmt.Println(relatedParty)
	context.JSON(http.StatusOK, gin.H{"data": relatedParty})
}
