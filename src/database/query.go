package database

import (
	"database/sql"
	"github.com/ffmoyano/gofer/logger"
	"recipes/src/entity"
)

var err error

func GetUserByIdAndActive(id int) (entity.User, error) {
	var user entity.User
	err = db.QueryRow(
		"Select id, name, email, created_at, verified from user where id = ? and verified = 1",
		id).Scan(&user.Id, &user.Name, &user.Email, &user.CreatedAt, &user.Verified)
	if err != nil {
		logger.Error(err.Error())
		return entity.User{}, err
	}
	user.Roles, err = GetRolesByUser(user)
	if err != nil {
		logger.Error(err.Error())
		return entity.User{}, err
	}
	return user, nil
}

func GetRolesByUser(user entity.User) ([]entity.Role, error) {
	var roles []entity.Role
	var rows *sql.Rows

	rows, err = db.Query(
		"select r.id, r.name from role r inner join user_role ur on r.id = ur.role_id where ur.user_id = ?",
		user.Id)
	if err != nil {
		logger.Error(err.Error())
		return roles, err
	}
	for rows.Next() {
		var role entity.Role
		err = rows.Scan(&role.Id, &role.Name)
		if err != nil {
			logger.Error(err.Error())
			return roles, err
		}
		roles = append(roles, role)
	}

	if rows != nil {
		err = rows.Close()
		if err != nil {
			logger.Warn(err.Error())
			return roles, err
		}
	}

	return roles, nil
}

func GetUsers() ([]entity.User, error) {

	var users []entity.User
	var rows *sql.Rows

	rows, err = db.Query(
		"Select id, name, email, created_at from user where verified = 1")
	if err != nil {
		logger.Error(err.Error())
		return users, err
	}

	for rows.Next() {
		var user entity.User
		err = rows.Scan(&user.Id, &user.Name, &user.Email, &user.CreatedAt)
		if err != nil {
			logger.Error(err.Error())
			return users, err
		}

		user.Roles, err = GetRolesByUser(user)
		if err != nil {
			return users, err
		}
		users = append(users, user)
	}

	err = rows.Close()
	if err != nil {
		logger.Warn(err.Error())
	}

	return users, nil
}
