package entity

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Id       int32  `gorm:"primaryKey;auto_increment;column:product_id"`
	Name     string `gorm:"index;column:name;type:varchar(100)"`
	Price    int64  `gorm:"column:price"`
	Quantity int32  `gorm:"column:quantity"`
}

func (Product) TableName() string {
	return "tb_product"
}
