package models

import (
	"context"
	"errors"
	"time"

	"github.com/SkYNewZ/aloesia/config"
	"github.com/mitchellh/mapstructure"
	uuid "github.com/satori/go.uuid"
	"google.golang.org/api/iterator"
)

// User model struct
type User struct {
	ID                 string    `json:"id" firestore:"id" valid:"uuidv4"`
	Email              string    `json:"email" firestore:"email" valid:"email,required"`
	Password           string    `json:"password" firestore:"password" valid:"alphanum,required"`
	LastConnectionDate string    `json:"last_connection_date" firestore:"last_connection_date"`
	CreatedAt          time.Time `json:"created_at" mapstructure:"created_at" firestore:"created_at"`
	UpdatedAt          time.Time `json:"updated_at" mapstructure:"updated_at" firestore:"updated_at"`
}

// Users many users
type Users []User

// CreateUser creates user on Firestore
func CreateUser(u *User) error {
	// check if user already exist
	if found, _ := userAlreadyExist(u.Email); found == true {
		return errors.New("User already exist")
	}

	u.ID = uuid.NewV4().String()

	// date
	u.CreatedAt = time.Now()
	u.UpdatedAt = time.Now()

	_, err := config.Firestore().Collection(firestoreUsersCollectionName).Doc(u.ID).Set(context.Background(), &u)
	return err
}

// GetAllUsers returns all users
func GetAllUsers() (*Users, error) {
	iter := config.Firestore().Collection(firestoreUsersCollectionName).Documents(context.Background())
	var data Users
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, err
		}
		var u User
		mapstructure.Decode(doc.Data(), &u)
		data = append(data, u)
	}
	return &data, nil
}

// GetUser return user by given id
func GetUser(id string) (*User, error) {
	dsnap, err := config.Firestore().Collection(firestoreUsersCollectionName).Doc(id).Get(context.Background())
	if err != nil {
		return &User{}, err
	}
	var u User
	mapstructure.Decode(dsnap.Data(), &u)
	return &u, nil
}

// DeleteUser delete user by given id
func DeleteUser(id string) error {
	_, err := config.Firestore().Collection(firestoreUsersCollectionName).Doc(id).Delete(context.Background())
	return err
}

// ensure user does not exist
func userAlreadyExist(email string) (bool, error) {
	iter := config.Firestore().Collection(firestoreUsersCollectionName).Where("email", "==", email).Documents(context.Background())
	found := false
	for {
		_, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return false, err
		}
		found = true
	}
	return found, nil
}
