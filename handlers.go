package main

import (
    "encoding/json"
    "net/http"
    "github.com/gorilla/mux"
)

var users = make(map[string]User)

func CreateUser(w http.ResponseWriter, r *http.Request) {
    var user User
    _ = json.NewDecoder(r.Body).Decode(&user)
    users[user.Username] = user
    json.NewEncoder(w).Encode(user)
}

func GetUsers(w http.ResponseWriter, r *http.Request) {
    var userList []User
    for _, user := range users {
        userList = append(userList, user)
    }
    json.NewEncoder(w).Encode(userList)
}

func GetUser(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    user, exists := users[params["username"]]
    if !exists {
        http.Error(w, "User not found", http.StatusNotFound)
        return
    }
    json.NewEncoder(w).Encode(user)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    var updatedUser User
    _ = json.NewDecoder(r.Body).Decode(&updatedUser)
    user, exists := users[params["username"]]
    if !exists {
        http.Error(w, "User not found", http.StatusNotFound)
        return
    }
    user.Password = updatedUser.Password
    user.Active = updatedUser.Active
    users[params["username"]] = user
    json.NewEncoder(w).Encode(user)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    _, exists := users[params["username"]]
    if !exists {
        http.Error(w, "User not found", http.StatusNotFound)
        return
    }
    delete(users, params["username"])
    json.NewEncoder(w).Encode(users)
}
