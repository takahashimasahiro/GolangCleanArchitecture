package database

import (
	"../../domain"
	"log"
)

type userRepository struct {
	db ConnectedDB
}

type UserRepository interface {
	Store(user domain.User) error
	FindByAuthToken(authToken string) (*domain.User, error)
	FindByUserID(userID string) (*domain.User, error)
	UpdateByUserID(userID string, name string) error
}

func NewUserRepository(db ConnectedDB) UserRepository {
	return &userRepository{db}
}

// DBにUserを登録
func (userRepository *userRepository) Store(user domain.User) error {
	_, err := userRepository.db.Exec("INSERT INTO user(user_id, auth_token, name) VALUES (?, ?, ?)", user.UserID, user.AuthToken, user.Name)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

// AuthTokenを条件にレコード取得
func (userRepository *userRepository) FindByAuthToken(authToken string) (*domain.User, error) {
	row := userRepository.db.QueryRow("SELECT * FROM user WHERE auth_token=?", authToken)
	return ConvertToUser(row)
}

// UserIDを条件にレコードを取得
func (userRepository *userRepository) FindByUserID(userID string) (*domain.User, error) {
	row := userRepository.db.QueryRow("SELECT * FROM user WHERE user_id=?", userID)
	return ConvertToUser(row)
}

// UserIDを条件にレコードを更新する
func (userRepository *userRepository) UpdateByUserID(userID string, name string) error {
	_, err := userRepository.db.Exec("UPDATE user SET name=? WHERE user_id=?", name, userID)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

// rowをUserへ変換
func ConvertToUser(row Row) (*domain.User, error) {
	user := domain.User{}
	err := row.Scan(&user.UserID, &user.AuthToken, &user.Name)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return &user, nil
}