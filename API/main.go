package main

import (
	"API/rotas"
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("teste")
	r := rotas.ConfigurarRotas()

	log.Println("servidor iniciado em: http://localhost:1000")
	log.Fatal(http.ListenAndServe(":1000", r))
}

//add para teste