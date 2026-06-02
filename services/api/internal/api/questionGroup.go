package api

import (
	"MockExamLab/internal/mapper"
	"MockExamLab/internal/models/DTO"
	"MockExamLab/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
)

type QuestionGroupAPI struct {
	s service.QuestionGroupService
}

func ProvideQuestionGroupAPI(s service.QuestionGroupService) QuestionGroupAPI {
	return QuestionGroupAPI{s: s}
}

// Create ... Create Question-Group Group
// @Summary Create Question-Group Group
// @Description Create Question Group
// @Tags Question-Groups
// @Accept json
// @Param user body DTO.QuestionGroupCreateRequest true "Create Question Group"
// @Success 200 {object} DTO.QuestionGroupCreateResponse
// @Failure 400 "error message detail"
// @Router /question-group [post]
// @Security ApiKeyAuth
func (api *QuestionGroupAPI) Create(c *gin.Context) {
	var r DTO.QuestionGroupCreateRequest
	u := getUserClaimsFromSession(c)
	if u == nil {
		c.JSON(http.StatusBadRequest, "User not Found")
		return
	}
	if err := c.ShouldBind(&r); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	data, reqConvertErr := mapper.QuestionGroupCreateRequestToModel(&r, u.UserId)
	if reqConvertErr != nil {
		c.JSON(http.StatusBadRequest, reqConvertErr)
		return
	}
	if res, err := api.s.Create(*data); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	} else {
		convertedData, resConvertErr := mapper.QuestionGroupModelToCreateResponse(res)
		if resConvertErr != nil {
			c.JSON(http.StatusBadRequest, resConvertErr)
			return
		}
		c.JSON(http.StatusOK, &convertedData)
	}
}

// DeleteById ... Question-Group by id
// @Summary Delete Question-Group by id
// @Description Delete Question-Group by id
// @Tags Question-Groups
// @Accept json
// @Param id path string true "Question-Group ID"
// @Success 200 "Success"
// @Failure 400 "error message detail"
// @Router /question-group/{id} [delete]
// @Security ApiKeyAuth
func (api *QuestionGroupAPI) DeleteById(c *gin.Context) {
	id, ParsError := uuid.Parse(c.Param("id"))
	if ParsError != nil {
		c.JSON(http.StatusBadRequest, "Invalid Id")
		return
	}
	if err := api.s.DeleteById(id); err != nil {
		c.JSON(http.StatusBadRequest, err)
	}
	c.AbortWithStatus(http.StatusOK)
}

// Update ... Update Question-Group
// @Summary Update Question-Group
// @Description Update Question-Group
// @Tags Question-Groups
// @Accept json
// @Param user body DTO.QuestionGroupUpdateRequest true "Update Question-Group"
// @Success 200 {object} DTO.QuestionGroupUpdateResponse
// @Failure 400 "error message detail"
// @Router /question-group [put]
// @Security ApiKeyAuth
func (api *QuestionGroupAPI) Update(c *gin.Context) {
	var r DTO.QuestionGroupUpdateRequest
	u := getUserClaimsFromSession(c)
	if u == nil {
		c.JSON(http.StatusBadRequest, "User not Found")
		return
	}
	if err := c.ShouldBind(&r); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	data, reqConvertErr := mapper.QuestionGroupUpdateRequestToModel(&r, u.UserId)
	if reqConvertErr != nil {
		c.JSON(http.StatusBadRequest, reqConvertErr)
		return
	}
	if res, err := api.s.Update(*data); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	} else {
		convertedData, resConvertErr := mapper.QuestionGroupModelToUpdateResponse(res)
		if resConvertErr != nil {
			c.JSON(http.StatusBadRequest, resConvertErr)
			return
		}
		c.JSON(http.StatusOK, convertedData)
	}
}

// FindById ... Find Question-Group by id
// @Summary Find Question-Group by id
// @Description Find Question-Group by id
// @Tags Question-Groups
// @Accept json
// @Param id path string true "Question-Group ID"
// @Success 200 {object} DTO.QuestionGroupResponse
// @Failure 400 "error message detail"
// @Router /question-group/{id} [get]
// @Security ApiKeyAuth
func (api *QuestionGroupAPI) FindById(c *gin.Context) {
	id, ParsError := uuid.Parse(c.Param("id"))
	if ParsError != nil {
		c.JSON(http.StatusBadRequest, "Invalid Id")
		return
	}
	if res, err := api.s.FindById(id); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	} else {
		convertedData, resConvertErr := mapper.QuestionGroupModelToResponse(res)
		if resConvertErr != nil {
			c.JSON(http.StatusBadRequest, resConvertErr)
			return
		}
		c.JSON(http.StatusOK, convertedData)
	}
}

// FindAllBySectionId ... Find Group by Section id
// @Summary Find Group by Section id
// @Description Find Group by Section id
// @Tags Question-Groups
// @Accept json
// @Param section_id query string true "section ID"
// @Success 200 {array} DTO.QuestionGroupResponse
// @Failure 400 "error message detail"
// @Router /question-group/all [get]
// @Security ApiKeyAuth
func (api *QuestionGroupAPI) FindAllBySectionId(c *gin.Context) {
	SectionId, ParsError := uuid.Parse(c.Query("section_id"))
	if ParsError != nil {
		c.JSON(http.StatusBadRequest, "Invalid Id")
		return
	}
	if res, err := api.s.FindAllBySectionId(SectionId); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	} else {
		convertedData, resConvertErr := mapper.QuestionGroupModelsToArrayResponse(res)
		if resConvertErr != nil {
			c.JSON(http.StatusBadRequest, resConvertErr)
			return
		}
		c.JSON(http.StatusOK, convertedData)
	}
}
