package handler

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/prrng/dealls/domain/usecase"
	"github.com/prrng/dealls/dto"
	"github.com/prrng/dealls/libs/auth"
	"github.com/prrng/dealls/libs/httpresp"
	"github.com/prrng/dealls/libs/logger"
)

type UserHandler struct {
	userUseCase usecase.UserUseCase
	jwtService  *auth.JWTService
	log         logger.Logger
}

func NewUserHandler(userUseCase usecase.UserUseCase, jwtService *auth.JWTService, log logger.Logger) *UserHandler {
	return &UserHandler{
		userUseCase: userUseCase,
		jwtService:  jwtService,
		log:         log,
	}
}

func (h *UserHandler) Login(c *gin.Context) {
	ctx := c.Request.Context()

	var req dto.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, httpresp.Err(err))
		return
	}

	employee, err := h.userUseCase.Authenticate(ctx, req.Email, req.Password)
	if err != nil {
		h.log.Error("Failed to authenticate user", err)
		c.JSON(http.StatusUnauthorized, httpresp.Err(errors.New("invalid credentials")))
		return
	}

	var role string
	if employee.IsAdmin {
		role = "admin" //TODO: set const
	}

	token, err := h.jwtService.GenerateToken(employee.ID, employee.Email, role)
	if err != nil {
		h.log.Error("Failed to generate token", err)
		c.JSON(http.StatusInternalServerError, httpresp.Err(errors.New(http.StatusText(http.StatusInternalServerError))))
		return
	}

	c.JSON(http.StatusOK, httpresp.OK(dto.LoginResponse{
		Token: token,
	}))
}
