package repo

import (
	"database/sql"

	"github.com/jmoiron/sqlx"
)

type Product struct {
	ID          int     `json:"id" db:"id"`
	Title       string  `json:"title" db:"title"`
	Description string  `json:"description" db:"description"`
	Price       float64 `json:"price" db:"price"`
	ImgUrl      string  `json:"imageUrl" db:"img_url"`
}

type productRepo struct {
	db *sqlx.DB
}

type ProductRepo interface {
	Create(Product) (*Product, error)
	Get(int) (*Product, error)
	List() ([]*Product, error)
	Update(int, Product) (*Product, error)
	Delete(productID int) error
}

func NewProductRepo(db *sqlx.DB) ProductRepo {
	return &productRepo{
		db: db,
	}
}

func (r *productRepo) Create(p Product) (*Product, error) {
	query := `
	INSERT INTO products (title, description,price,img_url)
	VALUES($1,$2,$3,$4)
	RETURNING id
	`

	row := r.db.QueryRow(query, p.Title, p.Description, p.Price, p.ImgUrl)
	err := row.Scan(&p.ID)
	if err != nil {
		return nil, err
	}
	return &p, nil
}

func (r *productRepo) Get(productID int) (*Product, error) {
	var prd Product

	query := `
	SELECT id, title, description, price, img_url FROM products 
	WHERE id=$1
	`
	err := r.db.Get(&prd, query, productID)
	if err != nil {
		if err == sql.ErrNoRows || err.Error() == "sql: no rows in result set" {
			return nil, nil // product not found
		}
		return nil, err
	}
	return &prd, nil
}

func (r *productRepo) List() ([]*Product, error) {
	var prdList []*Product
	query := `
	SELECT id, title, description, price, img_url FROM products
	`
	err := r.db.Select(&prdList, query)
	if err != nil {
		if err == sql.ErrNoRows || err.Error() == "sql: no rows in result set" {
			return nil, nil // product not found
		}
		return nil, err
	}

	return prdList, nil
}

func (r *productRepo) Update(pid int, p Product) (*Product, error) {

	query := `
	UPDATE products SET
	title=$1, description=$2, price=$3, img_url=$4
	WHERE id=$5
	`

	row := r.db.QueryRow(query, p.Title, p.Description, p.Price, p.ImgUrl, pid)

	err := row.Err()

	if err != nil {
		println("Error from update", err.Error())
		return nil, err
	}
	p.ID = pid
	return &p, nil

}

func (r *productRepo) Delete(productID int) error {
	query := `
	DELETE FROM products
	WHERE id=$1
	`
	_, err := r.db.Exec(query, productID)
	if err != nil {
		return err
	}
	return nil
}
