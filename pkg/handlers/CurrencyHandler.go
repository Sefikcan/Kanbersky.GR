package handlers

import (
	"GolangRelational/pkg/instrastructure/entities"
	"GolangRelational/pkg/internal/utils"
	"GolangRelational/pkg/services"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

type CurrencyHandler interface {
	Add(e echo.Context) error
	Delete(e echo.Context) error
	GetById(e echo.Context) error
	Update(e echo.Context) error
	GetAll(e echo.Context) error
}

type currencyHandler struct {
	currencyService services.CurrencyService
}

func (c currencyHandler) GetAll(e echo.Context) error {
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

	response, err := c.currencyService.GetAll(paginateModel)
	if err != nil {
		return e.JSON(http.StatusInternalServerError,echo.Map{
			"error": err.Error(),
		})
	}

	return e.JSON(http.StatusOK, echo.Map{
		"data": response,
	})
}

func (c currencyHandler) Update(e echo.Context) error {
	var currency entities.Currency

	id, err := strconv.Atoi(e.Param("id"))
	if err != nil {
		return e.JSON(http.StatusBadRequest,echo.Map{
			"error": err.Error(),
		})
	}

	currency.ID = id

	resp, err := c.currencyService.Update(currency)
	if err != nil {
		return e.JSON(http.StatusInternalServerError,echo.Map{
			"error": err.Error(),
		})
	}

	return e.JSON(http.StatusOK, echo.Map{
		"data": resp,
	})
}

func (c currencyHandler) GetById(e echo.Context) error {
	id, err := strconv.Atoi(e.Param("id"))
	if err != nil {
		return e.JSON(http.StatusBadRequest,echo.Map{
			"error": err.Error(),
		})
	}

	response, err := c.currencyService.GetById(id)
	if err != nil {
		return e.JSON(http.StatusInternalServerError,echo.Map{
			"error": err.Error(),
		})
	}

	return e.JSON(http.StatusOK, echo.Map{
		"data": response,
	})
}

func (c currencyHandler) Delete(e echo.Context) error {
	id, err := strconv.Atoi(e.Param("id"))
	if err != nil {
		return e.JSON(http.StatusBadRequest,echo.Map{
			"error": err.Error(),
		})
	}

	err = c.currencyService.Delete(id)
	if err != nil {
		return e.JSON(http.StatusInternalServerError,echo.Map{
			"error": err.Error(),
		})
	}

	return e.JSON(http.StatusNoContent, "")
}

func (c currencyHandler) Add(e echo.Context) error {
	var currency entities.Currency
	if err := e.Bind(&currency); err != nil {
		return e.JSON(http.StatusBadRequest,echo.Map{
			"error": err.Error(),
		})
	}
	
	response, err := c.currencyService.Add(currency)
	if err != nil {
		return e.JSON(http.StatusInternalServerError,echo.Map{
			"error": err.Error(),
		})
	}

	return e.JSON(http.StatusCreated, echo.Map{
		"data": response,
	})
}

func NewCurrencyHandler(currencyService services.CurrencyService) CurrencyHandler {
	return &currencyHandler{currencyService: currencyService}
}
