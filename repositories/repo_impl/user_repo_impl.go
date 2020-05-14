package repo_impl

import (
	"context"
	"errors"
	"golang_api/databases"
	"golang_api/models"
	"golang_api/repositories"
)

type UserRepoImpl struct {
	sql *databases.PostgresDB
}

func NewUserRepo(sql *databases.PostgresDB) repositories.UserRepo {
	return UserRepoImpl{sql:sql}
}

func (u UserRepoImpl) Add(c context.Context, user models.User) (models.User, error) {
	sqlString := `
		INSERT INTO users (id, name, email, password)
		VALUES (:id, :name, :email, :password)`

	_, err := u.sql.Db.NamedExecContext(c, sqlString, user)
	if err != nil {
		return user, err
	}
	return user, nil
}

func (u UserRepoImpl) GetList(c context.Context) ([]models.User, error) {
	sqlString := `SELECT * FROM users`

	rows, err := u.sql.Db.Query(sqlString)
	if err != nil {
		return make([]models.User, 0), err
	}

	users := make([]models.User, 0)
	for rows.Next() {
		user := models.User{}
		err := rows.Scan(&user.ID, &user.Name, &user.Email, &user.Password)
		if err != nil {
			break
		}

		users = append(users, user)
	}

	err = rows.Err()
	if err != nil {
		return make([]models.User, 0), nil
	}

	return users, nil
}

func (u UserRepoImpl) GetUserById(c context.Context, id string) (models.User, error) {
	sqlString := `SELECT * FROM users WHERE id=$1`

	user := models.User{}
	err := u.sql.Db.GetContext(c, &user, sqlString, id)
	if err != nil {
		return user, err
	}

	return user, nil
}

func (u UserRepoImpl) Update(c context.Context, user models.User) (models.User, error) {
	sqlString := `
	UPDATE users
	SET
		name = (CASE WHEN LENGTH(:name) = 0 THEN name ELSE :name END),
		email = (CASE WHEN LENGTH(:email) = 0 THEN email ELSE :email END)
	WHERE id=:id
`
	row, err := u.sql.Db.NamedExecContext(c, sqlString, user)
	if err != nil {
		return user, err
	}

	count, err := row.RowsAffected()
	if err != nil {
		return user, err
	}

	if count == 0 {
		return user, errors.New("Cannot update user")
	}

	return user, nil


	return user, nil
}

func (u UserRepoImpl) Delete(c context.Context, id string) (models.User, error) {
	//sqlString := `DELETE FROM users WHERE id:$1`
	return models.User{}, nil
}
