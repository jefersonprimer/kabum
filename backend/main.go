package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv" // Importe o pacote godotenv
	_ "github.com/lib/pq"
)

var db *sql.DB

func init() {
	// Carrega o arquivo .env
	err := godotenv.Load()
	if err != nil {
		log.Println("Atenção: Erro ao carregar o arquivo .env, usando variáveis de ambiente do sistema ou hardcoded.")
		// Não é um erro fatal se o .env não existir (por exemplo, em produção onde as vars de ambiente vêm do orquestrador)
	}

	connStr := os.Getenv("DATABASE_URL")
	if connStr == "" {
		log.Fatalf("Erro: Variável de ambiente DATABASE_URL não definida.")
	}

	var dbErr error
	db, dbErr = sql.Open("postgres", connStr)
	if dbErr != nil {
		log.Fatalf("Erro ao abrir a conexão com o banco de dados: %v", dbErr)
	}

	dbErr = db.Ping()
	if dbErr != nil {
		log.Fatalf("Erro ao conectar ao banco de dados: %v", dbErr)
	}

	log.Println("Conexão com o banco de dados estabelecida com sucesso!")
}

func main() {
	defer db.Close()

	router := gin.Default()

	router.GET("/pingdb", func(c *gin.Context) {
		var result string
		err := db.QueryRow("SELECT 'pong'").Scan(&result)
		if err != nil {
			c.JSON(500, gin.H{"message": "Erro ao consultar o banco de dados", "error": err.Error()})
			return
		}
		c.JSON(200, gin.H{"message": fmt.Sprintf("Database says: %s", result)})
	})

	// Pega a porta da variável de ambiente, padrão para 8080 se não definida
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	router.Run(":" + port)
}
