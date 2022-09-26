package handlers

import (
	"GolangRelational/pkg/instrastructure/entities"
	"GolangRelational/pkg/internal/utils"
	"GolangRelational/pkg/services"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

type OrderHandler interface {
	Add(e echo.Context) error
	Delete(e echo.Context) error
	GetById(e echo.Context) error
	Update(e echo.Context) error
	GetAll(e echo.Context) error
}

type orderHandler struct {
	orderService services.OrderService
}

func (o orderHandler) Add(e echo.Context) error {
	var order entities.Order
	if err := e.Bind(&order); err != nil {
		return e.JSON(http.StatusBadRequest,echo.Map{
			"error": err.Error(),
		})
	}

	response, err := o.orderService.Add(order)
	if err != nil {
		return e.JSON(http.StatusInternalServerError,echo.Map{
			"error": err.Error(),
		})
	}

	return e.JSON(http.StatusCreated, echo.Map{
		"data": response,
	})
}

func (o orderHandler) Delete(e echo.Context) error {
	id, err := strconv.Atoi(e.Param("id"))
	if err != nil {
		return e.JSON(http.StatusBadRequest,echo.Map{
			"error": err.Error(),
		})
	}

	err = o.orderService.Delete(id)
	if err != nil {
		return e.JSON(http.StatusInternalServerError,echo.Map{
			"error": err.Error(),
		})
	}

	return e.JSON(http.StatusNoContent, "")
}

func (o orderHandler) GetById(e echo.Context) error {
	id, err := strconv.Atoi(e.Param("id"))
	if err != nil {
		return e.JSON(http.StatusBadRequest,echo.Map{
			"error": err.Error(),
		})
	}

	response, err := o.orderService.GetById(id)
	if err != nil {
		return e.JSON(http.StatusInternalServerError,echo.Map{
			"error": err.Error(),
		})
	}

	return e.JSON(http.StatusOK, echo.Map{
		"data": response,
	})
}

func (o orderHandler) Update(e echo.Context) error {
	var order entities.Order

	id, err := strconv.Atoi(e.Param("id"))
	if err != nil {
		return e.JSON(http.StatusBadRequest,echo.Map{
			"error": err.Error(),
		})
	}

	if err := e.Bind(&order); err != nil {
		return e.JSON(http.StatusBadRequest,echo.Map{
			"error": err.Error(),
		})
	}

	order.Id = id

	resp, err := o.orderService.Update(order)
	if err != nil {
		return e.JSON(http.StatusInternalServerError,echo.Map{
			"error": err.Error(),
		})
	}

	return e.JSON(http.StatusOK, echo.Map{
		"data": resp,
	})
}

func (o orderHandler) GetAll(e echo.Context) error {
	var paginateModel utils.Pagination
	if e.QueryParam("page") != "" {
		resp, err := strconv.Atoi(e.QueryParam("page"))
		if err == nil {
			paginateModel.Page = resp
		}
	} else {
		paginateModel.Page = 1
	}

	if e.QueryParam("limit") != "" {
		resp, err := strconv.Atoi(e.QueryParam("limit"))
		if err == nil {
			paginateModel.Limit = resp
		}
	} else {
		paginateModel.Limit = 10
	}

	if e.QueryParam("sort") != "" {
		paginateModel.Sort = e.QueryParam("sort")
	}

	response, err := o.orderService.GetAll(paginateModel)
	if err != nil {
		return e.JSON(http.StatusInternalServerError,echo.Map{
			"error": err.Error(),
		})
	}

	return e.JSON(http.StatusOK, echo.Map{
		"data": response,
	})
}

func NewOrderHandler(orderService services.OrderService) OrderHandler {
	return &orderHandler{orderService: orderService}
}
