package form

import "fmt"

func logErr(key string, err error) {
	fmt.Printf(
		"error encountered while getting [%s] from the form: %s",
		key,
		err.Error(),
	)
}
