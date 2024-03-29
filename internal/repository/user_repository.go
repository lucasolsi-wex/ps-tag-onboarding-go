package repository

import (
	"context"
	"github.com/lucasolsi-wex/go-crud/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	MongoDBUserDb = "users"
)

func NewUserRepository(database *mongo.Database) UserRepository {
	return &userRepository{database}
}

type userRepository struct {
	databaseConnection *mongo.Database
}

func (userRepo *userRepository) ExistsByFirstNameAndLastName(firstName, lastName string, ctx context.Context) (bool, error) {
	collection := userRepo.databaseConnection.Collection(MongoDBUserDb)

	cursor, err := collection.Find(ctx, bson.M{"firstName": firstName, "lastName": lastName})
	if err != nil {
		return false, err
	}
	defer cursor.Close(ctx)

	return cursor.Next(ctx), nil
}

func (userRepo *userRepository) FindUserById(id primitive.ObjectID, ctx context.Context) (*models.UserModel, error) {
	collection := userRepo.databaseConnection.Collection(MongoDBUserDb)

	existingUser := &models.UserModel{}

	filter := bson.D{{Key: "_id", Value: id}}
	err := collection.FindOne(ctx, filter).Decode(existingUser)

	if err != nil {
		return nil, err
	}

	return existingUser, err
}

func (userRepo *userRepository) CreateUser(entity models.UserModel, ctx context.Context) (*models.UserModel, error) {
	collection := userRepo.databaseConnection.Collection(MongoDBUserDb)

	result, err := collection.InsertOne(ctx, entity)
	if err != nil {
		return nil, err
	}
	entity.Id = result.InsertedID.(primitive.ObjectID)

	return &entity, err
}

type UserRepository interface {
	CreateUser(request models.UserModel, ctx context.Context) (*models.UserModel, error)
	FindUserById(id primitive.ObjectID, ctx context.Context) (*models.UserModel, error)
	ExistsByFirstNameAndLastName(firstName, lastName string, ctx context.Context) (bool, error)
}
