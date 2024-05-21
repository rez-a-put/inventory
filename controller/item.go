package controller

import (
	"errors"
	m "inventory/model"
	r "inventory/repository"
	"strings"
	"time"
)

// GetItems : to get data items from db based on filter
func GetItems(sku, name, description, status, orderBy string) (retData []*m.Item, err error) {
	var statusArr, orderByArr []string

	// set status filter
	if status == "" {
		statusArr = []string{"1"}
	} else {
		statusArr = strings.Split(status, ",")
	}

	// set order data
	orderByArr = strings.Split(orderBy, ";")
	orderBy = ""
	if len(orderByArr) > 1 {
		for _, v := range orderByArr {
			if string(v[0]) == "-" {
				orderBy += strings.TrimPrefix(v, "-") + " desc,"
			} else {
				orderBy += v + " asc,"
			}
		}
	}
	orderBy = strings.TrimSuffix(orderBy, ",")

	retData, err = r.GetItems(sku, name, description, orderBy, statusArr)
	if err != nil {
		return nil, err
	}

	return retData, nil
}

// AddItem : to add new item
func AddItem(item *m.Item) (err error) {
	err = r.AddItem(item)
	if err != nil {
		return err
	}

	return nil
}

// ModifyItem : to change data of specific item
func ModifyItem(id string, newItem *m.Item) (err error) {
	var (
		item *m.Item
	)

	item, err = r.GetItemById(id)
	if err != nil {
		return err
	}

	if item.Status != 1 {
		err = errors.New("item status is not active")
		return err
	}

	item.Name = newItem.Name
	item.Description = newItem.Description
	item.UnitPrice = newItem.UnitPrice

	err = r.UpdateItem(item)
	if err != nil {
		return err
	}

	return nil
}

// RemoveItem : to delete existing item. deleted item will be flagged by status value 2 and will have value in deleted_at column
func RemoveItem(id string) (err error) {
	var (
		item *m.Item
	)

	item, err = r.GetItemById(id)
	if err != nil {
		return err
	}

	if item.Status != 1 {
		err = errors.New("item status is not active")
		return err
	}

	currTime := time.Now()
	item.Status = 2
	item.DeletedAt = &currTime

	err = r.UpdateItem(item)
	if err != nil {
		return err
	}

	return nil
}
