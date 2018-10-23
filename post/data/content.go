package data

import (
	"time"

	"github.com/lukma/sample-go-clean/common"
	"github.com/lukma/sample-go-clean/post/domain"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type data struct {
	collection *mgo.Collection
}

func NewContentRepository() domain.ContentRepository {
	_, database := common.GetConnection()

	collection := database.C("post_contents")

	return &data{
		collection: collection,
	}
}

func (data *data) CountContent(query string) (int, error) {
	result, err := data.collection.Find(bson.M{
		"$or": []bson.M{
			bson.M{"title": bson.M{"$regex": bson.RegEx{Pattern: ".*" + query + ".*", Options: "i"}}},
			bson.M{"content": bson.M{"$regex": bson.RegEx{Pattern: ".*" + query + ".*", Options: "i"}}},
		},
	}).Count()

	return result, err
}

func (data *data) GetContents(limit int, offset int, query string, sort string) ([]domain.ContentEntity, error) {
	var result []domain.ContentEntity
	err := data.collection.Find(bson.M{
		"$or": []bson.M{
			bson.M{"title": bson.M{"$regex": bson.RegEx{Pattern: ".*" + query + ".*", Options: "i"}}},
			bson.M{"content": bson.M{"$regex": bson.RegEx{Pattern: ".*" + query + ".*", Options: "i"}}},
		},
	}).Sort(sort).Skip(offset).Limit(limit).Iter().All(&result)

	return result, err
}

func (data *data) GetContent(id string) (domain.ContentEntity, error) {
	var result domain.ContentEntity
	err := data.collection.FindId(bson.ObjectIdHex(id)).One(&result)

	return result, err
}

func (data *data) CreateContent(obj domain.ContentEntity) (string, error) {
	obj.ID = bson.NewObjectId()
	obj.CreatedDate = time.Now()
	err := data.collection.Insert(obj)

	return obj.ID.Hex(), err
}

func (data *data) UpdateContent(id string, obj domain.ContentEntity) error {
	var selectedObj domain.ContentEntity
	err := data.collection.FindId(bson.ObjectIdHex(id)).One(&selectedObj)
	if err == nil {
		err = data.collection.UpdateId(selectedObj.ID, bson.M{"$set": obj})
	}

	return err
}

func (data *data) DeleteContent(id string) error {
	return data.collection.RemoveId(bson.ObjectIdHex(id))
}
