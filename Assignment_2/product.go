package main

type Product struct {
	ID            int64
	Name          string
	Price         float64
	Rating        float64
	SizeOfRatings int
}

var Products = []Product{
	Product{1, "Apple Iphone 14 Pro", 700000, 0, 0},
	Product{2, "Samsung Galaxy A71", 210000, 0, 0},
	Product{3, "Apple Iphone 11 SlimBox", 264000, 0, 0},
	Product{4, "Oppo Reno 7", 150000, 0, 0},
	Product{5, "Samsung Galaxy S22 Ultra", 600000, 0, 0},
}

func GetListOfProducts() []Product {
	return Products
}
