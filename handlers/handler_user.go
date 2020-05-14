package handlers

import (
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"golang_api/models"
	"golang_api/models/req"
	"golang_api/repositories"
	"net/http"
)

type HandlerUser struct {
	UserRepo repositories.UserRepo
}

func (handler HandlerUser) Add(c echo.Context) error {
	body := req.ReqSignUp{}

	if err := c.Bind(&body); err != nil {
		return c.JSON(http.StatusBadRequest, models.ResOK{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
			Data:    nil,
		})
	}

	if err := c.Validate(body); err != nil {
		return c.JSON(http.StatusBadRequest, models.ResOK{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
			Data:    nil,
		})
	}

	userID, err := uuid.NewUUID()
	if err != nil {
		return c.JSON(http.StatusForbidden, models.ResOK{
			Status:  http.StatusForbidden,
			Message: err.Error(),
			Data:    nil,
		})
	}
	
	user := models.User{
		ID:       userID.String(),
		Name:     body.Name,
		Email:    body.Email,
		Password: body.Password,
	}
	
	_, err = handler.UserRepo.Add(c.Request().Context(), user)
	if err != nil {
		return c.JSON(http.StatusConflict, models.ResOK{
			Status:  http.StatusConflict,
			Message: err.Error(),
			Data:    nil,
		})
	}
	
	return c.JSON(http.StatusOK, models.ResOK{
		Status:  http.StatusOK,
		Message: "OK",
		Data:    user,
	})
}

func (handler HandlerUser) GetList(c echo.Context) error {
	users, err := handler.UserRepo.GetList(c.Request().Context())
	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.ResOK{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
			Data:    nil,
		})
	}
	
	return c.JSON(http.StatusOK, models.ResOK{
		Status:  http.StatusOK,
		Message: "Ok",
		Data:    users,
	})
}

func (handler HandlerUser) GetUserById(c echo.Context) error {
	id := c.Param("id")
	
	user, err := handler.UserRepo.GetUserById(c.Request().Context(), id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, models.ResOK{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
			Data:    nil,
		})
	}
	
	return c.JSON(http.StatusOK, models.ResOK{
		Status:  http.StatusOK,
		Message: "Ok",
		Data:    user,
	})
}

func (handler HandlerUser) Update(c echo.Context) error {
	id := c.Param("id")
	body := req.ReqUpdateUser{}
	
	if err := c.Bind(&body); err != nil {
		return c.JSON(http.StatusBadRequest, models.ResOK{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
			Data:    nil,
		})
	}
	
	if err := c.Validate(body); err != nil {
		return c.JSON(http.StatusBadRequest, models.ResOK{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
			Data:    nil,
		})
	}
	
	user := models.User{
		ID:       id,
		Name:     body.Name,
		Email:    body.Email,
	}
	
	_, err := handler.UserRepo.Update(c.Request().Context(), user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.ResOK{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
			Data:    nil,
		})
	}
	
	return c.JSON(http.StatusOK, models.ResOK{
		Status:  http.StatusOK,
		Message: "Ok",
		Data:    user,
	})
}