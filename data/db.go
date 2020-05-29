package data

// Products holds multiple product
type Products []*Product

var productList = Products{
	{
		ID:          1,
		Name:        "Latte",
		Description: "Frothy milky coffee",
		Price:       2.45,
		SKU:         "abc323",
	},
	{
		ID:          2,
		Name:        "Espresso",
		Description: "Short and strong coffee without milk",
		Price:       1.99,
		SKU:         "fjd34",
	},
}

// GetProducts returns slice of products
func GetProducts() Products {
	return productList
}

// GetProduct returns the product with given id
func GetProduct(id int) (*Product, error) {
	i := findIndex(id)

	if i == -1 {
		return nil, ErrProductNotFound
	}

	return productList[i], nil
}

// AddProduct adds a product to product list
func AddProduct(p *Product) {
	p.ID = getNextID()
	productList = append(productList, p)
}

// UpdateProduct updates the product to given product
func UpdateProduct(p *Product) error {
	i := findIndex(p.ID)

	if i == -1 {
		return ErrProductNotFound
	}

	productList[i] = p
	return nil
}

// DeleteProduct deletes a product with given id from db
func DeleteProduct(id int) error {
	i := findIndex(id)

	if i == -1 {
		return ErrProductNotFound
	}

	productList = append(productList[:i], productList[i+1:]...)
	return nil
}

func findIndex(id int) int {
	for i, p := range productList {
		if id == p.ID {
			return i
		}
	}
	return -1
}

func getNextID() int {
	lastProduct := productList[len(productList)-1]
	return lastProduct.ID + 1
}
