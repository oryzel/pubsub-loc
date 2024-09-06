package model

type SyncOrderMessage struct {
	SyncOrderData SyncOrderData `json:"Data"`
}

type SyncOrderData struct {
	Shop        Shop     `json:"shop"`
	OrderSNList []string `json:"order_sn_list"`
}
