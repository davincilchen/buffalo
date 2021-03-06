package refresh

import (
	"github.com/gobuffalo/genny"
	"github.com/gobuffalo/packr/v2"
	"github.com/gobuffalo/plush"
	"github.com/gobuffalo/plushgen"
	"github.com/pkg/errors"
)

// New generator to generate refresh templates
func New(opts *Options) (*genny.Generator, error) {
	g := genny.New()
	if err := opts.Validate(); err != nil {
		return g, errors.WithStack(err)
	}
	g.Box(packr.New("buffalo:genny:refresh", "../refresh/templates"))

	ctx := plush.NewContext()
	ctx.Set("app", opts.App)
	g.Transformer(plushgen.Transformer(ctx))
	g.Transformer(genny.Dot())
	return g, nil
}
