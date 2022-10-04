package api

import (
	"MockExamLab/internal/mapper"
	"MockExamLab/internal/models/DTO"
	"MockExamLab/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
)

type SectionAPI struct {
	s service.SectionService
}

func ProvideSectionAPI(s service.SectionService) SectionAPI {
	return SectionAPI{s: s}
}

// Create ... Create Section
// @Summary Create Section
// @Description Create Section
// @Tags Sections
// @Accept json
// @Param Section body DTO.SectionCreateRequest true "Create section"
// @Success 200 {object} DTO.SectionCreateResponse
// @Failure 400 "error message detail"
// @Router /section [post]
// @Security ApiKeyAuth
func (api *SectionAPI) Create(c *gin.Context) {
	var r DTO.SectionCreateRequest
	u := getUserClaimsFromSession(c)
	if u == nil {
		c.JSON(http.StatusBadRequest, "User not Found")
		return
	}
	if err := c.ShouldBind(&r); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	data, reqConvertErr := mapper.SectionCreateRequestToModel(&r, u.UserId)
	if reqConvertErr != nil {
		c.JSON(http.StatusBadRequest, reqConvertErr)
		return
	}
	if res, err := api.s.Create(*data); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	} else {
		convertedData, resConvertErr := mapper.SectionModelToCreateResponse(res)
		if resConvertErr != nil {
			c.JSON(http.StatusBadRequest, resConvertErr)
			return
		}
		c.JSON(http.StatusOK, &convertedData)
	}
}

// DeleteById ... Delete Section by id
// @Summary Delete Section by id
// @Description Delete Section by id
// @Tags Sections
// @Accept json
// @Param id path string true "Section ID"
// @Success 200 "Success"
// @Failure 400 "error message detail"
// @Router /section/{id} [delete]
// @Security ApiKeyAuth
func (api *SectionAPI) DeleteById(c *gin.Context) {
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

// Update ... Update Section
// @Summary Update Section
// @Description Update Section
// @Tags Sections
// @Accept json
// @Param user body DTO.SectionUpdateRequest true "Update test"
// @Success 200 {object} DTO.SectionUpdateResponse
// @Failure 400 "error message detail"
// @Router /section [put]
// @Security ApiKeyAuth
func (api *SectionAPI) Update(c *gin.Context) {
	var r DTO.SectionUpdateRequest
	u := getUserClaimsFromSession(c)
	if u == nil {
		c.JSON(http.StatusBadRequest, "User not Found")
		return
	}
	if err := c.ShouldBind(&r); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	data, reqConvertErr := mapper.SectionUpdateRequestToModel(&r, u.UserId)
	if reqConvertErr != nil {
		c.JSON(http.StatusBadRequest, reqConvertErr)
		return
	}
	if res, err := api.s.Update(*data); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	} else {
		convertedData, resConvertErr := mapper.SectionModelToUpdateResponse(res)
		if resConvertErr != nil {
			c.JSON(http.StatusBadRequest, resConvertErr)
			return
		}
		c.JSON(http.StatusOK, convertedData)
	}
}

// FindById ... Find Section by id
// @Summary Find Section by id
// @Description Find Section by id
// @Tags Sections
// @Accept json
// @Param id path string true "Section ID"
// @Success 200 {object} DTO.SectionResponse
// @Failure 400 "error message detail"
// @Router /section/{id} [get]
// @Security ApiKeyAuth
func (api *SectionAPI) FindById(c *gin.Context) {
	id, ParsError := uuid.Parse(c.Param("id"))
	if ParsError != nil {
		c.JSON(http.StatusBadRequest, "Invalid Id")
		return
	}
	if res, err := api.s.FindById(id); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	} else {
		convertedData, resConvertErr := mapper.SectionModelToResponse(res)
		if resConvertErr != nil {
			c.JSON(http.StatusBadRequest, resConvertErr)
			return
		}
		c.JSON(http.StatusOK, convertedData)
	}
}

// FindAll ... Find All
// @Summary Find All
// @Description Find all Sections
// @Tags Sections
// @Accept json
// @Param test_id query string true "section ID"
// @Success 200 {array} DTO.SectionResponse
// @Failure 400 "error message detail"
// @Router /section/all [get]
// @Security ApiKeyAuth
func (api *SectionAPI) FindAll(c *gin.Context) {
	testId, ParsError := uuid.Parse(c.Query("test_id"))
	if ParsError != nil {
		c.JSON(http.StatusBadRequest, "Invalid Id")
		return
	}
	if res, err := api.s.FindAllByTestId(testId); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	} else {
		convertedData, resConvertErr := mapper.SectionModelsToArrayResponse(res)
		if resConvertErr != nil {
			c.JSON(http.StatusBadRequest, resConvertErr)
			return
		}
		c.JSON(http.StatusOK, convertedData)
	}
}
