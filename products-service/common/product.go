package common

type Product struct {
	ID          string  `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
}

func ListProducts() []Product {
	p := []Product{
		{ID: "498076ac-7b2f-47b3-b7b6-719f700a3ce4", Title: "product-1", Description: "product-description-1", Price: 18},
		{ID: "a3a1bf66-6f08-414c-812b-37e0e8be7ae9", Title: "product-2", Description: "product-description-2", Price: 115},
		{ID: "73b871ba-0b6d-4b0f-96f4-52a5a86659cb", Title: "product-3", Description: "product-description-3", Price: 28},
		{ID: "f334aa4c-8777-44f5-91df-8b0c9b567b02", Title: "product-4", Description: "product-description-4", Price: 56},
		{ID: "a60a6e3b-3af5-4134-90f0-856892f721b5", Title: "product-5", Description: "product-description-5", Price: 104},
		{ID: "13234fcf-837d-4aa4-8c0e-9b8f5ec1a07b", Title: "product-6", Description: "product-description-6", Price: 57},
		{ID: "3d8a8b2b-f8f7-4a5e-84d6-237b7d54aff5", Title: "product-7", Description: "product-description-7", Price: 59},
		{ID: "1d667f2b-d655-44aa-945e-f9ac331927f1", Title: "product-8", Description: "product-description-8", Price: 17},
		{ID: "dc3fc9be-70d0-4014-9b2e-7fd2440c06da", Title: "product-9", Description: "product-description-9", Price: 130},
		{ID: "8c0576dd-7ebe-412b-a4e7-d3a89d3f2d61", Title: "product-10", Description: "product-description-10", Price: 29},
	}

	return p
}
