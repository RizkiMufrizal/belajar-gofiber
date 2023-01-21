package repository

import (
	"context"
	"github.com/RizkiMufrizal/belajar-gofiber/entity"
)

type ProductRepository interface {
	Insert(ctx context.Context, product entity.Product) entity.Product
	Update(ctx context.Context, product entity.Product) entity.Product
	Delete(ctx context.Context, product entity.Product)
	FindById(ctx context.Context, id int32) (entity.Product, error)
	FindAl(ctx context.Context) []entity.Product
}
