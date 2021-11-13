package controller

type ResponseData map[string]interface{}

type Context interface {
	JSON(code int, i interface{}) error
	Bind(i interface{}) error
	Param(name string) string
}
