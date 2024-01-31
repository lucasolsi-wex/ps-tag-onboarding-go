package service

import (
	"context"
	"errors"
	"fmt"
	"github.com/lucasolsi-wex/go-crud/internal/models"
	"github.com/lucasolsi-wex/go-crud/internal/repository"
	"github.com/lucasolsi-wex/go-crud/internal/validation"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserService struct {
	Repository repository.UserRepository
}

func (us UserService) FindUserById(id string, ctx context.Context) (*models.UserResponse, *models.CustomErr) {

	convertedId, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		message := fmt.Sprintf("Coudn't convert ID: %s to ObjectID", id)
		errorMessage := models.NewBadRequestError(message)
		return nil, errorMessage
	}

	existingUser, err := us.Repository.FindUserById(convertedId, ctx)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			errorMessage := fmt.Sprintf("User not found with ID: %s", id)

			return nil, models.NewUserNotFoundError(errorMessage)
		}
		errorMessage := "Error in Find User By Id"
		return nil, models.NewInternalServerError(errorMessage)
	}

	return models.FromEntity(*existingUser), nil
}

func (us UserService) CreateUser(request models.UserRequest, ctx context.Context) (*models.UserResponse, *models.CustomErr) {

	alreadyExists, err := us.Repository.ExistsByFirstNameAndLastName(request.FirstName, request.LastName, ctx)
	if err != nil {
		return nil, models.NewInternalServerError("Error while searching for firstName and lastName")
	}

	if alreadyExists {
		return nil, validation.NewNotUniqueNameError()
	}

	userToRepo := models.NewUser(request.FirstName, request.LastName, request.Email, request.Age)
	userFromRepo, err := us.Repository.CreateUser(userToRepo, ctx)

	if err != nil {
		return nil, models.NewInternalServerError(err.Error())
	}

	return models.FromEntity(*userFromRepo), nil
}
