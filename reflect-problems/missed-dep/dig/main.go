package main

import (
	"fmt"
	"log"
	"net/http"

	"go.uber.org/dig"
)

type DB struct {
	data string
}

func NewDB() *DB {
	// POTENTIAL ERROR
	return nil
}

func (db *DB) Get() string {
	return db.data
}

type UserService struct {
	db *DB
}

func NewUserService(db *DB) *UserService {
	return &UserService{db: db}
}

func (s *UserService) Handler(w http.ResponseWriter, r *http.Request) {
	str := s.db.Get()

	_, err := w.Write([]byte(str))
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	c := dig.New()

	if err := c.Provide(NewDB); err != nil {
		log.Fatal(err)
	}

	if err := c.Provide(NewUserService); err != nil {
		log.Fatal(err)
	}

	err := c.Invoke(func(s *UserService) {
		http.HandleFunc("/users", s.Handler)
		fmt.Println("server started on :8090")
		log.Fatal(http.ListenAndServe(":8090", nil))
	})
	if err != nil {
		panic(err)
	}
}
