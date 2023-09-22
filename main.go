package main

import (
    "fmt"
	"net/http"
    "github.com/gorilla/mux"

    "github.com/AnuragChaubey2/Simple_Blogging.git/driver"
    store "github.com/AnuragChaubey2/Simple_Blogging.git/store/posts"
    service "github.com/AnuragChaubey2/Simple_Blogging.git/services/posts"
   handler "github.com/AnuragChaubey2/Simple_Blogging.git/handler/posts"

    _"github.com/go-sql-driver/mysql"
)


func main() {

    db := driver.ConnectToSQL()

    st := store.New(db)

    svc := service.NewPostsService(st)
    
   list := handler.NewPostsHandler(svc)

    r := mux.NewRouter()

    r.HandleFunc("/posts", list.GetAllPosts).Methods("GET")
    r.HandleFunc("/posts/{id:[0-9]+}", list.GetPostByID).Methods("GET")
    r.HandleFunc("/posts", list.CreatePost).Methods("POST")
    r.HandleFunc("/posts/{id:[0-9]+}", list.UpdatePost).Methods("PUT")
    r.HandleFunc("/posts/{id:[0-9]+}", list.DeletePost).Methods("DELETE")

    http.Handle("/", r)

    fmt.Println("Server is running on :8080...")
    http.ListenAndServe(":8080", nil)
}

