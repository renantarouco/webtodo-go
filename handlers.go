package main

import (
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

var templates = map[string]*template.Template{}

func init() {
	// O template de layout, mais externo.
	layoutPath := filepath.Join("templates", "layout.html")
	// Lista com os templates internos.
	templatesPath := map[string]string{
		"list":     filepath.Join("templates", "list.html"),
		"form":     filepath.Join("templates", "form.html"),
		"notfound": filepath.Join("templates", "notfound.html"),
	}
	// Compilar cada um dos templates préviamente e salvar a estrutura.
	for name, templatePath := range templatesPath {
		tmpl, err := template.ParseFiles(layoutPath, templatePath)
		if err != nil {
			log.Println(err.Error())
			os.Exit(1)
		}
		templates[name] = tmpl
	}
}

func RetrieveItemsHandler(w http.ResponseWriter, r *http.Request) {
	items, err := RetrieveItems()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	templates["list"].ExecuteTemplate(w, "layout", items)
}

func CreateItemHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		templates["form"].ExecuteTemplate(w, "layout", nil)
	case http.MethodPost:
		title := r.FormValue("title")
		deadlineStr := r.FormValue("deadline")
		_, err := CreateItem(title, deadlineStr)
		if err != nil {
			log.Println(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		http.Redirect(w, r, "/todos", http.StatusSeeOther)
	}
}

func UpdateItemHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr, ok := vars["id"]
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Redirect(w, r, "/todos/notfound", http.StatusSeeOther)
		return
	}
	item, err := RetrieveItem(id)
	if err != nil {
		http.Redirect(w, r, "/todos/notfound", http.StatusSeeOther)
		return
	}
	switch r.Method {
	case http.MethodGet:
		templates["form"].ExecuteTemplate(w, "layout", item)
	case http.MethodPost:
		deadlineStr := r.FormValue("deadline")
		deadline, err := time.Parse("2006-01-02", deadlineStr)
		if err != nil {
			log.Println(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		_, err = UpdateItem(id, Item{
			Title:    r.FormValue("title"),
			Deadline: deadline,
		})
		http.Redirect(w, r, "/todos", http.StatusSeeOther)
	}
}

func DeleteItemHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr, ok := vars["id"]
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Redirect(w, r, "/todos/notfound", http.StatusSeeOther)
		return
	}
	if err := DeleteItem(id); err != nil {
		log.Println(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/todos", http.StatusSeeOther)
}

func NotFoundHandler(w http.ResponseWriter, r *http.Request) {
	templates["notfound"].ExecuteTemplate(w, "layout", nil)
}
