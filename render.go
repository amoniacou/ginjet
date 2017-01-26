package ginjet

import (
	"github.com/CloudyKit/jet"
	"github.com/gin-gonic/gin/render"
	"net/http"
)

// JetRender is a custom Gin template renderer using Jet
type JetRender struct {
	Options  *RenderOptions
	Template *jet.Template
	Data     interface{}
}

// New creates a new JetRender instance with custom Options.
func New(options *RenderOptions) *JetRender {
	return &JetRender{
		Options: options,
	}
}

// Default creates a JetRender instance with default options.
func Default() *JetRender {
	return New(DefaultOptions())
}

func (r JetRender) Instance(name string, data interface{}) render.Render {
	set := jet.NewHTMLSet(r.Options.TemplateDir)
	t, err := set.GetTemplate(name)

	if err != nil {
		panic(err)
	}

	return JetRender{
		Data:     data,
		Options:  r.Options,
		Template: t,
	}
}

func (r JetRender) Render(w http.ResponseWriter) error {
	// Unless already set, write the Content-Type header.
	header := w.Header()
	if val := header["Content-Type"]; len(val) == 0 {
		header["Content-Type"] = []string{r.Options.ContentType}
	}
	if err := r.Template.Execute(w, nil, r.Data); err != nil {
		return err
	}
	return nil
}
