package usecase

import (
	"errors"
	"rest-api-go/domain/request"
	"rest-api-go/domain/response"
	"rest-api-go/models"
	"rest-api-go/repositories"
	"time"
)

type ProductUsecase interface {
	ProductList() ([]response.ProductListResponse, error)
	ProductDetail(productId int64) (*response.ProductDetailResponse, error)
	ProductCreate(req request.ProductCreateReq, userId int64) (*response.ProductDetailResponse, error)
	ProductUpdate(req request.ProductUpdateReq, userId int64) (*response.ProductDetailResponse, error)
	ProductDelete(productId, userId int64) error
}

type ProductUsecaseInteractor struct {
	ProductRepository repositories.ProductRepository
}

func ProductUsecaseImpl(
	productRepository repositories.ProductRepository,
) ProductUsecase {
	return &ProductUsecaseInteractor{
		ProductRepository: productRepository,
	}
}

func (u *ProductUsecaseInteractor) ProductList() ([]response.ProductListResponse, error) {
	products, err := u.ProductRepository.GetList()
	if err != nil {
		return nil, errors.New("Failed get list.")
	}

	var productsList []response.ProductListResponse
	for _, p := range products {
		productsList = append(productsList, response.ProductListResponse{
			ID:          p.ID,
			Name:        p.Name,
			Description: p.Description,
			Price:       p.Price,
			ImageURL:    p.ImageURL,
			CreatedAt:   p.CreatedAt,
			CreatedBy:   p.CreatedBy,
		})
	}

	return productsList, nil
}

func (u *ProductUsecaseInteractor) ProductDetail(productId int64) (*response.ProductDetailResponse, error) {
	product, err := u.ProductRepository.GetByID(productId)
	if err != nil {
		return nil, errors.New("Faied get product by ID")
	}

	var result response.ProductDetailResponse = response.ProductDetailResponse{
		ID:          product.ID,
		Name:        product.Name,
		Description: product.Description,
		Price:       product.Price,
		ImageURL:    product.ImageURL,
		CreatedAt:   product.CreatedAt,
		CreatedBy:   product.CreatedBy,
	}

	return &result, nil
}

func (u *ProductUsecaseInteractor) ProductCreate(req request.ProductCreateReq, userId int64) (*response.ProductDetailResponse, error) {
	product := models.Product{
		Name:        req.Name,
		Description: req.Description,
		Price:       req.Price,
		ImageURL:    req.ImageURL,
		CreatedAt:   time.Now(),
		CreatedBy:   userId,
	}

	data, err := u.ProductRepository.Create(&product)
	if err != nil {
		return nil, errors.New("Failed create new product.")
	}

	return &response.ProductDetailResponse{
		ID:        data.ID,
		Name:      data.Name,
		Price:     data.Price,
		ImageURL:  data.ImageURL,
		CreatedAt: data.CreatedAt,
		CreatedBy: data.CreatedBy,
	}, nil
}

func (u *ProductUsecaseInteractor) ProductUpdate(req request.ProductUpdateReq, userId int64) (*response.ProductDetailResponse, error) {
	product, err := u.ProductRepository.GetByID(req.ID)
	if err != nil {
		return nil, errors.New("Failed fetch product by ID.")
	}

	if product.CreatedBy != userId {
		return nil, errors.New("Unauthorized")
	}

	product.Name = req.Name
	product.Description = req.Description
	product.Price = req.Price
	product.ImageURL = req.ImageURL

	data, err := u.ProductRepository.Update(product)
	if err != nil {
		return nil, errors.New("Failed update product.")
	}

	return &response.ProductDetailResponse{
		ID:        data.ID,
		Name:      data.Name,
		Price:     data.Price,
		ImageURL:  data.ImageURL,
		CreatedAt: data.CreatedAt,
		CreatedBy: data.CreatedBy,
	}, nil
}

func (u *ProductUsecaseInteractor) ProductDelete(productId, userId int64) error {
	product, err := u.ProductRepository.GetByID(productId)
	if err != nil {
		return errors.New("Failed get product by ID")
	}

	if product.CreatedBy != userId {
		return errors.New("Unauthorized")
	}

	err = u.ProductRepository.Delete(productId)
	if err != nil {
		return errors.New("Failed delete product.")
	}

	return err
}
