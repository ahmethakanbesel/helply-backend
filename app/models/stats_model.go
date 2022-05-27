package models

type Stats struct {
	WidgetData   WidgetData     `json:"widget_data"`
	LicenseStats []LicenseStats `json:"license_stats"`
}

type WidgetData struct {
	TicketCount  int64 `json:"ticket_count"`
	UserCount    int64 `json:"user_count"`
	ProductCount int64 `json:"product_count"`
	LicenseCount int64 `json:"license_count"`
}

type LicenseStats struct {
	ID    string `json:"id"`
	Label string `json:"label"`
	Value int64  `json:"value"`
}
