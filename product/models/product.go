package models

type Product struct {
	Id                int64            `json:"id" xorm:"pk autoincr"`
	Name              string           `json:"name"`
	Stock             int64            `json:"stock"`
	Price             int64            `json:"price"`
	StockDecreaseLogs StockDecreaseLog `xorm:"-"`
}
