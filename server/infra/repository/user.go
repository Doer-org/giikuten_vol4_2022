package repository

import (
	"database/sql"

	"github.com/Doer-org/geekten_vol4_2022/domain/entity"
	"github.com/Doer-org/geekten_vol4_2022/domain/repository"
	_ "github.com/lib/pq"
)

type userRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) repository.UserRepository {
	return &userRepository{
		db: db,
	}
}

func (ur userRepository) CreateUser(id string, name string) (*entity.User, error) {
	statement := "INSERT INTO users VALUES($1,$2) returning id, name"

	stmt, err := ur.db.Prepare(statement)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	user := &entity.User{}
	err = stmt.QueryRow(id, name).Scan(&user.Id, &user.Name)

	return user, nil
}
