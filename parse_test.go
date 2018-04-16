package form

import (
	"errors"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestEnableDebugMode(t *testing.T) {
	form := Parse(nil)
	assert.Equal(t, false, form.debug)
	form.EnableDebugMode()
	assert.Equal(t, true, form.debug)
}

func TestParse(t *testing.T) {
	c := &gin.Context{}
	form := Parse(c)
	assert.Equal(t, false, form.debug)
	assert.Equal(t, c, form.c)
	assert.Equal(t, 0, len(form.errs))
}

func TestErrors(t *testing.T) {
	form := Parse(nil)
	err := errors.New("test")
	form.errs = append(form.errs, err)
	assert.Equal(t, err, form.Errors()[0])
}
