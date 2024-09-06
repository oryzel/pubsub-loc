package model

type Shop struct {
	Id                int64  `json:"id"`
	UserId            int64  `json:"user_id"`
	WmsCustomerId     string `json:"wms_customer_id"`
	Username          string `json:"username"`
	MarketplaceShopId int64  `json:"marketplace_shop_id"`
	PortalShopId      int64  `json:"portal_shop_id"`
	AccessToken       string `json:"access_token"`
	RefreshToken      string `json:"refresh_token"`
	TokenExpiryTime   string `json:"token_expiry_time"`
	WarehouseId       string `json:"warehouse_id"`
	WarehouseName     string `json:"warehouse_name"`
	Name              string `json:"name"`
	Description       string `json:"description"`
	Address           string `json:"address"`
	Province          string `json:"province"`
	City              string `json:"city"`
	District          string `json:"district"`
}
