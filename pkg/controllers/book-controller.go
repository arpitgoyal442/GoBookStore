package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/arpit/go-bookstore-mysql/pkg/models"
	"github.com/arpit/go-bookstore-mysql/pkg/utils"
	"github.com/gorilla/mux"
)

var NewBook models.Book

func GetBook(w http.ResponseWriter, r *http.Request){

	newBooks:= models.GetAllBooks()
	res,_ :=json.Marshal(newBooks)
	w.Header().Set("Content-Type","pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


func GetBookById(w http.ResponseWriter,r *http.Request){

	vars:=mux.Vars(r)
	bookId :=vars["Id"]
	ID,err:=strconv.ParseInt(bookId,0,0)

	if err!=nil{
		fmt.Println("Error While Parsing ")
	}

	book,_:=models.GetBookById(ID)
	res,_ :=json.Marshal(book)
	w.Header().Set("Content-Type","pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreateBook(w http.ResponseWriter,r *http.Request){

	book :=&models.Book{}
    utils.ParseBody(r,book)
	b:=book.CreateBook()
	res,_ :=json.Marshal(b)
	w.Header().Set("Content-Type","pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)


}

func DeleteBook(w http.ResponseWriter,r *http.Request){
     
	vars:=mux.Vars(r)
	bookId :=vars["Id"]
	ID,err:=strconv.ParseInt(bookId,0,0)

	if err!=nil{
		fmt.Println("Error While Parsing ")
	}

	book:=models.DeleteBookById(ID)
	res,_ :=json.Marshal(book)
	w.Header().Set("Content-Type","pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)


}

func UpdateBook(w http.ResponseWriter,r *http.Request){

	var updateBook= &models.Book{}
	utils.ParseBody(r,updateBook)
	vars :=mux.Vars(r)
	bookId:=vars["Id"]
	Id,err:=strconv.ParseInt(bookId,0,0)
	if err!=nil{
		fmt.Println("Error While Parsing ")
	}

	bookDetails,db:=models.GetBookById(Id)

	bookDetails.Name=updateBook.Name
	bookDetails.Author=updateBook.Author
	bookDetails.Publication=updateBook.Publication
	db.Save(&bookDetails)
	res,_ :=json.Marshal(bookDetails)
	w.Header().Set("Content-Type","pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
	


}


