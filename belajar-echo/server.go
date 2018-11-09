package main

import (
	"fmt"
	"html/template"
	"io"
	"net/http"

	"github.com/labstack/echo"
	"gopkg.in/go-playground/validator.v9"
)

// map untuk menyimpan pesan
type M map[string]interface{}

// struct user
type User struct {
	Name  string `json:"name" form:"name" query:"name" validate:"required"`
	Email string `json:"email" form:"email" query:"email" validate:"required,email"`
	Age   int    `json:"age" form:"age" query:"age" validate:"gte=0,lte=35"`
}

// struct custom validator
type CustomValidator struct {
	validator *validator.Validate
}

// method validate dari struct CustomValidator
func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}

// struct untuk render template
type Renderer struct {
	template *template.Template
	debug    bool
	location string
}

// function NewRenderer untuk mempermudah inisialisasi object render
func NewRenderer(location string, debug bool) *Renderer {
	tpl := new(Renderer)
	tpl.location = location
	tpl.debug = debug

	tpl.ReloadTemplates()

	return tpl
}

// method ReloadTemplates bertugas untuk parsing template
func (t *Renderer) ReloadTemplates() {
	t.template = template.Must(template.ParseGlob(t.location))
}

// method Render bertugas untuk render template setelah diparsing
func (t *Renderer) Render(
	w io.Writer,
	name string,
	data interface{},
	c echo.Context,
) error {
	if t.debug {
		t.ReloadTemplates()
	}

	return t.template.ExecuteTemplate(w, name, data)
}

// middleware
func middlewareOne(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		fmt.Println("From middleware one")
		return next(c)
	}
}

func middlewareTwo(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		fmt.Println("From middleware two")
		return next(c)
	}
}
func main() {
	e := echo.New()
	// pemanggilan function middleware
	e.Use(middlewareOne)
	e.Use(middlewareTwo)
	// pemanggilan method validator
	e.Validator = &CustomValidator{validator: validator.New()}
	// pemanggilan function render. ./template/*html merupakan path location template yang akan di
	e.Renderer = NewRenderer("./template/*html", true)
	e.GET("/", func(c echo.Context) error {
		data := "Hello, dudes"
		return c.String(http.StatusOK, data)
	})
	e.GET("/template", func(c echo.Context) error {
		data := M{"message": "Hallo, huda"}
		// index.html template yang akan di render
		return c.Render(http.StatusOK, "index.html", data)
	})
	e.POST("/user", func(c echo.Context) (err error) {
		// pointer user
		u := new(User)
		if err := c.Bind(u); err != nil {
			return err
		}
		if err := c.Validate(u); err != nil {
			return err
		}
		return c.JSON(http.StatusOK, u)
	})
	e.HTTPErrorHandler = func(err error, c echo.Context) {
		report, ok := err.(*echo.HTTPError)
		if !ok {
			report = echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
		if castedObject, ok := err.(validator.ValidationErrors); ok {
			for _, err := range castedObject {
				switch err.Tag() {
				case "required":
					report.Message = fmt.Sprintf("%s is required", err.Field())
				case "email":
					report.Message = fmt.Sprintf("%s is not valid email", err.Field())
				case "gte":
					report.Message = fmt.Sprintf("%s value must be greater than %s", err.Field(), err.Param())
				case "lte":
					report.Message = fmt.Sprintf("%s value must be lower than %s", err.Field(), err.Param())
				}
			}
		}
		c.Logger().Error(report)
		c.JSON(report.Code, report)
	}
	e.Start(":9000")
}
