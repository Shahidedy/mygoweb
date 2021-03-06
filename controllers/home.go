package controllers

import (
	"net/http"
	"viewmodels"
	"text/template"
)

type homeController struct {
	template *template.Template
}

func (this *homeController) get(w http.ResponseWriter, req *http.Request){
	vm := viewmodels.GetHome()

	w.Header().Add("Content-type", "text/html")
	this.template.Execute(w, vm)
}
