package main

import (
	"io"
	"log"
	"net/http"

	"github.com/tkyatg/ditest-api/repository"
	"github.com/tkyatg/ditest-api/usecase"
	dicontainer "github.com/tkyatg/example-golang-dicontainer"
)

func main() {

	container := dicontainer.NewContainer()
	if err := container.Register(repository.NewRepository); err != nil {
		log.Fatal(err)
	}
	if err := container.Register(usecase.NewUsecase); err != nil {
		log.Fatal(err)
	}
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		if err := container.Invoke(func(
			uc usecase.Usecase,
		) error {
			io.WriteString(w, uc.Hello())
			return nil
		}); err != nil {
			log.Fatal(err)
		}
	})
	log.Fatal(http.ListenAndServe(":8080", nil))
}
