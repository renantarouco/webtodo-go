package main

import (
	"errors"
	"time"
)

var data = map[int]Item{}

func RetrieveItems() (map[int]Item, error) {
	return data, nil
}

func CreateItem(title, deadlineStr string) (Item, error) {
	deadline, err := time.Parse("2006-01-02", deadlineStr)
	if err != nil {
		return Item{}, err
	}
	id := len(data)
	item := Item{id, title, deadline}
	data[id] = item
	return item, nil
}

func RetrieveItem(id int) (Item, error) {
	item, ok := data[id]
	if !ok {
		return Item{}, errors.New("item not found")
	}
	return item, nil
}

func UpdateItem(id int, item Item) (Item, error) {
	oldItem, ok := data[id]
	if !ok {
		return Item{}, errors.New("item not found")
	}
	newItem := Item{
		ID:       oldItem.ID,
		Title:    item.Title,
		Deadline: item.Deadline,
	}
	data[id] = newItem
	return item, nil
}

func DeleteItem(id int) error {
	delete(data, id)
	return nil
}
