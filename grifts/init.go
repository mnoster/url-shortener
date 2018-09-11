package grifts

import (
	"github.com/gobuffalo/buffalo"
	"github.com/mnoster/url_shortener/actions"
)

func init() {
	buffalo.Grifts(actions.App())
}
