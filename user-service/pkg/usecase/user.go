package usecase

import (
	"context"
	"errors"

	"github.com/nikhilnarayanan623/x-tention-crew/pkg/domain"
	repo "github.com/nikhilnarayanan623/x-tention-crew/pkg/repository/interfaces"
	"github.com/nikhilnarayanan623/x-tention-crew/pkg/usecase/interfaces"
	"github.com/nikhilnarayanan623/x-tention-crew/pkg/utils"
	"github.com/nikhilnarayanan623/x-tention-crew/pkg/utils/models/request"
)

type userUseCase struct {
	userRepo repo.UserRepo
}

var ErrUserAlreadyExist = errors.New("user already exist with given credentials")

func NewAuthUseCase(ur repo.UserRepo) interfaces.UserUseCase {
	return &userUseCase{
		userRepo: ur,
	}
}

func (u *userUseCase) CreateAccount(ctx context.Context, signUpDetails request.User) (uint32, error) {

	// first check the user already exist or not
	exist, err := u.userRepo.IsUserAlreadyExistWithThisEmail(ctx, signUpDetails.Email)
	if err != nil {
		return 0, utils.PrependMessageToError(err, "failed to check user already exist in db")
	}

	if exist {
		return 0, ErrUserAlreadyExist
	}

	// hash user password
	hashPass, err := utils.GenerateHashFromPassword(signUpDetails.Password)
	if err != nil {
		return 0, utils.PrependMessageToError(err, "failed to hash user password")
	}

	dbUser := domain.User{
		FirstName: signUpDetails.FirstName,
		LastName:  signUpDetails.LastName,
		Email:     signUpDetails.Email,
		Password:  hashPass,
	}
	// save user details on database
	userID, err := u.userRepo.SaveUser(ctx, dbUser)
	if err != nil {
		return 0, utils.PrependMessageToError(err, "failed to save user details on db")
	}

	return userID, nil
}
