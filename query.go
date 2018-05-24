package form

import (
	"strconv"
	"strings"
)

// GetQueryString returns query param as string
func (f *Form) GetQueryString(k string) string {
	return f.c.Query(k)
}

// GetQueryBool gets key in bool format
// 'true' and '1' will be considered to be true
// every other value will return false
func (f *Form) GetQueryBool(k string) bool {
	v := f.c.Query(k)
	if strings.ToLower(v) == "true" {
		return true
	}
	if strings.ToLower(v) == "1" {
		return true
	}
	return false
}

// GetQueryInt returns query param as string
func (f *Form) GetQueryInt(k string) int {
	i, err := strconv.Atoi(f.c.Query(k))
	if err != nil {
		f.errs = append(f.errs, err)
		if f.debug {
			logErr(k, err)
		}
	}
	return i
}

// GetQueryInt8 returns param from key as int8
func (f *Form) GetQueryInt8(k string) int8 {
	i, err := strconv.ParseInt(f.c.Query(k), 10, 8)
	if err != nil {
		f.errs = append(f.errs, err)
		if f.debug {
			logErr(k, err)
		}
	}
	return int8(i)
}

// GetQueryInt16 returns param from key as int16
func (f *Form) GetQueryInt16(k string) int16 {
	i, err := strconv.ParseInt(f.c.Query(k), 10, 16)
	if err != nil {
		f.errs = append(f.errs, err)
		if f.debug {
			logErr(k, err)
		}
	}
	return int16(i)
}

// GetQueryInt32 returns param from key as int32
func (f *Form) GetQueryInt32(k string) int32 {
	i, err := strconv.ParseInt(f.c.Query(k), 10, 32)
	if err != nil {
		f.errs = append(f.errs, err)
		if f.debug {
			logErr(k, err)
		}
	}
	return int32(i)
}

// GetQueryInt64 returns param from key as int64
func (f *Form) GetQueryInt64(k string) int64 {
	i, err := strconv.ParseInt(f.c.Query(k), 10, 64)
	if err != nil {
		f.errs = append(f.errs, err)
		if f.debug {
			logErr(k, err)
		}
	}
	return i
}
