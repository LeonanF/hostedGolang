package main

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	mongoClient *mongo.Client
	collection  *mongo.Collection
)

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello world")
	})

	port := os.Getenv("PORT")

	if port == "" {
		port = "3000"
	}

	connectDB()

	app.Get("/ping", func(c *fiber.Ctx) error {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		err := mongoClient.Ping(ctx, nil)
		if err != nil {
			return c.Status(500).SendString("Erro ao pingar o MongoDB")
		}

		return c.SendString("Ping bem-sucedido no MongoDB")
	})

	log.Fatal(app.Listen("0.0.0.0:" + port))

}

func connectDB() {
	//
	mongodbURI := os.Getenv("MONGODB_URI")

	// Primeiro é criado um cliente MongoDB e usado o método Connect para conectar-se ao banco de dados.
	// É fornecido um contexto de fundo neutro e configurada a URI de conexão (com a variável criada acima).
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(mongodbURI))
	if err != nil {
		panic(err)
	}

	// Teste de conexão com o MongoDB
	err = client.Ping(context.Background(), nil)
	if err != nil {
		panic("Erro ao conectar ao MongoDB: " + err.Error())
	}

	// O cliente criado é atribuído à variável global para que outras partes do código possam acessar e executar operações no banco de dados.
	mongoClient = client

	// É especificada a coleção de documentos a ser trabalhada em cima.
	collection = client.Database("rolegourmet").Collection("produtos")
}
