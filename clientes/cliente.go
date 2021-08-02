package domain

type Cliente struct {
	Nome, Cpf, Profissao string
}

func CriaCliente(nome string, Cpf string, Profissao string) Cliente {
	novoTitular := Cliente{}
	novoTitular.Nome = nome
	novoTitular.Cpf = Cpf
	novoTitular.Profissao = Profissao
	return novoTitular
}
