package services

import (
	"contacts-api-mongo/entity"
	"contacts-api-mongo/global"

	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func ContactCreate(contact entity.Contact) (entity.Contact, error) {
	res, err := global.DB.Collection("contacts").InsertOne(context.Background(), contact)
	if err != nil {
		log.Fatal(err)
		return entity.Contact{}, err
	}
	fmt.Println(res.InsertedID)
	return contact, nil
}

func ContactGet(contactID string) (entity.Contact, error) {
	var contact entity.Contact
	objectId, err := primitive.ObjectIDFromHex(contactID)
	if err != nil {
		log.Fatal(err)
		return contact, err
	}
	err = global.DB.Collection("contacts").FindOne(context.Background(), bson.M{"_id": objectId}).Decode(&contact)
	if err != nil {
		log.Fatal(err)
		return contact, err
	}
	return contact, nil
}

func ContactDelete(ID string) error {
	objectId, err := primitive.ObjectIDFromHex(ID)
	if err != nil {
		log.Fatal(err)
		return err
	}
	res, err := global.DB.Collection("contacts").DeleteOne(context.Background(), bson.M{"_id": objectId})
	if err != nil {
		log.Fatal(err)
		return err
	}
	fmt.Println("deleted: ", res.DeletedCount)
	return nil
}

func ContactUpdate(ID string, contact entity.Contact) (entity.Contact, error) {
	objectId, err := primitive.ObjectIDFromHex(ID)
	if err != nil {
		log.Fatal(err)
		return contact, err
	}
	contact.ID = objectId
	_, err = global.DB.Collection("contacts").UpdateOne(context.Background(), bson.M{"_id": objectId}, bson.M{"$set": contact})
	if err != nil {
		log.Fatal(err)
		return contact, err
	}
	return contact, nil
}

func ContactGetAll() ([]entity.Contact, error) {
	var contacts []entity.Contact
	cursor, err := global.DB.Collection("contacts").Find(context.Background(), bson.D{})
	if err = cursor.All(context.Background(), &contacts); err != nil {
		log.Fatal(err)
		return contacts, err
	}
	return contacts, nil
}
