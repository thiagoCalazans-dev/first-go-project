package main

import (
	"encoding/json"
	"net/http"

	db "github.com/thiagoCalazans-dev/first-go-project/database"
	"github.com/thiagoCalazans-dev/first-go-project/entity"
)

func main () {
	r := http.NewServeMux()
	r.HandleFunc("/", home)
	http.ListenAndServe(":3000", r)
}



func home (w http.ResponseWriter, r *http.Request) {
	
	user1 := entity.User{
	Name: "Thiago", 
	Email: "devt.calazans@gmail.com", 
	Status: true,}

	users := append(db.Users, user1)

json.NewEncoder(w).Encode(users)
}