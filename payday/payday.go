package payday

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	firebase "firebase.google.com/go"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
)

// Use the application default credentials
func main() {
	// Read From DB
	// Read From DB
}

type user struct {
	FirstNameEn string
	LastNameEn  string
	FirstNameTh string
	LastNameTh  string
	TitleEn     string
	TitleTh     string
	DisplayName string
	UserType    string
}

func New(data map[string]string) user {
	t := user{
		FirstNameEn: data["firstNameEn"],
		LastNameEn:  data["lastNameEn"],
		FirstNameTh: data["firstNameTh"],
		LastNameTh:  data["lastNameTh"],
		TitleEn:     data["titleEn"],
		TitleTh:     data["titleTh"],
		DisplayName: data["displayName"],
		UserType:    data["userType"],
	}
	return t
}
func (data user) AddUser(id string) string {
	ctx := context.Background()
	sa := option.WithCredentialsFile("./paydayconnect.json")
	app, err := firebase.NewApp(ctx, nil, sa)
	if err != nil {
		log.Fatalln(err)
	}

	client, err := app.Firestore(ctx)

	if err != nil {
		log.Fatalln(err)
	}
	defer client.Close()
	_, err = client.Collection("users").Doc(id).Set(ctx, map[string]interface{}{
		"first_name_en": data.FirstNameEn,
		"last_name_en":  data.LastNameEn,
		"first_name_th": data.FirstNameTh,
		"last_name_th":  data.FirstNameTh,
		"title_en":      data.TitleEn,
		"title_th":      data.TitleTh,
		"display_name":  data.DisplayName,
		"user_type":     data.UserType,
	})
	if err != nil {
		log.Fatalf("Failed adding aturing: %v", err)
	}
	return "Create Success"
}
func (data user) UpdateUser(id string) string {
	ctx := context.Background()
	sa := option.WithCredentialsFile("./paydayconnect.json")
	app, err := firebase.NewApp(ctx, nil, sa)
	if err != nil {
		log.Fatalln(err)
	}

	client, err := app.Firestore(ctx)

	if err != nil {
		log.Fatalln(err)
	}
	defer client.Close()
	_, err = client.Collection("users").Doc(id).Set(ctx, map[string]interface{}{
		"first_name_en": data.FirstNameEn,
		"last_name_en":  data.LastNameEn,
		"first_name_th": data.FirstNameTh,
		"last_name_th":  data.FirstNameTh,
		"title_en":      data.TitleEn,
		"title_th":      data.TitleTh,
		"display_name":  data.DisplayName,
		"user_type":     data.UserType,
	})
	if err != nil {
		log.Fatalf("Failed adding aturing: %v", err)
	}
	return "Create Success"
}
func (data user) GetUser(id string) map[string]interface{} {
	ctx := context.Background()
	sa := option.WithCredentialsFile("./paydayconnect.json")
	app, err := firebase.NewApp(ctx, nil, sa)
	if err != nil {
		log.Fatalln(err)
	}

	client, err := app.Firestore(ctx)

	if err != nil {
		log.Fatalln(err)
	}
	defer client.Close()
	dsnap, err := client.Collection("users").Doc(id).Get(ctx)
	if err != nil {
		log.Fatalf("Failed to iterate: %v", err)
	}
	user := dsnap.Data()
	return user
}
func (data user) ReadUser() string {
	ctx := context.Background()
	sa := option.WithCredentialsFile("./paydayconnect.json")
	app, err := firebase.NewApp(ctx, nil, sa)
	if err != nil {
		log.Fatalln(err)
	}

	client, err := app.Firestore(ctx)

	if err != nil {
		log.Fatalln(err)
	}
	defer client.Close()
	iter := client.Collection("users").Documents(ctx)
	var user []interface{}
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatalf("Failed to iterate: %v", err)
		}
		user = append(user, doc.Data())
	}

	jsonString, _ := json.Marshal(user)
	fmt.Println(string(jsonString))
	return string(jsonString)
}
