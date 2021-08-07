package controller

import (
	"bwastartup-crowdfunding/model"
	"bwastartup-crowdfunding/service"
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type campaignController struct {
	campaignService service.CampaignService
}

func NewCampaignController(campaignService service.CampaignService) *campaignController {
	return &campaignController{campaignService: campaignService}
}

// GetCampaigns godoc
// @Summary Get campaign data
// @Description Can use query parameter user_id, backer_id, or none
// @ID get-campaigns
// @Produce  json
// @Param user_id query integer false "UserID"
// @Param backer_id query integer false "BackerID"
// @Success 200 {object} model.WebResponse{data=model.GetCampaignResponse}
// @Failure 400 {object} model.WebResponse{data=string}
// @Router /campaigns [get]
func (campaignController *campaignController) GetCampaigns(c *gin.Context) {
	userId, err := strconv.Atoi(c.DefaultQuery("user_id", "0"))

	if err != nil {
		c.JSON(http.StatusBadRequest, model.WebResponse{
			Code:   http.StatusBadRequest,
			Status: http.StatusText(http.StatusBadRequest),
			Data:   err.Error(),
		})
		return
	}

	if userId < 0 {
		c.JSON(http.StatusBadRequest, model.WebResponse{
			Code:   http.StatusBadRequest,
			Status: http.StatusText(http.StatusBadRequest),
			Data:   "Wrong ID",
		})
		return
	}

	ctx := context.Background()
	var campaignResponses []model.GetCampaignResponse

	if userId > 0 {
		campaignResponses, err = campaignController.campaignService.FindByUserId(ctx, uint32(userId))
		if err != nil {
			c.JSON(http.StatusBadRequest, model.WebResponse{
				Code:   http.StatusBadRequest,
				Status: http.StatusText(http.StatusBadRequest),
				Data:   err.Error(),
			})
			return
		}
	} else {
		campaignResponses, err = campaignController.campaignService.FindAll(ctx)
		if err != nil {
			c.JSON(http.StatusBadRequest, model.WebResponse{
				Code:   http.StatusBadRequest,
				Status: http.StatusText(http.StatusBadRequest),
				Data:   err.Error(),
			})
			return
		}
	}

	c.JSON(http.StatusOK, model.WebResponse{
		Code:   http.StatusOK,
		Status: http.StatusText(http.StatusOK),
		Data:   campaignResponses,
	})
}

// GetCampaignDetails godoc
// @Summary Get campaign details with campaign id
// @Description must send campaign id in URI
// @ID get-campaign-details
// @Produce json
// @Param id path integer true "CampaignID"
// @Success 200 {object} model.WebResponse{data=model.GetCampaignDetailResponse}
// @Failure 400 {object} model.WebResponse{data=string}
// @Router /campaigns/:id [get]
func (campaignController *campaignController) GetCampaignDetails(c *gin.Context) {
	var campaignRequest model.GetCampaignDetailRequest
	err := c.ShouldBindUri(&campaignRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.WebResponse{
			Code:   http.StatusBadRequest,
			Status: http.StatusText(http.StatusBadRequest),
			Data:   err.Error(),
		})
		return
	}

	ctx := context.Background()
	response, err := campaignController.campaignService.FindById(ctx, campaignRequest.Id)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.WebResponse{
			Code:   http.StatusBadRequest,
			Status: http.StatusText(http.StatusBadRequest),
			Data:   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, model.WebResponse{
		Code:   http.StatusOK,
		Status: http.StatusText(http.StatusOK),
		Data:   response,
	})
}
