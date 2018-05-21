package form

import (
	"strconv"
	"strings"
)

// GetString gets key in string format
func (f *Form) GetString(k string) string {
	return f.c.PostForm(k)
}

// GetBool gets key in bool format
// 'true' and '1' will be considered to be true
// every other value will return false
func (f *Form) GetBool(k string) bool {
	v := f.c.PostForm(k)
	if strings.ToLower(v) == "true" {
		return true
	}
	if strings.ToLower(v) == "1" {
		return true
	}
	return false
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

// GetInt8 gets key in int8 format
func (f *Form) GetInt8(k string) int8 {
	i, err := strconv.ParseInt(f.c.PostForm(k), 10, 8)
	if err != nil {
		f.errs = append(f.errs, err)
		if f.debug {
			logErr(k, err)
		}
	}
	return int8(i)
}

// GetInt16 gets key in int16 format
func (f *Form) GetInt16(k string) int16 {
	i, err := strconv.ParseInt(f.c.PostForm(k), 10, 16)
	if err != nil {
		f.errs = append(f.errs, err)
		if f.debug {
			logErr(k, err)
		}
	}
	return int16(i)
}

// GetInt32 gets key in int32 format
func (f *Form) GetInt32(k string) int32 {
	i, err := strconv.ParseInt(f.c.PostForm(k), 10, 32)
	if err != nil {
		f.errs = append(f.errs, err)
		if f.debug {
			logErr(k, err)
		}
	}
	return int32(i)
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
