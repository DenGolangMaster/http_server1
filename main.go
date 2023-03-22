package main

import (
	"html/template"
	"net/http"
)

func main() {
	// file:///
	// ftp
	// http     192.168.2.12
	// https://example.com/path_url/awesome
	// ^         ^        ^
	// протокол хост      path URL
	// DNS domain name server

	// /path_url/awesome
	// public_dir -- путь где лежат файлы
	// /:template_file_name/:action_name/param/value

	temp, err := template.ParseGlob("templates/*")
	if err != nil {
		panic(err)
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		arr := []string{
			"одын",
			"два",
			"три",
		}
		_ = temp.ExecuteTemplate(w, "index.gohtml", map[string]any{
			"title": "Привет мир!!!!",
			"arr":   arr,
		})
	})
	http.HandleFunc("/articles", func(w http.ResponseWriter, r *http.Request) {
		_ = temp.ExecuteTemplate(w, "articles.gohtml", map[string]any{
			"title": "Статьи",
		})
	})
	http.HandleFunc("/afisha", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("AFISHA"))
	})

	_ = http.ListenAndServe(":3333", nil)
}
