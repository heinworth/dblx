/*
	api contains infrastructure for HTTP server, as well as handler functions
	for HTTP routes.
*/

package api

import (
	"github.com/julienschmidt/httprouter"

	"../database"
	"../user"

	"fmt"
	"net/http"
	"strconv"
	"math/rand"
	"encoding/json"
)

var users []user.User

func StartServer() error {
	
	fmt.Println("starting server...")

	router := httprouter.New()
	router.POST("/user/:name", addUser)
	router.GET("/user/:id", getUser)
	router.GET("/user", getAllUsers)

    return http.ListenAndServe(":8080", router)
}

func addUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	newUser := user.User{
		Name: ps.ByName("name"),
		ID: rand.Int(),
	}
	users = append(users, newUser)

	setDefaultHeaders(&w)
	w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(newUser)
}

func getUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	setDefaultHeaders(&w)

	id, err := strconv.Atoi(ps.ByName("id"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var res *user.User
	for _, v := range users {
		if v.ID == id {
			res = &v
		}
	}

	if res == nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
}

func getAllUsers(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	setDefaultHeaders(&w)

	res, err := database.GetAllUsers()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return 
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
}

func setDefaultHeaders(w *http.ResponseWriter) {
	// unsafe - for testing only!
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Content-Type", "application/json")
}


