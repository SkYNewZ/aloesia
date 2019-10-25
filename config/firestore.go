package config

import (
	"context"
	"log"
	"os"

	"cloud.google.com/go/firestore"
)

var firestoreClient *firestore.Client

// InitFirestoreDatabase init main firestore client
func InitFirestoreDatabase() {
	var err error
	firestoreClient, err = firestore.NewClient(context.Background(), os.Getenv("GOOGLE_CLOUD_PROJECT"))
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}
}

// Firestore Getter for firestore database
func Firestore() *firestore.Client {
	return firestoreClient
}

// CloseConnection clone connection with Firebase
func CloseConnection() {
	// Close client when done.
	defer firestoreClient.Close()
}
