package main

import (
	"html/template"
	"log"
	"net/http"
	"raspberry-pi/core"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("template/html/index.html")
	if err != nil {
		log.Println(err)
	}
	systemInfo := core.SystemInfo{}
	systemInfo.InitSystemInfo()

	_ = t.Execute(w, &systemInfo)
}
func main() {

	http.Handle("/css/", http.FileServer(http.Dir("template")))
	http.Handle("/js/", http.FileServer(http.Dir("template")))

	log.Println("Raspberry Pi Service Startup at http://localhost:9999")
	http.HandleFunc("/", indexHandler)
	err := http.ListenAndServe(":9999", nil)
	if err != nil {
		log.Println("Raspberry Pi Service Startup Failure")
	}
}
