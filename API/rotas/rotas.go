package rotas

import (
	"API/modelo"
	"encoding/json"
	"fmt"
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

func ConfigurarRotas() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/usuarios", ListarUsuarios).Methods("GET")

	r.HandleFunc("/usuarios/{id}", ObterUsuario).Methods("GET")

	// Rota para criar um novo usuário
	r.HandleFunc("/usuarios", CriarUsuario).Methods("POST")

	// Rota para atualizar um usuário existente
	r.HandleFunc("/usuarios/{id}", AtualizarUsuario).Methods("PUT")

	// Rota para excluir um usuário
	r.HandleFunc("/usuarios/{id}", DeletarUsuario).Methods("DELETE")

	return r
}

func ListarUsuarios(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(usuarios)
}

func ObterUsuario(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userID, err := strconv.Atoi(params["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "ID inválido")
		return
	}

	for _, user := range usuarios {
		if user.ID == userID {
			json.NewEncoder(w).Encode(user)
			return
		}
	}

	w.WriteHeader(http.StatusNotFound)
	fmt.Fprintf(w, "Usuário não encontrado")
}

func CriarUsuario(w http.ResponseWriter, r *http.Request) {
	var usuario modelo.Usuario
	_ = json.NewDecoder(r.Body).Decode(&usuario)

	mutex.Lock()
	usuario.ID = nextID
	nextID++
	usuarios = append(usuarios, usuario)
	mutex.Unlock()

	json.NewEncoder(w).Encode(usuario)
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
	_ = json.NewDecoder(r.Body).Decode(&updatedUser)

	mutex.Lock()
	for i, user := range usuarios {
		if user.ID == userID {
			usuarios[i] = updatedUser
			mutex.Unlock()
			json.NewEncoder(w).Encode(updatedUser)
			return
		}
	}
	mutex.Unlock()

	w.WriteHeader(http.StatusNotFound)
	fmt.Fprintf(w, "Usuário não encontrado")
}

func DeletarUsuario(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userID, err := strconv.Atoi(params["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "ID inválido")
		return
	}

	mutex.Lock()
	defer mutex.Unlock()

	for i, user := range usuarios {
		if user.ID == userID {
			usuarios = append(usuarios[:i], usuarios[i+1:]...)
			fmt.Fprintf(w, "Usuário removido com sucesso")
			return
		}
	}

	w.WriteHeader(http.StatusNotFound)
	fmt.Fprintf(w, "Usuário não encontrado")
}
