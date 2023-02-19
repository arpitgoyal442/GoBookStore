package utils

import (
	"net/http"
	"io/ioutil"
	"encoding/json"
)

func ParseBody(r *http.Request,x interface{}){

	 body,err := ioutil.ReadAll(r.Body)
	 
	 if err!=nil{
		panic(err)
	 }

	 err =json.Unmarshal([]byte(body),x)
	 if err!=nil{
		return
	 }
}