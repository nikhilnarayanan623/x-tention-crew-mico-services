package repository

import (
	"context"
	"time"

	"github.com/nikhilnarayanan623/x-tention-crew/pkg/domain"
	"github.com/nikhilnarayanan623/x-tention-crew/pkg/repository/interfaces"
	"gorm.io/gorm"
)

type userDB struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) interfaces.UserRepo {

	return &userDB{
		db: db,
	}
}
func (u *userDB) IsUserAlreadyExistWithThisEmail(ctx context.Context, email string) (exist bool, err error) {

	query := `SELECT EXISTS( SELECT 1 FROM users WHERE email = $1 )`

	err = u.db.Raw(query, email).Scan(&exist).Error

	return
}

func (u *userDB) SaveUser(ctx context.Context, user domain.User) (userID uint32, err error) {

	query := `INSERT INTO users (first_name, last_name, email, password, created_at) 
	VALUES($1, $2, $3, $4, $5) RETURNING id AS user_id`

	createdAt := time.Now()
	err = u.db.Raw(query, user.FirstName, user.LastName, user.Email, user.Password, createdAt).
		Scan(&userID).Error

	return
}

func (u *userDB) FindUserByID(ctx context.Context, id uint) (user domain.User, err error) {

	query := `SELECT first_name, last_name, phone, password, created_at, updated_at 
	FROM users WHERE id = $1`

	err = u.db.Raw(query, id).Scan(&user).Error

	return
}
