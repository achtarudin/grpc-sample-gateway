package logging

import (
	"bytes"
	"time"

	"github.com/fatih/color"
)

type Format struct{}

func (writer *Format) Write(result []byte) (n int, err error) {
	color.NoColor = false

	var c *color.Color

	lower := bytes.ToLower(result)

	if bytes.Contains(lower, []byte("error")) || bytes.Contains(lower, []byte("failed")) {
		c = color.New(color.FgRed)
	} else {
		c = color.New(color.FgGreen)
	}
	return c.Print(time.Now().UTC().Format("02/01/2006 15:04:05") + " " + string(result))
}
