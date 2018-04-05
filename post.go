package form

import "strconv"

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
