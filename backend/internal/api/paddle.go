package api

import (
	"MockExamLab/internal/models/DTO"
	"MockExamLab/internal/models/consts"
	"MockExamLab/internal/service"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
)

type PaddleAPI struct {
	s service.PaddleService
}

func ProvidePaddleAPI(s service.PaddleService) PaddleAPI {
	return PaddleAPI{s: s}
}

func (api *PaddleAPI) SubscriptionWebhook(c *gin.Context) {
	var b DTO.BasePaddleResponse

	/*	if err := c.ShouldBind(&b); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}*/

	jsonData, errJson := ioutil.ReadAll(c.Request.Body)
	if errJson != nil {
		c.JSON(http.StatusBadRequest, errJson)
		return
	}
	if err := json.Unmarshal(jsonData, &b); err != nil {
		c.JSON(http.StatusBadRequest, errJson)
		return
	}

	switch consts.PaddleAlert(b.AlertName) {
	case consts.SubscriptionCreated:
		{
			var p DTO.SubscriptionCreate
			if err := json.Unmarshal(jsonData, &p); err != nil {
				c.JSON(http.StatusBadRequest, errJson)
				return
			}
			if err := api.s.SubscriptionCreated(&b, &p); err != nil {
				c.JSON(http.StatusBadRequest, err)
				return
			}
		}
	case consts.SubscriptionUpdated:
		{
			var p DTO.SubscriptionUpdate
			if err := json.Unmarshal(jsonData, &p); err != nil {
				c.JSON(http.StatusBadRequest, errJson)
				return
			}
			if err := api.s.SubscriptionUpdate(&b, &p); err != nil {
				c.JSON(http.StatusBadRequest, err)
				return
			}
		}
	case consts.SubscriptionCancelled:
		{
			var p DTO.SubscriptionCanceled
			if err := json.Unmarshal(jsonData, &p); err != nil {
				c.JSON(http.StatusBadRequest, errJson)
				return
			}
			if err := api.s.SubscriptionCancelled(&b, &p); err != nil {
				c.JSON(http.StatusBadRequest, err)
				return
			}
		}
	case consts.SubscriptionPaymentRefunded:
		{
			var p DTO.SubscriptionPaymentRefunded
			if err := json.Unmarshal(jsonData, &p); err != nil {
				c.JSON(http.StatusBadRequest, errJson)
				return
			}
			if err := api.s.SubscriptionRefunded(&b, &p); err != nil {
				c.JSON(http.StatusBadRequest, err)
				return
			}
		}
	case consts.SubscriptionPaymentSucceeded:
		{
			var p DTO.SubscriptionPaymentSucceeded
			if err := json.Unmarshal(jsonData, &p); err != nil {
				c.JSON(http.StatusBadRequest, errJson)
				return
			}
			if err := api.s.SubscriptionPaymentSucceeded(&b, &p); err != nil {
				c.JSON(http.StatusBadRequest, err)
				return
			}
		}
	case consts.SubscriptionPaymentFailed:
		{
			var p DTO.SubscriptionPaymentFailed
			if err := json.Unmarshal(jsonData, &p); err != nil {
				c.JSON(http.StatusBadRequest, errJson)
				return
			}
			if err := api.s.SubscriptionPaymentFailed(&b, &p); err != nil {
				c.JSON(http.StatusBadRequest, err)
				return
			}
		}
	}
	c.Status(http.StatusOK)

}
