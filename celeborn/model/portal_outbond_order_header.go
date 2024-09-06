package model

type PortalOutboundOrderHeaderRequest struct {
	ID        string                        `json:"id"`
	Publisher string                        `json:"publisher"`
	Action    string                        `json:"action"`
	Body      PortalOutboundOrderHeaderData `json:"data"`
}

type PortalOutboundOrderHeaderData struct {
	DocNo       string `json:"doc_no"`
	OldStatus   string `json:"old_status"`
	Status      string `json:"status"`
	OrderNumber string `json:"order_number"`
	OrderSource string `json:"order_source"`
}
