package modelo

// Usuario representa um usuário no sistema
type Usuario struct {
	ID    int
	Nome  string
	Email string
}

// CriarNovoUsuario cria e retorna um novo usuário
func CriarNovoUsuario(ID int, nome, email string) Usuario {
	return Usuario{
		ID:    ID,
		Nome:  nome,
		Email: email,
	}
}

// AdicionarUsuario adiciona um novo usuário à lista de usuários
func AdicionarUsuario(listaUsuarios []Usuario, novoUsuario Usuario) []Usuario {
	listaUsuarios = append(listaUsuarios, novoUsuario)
	return listaUsuarios
}
