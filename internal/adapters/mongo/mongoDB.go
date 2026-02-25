package mongo

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/maryam-nokohan/go-article/internal/configs"
	"github.com/maryam-nokohan/go-article/internal/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoRepo struct {
	Client      *mongo.Client
	Collections *mongo.Collection
	Config      *configs.Config
}

func NewMongoRepo() (*MongoRepo, error) {
	config, err := configs.Newconfig()
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	c, err := mongo.Connect(ctx, options.Client().ApplyURI(config.URI))
	if err != nil {
		return nil, err
	}

	if err = c.Ping(ctx, nil); err != nil {
		return nil, err
	}

	collections := c.Database("hello").Collection("articles")
	_, err = collections.Indexes().CreateOne(ctx, mongo.IndexModel{
		Keys: bson.D{
			{Key: "title", Value: -1},
		},
		Options: options.Index().SetUnique(true),
	})
	if err != nil {
		return nil, err
	}

	return &MongoRepo{
		Client:      c,
		Collections: collections,
		Config:      config,
	}, nil

}

func (db *MongoRepo) Save(ctx context.Context, article *domain.Article) error {
	log.Println("In mongoRepo : Saving article with title:", article.Title)
	article.ID = primitive.NewObjectID().Hex()
	article.Created_at = time.Now()

	_, err := db.Collections.InsertOne(ctx, article)
	if err != nil {
		if mongo.IsDuplicateKeyError(err) {
			return fmt.Errorf("Duplicate Article , %v", err)
		}
		return err
	}
	return nil
}

func (db *MongoRepo) GetTopTags(ctx context.Context, limit int64) ([]domain.Tag, error) {
	log.Println("In mongoRepo : Getting top tags with limit:", limit)
	pipeline := mongo.Pipeline{
		bson.D{{"$unwind", "$tags"}},
		bson.D{{"$group", bson.D{
			{"_id", "$tags.word"},
			{"freq", bson.D{{"$sum", "$tags.freq"}}},
		}}},
		bson.D{{"$sort", bson.D{{"freq", -1}}}},
		bson.D{{"$limit", limit}},
	}

	cur, err := db.Collections.Aggregate(ctx, pipeline)
	if err != nil {
		return nil, err
	}
	defer cur.Close(ctx)

	var res []domain.Tag
	for cur.Next(ctx) {
		var doc struct {
			ID   string `bson:"_id"`
			Freq int64  `bson:"freq"`
		}
		if err := cur.Decode(&doc); err != nil {
			return nil, err
		}
		res = append(res, domain.Tag{
			Word: doc.ID,
			Freq: int64(doc.Freq),
		})
	}

	return res, nil
}
