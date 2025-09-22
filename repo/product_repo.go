package repo

type Product struct {
	ID          int     `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	ImgUrl      string  `json:"imageUrl"`
}

type productRepo struct {
	productList []*Product
}

type ProductRepo interface {
	Create(Product) (*Product, error)
	Get(productID int) (*Product, error)
	List() ([]*Product, error)
	Update(Product) (*Product, error)
	Delete(productID int) error
}

func NewProductRepo() ProductRepo {
	repo := &productRepo{}
	generateInitialProducts(repo)
	return repo
}

func (r *productRepo) Create(p Product) (*Product, error) {
	p.ID = len(r.productList) + 1
	r.productList = append(r.productList, &p)
	return &p, nil
}

func (r *productRepo) Get(productID int) (*Product, error) {
	for _, product := range r.productList {
		if productID == product.ID {
			return product, nil
		}
	}
	return nil, nil
}

func (r *productRepo) List() ([]*Product, error) {
	return r.productList, nil
}

func (r *productRepo) Update(product Product) (*Product, error) {
	for idx, p := range r.productList {
		if p.ID == product.ID {
			r.productList[idx] = &product
			return &product, nil
		}
	}
	return nil, nil
}

func (r *productRepo) Delete(productID int) error {
	var tempList []*Product = make([]*Product, 0)

	for _, p := range r.productList {
		if p.ID != productID {
			tempList = append(tempList, p)
		}
	}
	r.productList = tempList
	return nil
}

func generateInitialProducts(r *productRepo) {
	pd1 := Product{
		ID:          1,
		Title:       "Orange",
		Description: "Orage is great, I love it",
		Price:       100,
		ImgUrl:      "https://www.dole.com/sites/default/files/media/2025-01/oranges.png",
	}

	pd2 := Product{
		ID:          2,
		Title:       "Apple",
		Description: "Apple is great, I love it",
		Price:       300,
		ImgUrl:      "https://assets.clevelandclinic.org/transform/LargeFeatureImage/cd71f4bd-81d4-45d8-a450-74df78e4477a/Apples-184940975-770x533-1_jpg",
	}

	pd3 := Product{
		ID:          3,
		Title:       "Banana",
		Description: "Banana is great, I love it",
		Price:       60,
		ImgUrl:      "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcTAiE9BZy3SyV4Xb83a3V-e8Ywq6z1bVZlGsA&s",
	}

	pd4 := Product{
		ID:          4,
		Title:       "Grapes",
		Description: "Grapes is great, I love it",
		Price:       140,
		ImgUrl:      "https://images.everydayhealth.com/images/diet-nutrition/what-are-grapes-nutrition-health-benefits-risks-alt-1440x810.jpg?sfvrsn=f9f2dae1_3",
	}

	pd5 := Product{
		ID:          5,
		Title:       "Mango",
		Description: "Mango is great, I love it",
		Price:       150,
		ImgUrl:      "https://static.vecteezy.com/system/resources/previews/026/795/003/non_2x/mango-fruit-tropical-transparent-png.png",
	}

	pd6 := Product{
		ID:          6,
		Title:       "Guava",
		Description: "Guava is great, I love it",
		Price:       140,
		ImgUrl:      "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcQBePRa9mErJGZpF7C9ZuD4RehKCcFyySQOEw&s",
	}
	r.productList = append(r.productList, &pd1)
	r.productList = append(r.productList, &pd2)
	r.productList = append(r.productList, &pd3)
	r.productList = append(r.productList, &pd4)
	r.productList = append(r.productList, &pd5)
	r.productList = append(r.productList, &pd6)
}
