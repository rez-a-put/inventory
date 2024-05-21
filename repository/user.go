package repository

import (
	dbs "inventory/database"
	m "inventory/model"
)

// GetUsers : to get all user data from db
func GetUsers(name, email, orderBy string, statusArr []string) (users []*m.User, err error) {
	db := dbs.DB

	if name != "" {
		db = db.Where("name = ?", name)
	}

	if email != "" {
		db = db.Where("email = ?", email)
	}

	if len(statusArr) > 0 {
		db = db.Where("status in ?", statusArr)
	}

	if orderBy != "" {
		db = db.Order(orderBy)
	}

	res := db.Find(&users)

	return users, res.Error
}

// GetUserById : to get detail user data from db based on id
func GetUserById(id string) (user *m.User, err error) {
	res := dbs.DB.First(&user, id)

	return user, res.Error
}

// AddUser : to add new user data into db
func AddUser(user *m.User) (err error) {
	res := dbs.DB.Create(&user)

	return res.Error
}

// UpdateUser : to modify user data in db
func UpdateUser(user *m.User) (err error) {
	res := dbs.DB.Model(&user).Updates(user)

	return res.Error
}

// DeleteUser : to remove user data from db
func DeleteUser(id string) (err error) {
	res := dbs.DB.Delete(&m.User{}, id)

	return res.Error
}
