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

func New(FirstNameEn string,
	LastNameEn string,
	FirstNameTh string,
	LastNameTh string,
	TitleEn string,
	TitleTh string,
	DisplayName string,
	UserType string) user {
	t := user{
		FirstNameEn: FirstNameEn,
		LastNameEn:  LastNameEn,
		FirstNameTh: FirstNameTh,
		LastNameTh:  LastNameTh,
		TitleEn:     TitleEn,
		TitleTh:     TitleTh,
		DisplayName: DisplayName,
		UserType:    UserType,
	}
	return t
}
func (data user) AddUser() string {
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
	_, _, err = client.Collection("users").Add(ctx, map[string]interface{}{
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
