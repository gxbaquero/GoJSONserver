package main

import(
	"fmt"
    "io/ioutil"
    "strings"
)

//ttl del token una hora

func main(){
		servidor := NewServer(":9990")
		servidor.Handle("GET","/", HandleRoot)
		servidor.Handle("POST","/create", PostRequest)
		servidor.Handle("POST","/auth", UserPostRequest)
	 	files, err := ioutil.ReadDir("./services/")
    	if err != nil {
        	fmt.Printf("%v+",err)
	    }

	    for _, f := range files {
	    	servicio :=f.Name()
	    	servicio = strings.Replace(servicio,".json","",-1)
	    	if servicio=="buffer"{
	    		continue
	    	}
	    	servicio ="/api/"+servicio
	    	servidor.Handle("POST",servicio, servidor.AddMiddleware(HandleApi, CheckAuth()))//, LoG()))
	            //fmt.Println(f.Name())
	    }
		//servidor.Handle("POST","/api/", servidor.AddMiddleware(HandleApi, CheckAuth()))//, LoG()))
		servidor.Listen()
}