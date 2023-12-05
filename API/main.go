package main

import (
	"API/modelo"
	"API/rotas"
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("teste")

	novoUsuario := modelo.CriarNovoUsuario(1, "usuario", "usu@gmail.com")
	fmt.Println("Novo usuário criado:", novoUsuario)

	usuarios := []modelo.Usuario{}
	usuarios = modelo.AdicionarUsuario(usuarios, novoUsuario)
	fmt.Println("Lista de usuários:", usuarios)

	r := rotas.ConfigurarRotas()

	log.Println("Servidor iniciado em: http://localhost:1000")
	log.Fatal(http.ListenAndServe(":1000", r))
}
