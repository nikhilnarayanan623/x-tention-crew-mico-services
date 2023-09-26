package usecase

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/nikhilnarayanan623/x-tention-crew/pkg/domain"
	repo "github.com/nikhilnarayanan623/x-tention-crew/pkg/repository/interfaces"
	"github.com/nikhilnarayanan623/x-tention-crew/pkg/usecase/interfaces"
	"github.com/nikhilnarayanan623/x-tention-crew/pkg/utils"
	"github.com/nikhilnarayanan623/x-tention-crew/pkg/utils/models/request"
	"github.com/nikhilnarayanan623/x-tention-crew/pkg/utils/models/response"
)

const (
	cacheDuration = 2 * time.Hour
)

var ErrUserAlreadyExist = errors.New("user already exist with given credentials")

type userUseCase struct {
	userRepo  repo.UserRepo
	cacheRepo repo.CacheRepo
}

func NewAuthUseCase(ur repo.UserRepo, cr repo.CacheRepo) interfaces.UserUseCase {
	return &userUseCase{
		userRepo:  ur,
		cacheRepo: cr,
	}
}

func (u *userUseCase) CreateAccount(ctx context.Context, signUpDetails request.User) (response.User, error) {

	// first check the user already exist or not
	exist, err := u.userRepo.IsUserAlreadyExistWithThisEmail(ctx, signUpDetails.Email)
	if err != nil {
		return response.User{}, utils.PrependMessageToError(err, "failed to check user already exist in db")
	}

	if exist {
		return response.User{}, ErrUserAlreadyExist
	}

	// hash user password
	hashPass, err := utils.GenerateHashFromPassword(signUpDetails.Password)
	if err != nil {
		return response.User{}, utils.PrependMessageToError(err, "failed to hash user password")
	}

	user := domain.User{
		FirstName: signUpDetails.FirstName,
		LastName:  signUpDetails.LastName,
		Email:     signUpDetails.Email,
		Password:  hashPass,
	}
	// save user details on database
	user, err = u.userRepo.SaveUser(ctx, user)
	if err != nil {
		return response.User{}, utils.PrependMessageToError(err, "failed to save user details on db")
	}

	// save user to cache repo
	key := userIDToKey(user.ID)

	resUser := response.User{
		ID:        user.ID,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
	}

	// run save data on cache repo in another goroutine to avoid the time delay
	go u.saveDataToCacheRepo(key, resUser)

	return resUser, nil
}

// save any data to cache store by converting the data to json string(byte array)
func (u *userUseCase) saveDataToCacheRepo(key string, data interface{}) {
	jsonData, err := json.Marshal(data)
	if err != nil {
		log.Printf("failed to marshal data into json: for key %s and data: %+v\n", key, data)
		return
	}

	if err := u.cacheRepo.Set(context.Background(), key, jsonData, cacheDuration); err != nil {
		log.Println("failed to set data on cache repo: %", err)
	}
}

// convert user id of uint32 to string
func userIDToKey(userID uint32) string {
	return fmt.Sprintf("user-%d", userID)
}
