package main

import (
	"net/http"
)

type Router struct{
	rules map[string]map[string]http.HandlerFunc
}

func NewRouter() *Router{
	return &Router{
		rules: make(map[string]map[string]http.HandlerFunc),
	}
}

func (r *Router) FindHandler(path string, metodo string) (http.HandlerFunc, bool, bool) {
	_, exist := r.rules[path]
	handler,methodExist :=r.rules[path][metodo]
	return handler, methodExist, exist
}

//implementamos lo necesario para crear el servidor
func (r *Router) ServeHTTP(w http.ResponseWriter, request *http.Request){
	handler, methodExist, exist := r.FindHandler(request.URL.Path, request.Method)
	if !exist{
		w.WriteHeader(http.StatusNotFound)
		return  //sale de la funcion si no existe
	}

	if !methodExist{
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	handler(w,request)
	
}