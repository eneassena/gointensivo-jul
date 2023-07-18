package entity

type OrderRepositoryInterface interface {
	Save(o *Order) error
	GetTotalTransactions() (int, error)
}
