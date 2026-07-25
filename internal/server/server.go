package server

import (
	"html/template"
	"io/fs"
	"log/slog"
	"net/http"
	"time"

	"portfolio-website/web"
)

// Server holds the parsed templates and portfolio content.
type Server struct {
	tmpl      *template.Template
	portfolio Portfolio
	logger    *slog.Logger
}

// New builds a Server, parsing templates up-front so bad templates fail fast.
func New(logger *slog.Logger) (*Server, error) {
	tmpl, err := template.ParseFS(web.TemplatesFS, "templates/*.html")
	if err != nil {
		return nil, err
	}

	p := defaultPortfolio()
	p.Year = time.Now().Year()

	return &Server{
		tmpl:      tmpl,
		portfolio: p,
		logger:    logger,
	}, nil
}

// Routes wires up all handlers and returns the root http.Handler.
func (s *Server) Routes() (http.Handler, error) {
	mux := http.NewServeMux()

	staticSub, err := fs.Sub(web.StaticFS, "static")
	if err != nil {
		return nil, err
	}
	fileServer := http.FileServer(http.FS(staticSub))
	mux.Handle("/static/", http.StripPrefix("/static/", cacheControl(fileServer)))

	mux.HandleFunc("/healthz", s.handleHealth)
	mux.HandleFunc("/", s.handleIndex)

	return s.recoverPanic(s.logRequests(mux)), nil
}

func (s *Server) handleIndex(w http.ResponseWriter, r *http.Request) {
	// The "/" pattern is a catch-all in Go 1.21's mux, so filter to exact root.
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	if r.Method != http.MethodGet {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if err := s.tmpl.ExecuteTemplate(w, "index.html", s.portfolio); err != nil {
		s.logger.Error("template render failed", "error", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

func (s *Server) handleHealth(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte(`{"status":"ok"}`))
}
