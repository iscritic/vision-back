package service

type Product struct {
	ID    int     `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

type Service struct{}

func NewService() *Service {
	return &Service{}
}

func (s *Service) GetProducts() []Product {
	return []Product{
		{ID: 1, Name: "Product 1", Price: 10.0},
		{ID: 2, Name: "Product 2", Price: 20.0},
	}
}
