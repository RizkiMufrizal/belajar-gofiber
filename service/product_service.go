package service

import (
	"context"
	"github.com/RizkiMufrizal/belajar-gofiber/model"
)

type ProductService interface {
	Create(ctx context.Context, model model.ProductCreateOrUpdateModel) model.ProductCreateOrUpdateModel
	Update(ctx context.Context, productModel model.ProductCreateOrUpdateModel, id int32) model.ProductCreateOrUpdateModel
	Delete(ctx context.Context, id int32)
	FindById(ctx context.Context, id int32) model.ProductModel
	FindAll(ctx context.Context) []model.ProductModel
}
