package form

import "strconv"

// GetParamString returns param from key as string
func (f *Form) GetParamString(k string) string {
	return f.c.Param(k)
}

// GetParamInt returns param from key as int
func (f *Form) GetParamInt(k string) int {
	i, err := strconv.Atoi(f.c.PostForm(k))
	if err != nil {
		f.errs = append(f.errs, err)
		if f.debug {
			logErr(k, err)
		}
	}
	return i
}

// GetParamInt8 returns param from key as int8
func (f *Form) GetParamInt8(k string) int8 {
	i, err := strconv.ParseInt(f.c.PostForm(k), 10, 8)
	if err != nil {
		f.errs = append(f.errs, err)
		if f.debug {
			logErr(k, err)
		}
	}
	return int8(i)
}

// GetParamInt16 returns param from key as int16
func (f *Form) GetParamInt16(k string) int16 {
	i, err := strconv.ParseInt(f.c.Param(k), 10, 16)
	if err != nil {
		f.errs = append(f.errs, err)
		if f.debug {
			logErr(k, err)
		}
	}
	return int16(i)
}

// GetParamInt32 returns param from key as int32
func (f *Form) GetParamInt32(k string) int32 {
	i, err := strconv.ParseInt(f.c.Param(k), 10, 32)
	if err != nil {
		f.errs = append(f.errs, err)
		if f.debug {
			logErr(k, err)
		}
	}
	return int32(i)
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
