package restful

import (
	"basic-personal-financial-tracking-api/service/common"
	"basic-personal-financial-tracking-api/service/module/domain"

	"github.com/gin-gonic/gin"
)

type newRestfulHandler struct {
	usecase domain.PersonalFinancialTrackingUseCase
}

func NewHandler(restful *gin.Engine, usecase domain.PersonalFinancialTrackingUseCase) {

	restfulHandler := newRestfulHandler{usecase: usecase}

	api := restful.Group(common.APIGroup)
	{
		api.GET("/personal-financial-trackings", restfulHandler.GetPersonalFinancialTrackingsDelivery)
	}
}
