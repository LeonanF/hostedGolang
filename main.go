package main

import (
	"context"
	"log"
	"os"

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

	log.Fatal(app.Listen("0.0.0.0:" + port))

}

func connectDB() {
	mongodbURI := os.Getenv("MONGODB_URI")

	//Primeiro é criado um cliente MongoDB, e usada a função Connect para conectar com o banco de dados
	//É fornecido um contexto de fundo neutro, e configurado a URI de conexão (com a váriavel criada acima)
	//Devido à importância da conexão com o BD, se ocorrer algum erro, o programa inteiro é encerrado
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(mongodbURI))
	if err != nil {
		panic(err)
	}

	//O cliente criado é atribuído à váriavel global para que outras partes do código possam acessar e executar operações no banco de dados
	mongoClient = client

	//É especificada a coleção de documentos a ser trabalhado em cima.
	collection = client.Database("rolegourmet").Collection("produtos")
}
