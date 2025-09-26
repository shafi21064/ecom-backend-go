package product

import "e-com/domain"

type service struct {
	productRepo ProductRepo
}

func NewService(productRepo ProductRepo) Service {
	return &service{
		productRepo: productRepo,
	}
}

func (svc *service) Create(product domain.Product) (*domain.Product, error) {

	usr, err := svc.productRepo.Create(product)
	if err != nil {
		return nil, err
	}
	if usr == nil {
		return nil, nil
	}

	return usr, nil
}

func (svc *service) Get(productId int) (*domain.Product, error) {

	prd, err := svc.productRepo.Get(productId)
	if err != nil {
		return nil, err
	}
	if prd == nil {
		return nil, nil
	}

	return prd, nil
}

func (svc *service) List() ([]*domain.Product, error) {
	prd, err := svc.productRepo.List()
	if err != nil {
		return nil, err
	}
	if prd == nil {
		return nil, nil
	}

	return prd, nil
}

func (svc *service) Update(productID int, product domain.Product) (*domain.Product, error) {
	prd, err := svc.productRepo.Update(productID, product)
	if err != nil {
		return nil, err
	}
	if prd == nil {
		return nil, nil
	}

	return prd, nil

}

func (svc *service) Delete(productID int) error {
	err := svc.productRepo.Delete(productID)
	if err != nil {
		return err
	}
	return nil
}
