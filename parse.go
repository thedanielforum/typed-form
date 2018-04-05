package form

import (
	"strconv"

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

// GetString gets key in string format
func (f *Form) GetString(k string) string {
	return f.c.PostForm(k)
}

// GetInt gets key in int format
func (f *Form) GetInt(k string) int {
	i, err := strconv.Atoi(f.c.PostForm(k))
	if err != nil {
		f.errs = append(f.errs, err)
		if f.debug {
			logErr(k, err)
		}
	}
	return i
}

// GetInt64 gets key in int64 format
func (f *Form) GetInt64(k string) int64 {
	i, err := strconv.ParseInt(f.c.PostForm(k), 10, 64)
	if err != nil {
		f.errs = append(f.errs, err)
		if f.debug {
			logErr(k, err)
		}
	}
	return i
}
