package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/relay"
	"github.com/joho/godotenv"
	"github.com/olawolu/circle-payments/pkg/db"
	graph "github.com/olawolu/circle-payments/pkg/graphql"
)

func init() {
	if err := godotenv.Load("./.env"); err != nil {
		log.Fatalf("Failed to load environment variables: %v", err)
	}
}

func main() {
	// Read and parse the schema:
	bstr, err := ioutil.ReadFile("./pkg/graphql/schema.graphql")
	if err != nil {
		panic(err)
	}
	schemaString := string(bstr)

	// db connection
	dbUrl := os.Getenv("MONGO_URL")
	database := os.Getenv("MONGO_DB")
	timeout, _ := strconv.Atoi(os.Getenv("MONGO_TIMEOUT"))
	repo, err := db.NewMongoRepository(dbUrl, database, timeout)
	if err != nil {
		log.Fatal(err)
	}

	resolver := graph.RootResolver{}
	resolver.Client = http.Client{}
	resolver.APIKey = os.Getenv("API_KEY")
	resolver.URL = "https://api-sandbox.circle.com/v1"

	log.Println(resolver.APIKey)

	go updatePaymentData(*repo, resolver)

	graphqlSchema := graphql.MustParseSchema(schemaString, &resolver)

	fmt.Println(graphqlSchema)
	http.Handle("/graphql", &relay.Handler{Schema: graphqlSchema})
	log.Println("listening")
	log.Fatal(http.ListenAndServe(":8000", http.DefaultServeMux))
}

func updatePaymentData(repo db.Mongo, resolver graph.RootResolver) {
	// var payments []payments.PaymentResponse

	paymentList, err := repo.Get()
	if err != nil {
		log.Println("Database Get() operatiob error in database")
	}
	for i := 0; i < len(paymentList); i++ {
		v := paymentList[i]
		if v.Status != "confirmed" {
			response, err := resolver.GetPaymentCall(v.PaymentID)
			if err != nil {
				log.Println("updatePaymentData().GetPaymentCall() error ", err)
			}
			switch response.Status {
			case "confirmed":
				err = repo.Update(response)
				if err != nil {
					log.Println("Update error ", err)
				}
			default:
				time.Sleep(5 * time.Second)
				i = -1
			}
		}
	}
}
