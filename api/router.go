package api

import (
	"errors"
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
	"reflect"
)

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	if err := cv.validator.Struct(i); err != nil {
		var ve validator.ValidationErrors
		if errors.As(err, &ve) {
			out := make([]apiError, len(ve))
			for idx, fe := range ve {
				out[idx] = apiError{fe.Field(), msgForTag(fe)}
			}
			return echo.NewHTTPError(http.StatusBadRequest, &resBodyError{out})
		}
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return nil
}

type httpMsg struct {
	Message string `json:"message"`
}

type apiError struct {
	Param   string `json:"param"`
	Message string `json:"message"`
}

type resBodyError struct {
	Errors []apiError `json:"errors"`
}

func msgForTag(fe validator.FieldError) string {
	switch fe.Tag() {
	case "required":
		return "This field is required"
	case "email":
		return "Invalid email"
	case "containsany":
		fmt.Println(fe.Field())
		switch fe.Field() {
		case "Password":
			return "The password must contain at least one number"
		}
	case "max":
		switch fe.Type().Kind() {
		case reflect.String:
			return "Maximum length is exceeded"
		}
	}
	return fe.Error()
}

func AddRoutes(e *echo.Echo) error {
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Validator = &CustomValidator{validator: validator.New()}

	addStatusRoutes(e)
	addUserRoutes(e)
	addCountiesRoutes(e)
	addCountriesRoutes(e)

	return nil
}

func addUserRoutes(e *echo.Echo) {
	usersRouter := e.Group("/users")
	usersRouter.GET("/", getUsers)
	usersRouter.POST("/", createUser)
}

func addCountiesRoutes(e *echo.Echo) {
	countiesRouter := e.Group("/counties")
	countiesRouter.GET("/", getCounties)
}

func addCountriesRoutes(e *echo.Echo) {
	countriesRouter := e.Group("/countries")
	countriesRouter.GET("/", getCountries)
}

func addStatusRoutes(e *echo.Echo) {
	e.GET("/status", getStatus)
}
