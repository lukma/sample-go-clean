package data

import (
	"context"
	"os"
	"path"
	"time"

	"firebase.google.com/go/auth"
	"github.com/lukma/sample-go-clean/account/domain"
	"github.com/lukma/sample-go-clean/common"
	"golang.org/x/crypto/bcrypt"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type data struct {
	collection *mgo.Collection
	client     *auth.Client
}

func NewAuthRepository() domain.AuthRepository {
	app, err := common.GetFirebaseApp(path.Join(os.Getenv("GOPATH"), "/src/github.com/lukma/sample-go-clean/credentials/firebase-adminsdk.json"))
	if err != nil {
		panic(err.Error())
	}

	client, err := app.Auth(context.Background())
	if err != nil {
		panic(err.Error())
	}

	_, database := common.GetConnection()

	collection := database.C("account_auths")

	return &data{
		collection: collection,
		client:     client,
	}
}

func (data *data) GetAuthByID(id string) (domain.AuthEntity, error) {
	var result domain.AuthEntity
	err := data.collection.FindId(bson.ObjectId(id)).One(&result)
	return result, err
}

func (data *data) GetAuthByUsernameOrEmail(usernameOrEmail string, password string) (domain.AuthEntity, error) {
	var result domain.AuthEntity
	err := data.collection.Find(bson.M{
		"$or": []bson.M{
			bson.M{"username": usernameOrEmail},
			bson.M{"email": usernameOrEmail},
		},
	}).One(&result)

	if err == nil {
		err = bcrypt.CompareHashAndPassword([]byte(result.Password), []byte(password))
	}

	return result, err
}

func (data *data) GetAuthByThirdParty(thirdParty string, token string) (domain.AuthEntity, error) {
	var query bson.M
	if thirdParty == "facebook" {
		query = bson.M{"facebook_token": token}
	} else if thirdParty == "google" {
		query = bson.M{"google_token": token}
	}

	var result domain.AuthEntity
	err := data.collection.Find(query).One(&result)
	return result, err
}

func (data *data) CreateAuth(obj domain.AuthEntity) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(obj.Password), bcrypt.DefaultCost)
	if err == nil {
		obj.ID = bson.NewObjectId()
		obj.Password = string(hashedPassword)
		obj.CreatedDate = time.Now()
		err = data.collection.Insert(obj)
	}

	return obj.ID.Hex(), err
}

func (data *data) UpdateAuth(id string, obj domain.AuthEntity) error {
	var selectedObj domain.AuthEntity
	err := data.collection.FindId(bson.ObjectIdHex(id)).One(&selectedObj)
	if err == nil {
		err = data.collection.UpdateId(selectedObj.ID, bson.M{"$set": obj})
	}

	return err
}
