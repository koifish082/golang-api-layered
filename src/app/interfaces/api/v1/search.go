package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type SearchResources struct {
	RouterGroup *gin.RouterGroup
}

func NewSearchResources(rg *gin.RouterGroup) *SearchResources {
	return &SearchResources{
		RouterGroup: rg,
	}
}

func (rs SearchResources) Routes() {
	rs.RouterGroup.GET("/code", rs.getRepository)
}

func (rs SearchResources) getRepository(ctx *gin.Context) {
	ctx.String(http.StatusOK, "Hi")
}
