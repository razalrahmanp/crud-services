package main

import (
    "github.com/gorilla/mux"
    "net/http"
)

func InitRouter() *mux.Router {
    router := mux.NewRouter()
    router.HandleFunc("/users", CreateUser).Methods("POST")
    router.HandleFunc("/users", GetUsers).Methods("GET")
    router.HandleFunc("/users/{username}", GetUser).Methods("GET")
    router.HandleFunc("/users/{username}", UpdateUser).Methods("PUT")
    router.HandleFunc("/users/{username}", DeleteUser).Methods("DELETE")
    return router
}
