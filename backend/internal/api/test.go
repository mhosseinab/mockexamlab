package api

import (
	"MockExamLab/internal/mapper"
	"MockExamLab/internal/models/DTO"
	"MockExamLab/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
)

type TestAPI struct {
	testService service.TestService
}

func ProvideTestAPI(s service.TestService) TestAPI {
	return TestAPI{testService: s}
}

// Create ... Create Test
// @Summary Create Test
// @Description Create Test
// @Tags Tests
// @Accept json
// @Param user body DTO.TestCreateRequest true "Create test"
// @Success 200 {object} DTO.TestCreateResponse
// @Failure 400 "error message detail"
// @Router /test [post]
// @Security ApiKeyAuth
func (api *TestAPI) Create(c *gin.Context) {
	var t DTO.TestCreateRequest
	u := getUserClaimsFromSession(c)
	if u == nil {
		c.JSON(http.StatusBadRequest, "User not Found")
		return
	}
	if err := c.ShouldBind(&t); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	data, reqConvertErr := mapper.TestCreateRequestToModel(&t, u.UserId)
	if reqConvertErr != nil {
		c.JSON(http.StatusBadRequest, reqConvertErr)
		return
	}
	if res, err := api.testService.Create(*data); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	} else {
		convertedData, resConvertErr := mapper.TestModelToCreateResponse(res)
		if resConvertErr != nil {
			c.JSON(http.StatusBadRequest, resConvertErr)
			return
		}
		c.JSON(http.StatusOK, &convertedData)
	}
}

// DeleteById ... Delete test by id
// @Summary Delete test by id
// @Description Delete test by id
// @Tags Tests
// @Accept json
// @Param id path string true "Test ID"
// @Success 200 "Success"
// @Failure 400 "error message detail"
// @Router /test/{id} [delete]
// @Security ApiKeyAuth
func (api *TestAPI) DeleteById(c *gin.Context) {
	id, ParsError := uuid.Parse(c.Param("id"))
	if ParsError != nil {
		c.JSON(http.StatusBadRequest, "Invalid Id")
		return
	}
	if err := api.testService.DeleteById(id); err != nil {
		c.JSON(http.StatusBadRequest, err)
	}
	c.AbortWithStatus(http.StatusOK)
}

// Update ... Update Test
// @Summary Update Test
// @Description Update Test
// @Tags Tests
// @Accept json
// @Param user body DTO.TestUpdateRequest true "Update test"
// @Success 200 {object} DTO.TestUpdateResponse
// @Failure 400 "error message detail"
// @Router /test [put]
// @Security ApiKeyAuth
func (api *TestAPI) Update(c *gin.Context) {
	var t DTO.TestUpdateRequest
	u := getUserClaimsFromSession(c)
	if u == nil {
		c.JSON(http.StatusBadRequest, "User not Found")
		return
	}
	if err := c.ShouldBind(&t); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	data, reqConvertErr := mapper.TestUpdateRequestToModel(&t, u.UserId)
	if reqConvertErr != nil {
		c.JSON(http.StatusBadRequest, reqConvertErr)
		return
	}
	if res, err := api.testService.Update(*data); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	} else {
		convertedData, resConvertErr := mapper.TestModelToUpdateResponse(res)
		if resConvertErr != nil {
			c.JSON(http.StatusBadRequest, resConvertErr)
			return
		}
		c.JSON(http.StatusOK, &convertedData)
	}
}

// FindById ... Find Test by id
// @Summary Find Test by id
// @Description Find Test by id
// @Tags Tests
// @Accept json
// @Param id path string true "Test ID"
// @Success 200 {object} DTO.TestResponse
// @Failure 400 "error message detail"
// @Router /test/{id} [get]
// @Security ApiKeyAuth
func (api *TestAPI) FindById(c *gin.Context) {
	id, ParsError := uuid.Parse(c.Param("id"))
	if ParsError != nil {
		c.JSON(http.StatusBadRequest, "Invalid Id")
		return
	}
	if res, err := api.testService.FindById(id); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	} else {
		convertedData, resConvertErr := mapper.TestModelToResponse(res)
		if resConvertErr != nil {
			c.JSON(http.StatusBadRequest, resConvertErr)
			return
		}
		c.JSON(http.StatusOK, &convertedData)
	}
}

// FindByIdWithPreload ... Get test wih all detail
// @Summary Get test wih all detail
// @Description get test with sections and section's question groups
// @Tags Tests
// @Accept json
// @Param id path string true "Test ID"
// @Success 200 {object} DTO.TestResponse
// @Failure 400 "error message detail"
// @Router /test/full/{id} [get]
// @Security ApiKeyAuth
func (api *TestAPI) FindByIdWithPreload(c *gin.Context) {
	id, ParsError := uuid.Parse(c.Param("id"))
	if ParsError != nil {
		c.JSON(http.StatusBadRequest, "Invalid Id")
		return
	}
	if res, err := api.testService.FindByIdWithPreloads(id); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	} else {
		convertedData, resConvertErr := mapper.TestModelToResponse(res)
		if resConvertErr != nil {
			c.JSON(http.StatusBadRequest, resConvertErr)
			return
		}
		c.JSON(http.StatusOK, &convertedData)
	}
}

// FindAll ... Find All
// @Summary Find All
// @Description type of "data" field in response is array of "DTO.TestResponse"
// @Tags Tests
// @Accept json
// @Param search query string false "search"
// @Param filter query string false "module:0 is ielts general"
// @Param page query int true "page"
// @Param page_size query int true "page size"
// @Param order_by query string false "order by"
// @Param order_direction query string false "order direction"
// @Success 200 {object} DTO.PaginateResponse
// @Failure 400 "error message detail"
// @Router /test/all [get]
// @Security ApiKeyAuth
func (api *TestAPI) FindAll(c *gin.Context) {
	params := getQueryParameters(c)
	if res, count, err := api.testService.FindAll(params); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	} else {
		convertedData, resConvertErr := mapper.TestModelsToArrayResponse(res, count)
		if resConvertErr != nil {
			c.JSON(http.StatusBadRequest, resConvertErr)
			return
		}
		c.JSON(http.StatusOK, &convertedData)
	}
}
