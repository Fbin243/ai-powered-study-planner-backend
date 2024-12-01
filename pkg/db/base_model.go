package db

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type BaseModel struct {
	ID        primitive.ObjectID `json:"id" bson:"_id"`
	CreatedAt time.Time          `json:"createdAt" bson:"created_at"`
	UpdatedAt time.Time          `json:"updatedAt" bson:"updated_at"`
}

type IBaseModel interface {
	SetID()
	SetCreatedAtByNow()
	SetUpdatedAtByNow()
}

func (m *BaseModel) SetID() {
	m.ID = primitive.NewObjectID()
}

func (m *BaseModel) SetCreatedAtByNow() {
	m.CreatedAt = time.Now()
}

func (m *BaseModel) SetUpdatedAtByNow() {
	m.UpdatedAt = time.Now()
}

type IBaseRepo[M IBaseModel] interface {
	Insert(*M) (*M, error)
	FindById(ID primitive.ObjectID) (*M, error)
	UpdateById(ID primitive.ObjectID, m *M) (*M, error)
	DeleteById(ID primitive.ObjectID) (*M, error)
}

type BaseRepo[M IBaseModel] struct {
	Collection *mongo.Collection
}

func NewBaseRepo[M IBaseModel](collection *mongo.Collection) *BaseRepo[M] {
	return &BaseRepo[M]{
		Collection: collection,
	}
}

func (r *BaseRepo[M]) Insert(m *M) (*M, error) {
	_m := *m
	_m.SetID()
	_m.SetCreatedAtByNow()
	_m.SetUpdatedAtByNow()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := r.Collection.InsertOne(ctx, m)
	if err != nil {
		return nil, err
	}

	return m, nil
}

func (r *BaseRepo[M]) FindById(ID primitive.ObjectID) (*M, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var result M
	err := r.Collection.FindOne(ctx, bson.M{"_id": ID}).Decode(&result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (r *BaseRepo[M]) UpdateById(ID primitive.ObjectID, m *M) (*M, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_m := *m
	_m.SetUpdatedAtByNow()

	_, err := r.Collection.ReplaceOne(ctx, bson.M{"_id": ID}, _m)
	if err != nil {
		return nil, err
	}

	return &_m, nil
}

func (r *BaseRepo[M]) DeleteById(ID primitive.ObjectID) (*M, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var result M
	err := r.Collection.FindOneAndDelete(ctx, bson.M{"_id": ID}).Decode(&result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
