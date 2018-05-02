# typed-form

## badges
[![GoDoc](https://godoc.org/github.com/thedanielforum/typed-form?status.svg)](https://godoc.org/github.com/thedanielforum/typed-form) [![CircleCI](https://circleci.com/gh/thedanielforum/typed-form.svg?style=svg)](https://circleci.com/gh/thedanielforum/typed-form)

## How to use
```go
import form "github.com/thedanielforum/typed-form"

// c is the gin context
form := form.Parse(c)
s := service.GetData(
	form.GetParamInt64("id"),
	form.GetString("name"),
)
if form.Errors() != nil {
	err := form.Errors()[0]
	c.JSON(http.StatusBadRequest, b.errorMsg(c, err.Error()))
	return
}
```

## Avalible Get's
```go
form := form.Parse(gin.Context)

// Get post form data
form.GetString("key")
form.GetInt("key")
form.GetInt8("key")
form.GetInt16("key")
form.GetInt32("key")
form.GetInt64("key")

// Get url param
form.GetParamString("key")
form.GetParamInt("key")
form.GetParamInt8("key")
form.GetParamInt16("key")
form.GetParamInt32("key")
form.GetParamInt64("key")
```
