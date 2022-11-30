package bootstrap

import (
	"log"
	"os"
	"test/init/database/migrations"
	"test/init/database/storage"
	"test/init/repository"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
)

func InitializeApp(app *fiber.App){
	// _, ok := os.LookupEnv()
	err := godotenv.Load(".env")
	if err != nil{
		log.Fatal(err)
	}
	config := &storage.Config{
		Host: os.Getenv("DB_HOST"),
		Port: os.Getenv("DB_PORT"),
		Password: os.Getenv("DB_PASSWORD"),
		User: os.Getenv("DB_USER"),
		SSLMode: os.Getenv("DB_SSLMODE"),
		DBName: os.Getenv("DB_NAME"),
	}

	db, err := storage.NewConnection(config)

	if err != nil {
		log.Fatal("Could not load the database")
	}

	err = migrations.MigrateUsers(db)

	if err != nil{
		log.Fatal("COuld not migrate the db")
	}

	repo := repository.Repository{
		DB: db,
	}
	app.Use(cors.New(cors.Config{AllowCredentials: true}))
	repo.SetupRoutes(app)
	app.Listen(":8081")
}