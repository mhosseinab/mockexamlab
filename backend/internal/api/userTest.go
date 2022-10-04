package api

import (
	"MockExamLab/internal/mapper"
	"MockExamLab/internal/models/DTO"
	"MockExamLab/internal/models/customErrors"
	"MockExamLab/internal/service"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
)

type UserTestAPI struct {
	s service.UserTestService
}

func ProvideUserTestAPI(s service.UserTestService) UserTestAPI {
	return UserTestAPI{s: s}
}

// Create ... Create user-test
// @Summary Create user-test
// @Description Create user-test
// @Tags User-tests
// @Accept json
// @Param user body DTO.UserTestCreateRequest true "Create test"
// @Success 200 {object} DTO.UserTestCreateResponse
// @Failure 400 "error message detail"
// @Router /user-test [post]
// @Security ApiKeyAuth
func (api *UserTestAPI) Create(c *gin.Context) {
	var t DTO.UserTestCreateRequest
	u := getUserClaimsFromSession(c)
	if u == nil {
		c.JSON(http.StatusBadRequest, "User not Found")
		return
	}
	if err := c.ShouldBind(&t); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	if res, err := api.s.Create(mapper.UserTestCreateRequestToModel(&t, u.UserId)); err != nil {
		if errors.Is(err, &customErrors.SubscriptionError{}) {
			c.JSON(http.StatusMethodNotAllowed, err.Error())
		} else {
			c.JSON(http.StatusBadRequest, err)
			return
		}
	} else {
		c.JSON(http.StatusOK, mapper.UserTestModelToCreateResponse(res))
	}
}

// DeleteById ... Delete user-test by id
// @Summary Delete user-test by id
// @Description Delete user-test by id
// @Tags User-tests
// @Accept json
// @Param id path string true "user-test ID"
// @Success 200 "Success"
// @Failure 400 "error message detail"
// @Router /user-test/{id} [delete]
// @Security ApiKeyAuth
func (api *UserTestAPI) DeleteById(c *gin.Context) {
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

/*func (api *TestAPI) Update(c *gin.Context) {
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
	if res, err := api.testService.Update(mapper.TestUpdateRequestToModel(&t, u.UserId)); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	} else {
		c.JSON(http.StatusOK, mapper.TestModelToUpdateResponse(res))
	}
}*/

// FindById ... Find user-test by id
// @Summary Find user-test by id
// @Description Find user-test by id
// @Tags User-tests
// @Accept json
// @Param id path string true "user-test ID"
// @Success 200 {object} DTO.UserTestResponse
// @Failure 400 "error message detail"
// @Router /user-test/{id} [get]
// @Security ApiKeyAuth
func (api *UserTestAPI) FindById(c *gin.Context) {
	id, ParsError := uuid.Parse(c.Param("id"))
	if ParsError != nil {
		c.JSON(http.StatusBadRequest, "Invalid Id")
		return
	}
	if res, err := api.s.FindById(id); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	} else {
		c.JSON(http.StatusOK, mapper.UserTestModelToResponse(res))
	}
}

// FindAllByUserId ... Find user-test by user id
// @Summary Find all user's test with user id specification and pagination
// @Description type of "data" field in response is array of "DTO.UserTestResponse"
// @Tags User-tests
// @Accept json
// @Param id query string true "user-test ID"
// @Param page query int true "page"
// @Param page_size query int true "page size"
// @Success 200 {object} DTO.PaginateResponse
// @Failure 400 "error message detail"
// @Router /user-test/all [get]
// @Security ApiKeyAuth
func (api *UserTestAPI) FindAllByUserId(c *gin.Context) {
	params := getQueryParameters(c)
	userId, ParsError := uuid.Parse(c.Query("user_id"))
	if ParsError != nil {
		c.JSON(http.StatusBadRequest, "Invalid Id")
		return
	}
	if res, count, err := api.s.FindAllByUserId(userId, params); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	} else {
		c.JSON(http.StatusOK, mapper.UserTestModelsToArrayResponse(res, count))
	}
}

// SubmitTest ... Submit Test
// @Summary Submit Test
// @Description Submit Test
// @Tags User-tests
// @Accept json
// @Param test_id query string true "Test ID"
// @Param user_test_id query string true "User Test ID"
// @Success 200 {object} DTO.UserTestResultResponse
// @Failure 400 "error message detail"
// @Router /user-test/submit [put]
// @Security ApiKeyAuth
func (api *UserTestAPI) SubmitTest(c *gin.Context) {
	testId, errParsTestId := uuid.Parse(c.Query("test_id"))
	if errParsTestId != nil {
		c.JSON(http.StatusBadRequest, errParsTestId.Error())
		return
	}
	userTestId, errParsUserTestId := uuid.Parse(c.Query("user_test_id"))
	if errParsUserTestId != nil {
		c.JSON(http.StatusBadRequest, errParsUserTestId.Error())
		return
	}
	if res, err := api.s.SubmitTest(testId, userTestId); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	} else {
		c.JSON(http.StatusOK, res)
	}
}
