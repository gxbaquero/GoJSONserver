package main

import (
	"net/http"
	"encoding/json"
)

type Middleware func(http.HandlerFunc) http.HandlerFunc

type MetaData interface{}

type User struct{
	Name string `json:"name"`
	Email string `json:"email"`
	Day string `json:"day"`
	Hour string `json:"hour"`
}

type Servicio struct{
	Email string `json:"email"`
	//Method string `json:"method"`
	Key string `json:"key"`
	Value string `json:"value"`
}

func (u *User) toJson() ([]byte, error){
	return json.Marshal(u)
}


type Server struct{
	port string
	router *Router
}

func NewServer(puerto string) *Server{
	return &Server{
		port: puerto,
		router: NewRouter(),
	}
}

func (s *Server) Handle(metodo string, path string, handler http.HandlerFunc) {
	_, exist := s.router.rules[path]
	if !exist {
		s.router.rules[path] = make(map[string]http.HandlerFunc)
	}
	s.router.rules[path][metodo] = handler
}


func (s *Server) AddMiddleware(f http.HandlerFunc, middleware ...Middleware) http.HandlerFunc{
	for _, capas := range middleware{
		f = capas(f)
	}
	return f
}

func (s *Server) Listen() error{
	http.Handle("/", s.router)
	error := http.ListenAndServe(s.port,nil)
	if error!= nil{
		return error
	}else{
		return nil
	}
}