package api

import (
	"MockExamLab/internal/mapper"
	"MockExamLab/internal/models/DTO"
	"MockExamLab/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
)

type QuestionAPI struct {
	s service.QuestionService
}

func ProvideQuestionAPI(s service.QuestionService) QuestionAPI {
	return QuestionAPI{s: s}
}

// Create ... Create Question
// @Summary Create Question
// @Description Create Question
// @Tags Questions
// @Accept json
// @Param user body DTO.QuestionCreateRequest true "Create Question"
// @Success 200 {object} DTO.QuestionCreateResponse
// @Failure 400 "error message detail"
// @Router /question [post]
// @Security ApiKeyAuth
func (api *QuestionAPI) Create(c *gin.Context) {
	var t DTO.QuestionCreateRequest
	u := getUserClaimsFromSession(c)
	if u == nil {
		c.JSON(http.StatusBadRequest, "User not Found")
		return
	}
	if err := c.ShouldBind(&t); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	if res, err := api.s.Create(mapper.QuestionCreateRequestToModel(&t, u.UserId)); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	} else {
		c.JSON(http.StatusOK, mapper.QuestionModelToCreateResponse(res))
	}
}

// DeleteById ... question by id
// @Summary Delete question by id
// @Description Delete question by id
// @Tags Questions
// @Accept json
// @Param id path string true "Questions ID"
// @Success 200 "Success"
// @Failure 400 "error message detail"
// @Router /question/{id} [delete]
// @Security ApiKeyAuth
func (api *QuestionAPI) DeleteById(c *gin.Context) {
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

// Update ... Update Question
// @Summary Update Question
// @Description Update Question
// @Tags Questions
// @Accept json
// @Param user body DTO.QuestionUpdateRequest true "Update Question"
// @Success 200 {object} DTO.QuestionUpdateResponse
// @Failure 400 "error message detail"
// @Router /question [put]
// @Security ApiKeyAuth
func (api *QuestionAPI) Update(c *gin.Context) {
	var t DTO.QuestionUpdateRequest
	u := getUserClaimsFromSession(c)
	if u == nil {
		c.JSON(http.StatusBadRequest, "User not Found")
		return
	}
	if err := c.ShouldBind(&t); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	if res, err := api.s.Update(mapper.QuestionUpdateRequestToModel(&t, u.UserId)); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	} else {
		c.JSON(http.StatusOK, mapper.QuestionModelToUpdateResponse(res))
	}
}

// FindById ... Find Question by id
// @Summary Find Question by id
// @Description Find Question by id
// @Tags Questions
// @Accept json
// @Param id path string true "Question ID"
// @Success 200 {object} DTO.QuestionResponse
// @Failure 400 "error message detail"
// @Router /question/{id} [get]
// @Security ApiKeyAuth
func (api *QuestionAPI) FindById(c *gin.Context) {
	id, ParsError := uuid.Parse(c.Param("id"))
	if ParsError != nil {
		c.JSON(http.StatusBadRequest, "Invalid Id")
		return
	}
	if res, err := api.s.FindById(id); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	} else {
		c.JSON(http.StatusOK, mapper.QuestionModelToResponse(res))
	}
}

// FindAllByGroupId ... Find Questions by group id
// @Summary Find Questions by group id
// @Description Find Questions by group id
// @Tags Questions
// @Accept json
// @Param group_id query string true "group ID"
// @Success 200 {array} DTO.QuestionResponse
// @Failure 400 "error message detail"
// @Router /question/all [get]
// @Security ApiKeyAuth
func (api *QuestionAPI) FindAllByGroupId(c *gin.Context) {
	groupId, ParsError := uuid.Parse(c.Query("group_id"))
	if ParsError != nil {
		c.JSON(http.StatusBadRequest, "Invalid Id")
		return
	}
	if res, err := api.s.FindAllByGroupId(groupId); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	} else {
		c.JSON(http.StatusOK, mapper.QuestionModelsToArrayResponse(res))
	}
}

// FindAllByTestId ... Find Questions by test id
// @Summary Find Questions by test id with statistic data
// @Description Find Questions by test id with statistic data
// @Tags Questions
// @Accept json
// @Param test_id query string true "Test ID"
// @Param user_test_id query string true "User Test ID"
// @Success 200 {array} DTO.QuestionWithUserAnswerResponse
// @Failure 400 "error message detail"
// @Router /question/test/all [get]
// @Security ApiKeyAuth
func (api *QuestionAPI) FindAllByTestId(c *gin.Context) {
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
	if res, err := api.s.FindAllByTestId(testId, userTestId); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	} else {
		c.JSON(http.StatusOK, mapper.QuestionModelWithUserAnswerToArrayResponse(res))
	}
}
