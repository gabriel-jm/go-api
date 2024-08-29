package repositories

import (
	"database/sql"
	"fmt"
	"go-api/models"
	"log"
)

type ProductRepository struct {
	conn *sql.DB
}

func NewProductRepository(conn *sql.DB) ProductRepository {
	return ProductRepository{conn}
}

func (r *ProductRepository) GetProducts() ([]models.Product, error) {
	query := "select * from products"
	rows, err := r.conn.Query(query)

	if err != nil {
		log.Fatal(err)
		return []models.Product{}, err
	}

	var productList []models.Product
	var product models.Product

	for rows.Next() {
		err = rows.Scan(
			&product.Id,
			&product.Name,
			&product.Price,
		)

		if err != nil {
			log.Fatal(err)
			return []models.Product{}, err
		}

		productList = append(productList, product)
	}

	rows.Close()

	return productList, nil
}

func (r *ProductRepository) CreateProduct(product models.Product) (int, error) {
	query, err := r.conn.Prepare("insert into products values (default, $1, $2) returning id")

	if err != nil {
		fmt.Println(err)
		return 0, err
	}

	var id int
	err = query.QueryRow(product.Name, product.Price).Scan(&id)

	if err != nil {
		fmt.Println(err)
		return 0, err
	}

	query.Close()

	return id, nil
}

func (r *ProductRepository) GetProductById(id int) (*models.Product, error) {
	query, err := r.conn.Prepare("select * from products where id = $1")

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	var product models.Product
	err = query.QueryRow(id).Scan(&product.Id, &product.Name, &product.Price)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}

		return nil, err
	}

	query.Close()

	return &product, nil
}
