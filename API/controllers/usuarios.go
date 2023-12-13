package controllers

import (
	"API/modelo"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"sync"

	"github.com/gorilla/mux"
)

var (
	mutex    sync.Mutex
	usuarios []modelo.Usuario
	nextID   = 1
)

// Implemente as funções de manipulação de usuários (Listar, Criar, Atualizar, Deletar, Obter etc.)
// Você pode pegar os códigos das funções dos arquivos anteriores e organizá-los aqui.

func ListarUsuarios(w http.ResponseWriter, r *http.Request) {
	// Implementação para listar usuários
}

func ObterUsuario(w http.ResponseWriter, r *http.Request) {
	// Implementação para obter um usuário pelo ID
}

func CriarUsuario(w http.ResponseWriter, r *http.Request) {
	corpoRequest, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Erro ao ler o corpo da requisição")
		return
	}

	defer r.Body.Close()

	var novoUsuario modelo.Usuario
	if erro = json.Unmarshal(corpoRequest, &novoUsuario); erro != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Erro ao decodificar o JSON do corpo da requisição")
		return
	}

	if novoUsuario.Nome == "" || novoUsuario.Email == "" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Nome e email são obrigatórios")
		return
	}

	mutex.Lock()
	novoUsuario.ID = nextID
	nextID++
	usuarios = append(usuarios, novoUsuario)
	mutex.Unlock()

	respostaJSON, _ := json.Marshal(novoUsuario)
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	w.Write(respostaJSON)
}

func AtualizarUsuario(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userID, err := strconv.Atoi(params["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "ID inválido")
		return
	}

	var updatedUser modelo.Usuario
	if err := json.NewDecoder(r.Body).Decode(&updatedUser); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Erro ao decodificar o usuário")
		return
	}

	mutex.Lock()
	defer mutex.Unlock()

	for i, user := range usuarios {
		if user.ID == userID {
			// Atualiza os dados do usuário encontrado
			usuarios[i] = updatedUser

			// Retorna o usuário atualizado
			json.NewEncoder(w).Encode(updatedUser)
			return
		}
	}

	w.WriteHeader(http.StatusNotFound)
	fmt.Fprintf(w, "Usuário não encontrado")
}

func DeletarUsuario(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)
	userID, erro := strconv.Atoi(parametros["id"])
	if erro != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "id invalido")
		return
	}

	mutex.Lock()
	defer mutex.Unlock()

	for i, user := range usuarios {
		if user.ID == userID {
			usuarios = append(usuarios[:i], usuarios[i+1:]...)
			w.WriteHeader(http.StatusOK)
			fmt.Fprintf(w, "Usuário removido com sucesso")
			return
		}
	}
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprintf(w, "Usuário não encontrado")
}
