package main

import (
	//	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

var staticDir = "/home/seafooler/go-http-test/public"

func statusHandler(rw http.ResponseWriter, request *http.Request) {
	content := struct {
		Glustermnts map[string]bool
	}{
		Glustermnts: map[string]bool{
			"single": true,
			"double": false,
			"trible": true,
		},
	}	

	t := template.Must(template.ParseFiles(
		filepath.Join(staticDir, "template/header.html"),
		filepath.Join(staticDir, "template/footer.html"),
		filepath.Join(staticDir, "template/status.html")))
	t.ExecuteTemplate(rw, "main", &content)
}

func main() {
	fs := http.FileServer(http.Dir(staticDir))
	http.Handle("/public/",http.StripPrefix("/public", fs))
	http.HandleFunc("/status", statusHandler)
	err := http.ListenAndServe(":"+"8080", nil)
	if err != nil {
		log.Fatal("Error: ListenAndServe", 8080)
	}
	/*
		content := struct {
			glustermnts []string
		}{
			glustermnts: []string{
				"single",
				"double",
				"trible"},
		}

		//	glustermnts := []string{"single", "double", "trible"}
		for _, mnt := range content.glustermnts {
			fmt.Println(mnt)
		}
	*/
}
