// Package web exposes the embedded templates and static assets so the compiled
// binary is fully self-contained with no external file dependencies.
package web

import "embed"

//go:embed all:templates
var TemplatesFS embed.FS

//go:embed all:static
var StaticFS embed.FS
