package main

import (
	"example/hello/Projects/MongoGolang/controllers"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2"
)

func main() {

	r := httprouter.New()
	uc := controllers.NewUserController(getSession())
	r.GET("/user/:id", uc.GetUser)
	r.POST("/user", uc.CreateUser)
	r.DELETE("/user/:id", uc.DeleteUser)
	http.ListenAndServe(":8080", r)

}

func getSession() *mgo.Session {

	s, err := mgo.Dial("localhost:27017")
	if err != nil {
		log.Fatal("Failed to connect to MongoDB:", err)
	}
	return s
}