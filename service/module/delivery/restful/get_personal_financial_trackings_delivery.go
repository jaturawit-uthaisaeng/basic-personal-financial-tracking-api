package restful

import (
	"net/http"
"basic-personal-financial-tracking-api/service/module/delivery/restful/model"
	"github.com/gin-gonic/gin"
)

func (newRestfulHandler *newRestfulHandler) GetPersonalFinancialTrackingsDelivery(c *gin.Context) {
	c.JSON(http.StatusOK, model.GetPersonalFinancialTrackingsDeliveryResponse{
		Message: "Get Personal Financial Trackings Delivery",
	})
}
