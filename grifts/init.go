package grifts

import (
	"github.com/gobuffalo/buffalo"
	"github.com/hurtownia/actions"
)

func init() {
	buffalo.Grifts(actions.App())
}
