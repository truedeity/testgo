package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type UserController struct {
	session *mgo.Session
}

type User struct {
	ID        bson.ObjectId
	FirstName string
	LastName  string
}

func NewUserController(s *mgo.Session) *UserController {
	return &UserController{s}
}

func main() {
	http.HandleFunc("/", hello)
	http.HandleFunc("/userlist", userlist)
	http.HandleFunc("/saveinfo", saveinfo)
	http.ListenAndServe(":1233", nil)
}

func hello(writer http.ResponseWriter, request *http.Request) {
	dat, _ := ioutil.ReadFile("/home/dan/go/src/github.com/truedeity/testgo2/test.html")

	writer.Write(dat)
}

func userlist(writer http.ResponseWriter, request *http.Request) {
	//uc := NewUserController(getSession())
	//query := uc.session.DB("sampledb1").C("users").Find()

}

func saveinfo(writer http.ResponseWriter, request *http.Request) {

	firstname := request.PostFormValue("firstname")
	lastname := request.PostFormValue("lastname")

	fmt.Println(firstname)
	fmt.Println(lastname)

	u := User{}
	u.ID = bson.NewObjectId()
	u.FirstName = firstname
	u.LastName = lastname

	uc := NewUserController(getSession())

	uc.session.DB("sampledb1").C("users").Insert(u)

	uj, _ := json.Marshal(u)

	writer.WriteHeader(201)
	writer.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(writer, "%s", uj)
	//	fmt.Fprint(writer, "%s", uj)

}

func getSession() *mgo.Session {
	s, err := mgo.Dial("mongodb://localhost")

	if err != nil {
		panic(err)
	}

	return s
}
