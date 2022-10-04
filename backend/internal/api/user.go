package api

import (
	"MockExamLab/internal/mapper"
	"MockExamLab/internal/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserAPI struct {
	userService service.UserService
}

func ProvideUserAPI(s service.UserService) UserAPI {
	return UserAPI{userService: s}
}

/*// UserCreate ... Create User
// @Summary Create new user based on parameters
// @Description Create new user
// @Tags Users
// @Accept json
// @Param user body DTO.UserSignUpRequest true "Create user"
// @Success 200 {object} DTO.UserSignUpOrLoginUserResponse
// @Failure 400 "error message detail"
// @Router /user/signUp [post]
func (api *UserAPI) UserCreate(c *gin.Context) {
	var u DTO.UserSignUpRequest
	if err := c.ShouldBind(&u); err != nil {
		c.JSON(400, err)
		return
	}

	if res, err := api.userService.SignUpUser(mapper.UserSignUpRequestToModel(&u)); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
	} else {
		token := SignJWT(res)
		c.JSON(http.StatusOK, mapper.UserModelToSignupOrLoginResponse(res, token))
	}
}

// UserLogin ... Login User
// @Summary Login user with email and password
// @Description Login User
// @Tags Users
// @Accept json
// @Param user body DTO.UserLoginRequest true "Login Data"
// @Success 200 {object} DTO.UserSignUpOrLoginUserResponse
// @Failure 400 "error message detail"
// @Router /user/login [post]
func (api *UserAPI) UserLogin(c *gin.Context) {
	var req DTO.UserLoginRequest
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	if res, err := api.userService.LoginUser(req); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
	} else {
		token := SignJWT(res)
		c.JSON(http.StatusOK, mapper.UserModelToSignupOrLoginResponse(res, token))
	}

}*/

// UserLoginOrSignUp ... Login or signup user
// @Summary Login or signup user
// @Description Login or signup user with firebase token
// @Tags Users
// @Accept json
// @Success 200 {object} DTO.UserSignUpOrLoginUserResponse
// @Failure 400 "error message detail"
// @Router /user/auth [post]
// @Security ApiKeyAuth
func (api *UserAPI) UserLoginOrSignUp(c *gin.Context) {
	u := getUserClaimsFromSession(c)
	if u == nil {
		c.JSON(http.StatusBadRequest, "User not Found")
		return
	}

	if res, err := api.userService.LoginOrSignUpUser(*u); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
	} else {
		c.JSON(http.StatusOK, mapper.UserModelToSignupOrLoginResponse(res))
	}
}

// FindWithFilterAndPagination ... Get all user
// @Summary Get all user with pagination
// @Description type of "data" field in response is array of "DTO.UserSignUpOrLoginUserResponse"
// @Tags Users
// @Accept json
// @Param search query string false "search"
// @Param filter query string false "filter"
// @Param page query int true "page"
// @Param page_size query int true "page size"
// @Param order_by query string false "order by"
// @Param order_direction query string false "order direction"
// @Success 200 {object} DTO.PaginateResponse
// @Failure 400 "error message detail"
// @Router /user [get]
// @Security ApiKeyAuth
func (api *UserAPI) FindWithFilterAndPagination(c *gin.Context) {
	params := getQueryParameters(c)
	if res, count, err := api.userService.FindAllFilterPagination(params); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
	} else {
		c.JSON(http.StatusOK, mapper.UserModelToArrayResponse(res, count))
	}

}
