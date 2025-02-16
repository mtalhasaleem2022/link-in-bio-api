package repositories

import (
	"context"
	"link-in-bio-api/internal/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// LinkRepository implements the LinkRepositoryInterface.
type LinkRepository struct {
	client *mongo.Client
	db     *mongo.Database
}

func NewLinkRepository(uri string) *LinkRepository {
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}
	return &LinkRepository{
		client: client,
		db:     client.Database("linkinbio"),
	}
}

func (r *LinkRepository) CreateLink(ctx context.Context, link *models.Link) error {
	collection := r.db.Collection("links")
	_, err := collection.InsertOne(ctx, link)
	return err
}

func (r *LinkRepository) UpdateLink(ctx context.Context, link *models.Link) error {
	collection := r.db.Collection("links")
	filter := bson.M{"_id": link.ID}
	update := bson.M{"$set": link}
	_, err := collection.UpdateOne(ctx, filter, update)
	return err
}

func (r *LinkRepository) DeleteLink(ctx context.Context, id string) error {
	collection := r.db.Collection("links")
	_, err := collection.DeleteOne(ctx, bson.M{"_id": id})
	return err
}

func (r *LinkRepository) GetLink(ctx context.Context, id string) (*models.Link, error) {
	collection := r.db.Collection("links")
	var link models.Link
	err := collection.FindOne(ctx, bson.M{"_id": id}).Decode(&link)
	if err != nil {
		return nil, err
	}
	return &link, nil
}

func (r *LinkRepository) IncrementClicks(ctx context.Context, linkID string, ipAddress string) error {

	// Increment click count
	linkCollection := r.db.Collection("links")
	_, err := linkCollection.UpdateOne(
		ctx,
		bson.M{"_id": linkID},
		bson.M{"$inc": bson.M{"clicks": 1}},
	)
	if err != nil {
		return err
	}

	// Log visit
	visitCollection := r.db.Collection("visits")
	visit := models.Visit{
		LinkID:    linkID,
		IPAddress: ipAddress,
		Timestamp: time.Now(),
	}
	_, err = visitCollection.InsertOne(ctx, visit)
	return err
}
