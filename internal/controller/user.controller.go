package controller

import (
	"net/http"
	"strconv"

	"github.com/albertopformoso/inventory/encryption"
	"github.com/albertopformoso/inventory/internal/controller/dto"
	"github.com/albertopformoso/inventory/internal/helper"
	"github.com/albertopformoso/inventory/internal/service"

	"github.com/labstack/echo/v4"
)

func (ctrl *controller) RegisterUser(c echo.Context) error {
	ctx := c.Request().Context()
	params := dto.RegisterUser{}

	if err := c.Bind(&params); err != nil {
		res := helper.BuildErrorResopnse("invalid fields", ErrInvalidRequest.Error(), helper.EmptyObj{})
		return c.JSONPretty(http.StatusBadRequest, res, "  ")
	}

	if err := ctrl.dataValidator.Struct(params); err != nil {
		res := helper.BuildErrorResopnse("invalid data", err.Error(), helper.EmptyObj{})
		return c.JSONPretty(http.StatusBadRequest, res, "  ")
	}

	if err := ctrl.service.RegisterUser(ctx, params.Email, params.Name, params.Password); err != nil {
		if err == service.ErrUserAlreadyExists {
			res := helper.BuildErrorResopnse("invalid request", err.Error(), helper.EmptyObj{})
			return c.JSONPretty(http.StatusConflict, res, "  ")
		}

		res := helper.BuildErrorResopnse("", ErrInternalServerError.Error(), helper.EmptyObj{})
		return c.JSONPretty(http.StatusInternalServerError, res, "  ")
	}

	res := helper.BuildResponse("user created successfully", map[string]any{"email": params.Email, "name": params.Email})
	return c.JSONPretty(http.StatusCreated, res, "  ")
}

func (ctrl *controller) LoginUser(c echo.Context) error {
	ctx := c.Request().Context()
	params := dto.LoginUser{}

	if err := c.Bind(&params); err != nil {
		ctrl.log.Err(err).Msg("user params binding failed")
		res := helper.BuildErrorResopnse("invalid fields", ErrInvalidRequest.Error(), helper.EmptyObj{})
		return c.JSONPretty(http.StatusBadRequest, res, "  ")
	}

	if err := ctrl.dataValidator.Struct(params); err != nil {
		ctrl.log.Err(err).Msg("data validator: data isn't correct")
		res := helper.BuildErrorResopnse("invalid data", ErrInvalidRequest.Error(), helper.EmptyObj{})
		return c.JSONPretty(http.StatusBadRequest, res, "  ")
	}

	u, err := ctrl.service.LoginUser(ctx, params.Email, params.Password)
	if err != nil {
		ctrl.log.Err(err).Msg("login failed")
		res := helper.BuildErrorResopnse("login failed", ErrInternalServerError.Error(), helper.EmptyObj{})
		return c.JSONPretty(http.StatusInternalServerError, res, "  ")
	}

	token, err := encryption.SignedLoginToken(u)
	if err != nil {
		ctrl.log.Err(err).Msg("faild to generate token")
		res := helper.BuildErrorResopnse("failed to generate token", ErrInternalServerError.Error(), helper.EmptyObj{})
		return c.JSONPretty(http.StatusInternalServerError, res, "  ")
	}

	cookie := &http.Cookie{
		Name:     "Authorization",
		Value:    token,
		Secure:   true,
		SameSite: http.SameSiteNoneMode,
		HttpOnly: true,
		Path:     "/",
	}

	c.SetCookie(cookie)

	response := map[string]interface{}{
		"message": "success login",
		// "cookie":   cookie,
	}

	return c.JSONPretty(http.StatusOK, response, "  ")
}

func (ctrl *controller) AddUserRole(c echo.Context) error {
	ctx := c.Request().Context()

	userId := c.QueryParam("user_id")
	roleId := c.QueryParam("role_id")

	uid, err := strconv.ParseInt(userId, 0, 64)
	if err != nil {
		ctrl.log.Err(err).Msg("string to int conversion failed")
		res := helper.BuildErrorResopnse("invalid user id", err.Error(), helper.EmptyObj{})
		return c.JSONPretty(http.StatusBadRequest, res, "  ")
	}

	rid, err := strconv.ParseInt(roleId, 0, 64)
	if err != nil {
		ctrl.log.Err(err).Msg("string to int conversion failed")
		res := helper.BuildErrorResopnse("invalid role id", err.Error(), helper.EmptyObj{})
		return c.JSONPretty(http.StatusBadRequest, res, "  ")
	}

	if err := ctrl.service.AddUserRole(ctx, uid, rid); err != nil {
		ctrl.log.Err(err).Msg("adding user role failed")
		res := helper.BuildErrorResopnse("failed to add user role", ErrInternalServerError.Error(), helper.EmptyObj{})
		return c.JSONPretty(http.StatusInternalServerError, res, "  ")
	}

	res := helper.BuildResponse("user role added successfully", map[string]any{"user_id": uid, "role_id": rid})
	return c.JSONPretty(http.StatusOK, res, "  ")
}

func (ctrl *controller) RemoveUserRole(c echo.Context) error {
	ctx := c.Request().Context()

	userId := c.QueryParam("user_id")
	roleId := c.QueryParam("role_id")

	uid, err := strconv.ParseInt(userId, 0, 64)
	if err != nil {
		ctrl.log.Err(err).Msg("string to int conversion failed")
		res := helper.BuildErrorResopnse("invalid user id", err.Error(), helper.EmptyObj{})
		return c.JSONPretty(http.StatusBadRequest, res, "  ")
	}

	rid, err := strconv.ParseInt(roleId, 0, 64)
	if err != nil {
		ctrl.log.Err(err).Msg("string to int conversion failed")
		res := helper.BuildErrorResopnse("invalid role id", err.Error(), helper.EmptyObj{})
		return c.JSONPretty(http.StatusBadRequest, res, "  ")
	}

	if err := ctrl.service.RemoveUserRole(ctx, uid, rid); err != nil {
		ctrl.log.Err(err).Msg("removing user role failed")
		res := helper.BuildErrorResopnse("failed to remove user role", ErrInternalServerError.Error(), helper.EmptyObj{})
		return c.JSONPretty(http.StatusInternalServerError, res, "  ")
	}

	res := helper.BuildResponse("user role removed successfully", map[string]any{"user_id": uid, "role_id": rid})
	return c.JSONPretty(http.StatusOK, res, "  ")
}
