package repo

import (
	"database/sql"
	"e-com/domain"
	"e-com/product"

	"github.com/jmoiron/sqlx"
)

type productRepo struct {
	db *sqlx.DB
}

type ProductRepo interface {
	product.ProductRepo
}

func NewProductRepo(db *sqlx.DB) ProductRepo {
	return &productRepo{
		db: db,
	}
}

func (r *productRepo) Create(p domain.Product) (*domain.Product, error) {
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

func (r *productRepo) Get(productID int) (*domain.Product, error) {
	var prd domain.Product

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

func (r *productRepo) List(page, limit int) ([]*domain.Product, error) {

	offset := ((page - 1) * limit) + 1

	var prdList []*domain.Product

	query := `
	SELECT id, title, description, price, img_url 
	FROM products
	ORDER BY id
	LIMIT $1 OFFSET $2
	`
	err := r.db.Select(&prdList, query, limit, offset)
	if err != nil {
		if err == sql.ErrNoRows || err.Error() == "sql: no rows in result set" {
			return nil, nil // product not found
		}
		return nil, err
	}

	return prdList, nil
}

func (r *productRepo) Count() (int, error) {
	var count int

	query := `SELECT COUNT(*) FROM products`

	err := r.db.Get(&count, query)
	if err != nil {
		return 0, err
	}

	return count, nil
}

func (r *productRepo) Update(pid int, p domain.Product) (*domain.Product, error) {

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
