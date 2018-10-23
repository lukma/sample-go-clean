package data

import (
	"context"
	"os"
	"path"
	"time"

	"firebase.google.com/go/auth"
	"github.com/lukma/sample-go-clean/account/domain"
	"github.com/lukma/sample-go-clean/common"
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

func (data *data) GetAuth(faID string) (domain.AuthEntity, error) {
	var result domain.AuthEntity
	err := data.collection.Find(bson.M{"fa_id": faID}).One(&result)
	return result, err
}

func (data *data) CreateAuth(obj domain.AuthEntity) (string, error) {
	user, err := data.client.GetUser(context.Background(), obj.FaID)

	obj.ID = bson.NewObjectId()
	obj.Email = user.Email
	obj.Phone = user.PhoneNumber
	obj.CreatedDate = time.Now()
	err = data.collection.Insert(obj)

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
