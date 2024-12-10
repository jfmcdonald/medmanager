package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func ConnectToMongoDB() (*mongo.Client, error) {
	// Set client options
	clientOpts := options.Client().ApplyURI(mongoURI).SetServerSelectionTimeout(5 * time.Second)

	// Create a new client and connect to the server
	client, err := mongo.Connect(context.Background(), clientOpts)
	if err != nil {
		return nil, fmt.Errorf("failed to create client: %w", err)
	}

	// Ping the primary to verify connectivity
	if err := client.Ping(context.Background(), readpref.Primary()); err != nil {
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	return client, nil
}

func insertPatient(p *Patient) (bool, error) {

	mongoclient, err := ConnectToMongoDB()
	if err != nil {
		return false, fmt.Errorf("unable to connect to db")

	}

	collection := mongoclient.Database("medmanager").Collection("patients")

	_, err = collection.InsertOne(context.TODO(), p)
	if err != nil {
		return false, fmt.Errorf("unable to save patient data")
	} else {
		return true, nil
	}

}

func getPatients() ([]Patient, error) {
	var dbpatients []Patient

	mongoclient, err := ConnectToMongoDB()
	if err != nil {
		return dbpatients, fmt.Errorf("unable to connect to db")

	}

	collection := mongoclient.Database("medmanager").Collection("patients")

	// Retrieve all documents
	cursor, err := collection.Find(context.TODO(), bson.D{})
	if err != nil {
		log.Fatal("failed to find documents:", err)
	}
	defer cursor.Close(context.TODO())

	if err = cursor.All(context.TODO(), &dbpatients); err != nil {
		log.Fatal("failed to decode documents:", err)
	}
	return dbpatients, nil
}
