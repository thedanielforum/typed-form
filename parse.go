package form

import (
	"github.com/gin-gonic/gin"
)

// Form instance
type Form struct {
	c     *gin.Context
	errs  []error
	debug bool
}

// EnableDebugMode enables debug mode
// is false by default
func (f *Form) EnableDebugMode() {
	f.debug = true
}

// Parse new request from gin.Context
func Parse(c *gin.Context) *Form {
	return &Form{c: c}
}

// Errors returns all errors if any from the form
func (f *Form) Errors() []error {
	return f.errs
}
