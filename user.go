package main //Import the main package
import (
	"context"
	// "go.mongodb.org/mongo-driver/mongo"
	// "go.mongodb.org/mongo-driver/mongo/options"
	// "go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

type User struct {
	FullName         string `bson:"name"`
	ID               string `bson:"id,omitempty"`
	PhoneNumber      int    `bson:"phone"`
	Email            string `bson:"email"`
	Address          string `bson:"address,omitempty"`
	SmartMeterNumber int    `bson:"smart-meter-number,omitempty"`
	PasswordHash     string `bson:"passowrd-hash,omitempty"`
	AddressHash      string `bson:"address-hash,omitempty"`
	EnergyAmount     int    `bson:"energy-amount,omitempty"`
	CurrencyAmount   int    `bson:"currency-amount,omitempty"`
}

func ValidateUserLogin(Email string, PasswordHash string) bool {
	var result bool
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	UserDatabase := connectToDb("Users")
	UserCollection := UserDatabase.Collection("Users")
	filterCursor, err := UserCollection.Find(ctx, bson.M{"email": Email})
	if err != nil {
		log.Fatal(err)
		result = false
	}
	var Profiles []User
	if err = filterCursor.All(ctx, &Profiles); err != nil {
		log.Fatal(err)
		result = false
	}
	Profile := Profiles[0]
	if Email == Profile.Email && PasswordHash == Profile.PasswordHash {
		result = true
	} else {
		result = false
	}

	return result
}
