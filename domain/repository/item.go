package repository

import (
	"awesomeProject/domain/model"
)

// ItemRepository : Item における Repository のインターフェース
type ItemRepository interface {
	FindAll() (items []*model.Item, err error)
	Create(item *model.Item) (err error)
	Update(item *model.Item) (err error)
}
