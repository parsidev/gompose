package utils

import pl "github.com/gertd/go-pluralize"

var plc = pl.NewClient()

func Pluralize(s string) string {
	return plc.Plural(s)
}
