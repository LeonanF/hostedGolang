package main

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	mongoClient *mongo.Client
	collection  *mongo.Collection
)

type Produto struct {
	ID   primitive.ObjectID `bson:"_id,omitempty"`
	Name string             `bson:"name"`
}

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

	// Novo endpoint para verificar a conexão com o MongoDB e retornar um documento
	app.Get("/ping", func(c *fiber.Ctx) error {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		// Realiza uma consulta simples ao MongoDB para buscar um documento
		var produto Produto
		err := collection.FindOne(ctx, bson.M{}).Decode(&produto)
		if err != nil {
			return c.Status(500).SendString("Erro ao buscar documento no MongoDB")
		}

		return c.JSON(produto)
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
