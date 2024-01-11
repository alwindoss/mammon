package main

import (
	"log"
	"net/http"
	"os"

	"github.com/alwindoss/mammon/cmd/mammon/views"
	"github.com/alwindoss/thea"
)

func main() {
	cfg := thea.Config{
		FS: views.FS,
	}
	vm, err := thea.New(&cfg)
	if err != nil {
		log.Printf("unable to initialize thea view manager")
		os.Exit(-1)
	}
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		vm.Render(w, "index.page.html", nil)
	})
	http.HandleFunc("/about", func(w http.ResponseWriter, r *http.Request) {
		vm.Render(w, "about.page.html", nil)
	})
	http.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		vm.Render(w, "login.page.html", nil)
	})
	http.HandleFunc("/admin/", func(w http.ResponseWriter, r *http.Request) {
		vm.Render(w, "home.page.html", nil)
	})
	http.HandleFunc("/admin/profile", func(w http.ResponseWriter, r *http.Request) {
		vm.Render(w, "profile.page.html", nil)
	})
	http.ListenAndServe(":8080", nil)
}
