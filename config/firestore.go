package config

// TODO: Implement this file

var firestore int

// InitFirestoreDatabase init main firestore client
func InitFirestoreDatabase() {
	firestore = 1
}

// Firestore Getter for firestore database
func Firestore() int {
	return firestore
}
