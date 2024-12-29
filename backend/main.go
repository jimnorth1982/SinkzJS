package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"strings"
	"syscall"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"golang.org/x/term"
	itemsController "sinkzjs.org/m/v2/items/controller"
	itemsDb "sinkzjs.org/m/v2/items/db"
	itemsRoutes "sinkzjs.org/m/v2/items/routes"

	exilesController "sinkzjs.org/m/v2/exiles/controller"
	exilesDb "sinkzjs.org/m/v2/exiles/db"
	exilesRoutes "sinkzjs.org/m/v2/exiles/routes"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
		panic(err)
	}
	restAPI()
}

func restAPI() {

	mongodbProvider := itemsDb.NewMongoDBProvider()

	u, p := login()

	setup(mongodbProvider, u, p)

	e := echo.New()

	itemsProvider := itemsDb.NewInMemoryProvider("items/db/data/item_data.json")
	itemsController := *itemsController.NewController(itemsProvider)
	itemsRoutes.Routes(itemsController, e)

	exilesProvider := exilesDb.NewInMemoryProvider("exiles/db/data/exile_data.json")
	exilesController := *exilesController.NewController(exilesProvider)
	exilesRoutes.Routes(exilesController, e)

	e.Logger.Fatal(e.Start(":8080"))
}

func login() (u string, p string) {
	if os.Args != nil {
		return os.Args[1], os.Args[2]
	}

	username, password, err := credentials()

	if err != nil {
		log.Fatalf("Bad input: %v", err)
		panic(err)
	}

	return username, password
}
func setup(mongodbProvider *itemsDb.MongoDBProvider, username string, password string) {
	client, err := mongodbProvider.GetConnection(os.Getenv("MONGODB_URI"), username, password)
	defer client.Disconnect(context.TODO())

	if err != nil {
		log.Fatalf("Error getting connection to MongoDB: %v", err)
		panic(err)
	}

	if err := itemsDb.ClearAndLoadDataFromSJON(client); err != nil {
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
