/*
	Contains functions for accessing database, as well as a mocking framework
	for future tests.
	Prior to usage, requires package variable db to be set.
*/

package database

import (
	"../user"
	"math/rand"
)

type Database interface {
	getAllUsers() ([]user.User, error)
}

var db Database

type DBMock struct{}
type DBImplementation struct{}

func SetDatabase(DB Database) {
	db = DB
}

func GetAllUsers() ([]user.User, error) {
	return db.getAllUsers()
}

func (DBMock) getAllUsers() ([]user.User, error) {
	return []user.User{
		user.User{
			Name: "Alex Whitehorn",
			ID: rand.Int(),
		},
		user.User{
			Name: "User McUserface",
			ID: rand.Int(),
		},
		user.User{
			Name: "John Doe",
			ID: rand.Int(),
		},
	}, nil
}

// TODO: implement DB
func (DBImplementation) getAllUsers() ([]user.User, error) {
	return nil, nil
}

