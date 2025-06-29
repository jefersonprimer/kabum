// main.go
package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv" // Para carregar variáveis de ambiente
)

func main() {
	// Carrega as variáveis de ambiente do arquivo .env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Erro ao carregar arquivo .env")
	}

	// Inicializa o Gin
	router := gin.Default()

	// Exemplo de rota de teste
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// TODO: Configurar rotas para interagir com a API do Mercado Livre e Supabase aqui

	// Inicia o servidor
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Porta padrão se não estiver definida
	}
	log.Printf("Servidor rodando na porta %s", port)
	router.Run(":" + port)
}
