package api

import (
	"MockExamLab/internal/mapper"
	"MockExamLab/internal/models/DTO"
	"MockExamLab/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	"path"
	"path/filepath"
	"strconv"
	"time"
)

type UserAnswerAPI struct {
	s service.UserAnswerService
}

func ProvideUserAnswerAPI(s service.UserAnswerService) UserAnswerAPI {
	return UserAnswerAPI{s: s}
}

// Create ... Create user-answer
// @Summary Create user-answer
// @Description Create user-answer
// @Tags User-answers
// @Accept json
// @Param user body DTO.UserAnswerCreateRequest true "Create test"
// @Success 200 {object} DTO.UserAnswerCreateResponse
// @Failure 400 "error message detail"
// @Router /user-answer [post]
// @Security ApiKeyAuth
func (api *UserAnswerAPI) Create(c *gin.Context) {
	var t DTO.UserAnswerCreateRequest
	u := getUserClaimsFromSession(c)
	if u == nil {
		c.JSON(http.StatusBadRequest, "User not Found")
		return
	}
	if err := c.ShouldBind(&t); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	if res, err := api.s.Create(mapper.UserAnswerCreateRequestToModel(&t)); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	} else {
		c.JSON(http.StatusOK, mapper.UserAnswerModelToCreateResponse(res))
	}
}

// CreateSpeaking ... Create speaking user-answer
// @Summary Create speaking user-answer
// @Description Create speaking user-answer
// @Tags User-answers
// @Accept json
// @Param user body DTO.UserAnswerCreateRequest true "Create test"
// @Success 200 {object} DTO.UserAnswerCreateResponse
// @Failure 400 "error message detail"
// @Router /user-answer/speaking [post]
// @Security ApiKeyAuth
func (api *UserAnswerAPI) CreateSpeaking(c *gin.Context) {

	file, errForm := c.FormFile("upload")
	if errForm != nil {
		c.JSON(http.StatusBadRequest, errForm.Error())
	}
	extension := filepath.Ext(file.Filename)
	fileNamePath := path.Join("./upload/answer/speaking", strconv.Itoa(time.Now().Year()), "/", time.Now().Month().String(), "/", uuid.New().String()+extension)
	if err := c.SaveUploadedFile(file, fileNamePath); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
	}
	var t DTO.UserAnswerCreateRequest
	u := getUserClaimsFromSession(c)
	if u == nil {
		c.JSON(http.StatusBadRequest, "User not Found")
		return
	}
	if err := c.ShouldBind(&t); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	t.Answer = fileNamePath
	if res, err := api.s.Create(mapper.UserAnswerCreateRequestToModel(&t)); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	} else {
		c.JSON(http.StatusOK, mapper.UserAnswerModelToCreateResponse(res))
	}
}

// BatchCreate ... Batch Create user-answer
// @Summary BatchCreate user-answer
// @Description BatchCreate user-answer
// @Tags User-answers
// @Accept json
// @Param user body DTO.UserAnswerCreateRequest true "Create user-answer"
// @Success 200 {array} DTO.UserAnswerCreateResponse
// @Failure 400 "error message detail"
// @Router /user-answer/batch [post]
// @Security ApiKeyAuth
func (api *UserAnswerAPI) BatchCreate(c *gin.Context) {
	var t []DTO.UserAnswerCreateRequest
	u := getUserClaimsFromSession(c)
	if u == nil {
		c.JSON(http.StatusBadRequest, "User not Found")
		return
	}
	if err := c.ShouldBind(&t); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	if res, err := api.s.BatchCreate(mapper.ArrayUserAnswerCreateRequestToModel(&t)); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	} else {
		c.JSON(http.StatusOK, mapper.ArrayUserAnswerModelToCreateResponse(res))
	}
}

// DeleteById ... Delete user-answer by id
// @Summary Delete user-answer by id
// @Description Delete user-answer by id
// @Tags User-answers
// @Accept json
// @Param id path string true "user-answer ID"
// @Success 200 "Success"
// @Failure 400 "error message detail"
// @Router /user-answer/{id} [delete]
// @Security ApiKeyAuth
func (api *UserAnswerAPI) DeleteById(c *gin.Context) {
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

// UpdateAnswer ... Update user answer
// @Summary Update user answer
// @Description Update User answer
// @Tags User-answers
// @Accept json
// @Param user body DTO.UserAnswerUpdateAnswerRequest true "Update User-Answer"
// @Success 200 {object} DTO.UserAnswerUpdateAnswerResponse
// @Failure 400 "error message detail"
// @Router /user-answer/answer [put]
// @Security ApiKeyAuth
func (api *UserAnswerAPI) UpdateAnswer(c *gin.Context) {

	var t DTO.UserAnswerUpdateAnswerRequest
	u := getUserClaimsFromSession(c)
	if u == nil {
		c.JSON(http.StatusBadRequest, "User not Found")
		return
	}
	if err := c.ShouldBind(&t); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	if res, err := api.s.UpdateAnswer(t.Id, t.Answer); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	} else {
		c.JSON(http.StatusOK, mapper.UserAnswerModelToUpdateAnswerResponse(res))
	}
}

// UpdateSpeakingAnswer ... Update speaking user answer
// @Summary Update speaking user answer
// @Description Update speaking User answer
// @Tags User-answers
// @Accept json
// @Param user body DTO.UserAnswerUpdateAnswerRequest true "Update User-Answer"
// @Success 200 {object} DTO.UserAnswerUpdateAnswerResponse
// @Failure 400 "error message detail"
// @Router /user-answer/answer/speaking [put]
// @Security ApiKeyAuth
func (api *UserAnswerAPI) UpdateSpeakingAnswer(c *gin.Context) {
	file, errForm := c.FormFile("upload")
	if errForm != nil {
		c.JSON(http.StatusBadRequest, errForm.Error())
	}
	extension := filepath.Ext(file.Filename)
	fileNamePath := path.Join("./upload/answer/speaking", strconv.Itoa(time.Now().Year()), "/", time.Now().Month().String(), "/", uuid.New().String()+extension)
	if err := c.SaveUploadedFile(file, fileNamePath); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
	}
	var t DTO.UserAnswerUpdateAnswerRequest
	u := getUserClaimsFromSession(c)
	if u == nil {
		c.JSON(http.StatusBadRequest, "User not Found")
		return
	}
	if err := c.ShouldBind(&t); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	t.Answer = fileNamePath
	if res, err := api.s.UpdateAnswer(t.Id, t.Answer); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	} else {
		c.JSON(http.StatusOK, mapper.UserAnswerModelToUpdateAnswerResponse(res))
	}
}

// UpdateMarker ... Update user answer score
// @Summary Update user answer score
// @Description Update user answer score
// @Tags User-answers
// @Accept json
// @Param user body DTO.UserAnswerUpdateMarkerRequest true "Update User-Answer"
// @Success 200 {object} DTO.UserAnswerUpdateMarkerResponse
// @Failure 400 "error message detail"
// @Router /user-answer/marker [put]
// @Security ApiKeyAuth
func (api *UserAnswerAPI) UpdateMarker(c *gin.Context) {
	var t DTO.UserAnswerUpdateMarkerRequest
	u := getUserClaimsFromSession(c)
	if u == nil {
		c.JSON(http.StatusBadRequest, "User not Found")
		return
	}
	if err := c.ShouldBind(&t); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	if res, err := api.s.UpdateMarker(t.Id, t.MarkerScore, t.MarkerComment); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	} else {
		c.JSON(http.StatusOK, mapper.UserAnswerModelToUpdateMarkerResponse(res))
	}
}

// FindById ... Find user-answer by id
// @Summary Find user-answer by id
// @Description Find user-answer by id
// @Tags User-answers
// @Accept json
// @Param id path string true "User answer ID"
// @Success 200 {object} DTO.UserAnswerResponse
// @Failure 400 "error message detail"
// @Router /user-answer/{id} [get]
// @Security ApiKeyAuth
func (api *UserAnswerAPI) FindById(c *gin.Context) {
	id, ParsError := uuid.Parse(c.Param("id"))
	if ParsError != nil {
		c.JSON(http.StatusBadRequest, "Invalid Id")
		return
	}
	if res, err := api.s.FindById(id); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	} else {
		c.JSON(http.StatusOK, mapper.UserAnswerModelToResponse(res))
	}
}

// FindAllByUserTestId ... Find user-answer by user test id
// @Summary Find user-answer by user test id
// @Description Find user-answer by user test id
// @Tags User-answers
// @Accept json
// @Param id query string true "User Test ID"
// @Success 200 {array} DTO.UserAnswerResponse
// @Failure 400 "error message detail"
// @Router /user-answer/all [get]
// @Security ApiKeyAuth
func (api *UserAnswerAPI) FindAllByUserTestId(c *gin.Context) {
	userTestId, ParsError := uuid.Parse(c.Query("user_test_id"))
	if ParsError != nil {
		c.JSON(http.StatusBadRequest, "Invalid Id")
		return
	}
	if res, err := api.s.FindAllByUserTestId(userTestId); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	} else {
		c.JSON(http.StatusOK, mapper.UserAnswerModelsToArrayResponse(res))
	}
}
