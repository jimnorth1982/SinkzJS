package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"syscall"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/term"
	itemsController "sinkzjs.org/m/v2/items/controller"
	itemsDb "sinkzjs.org/m/v2/items/db"
	itemsRoutes "sinkzjs.org/m/v2/items/routes"

	exilesController "sinkzjs.org/m/v2/exiles/controller"
	exilesDb "sinkzjs.org/m/v2/exiles/db"
	exilesRoutes "sinkzjs.org/m/v2/exiles/routes"
)

func main() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
		panic(err)
	}
	restAPI()
}

func restAPI() {

	mongodbProvider := itemsDb.NewMongoDBProvider("items", "items")
	setup(mongodbProvider, login())

	e := echo.New()

	itemsController := *itemsController.NewController(mongodbProvider)
	itemsRoutes.Routes(itemsController, e)

	exilesProvider := exilesDb.NewInMemoryProvider("exiles/db/data/exile_data.json")
	exilesController := *exilesController.NewController(exilesProvider)
	exilesRoutes.Routes(exilesController, e)

	e.Logger.Fatal(e.Start(":8080"))
}

func login() options.Credential {
	if os.Args != nil {
		return options.Credential{Username: os.Args[1], Password: os.Args[2]}
	}

	username, password, err := credentials()

	if err != nil {
		log.Fatalf("Bad input: %v", err)
		panic(err)
	}

	return options.Credential{Username: username, Password: password}
}

func setup(mongodbProvider *itemsDb.MongoDBProvider, creds options.Credential) {
	err := mongodbProvider.Connect(os.Getenv("MONGODB_URI"), creds, 4)

	if err != nil {
		log.Fatalf("Error getting connection to MongoDB: %v", err)
		panic(err)
	}

	if err := mongodbProvider.ClearAndLoadDataFromJSON(); err != nil {
		log.Fatalf("Error loading data: %v", err)
		panic(err)
	}
}

func credentials() (string, string, error) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter Username: ")
	username, err := reader.ReadString('\n')
	if err != nil {
		return "", "", err
	}

	fmt.Print("Enter Password: ")
	bytePassword, err := term.ReadPassword(int(syscall.Stdin))
	if err != nil {
		return "", "", err
	}

	password := string(bytePassword)
	return strings.TrimSpace(username), strings.TrimSpace(password), nil
}
