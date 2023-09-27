package usecase

import (
	"context"
	"time"

	clientinterface "github.com/nikhilnarayanan623/x-tention-crew/service2/pkg/client/interfaces"
	"github.com/nikhilnarayanan623/x-tention-crew/service2/pkg/usecase/interfaces"
	"github.com/nikhilnarayanan623/x-tention-crew/service2/pkg/utils/models"
)

type UseCase struct {
	userClient clientinterface.UserServiceClient
	block      chan struct{}
}

func NewUseCase(client clientinterface.UserServiceClient) interfaces.UseCase {
	return &UseCase{
		userClient: client,
		block:      make(chan struct{}, 1), // 1 sized chan to block the function call to only for one
	}
}

// call to this function will get the users data sequentially
func (c *UseCase) GetSequentially(waitSecond int) (models.AllUserDetails, error) {
	// block the channel
	c.block <- struct{}{}
	// always execute this function at then end
	defer func() {
		// wait for the given waitSecond and
		time.Sleep(time.Duration(waitSecond) * time.Second)
		//free up the channel after the whole function execution
		<-c.block
	}()

	// get the data from client
	allUsers, err := c.userClient.GetAllUsers(context.Background())
	if err != nil {
		return models.AllUserDetails{}, err
	}

	return allUsers, nil
}

// call to thi function will get the data parallel(actually in here it's concurrently)
func (c *UseCase) GetParallel(waitSecond int) (models.AllUserDetails, error) {
	// wait for the wait time at the end of this function exec
	defer func() {
		time.Sleep(time.Duration(waitSecond) * time.Second)
	}()

	// get the data from client
	allUsers, err := c.userClient.GetAllUsers(context.Background())
	if err != nil {
		return models.AllUserDetails{}, err
	}

	return allUsers, nil
}
