package  services

import (
	"../domains"
	"../utils"
)

type itemService struct {}

var (
	ItemService itemService
)

func (i *itemService) GetItem(itemId string)(*domains.Item, *utils.ApplicationError)  {
	return domains.ItemDao.GetItem(itemId)
}