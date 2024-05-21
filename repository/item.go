package repository

import (
	dbs "inventory/database"
	m "inventory/model"
	"strconv"
)

// GetItems : to get all item data from db
func GetItems(sku, name, description, orderBy string, statusArr []string) (items []*m.Item, err error) {
	db := dbs.DB

	if sku != "" {
		db = db.Where("sku = ?", sku)
	}

	if name != "" {
		db = db.Where("name like ?", "%"+name+"%")
	}

	if description != "" {
		db = db.Where("description like ?", "%"+description+"%")
	}

	if len(statusArr) > 0 {
		db = db.Where("status in ?", statusArr)
	}

	if orderBy != "" {
		db = db.Order(orderBy)
	}

	res := db.Find(&items)

	return items, res.Error
}

// GetItemById : to get detail item data from db based on id
func GetItemById(id string) (item *m.Item, err error) {
	res := dbs.DB.First(&item, id)

	return item, res.Error
}

// AddItem : to add new item data into db
func AddItem(item *m.Item) (err error) {
	res := dbs.DB.Create(&item)

	return res.Error
}

// UpdateItem : to modify item data in db
func UpdateItem(item *m.Item) (err error) {
	res := dbs.DB.Model(&item).Updates(item)

	return res.Error
}

// DeleteItem : to remove item data from db
func DeleteItem(id string) error {
	// start transaction
	db := dbs.DB.Begin()

	idInt, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return err
	}

	item := &m.Item{ID: idInt}

	// update status of item
	res := db.Model(&item).Update("status", 2)
	if res.Error != nil {
		db.Rollback()
		return res.Error
	}

	// set deleted at
	res = dbs.DB.Delete(&m.Item{}, id)
	if res.Error != nil {
		db.Rollback()
		return res.Error
	}

	// commit transaction
	db.Commit()
	return res.Error
}
