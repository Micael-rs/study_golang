package atividade2

type Usuario struct {
	Nome  string
	Email string
}

func CriarUsuario(nome, email string) *Usuario {

	return &Usuario{
		Nome:  nome,
		Email: email,
	}

}
