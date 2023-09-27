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

func (u *userDB) SaveUser(ctx context.Context, user domain.User) (domain.User, error) {

	query := `INSERT INTO users (first_name, last_name, email, password, created_at) 
	VALUES($1, $2, $3, $4, $5) RETURNING *`

	createdAt := time.Now()
	err := u.db.Raw(query, user.FirstName, user.LastName, user.Email, user.Password, createdAt).Scan(&user).Error

	return user, err
}

func (u *userDB) FindUserByID(ctx context.Context, id uint32) (user domain.User, err error) {

	query := `SELECT id, first_name, last_name, email, password, created_at, updated_at 
	FROM users WHERE id = $1`

	err = u.db.Raw(query, id).Scan(&user).Error

	return
}

func (u *userDB) UpdateUser(ctx context.Context, user domain.User) (domain.User, error) {

	query := `UPDATE users SET first_name = $1, last_name = $2 , email = $3 , password = $4, updated_at = $5  
	WHERE id = $6 RETURNING *`

	updatedAt := time.Now()
	err := u.db.Raw(query, user.FirstName, user.LastName, user.Email,
		user.Password, updatedAt, user.ID).Scan(&user).Error

	return user, err
}
