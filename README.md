# typed-form

## badges
[![GoDoc](https://godoc.org/github.com/thedanielforum/typed-form?status.svg)](https://godoc.org/github.com/thedanielforum/typed-form)

## How to use
```go
import form "github.com/thedanielforum/typed-form"

form := form.Parse(c)
s := page.NewGetPage(
    form.GetParamInt64("id"),
)
if form.Errors() != nil {
	err := form.Errors()[0]
	c.JSON(http.StatusBadRequest, b.errorMsg(c, err.Error()))
	return
}
```
