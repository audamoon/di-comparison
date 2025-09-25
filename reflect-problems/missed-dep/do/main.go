package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/samber/do"
)

type DB struct {
	data string
}

func NewDB(i *do.Injector) (*DB, error) {
	return nil, errors.New("db init failed")
}

func (db *DB) Get() string {
	return db.data
}

type UserService struct {
	db *DB
}

func NewUserService(i *do.Injector) (*UserService, error) {
	db := do.MustInvoke[*DB](i)

	return &UserService{db: db}, nil
}

func (s *UserService) Handler(w http.ResponseWriter, r *http.Request) {
	str := s.db.Get()

	_, err := w.Write([]byte(str))
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	i := do.New()

	do.Provide(i, NewDB)
	do.Provide(i, NewUserService)

	// Ошибка при запуске — сервер даже не поднимется
	_, err := do.Invoke[*UserService](i)
	if err != nil {
		log.Fatalf("failed to init: %v", err)
	}

	// Если бы ошибки не было:
	s := do.MustInvoke[*UserService](i)
	http.HandleFunc("/users", s.Handler)
	fmt.Println("server started on :8081")
	log.Fatal(http.ListenAndServe(":8081", nil))
}
