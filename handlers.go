package main

import (
	"net/http"
	"fmt"
	"encoding/json"
	"encoding/base64"
	"time"
	"strconv"
	"log"
	"io/ioutil"
	"strings"
)

func HandleRoot(w http.ResponseWriter, request *http.Request){
	fmt.Fprintf(w,"Servidor Online !")
	return
}

func HandleApi(w http.ResponseWriter, request *http.Request){ // , request *http.Request
	fileName:=request.RequestURI
	fileName=strings.Replace(fileName,"/api/","services/",-1)

	fileName=fileName+".json"
	fmt.Println(fileName)
	fileContent, ierr :=ioutil.ReadFile(fileName)
	if ierr!=nil{
		log.Fatal(ierr)
	}	
	w.Header().Set("Content-Type","application/json")
	w.Write([]byte(fileContent))
		//`{"tmp":"`+fileName+`"}`))

}

func PostRequest(w http.ResponseWriter, request *http.Request){
	decoder := json.NewDecoder(request.Body)
	var metadata MetaData
	err := decoder.Decode(&metadata)
	if err != nil{
		fmt.Fprintf(w, "error: %v",err)
		return
	}

	fmt.Fprintf(w, "PayLoad: %v\n", metadata)
}

func UserPostRequest(w http.ResponseWriter, request *http.Request){
	decoder := json.NewDecoder(request.Body)
	var user User
	err := decoder.Decode(&user)
	if err != nil{
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	
	timestamp := strconv.FormatInt(time.Now().UTC().UnixNano(), 10)
	toke:=timestamp+"="+user.Email
	toke=base64.StdEncoding.EncodeToString([]byte(toke))
	fmt.Println(toke)

	fileName :="services/buffer.json"
	//Aca empieza la lectura de un buffer basico
	fileContent, ierr :=ioutil.ReadFile(fileName)
	if ierr!=nil{
		log.Fatal(ierr)
	}

	fileContentS :=string(fileContent)

	var stringLine string =fileContentS+"\n"+toke+";"+user.Email+";"+timestamp+"\n"

	b:=[]byte(stringLine)
	err = ioutil.WriteFile(fileName,b,0755)
	if err!=nil{
		log.Fatal(err)
	}
	//Ya almacenó el token en el buffer

	w.Header().Set("Content-Type","application/json")
	w.Write([]byte(`{"token":"`+toke+`"}`))
}