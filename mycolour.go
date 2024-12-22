package mycolour

import (
	"fmt"
)

const (
	escape  string = "\033["
	normal  string = "0"
	bold    string = "1"
	black   string = ";90m"
	red     string = ";91m"
	green   string = ";92m"
	yellow  string = ";93m"
	blue    string = ";94m"
	magenta string = ";95m"
	cyan    string = ";96m"
	white   string = ";97m"
	end     string = "0m"
)

type format string

const (
	Normal format = "0"
	Bold format = "1"
	Underline format = "4"
	Negative format = "7"
)

func Black(format format, s string) string {
	return fmt.Sprintf("%s%s%s%s%s%s", escape, format, black, s, escape, end)
}

func Red(format format, s string) string {
	return fmt.Sprintf("%s%s%s%s%s%s", escape, format, red, s, escape, end)
}

func Green(format format, s string) string {
	return fmt.Sprintf("%s%s%s%s%s%s", escape, format, green, s, escape, end)
}

func Yellow(format format, s string) string {
	return fmt.Sprintf("%s%s%s%s%s%s", escape, format, yellow, s, escape, end)
}

func Blue(format format, s string) string {
	return fmt.Sprintf("%s%s%s%s%s%s", escape, format, blue, s, escape, end)
}

func Magenta(format format, s string) string {
	return fmt.Sprintf("%s%s%s%s%s%s", escape, format, magenta, s, escape, end)
}

func Cyan(format format, s string) string {
	return fmt.Sprintf("%s%s%s%s%s%s", escape, format, cyan, s, escape, end)
}

func White(format format, s string) string {
	return fmt.Sprintf("%s%s%s%s%s%s", escape, format, white, s, escape, end)
}
