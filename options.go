package ginjet

// Options for JetRender
type RenderOptions struct {
	TemplateDir string
	ContentType string
}

// Default options
func DefaultOptions() *RenderOptions {
	return &RenderOptions{
		TemplateDir: "./templates",
		ContentType: "text/html; charset=utf-8",
	}
}
