package web

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strings"

	"ascii/data"
	"ascii/output"
	"ascii/xerrors"

	"ascii/args"
	"ascii/graphics"
)

type PageData struct {
	Result   string
	Error    string
	Terminal string
	Text     string
	Banner   string
}

func RenderTemplate(w http.ResponseWriter, tmpl string, data interface{}) {
	t, err := template.ParseFiles("templates/" + tmpl)
	if err != nil {
		http.Error(w, "Template not found: "+err.Error(), http.StatusNotFound)
		return
	}

	err = t.Execute(w, data)
	if err != nil {
		http.Error(w, "Error rendering template: "+err.Error(), http.StatusInternalServerError)
	}
}

// renderError a helper function to render the custom error page
func renderError(w http.ResponseWriter, statusCode int, errorMessage string) {
	w.WriteHeader(statusCode)
	// s := http.Flusher
	RenderTemplate(
		w, "error.html", map[string]interface{}{
			"StatusCode": statusCode,
			"StatusText": http.StatusText(statusCode),
			"Message":    errorMessage,
		},
	)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		renderError(w, http.StatusNotFound, "Page Not Found")
		return
	}
	if r.Method != http.MethodGet {
		renderError(w, http.StatusMethodNotAllowed, "Method Not Allowed")
		return
	}

	pageData := PageData{}
	RenderTemplate(w, "index.html", &pageData)
}

// handler for processing ascii art
func asciiArtHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		renderError(w, http.StatusMethodNotAllowed, "Method Not Allowed")
		return
	}
	if r.URL.Path != "/ascii-art" {
		renderError(w, http.StatusNotFound, "Page Not Found")
		return
	}
	text := r.FormValue("text")
	banner := r.FormValue("banner")
	terminal := r.FormValue("terminal")

	isTerminal := terminal == "on"
	if banner == "" {
		renderError(w, http.StatusBadRequest, "Please provide a banner to be used.")
		return
	}

	validBanner := make(map[string]bool)
	validBanner["shadow"] = true
	validBanner["standard"] = true
	validBanner["thinkertoy"] = true

	// when an invalid banner is passed, return internal server error
	if !validBanner[banner] {
		renderError(w, http.StatusNotFound, "Sorry! Banner not supported")
		return
	}

	asciiArt := ""
	var err xerrors.AsciiError
	if isTerminal {
		escapedText := args.Escape(text)
		escapedText, _ = output.HandleSpecial(escapedText)
		draws := data.DrawInfo{
			Text:  escapedText,
			Style: banner,
		}
		asciiArt, err = graphics.Draw(draws)
	} else {
		draws := data.DrawInfo{
			Text:  text,
			Style: banner,
		}
		asciiArt, err = graphics.Draw(draws)
	}

	if err != nil {
		switch err.Type() {
		case xerrors.TypeInvalidAscii:
			renderError(w, http.StatusBadRequest, err.Error())
		case xerrors.TypeInvalidBanner:
			renderError(w, http.StatusNotFound, err.Error())
		case xerrors.TypeInvalidGraphics:
			renderError(w, http.StatusNotFound, err.Error())
		default:
			renderError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}

	RenderTemplate(w, "index.html", &PageData{Result: asciiArt, Terminal: terminal, Text: text, Banner: banner})
}

func Server() {
	port := ":9090"
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/ascii-art", asciiArtHandler)

	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static", fs))

	fmt.Printf("server started on port %s \n", strings.TrimPrefix(port, ":"))
	log.Fatal(http.ListenAndServe(port, nil))
}
