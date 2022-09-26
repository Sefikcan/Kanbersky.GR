package pkg

import (
	"GolangRelational/pkg/dto/currency"
	"GolangRelational/pkg/dto/order"
	"GolangRelational/pkg/handlers"
	"GolangRelational/pkg/instrastructure"
	"GolangRelational/pkg/instrastructure/entities"
	"GolangRelational/pkg/services"
	"context"
	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"github.com/pangpanglabs/echoswagger/v2"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

type server struct {
	echo *echo.Echo
	v *validator.Validate
}

func NewServer() *server {
	return &server{echo: echo.New(), v: validator.New()}
}

func (s *server) Run() error {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM, syscall.SIGINT)
	defer cancel()

	se := echoswagger.New(s.echo, "doc/", &echoswagger.Info{
		Title:          "Golang Relational Api",
		Description:    "This is a sample Postgresql relational crud operation",
		Version:        "1.0.0",
		TermsOfService: "http://swagger.io/terms/",
		Contact: &echoswagger.Contact{
			Name: "Şefik Can Kanber",
		},
		License: &echoswagger.License{
			Name: "Apache 2.0",
			URL:  "http://www.apache.org/licenses/LICENSE-2.0.html",
		},
	})

	se.SetExternalDocs("Find out more about Swagger", "http://swagger.io").
		SetResponseContentType("application/xml", "application/json").
		SetUI(echoswagger.UISetting{DetachSpec: true, HideTop: true})

	currencyRepository := instrastructure.NewCurrencyRepository(Connect())
	currencyService := services.NewCurrencyService(currencyRepository)
	currencyHandler := handlers.NewCurrencyHandler(currencyService)
	orderRepository := instrastructure.NewOrderRepository(Connect())
	orderService := services.NewOrderService(orderRepository, currencyService)
	orderHandler := handlers.NewOrderHandler(orderService)

	c := se.Group("currency","/api/v1/currencies")

	c.POST("", currencyHandler.Add).
		AddParamBody(currency.Request{}, "body", "Currency object that needs to be added to the database", true).
		AddResponse(http.StatusCreated, "success", nil, nil).
		AddResponse(http.StatusBadRequest, "invalid input", nil, nil).
		AddResponse(http.StatusInternalServerError, "Db operation failed", nil, nil).
		SetRequestContentType("application/json", "application/xml").
		SetSummary("Add a new currency to the database")

	c.PUT("/:id", currencyHandler.Update).
		AddParamBody(currency.Request{},"body", "Currency object that needs to be update to the store", true).
		AddParamPath(int64(0), "id", "Currency id to update").
		AddResponse(http.StatusBadRequest, "Invalid ID supplied", nil, nil).
		AddResponse(http.StatusNotFound, "Currency not found", nil, nil).
		AddResponse(http.StatusOK, "Success", nil, nil).
		AddResponse(http.StatusInternalServerError, "Db operation failed", nil, nil).
		SetSummary("Update a Currency")

	c.DELETE("/:id", currencyHandler.Delete).
		AddParamPath(int64(0), "id", "Currency id to delete").
		AddResponse(http.StatusBadRequest, "Invalid ID supplied", nil, nil).
		AddResponse(http.StatusNotFound, "Currency not found", nil, nil).
		AddResponse(http.StatusNoContent, "Success", nil, nil).
		AddResponse(http.StatusInternalServerError, "Db operation failed", nil, nil).
		SetSummary("Delete a Currency")

	c.GET("/:id", currencyHandler.GetById).
		AddParamPath(int64(0), "id", "Currency id to get").
		AddResponse(http.StatusBadRequest, "Invalid ID supplied", nil, nil).
		AddResponse(http.StatusNotFound, "Currency not found", nil, nil).
		AddResponse(http.StatusOK, "Success", nil, nil).
		AddResponse(http.StatusInternalServerError, "Db operation failed", nil, nil).
		SetSummary("Get single Currency")

	c.GET("", currencyHandler.GetAll).
		AddParamQuery(int64(0), "page", "Page query param", false).
		AddParamQuery(int64(0), "limit", "Limit query param", false).
		AddParamQuery("", "sort", "Sort query param", false).
		AddResponse(http.StatusOK, "Success", nil, nil).
		AddResponse(http.StatusInternalServerError, "Db operation failed", nil, nil).
		SetSummary("Get All Currency")

	o := se.Group("order","/api/v1/orders")

	o.POST("", orderHandler.Add).
		AddParamBody(order.OrderRequest{}, "body", "Order object that needs to be added to the database", true).
		AddResponse(http.StatusCreated, "success", nil, nil).
		AddResponse(http.StatusBadRequest, "invalid input", nil, nil).
		AddResponse(http.StatusInternalServerError, "Db operation failed", nil, nil).
		SetRequestContentType("application/json", "application/xml").
		SetSummary("Add a new order to the database")

	o.PUT("/:id", orderHandler.Update).
		AddParamBody(order.OrderRequest{},"body", "Order object that needs to be update to the store", true).
		AddParamPath(int64(0), "id", "Order id to update").
		AddResponse(http.StatusBadRequest, "Invalid ID supplied", nil, nil).
		AddResponse(http.StatusNotFound, "Order not found", nil, nil).
		AddResponse(http.StatusOK, "Success", nil, nil).
		AddResponse(http.StatusInternalServerError, "Db operation failed", nil, nil).
		SetSummary("Update a Order")

	o.DELETE("/:id", orderHandler.Delete).
		AddParamPath(int64(0), "id", "Order id to delete").
		AddResponse(http.StatusBadRequest, "Invalid ID supplied", nil, nil).
		AddResponse(http.StatusNotFound, "Order not found", nil, nil).
		AddResponse(http.StatusNoContent, "Success", nil, nil).
		AddResponse(http.StatusInternalServerError, "Db operation failed", nil, nil).
		SetSummary("Delete a Order")

	o.GET("/:id", orderHandler.GetById).
		AddParamPath(int64(0), "id", "Order id to get").
		AddResponse(http.StatusBadRequest, "Invalid ID supplied", nil, nil).
		AddResponse(http.StatusNotFound, "Order not found", nil, nil).
		AddResponse(http.StatusOK, "Success", nil, nil).
		AddResponse(http.StatusInternalServerError, "Db operation failed", nil, nil).
		SetSummary("Get single Order")

	o.GET("", orderHandler.GetAll).
		AddParamQuery(int64(0), "page", "Page query param", false).
		AddParamQuery(int64(0), "limit", "Limit query param", false).
		AddParamQuery("", "sort", "Sort query param", false).
		AddResponse(http.StatusOK, "Success", nil, nil).
		AddResponse(http.StatusInternalServerError, "Db operation failed", nil, nil).
		SetSummary("Get All Order")

	se.Echo().Start(":7611")

	<-ctx.Done()
	if err := s.echo.Server.Shutdown(ctx); err != nil {
		panic(any("echo.server shutdown"))
	}

	return nil
}

func Connect() *gorm.DB {
	db, err := gorm.Open(postgres.Open("postgres://pg:admin@localhost:5432/ecom"),&gorm.Config{})
	if err != nil {
		panic(any("Could not connect to the database"))
	}

	db.AutoMigrate(&entities.Currency{}, &entities.Order{}, &entities.OrderItem{})

	return db
}
