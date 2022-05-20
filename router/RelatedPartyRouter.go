package router

import (
	"github.com/gin-gonic/gin"
	"learn-go/controller"
)

func LoadRelatedPartyRouter(engine *gin.Engine) {
	RelatedPartyGroup := engine.Group("/related_party")
	{
		RelatedPartyGroup.GET("/:id", controller.FindRelatedParty)
	}

}
