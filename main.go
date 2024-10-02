package main

import (
	"github.com/alura-cursos/Curso_CI/database"
	"github.com/alura-cursos/Curso_CI/routes"
)

func main() {
	database.ConectaComBancoDeDados()
	routes.HandleRequest()
}
