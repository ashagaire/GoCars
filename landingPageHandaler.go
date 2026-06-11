package main

import (
	"html/template"
	"net/http"
	
)


type carsDataset struct {
	DummyDataset []TemplateData 
}
func landingPageHandaler(w http.ResponseWriter, req *http.Request) {
	// showData := []TemplateData {}
	var data carsDataset
	//parsing all template files 
	files := getTemplateFiles()
	if len(files) == 0 {
		http.Error(w, "No Template Files Found", http.StatusInternalServerError)
		return
	}

	tpl, err1 := template.ParseFiles(files...)
	if  err1!= nil {
		http.Error(w, err1.Error(), http.StatusInternalServerError)
		return
	}
	data.DummyDataset = dataset()
	err := tpl.ExecuteTemplate(w,"landingPage",data)
	if err != nil {
		tpl.ExecuteTemplate(w, err.Error(), data)
		return
	}
}