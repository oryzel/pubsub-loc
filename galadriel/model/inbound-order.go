package model

type InboundOrderStatusNotificationRequest struct {
	ID        string                             `json:"id"`
	Publisher string                             `json:"publisher"`
	Action    string                             `json:"action"`
	Data      InboundOrderStatusNotificationData `json:"data"`
}

type InboundOrderStatusNotificationData struct {
	UserID int64                          `json:"user_id"`
	DocNo  string                         `json:"doc_no"`
	Status string                         `json:"status"`
	SKUS   []InboundStatusNotificationSKU `json:"skus"`
}

type InboundStatusNotificationSKU struct {
	Code   string
	Qty    int64
	Status string
}
