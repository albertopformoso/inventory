package controller

import (
	"net/http"
	"strconv"

	"github.com/albertopformoso/inventory/internal/controller/dto"
	"github.com/albertopformoso/inventory/internal/helper"
	"github.com/albertopformoso/inventory/internal/model"
	"github.com/albertopformoso/inventory/internal/service"
	"github.com/labstack/echo/v4"
)

func (ctrl *controller) AddProduct(c echo.Context) error {
	// Get the payload from the request
	ctx := c.Request().Context()
	email := c.Get("email").(string)
	params := dto.AddProduct{}

	if err := c.Bind(&params); err != nil {
		ctrl.log.Err(err).Msg("failed to bind params")
		res := helper.BuildErrorResopnse("invalid parameters", ErrInvalidRequest.Error(), helper.EmptyObj{})
		return c.JSONPretty(http.StatusBadRequest, res, "  ")
	}

	if err := ctrl.dataValidator.Struct(params); err != nil {
		ctrl.log.Err(err).Msg("data validator: data isn't correct")
		res := helper.BuildErrorResopnse("invalid data", ErrInvalidRequest.Error(), helper.EmptyObj{})
		return c.JSONPretty(http.StatusBadRequest, res, "  ")
	}

	p := model.Product{
		Name:        params.Name,
		Description: params.Description,
		Price:       params.Price,
	}

	if err := ctrl.service.AddProduct(ctx, p, email); err != nil {
		ctrl.log.Err(err).Msg("add product failed")

		if err == service.ErrInvalidPermissions {
			res := helper.BuildErrorResopnse("Unauthorized", err.Error(), helper.EmptyObj{})
			return c.JSONPretty(http.StatusForbidden, res, "  ")
		}

		res := helper.BuildErrorResopnse("server error", ErrInternalServerError.Error(), helper.EmptyObj{})
		return c.JSONPretty(http.StatusInternalServerError, res, "  ")
	}

	res := helper.BuildResponse("product added successfully", params)
	return c.JSONPretty(http.StatusCreated, res, "  ")
}

func (ctrl *controller) GetProducts(c echo.Context) error {
	ctx := c.Request().Context()
	ctrl.log.Info().Msgf("Cookie: %v", c.Get("email"))

	pp, err := ctrl.service.GetProducts(ctx)
	if err != nil {
		ctrl.log.Err(err).Msg("faild to get products")
		res := helper.BuildErrorResopnse("cannot get products", ErrInternalServerError.Error(), helper.EmptyObj{})
		return c.JSONPretty(http.StatusInternalServerError, res, "  ")
	}

	res := helper.BuildResponse("products list", pp)
	return c.JSONPretty(http.StatusOK, res, "  ")
}

func (ctrl *controller) GetProduct(c echo.Context) error {
	ctx := c.Request().Context()

	id := c.Param("id")
	pid, err := strconv.ParseInt(id, 0, 64)
	if err != nil {
		ctrl.log.Err(err).Msg("invalid product id")
		res := helper.BuildErrorResopnse("invalid product id", ErrInvalidRequest.Error(), helper.EmptyObj{})
		return c.JSONPretty(http.StatusBadRequest, res, "  ")
	}

	p, err := ctrl.service.GetProduct(ctx, pid)
	if err != nil {
		ctrl.log.Err(err).Msg("faild to get the product")
		res := helper.BuildErrorResopnse("cannot get the product", ErrInternalServerError.Error(), helper.EmptyObj{})
		return c.JSONPretty(http.StatusInternalServerError, res, "  ")
	}

	res := helper.BuildResponse("product retrieved successfully", p)
	return c.JSONPretty(http.StatusOK, res, "  ")
}
