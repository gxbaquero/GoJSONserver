package main

import (
	"net/http"
	"fmt"
	"log"
	"encoding/json"
	"encoding/base64"
	"io/ioutil"
	"time"
	"strings"
	"strconv"	
)

func CheckAuth() Middleware {
	return func(f http.HandlerFunc) http.HandlerFunc{
		return func(w http.ResponseWriter, request *http.Request){			
			token :=request.Header["Authorization"]
			tokenS := token[0]
			decoder := json.NewDecoder(request.Body)
			var servicio Servicio
			err := decoder.Decode(&servicio)
			if err != nil{
				//fmt.Fprintf(w, "error: %v",err)
				log.Fatal(err)
				return
			}
			mail := servicio.Email
			fmt.Println(w,"check authentication")
			fmt.Println(token)
			fmt.Println(mail)

			fileName :="services/buffer.json"
			fileContent, ierr :=ioutil.ReadFile(fileName)
			if ierr!=nil{
				log.Fatal(ierr)
			}

			tokenS = strings.Replace(tokenS,"Bearer ","",-1)
			fileContentS :=string(fileContent)

			//validamos Token
			flag := strings.Contains(fileContentS, tokenS)
			fmt.Println(flag)			

			if flag{
				decoded, erro := base64.StdEncoding.DecodeString(tokenS)
				if erro!=nil{
					fmt.Println(erro)
					w.WriteHeader(http.StatusForbidden)
					w.Write([]byte(`{"error":"token no valido"}`))
				}else{
					tokenSArray :=strings.Split(string(decoded),"=")
					if mail==tokenSArray[1]{
						timestamp := strconv.FormatInt(time.Now().UTC().UnixNano(), 10)
						horaActual,_ := strconv.Atoi(timestamp)
						fmt.Println(horaActual)
						horaToken,_ :=  strconv.Atoi(tokenSArray[0])
						fmt.Println(horaToken)
						ttl := (horaActual - horaToken)/1000000000
						fmt.Println(ttl)
						if (ttl<3600){	
							//fmt.Printf("%+v", request)
							f(w, request)
						}else{
							w.WriteHeader(http.StatusForbidden)
							w.Write([]byte(`{"error":"token vencido"}`))
						}					
					}else{
						w.WriteHeader(http.StatusForbidden)
						w.Write([]byte(`{"error":"token no coincide con correo"}`))
					}
				}
			}else{
				w.WriteHeader(http.StatusForbidden)
				w.Write([]byte(`{"error":"token no valido"}`))
			}
		}
	}
}

func LoG() Middleware{
	return func(f http.HandlerFunc) http.HandlerFunc{
		return func(w http.ResponseWriter, request *http.Request){
			horaInicio :=time.Now()
			defer func(){
				log.Println(request.URL.Path, time.Since(horaInicio))
			}()
			f(w,request)
		}
	}
}