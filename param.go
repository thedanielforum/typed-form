package form

import "strconv"

// GetParamString returns param from key as string
func (f *Form) GetParamString(k string) string {
	return f.c.Param(k)
}

// GetParamInt64 returns param from key as int64
func (f *Form) GetParamInt64(k string) int64 {
	i, err := strconv.ParseInt(f.c.Param(k), 10, 64)
	if err != nil {
		f.errs = append(f.errs, err)
		if f.debug {
			logErr(k, err)
		}
	}
	return i
}
