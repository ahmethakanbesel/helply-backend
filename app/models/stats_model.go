package models

type Stats struct {
	WidgetData WidgetData `json:"widget_data"`
}

type WidgetData struct {
	TicketCount  int64 `json:"ticket_count"`
	UserCount    int64 `json:"user_count"`
	ProductCount int64 `json:"product_count"`
	LicenseCount int64 `json:"license_count"`
}
