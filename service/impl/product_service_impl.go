package impl

import (
	"context"
	"github.com/RizkiMufrizal/belajar-gofiber/entity"
	"github.com/RizkiMufrizal/belajar-gofiber/exception"
	"github.com/RizkiMufrizal/belajar-gofiber/model"
	"github.com/RizkiMufrizal/belajar-gofiber/repository"
	"github.com/RizkiMufrizal/belajar-gofiber/service"
)

func NewProductServiceImpl(productRepository *repository.ProductRepository) service.ProductService {
	return &productServiceImpl{ProductRepository: *productRepository}
}

type productServiceImpl struct {
	repository.ProductRepository
}

func (service *productServiceImpl) Create(ctx context.Context, productModel model.ProductCreateOrUpdateModel) model.ProductCreateOrUpdateModel {
	product := entity.Product{
		Name:     productModel.Name,
		Price:    productModel.Price,
		Quantity: productModel.Quantity,
	}
	service.ProductRepository.Insert(ctx, product)
	return productModel
}

func (service *productServiceImpl) Update(ctx context.Context, productModel model.ProductCreateOrUpdateModel, id int32) model.ProductCreateOrUpdateModel {
	product := entity.Product{
		Id:       id,
		Name:     productModel.Name,
		Price:    productModel.Price,
		Quantity: productModel.Quantity,
	}
	service.ProductRepository.Update(ctx, product)
	return productModel
}

func (service *productServiceImpl) Delete(ctx context.Context, id int32) {
	product, err := service.ProductRepository.FindById(ctx, id)
	if err != nil {
		panic(exception.NotFoundError{
			Message: err.Error(),
		})
	}
	service.ProductRepository.Delete(ctx, product)
}

func (service *productServiceImpl) FindById(ctx context.Context, id int32) model.ProductModel {
	productCache, err := service.ProductRepository.FindById(ctx, id)
	exception.PanicLogging(err)

	return model.ProductModel{
		Id:       productCache.Id,
		Name:     productCache.Name,
		Price:    productCache.Price,
		Quantity: productCache.Quantity,
	}
}

func (service *productServiceImpl) FindAll(ctx context.Context) (responses []model.ProductModel) {
	products := service.ProductRepository.FindAl(ctx)
	for _, product := range products {
		responses = append(responses, model.ProductModel{
			Id:       product.Id,
			Name:     product.Name,
			Price:    product.Price,
			Quantity: product.Quantity,
		})
	}
	if len(products) == 0 {
		return []model.ProductModel{}
	}
	return responses
}
