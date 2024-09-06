package model

type SalesOrderStatusNotificationRequest struct {
	ID        string                           `json:"Id"`
	Publisher string                           `json:"Publisher"`
	Action    string                           `json:"Action"`
	Body      SalesOrderStatusNotificationData `json:"Data"`
}

type SalesOrderStatusNotificationData struct {
	ShopID      int64  `json:"shop_id"`
	DocNo       string `json:"doc_no"`
	OldStatus   string `json:"old_status"`
	Status      string `json:"status"`
	OrderSource string `json:"order_source"`
	OrderNumber string `json:"order_number"`
}
