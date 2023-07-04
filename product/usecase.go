package product

type ProductUsecase struct {
	repo ProductRepositoryInterface
}

func NewProductUsecase(repo ProductRepositoryInterface) *ProductUsecase {
	return &ProductUsecase{
		repo: repo,
	}
}

func (p *ProductUsecase) Execute(id int) (Product, error) {
	product, err := p.repo.GetProduct(id)
	if err != nil {
		return Product{}, err
	}
	return *product, nil
}
