package model

type StockUpdateRequest struct {
	ID        string                   `json:"Id"`
	Publisher string                   `json:"Publisher"`
	Action    string                   `json:"Action"`
	Data      []StockUpdateRequestData `json:"Data"`
}

type StockUpdateRequestData struct {
	UserID      string `json:"user_id"`
	SkuID       int64  `json:"sku_id"`
	SKU         string `json:"sku"`
	WarehouseID string `json:"warehouse_id"`
	Stock       int64  `json:"stock"`
	UpdatedAt   string `json:"updated_at"`
}
