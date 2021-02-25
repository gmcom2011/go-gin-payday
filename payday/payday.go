package payday

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"github.com/gin-gonic/gin"
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
type app struct {
	ctx     context.Context
	storage *firestore.Client
}

func DataBody(c *gin.Context) map[string]string {
	buf := make([]byte, 1024)
	rawBody, _ := c.Request.Body.Read(buf)
	jsonBody := buf[0:rawBody]
	fmt.Println(jsonBody)
	var mapBody map[string]string
	json.Unmarshal(jsonBody, &mapBody)
	return mapBody
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
	updateData := map[string]interface{}{
		"first_name_en": data.FirstNameEn,
		"last_name_en":  data.LastNameEn,
		"first_name_th": data.FirstNameTh,
		"last_name_th":  data.FirstNameTh,
		"title_en":      data.TitleEn,
		"title_th":      data.TitleTh,
		"display_name":  data.DisplayName,
		"user_type":     data.UserType,
	}
	fmt.Println("Update Data:", updateData)
	fmt.Println("Update Data ID:", id)

	_, Updateerr := client.Collection("users").Doc(id).Set(ctx, map[string]interface{}{

		"first_name_en": data.FirstNameEn,
		"last_name_en":  data.LastNameEn,
		"first_name_th": data.FirstNameTh,
		"last_name_th":  data.FirstNameTh,
		"title_en":      data.TitleEn,
		"title_th":      data.TitleTh,
		"display_name":  data.DisplayName,
		"user_type":     data.UserType,
	}, firestore.MergeAll)
	if Updateerr != nil {
		log.Fatalf("Failed adding aturing: %v", Updateerr)
	}
	return "Update Success"
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

func DeleteUser(id string) {
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
	_, deleteErr := client.Collection("users").Doc(id).Delete(ctx)
	if deleteErr != nil {
		// Handle any errors in an appropriate way, such as returning them.
		log.Printf("An error has occurred: %s", err)
	}
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func (route *app) UploadProfile(w http.ResponseWriter, r *http.Request) {
	route.ctx = context.Background()
	file, handler, err := r.FormFile("image")
	r.ParseMultipartForm(10 << 20)
	if err != nil {
		respondWithJSON(w, http.StatusBadRequest, err.Error())
		return
	}
	defer file.Close()

	imagePath := handler.Filename

	bucket := "payday-e074e.appspot.com"

	wc := route.storage.Bucket(bucket).Object(imagePath).NewWriter(route.ctx)
	_, err = io.Copy(wc, file)
	if err != nil {
		respondWithJSON(w, http.StatusBadRequest, err.Error())
		return

	}
	if err := wc.Close(); err != nil {
		respondWithJSON(w, http.StatusBadRequest, err.Error())
		return
	}

}
