package db

import (
	"context"
	"log"
	"time"

	"github.com/olawolu/circle-payments/pkg/payments"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"gopkg.in/mgo.v2/bson"
)

type Mongo struct {
	client   *mongo.Client
	database string
	timeout  time.Duration
}

func newMongoClient(mongoURL string, timeout int) (*mongo.Client, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(timeout)*time.Second)
	defer cancel()

	// create client
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoURL))
	if err != nil {
		return nil, err
	}
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		return nil, err
	}
	return client, err
}

func NewMongoRepository(mongoURL, dbName string, timeout int) (*Mongo, error) {
	repo := &Mongo{
		timeout:  time.Duration(timeout) * time.Second,
		database: dbName,
	}
	client, err := newMongoClient(mongoURL, timeout)
	if err != nil {
		log.Println("monog.NewMongoRepository ", err)
		return nil, err
	}
	repo.client = client
	return repo, nil
}

func (m *Mongo) Get() ([]payments.PaymentResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), m.timeout)
	defer cancel()
	response := payments.PaymentResponse{}
	var paymentList []payments.PaymentResponse
	collection := m.client.Database(m.database).Collection("payments")
	filter := bson.M{}
	cursor, err := collection.Find(ctx, filter)
	if err != nil {
		log.Println("mongo.Add ", err)
		return nil, err
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		if err = cursor.Decode(&response); err != nil {
			return nil, err
		}
		paymentList = append(paymentList, response)
	}
	return paymentList, nil
}

func (m *Mongo) Add(paymentDetail *payments.PaymentResponse) error {
	ctx, cancel := context.WithTimeout(context.Background(), m.timeout)
	defer cancel()
	collection := m.client.Database(m.database).Collection("payments")
	_, err := collection.InsertOne(
		ctx,
		bson.M{
			"user":      paymentDetail.Metadata.Email,
			"paymentID": paymentDetail.PaymentID,
			"amount":    paymentDetail.Amount,
			"status":    paymentDetail.Status,
			"createdAt": paymentDetail.CreateDate,
			"updatedAt": paymentDetail.UpdateDate,
		},
	)
	if err != nil {
		log.Println("mongo.Add ", err)
		return err
	}
	return nil
}

func (m *Mongo) Update(paymentDetail *payments.PaymentResponse) error {
	ctx, cancel := context.WithTimeout(context.Background(), m.timeout)
	defer cancel()
	collection := m.client.Database(m.database).Collection("payments")
	result := collection.FindOneAndUpdate(
		ctx,
		bson.M{
			"paymentID": paymentDetail.PaymentID,
		},
		bson.M{
			"status":    paymentDetail.Status,
			"updatedAt": paymentDetail.UpdateDate,
		},
	)

	if result.Err() != nil {
		log.Println("mongo.Add ", result.Err().Error())
		return result.Err()
	}
	return nil
}
