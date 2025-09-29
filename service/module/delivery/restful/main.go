package restful

import (
	"basic-personal-financial-tracking-api/service/module/domain"

	"github.com/gin-gonic/gin"
)

type newRestfulHandler struct {
	usecase domain.PersonalFinancialTrackingUseCase
}

func NewHandler(restful *gin.Engine, usecase domain.PersonalFinancialTrackingUseCase) {

	// restfulHandler := newHandler{usecase: usecase}

	// api := restful.Group(common.APIGroup)
	// {

	// }
}
