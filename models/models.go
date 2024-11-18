package models

type Invoice struct {
	ID           uint          `gorm:"primaryKey" json:"id"`
	InvoiceID    string        `gorm:"unique;not null" json:"invoice_id"`
	IssueDate    string        `json:"issue_date"`
	DueDate      string        `json:"due_date"`
	Subject      string        `json:"subject"`
	FromCompany  string        `json:"from_company"`
	FromAddress  string        `json:"from_address"`
	ToCompany    string        `json:"to_company"`
	ToAddress    string        `json:"to_address"`
	Subtotal     float64       `json:"subtotal"`
	Tax          float64       `json:"tax"`
	Total        float64       `json:"total"`
	Status       string        `gorm:"default:UNPAID" json:"status"`
	InvoiceItems []InvoiceItem `gorm:"foreignKey:InvoiceID" json:"items"`
}

type InvoiceItem struct {
	ID          uint    `gorm:"primaryKey" json:"id"`
	InvoiceID   uint    `json:"invoice_id"`
	ItemType    string  `json:"item_type"`
	Description string  `json:"description"`
	Quantity    float64 `json:"quantity"`
	UnitPrice   float64 `json:"unit_price"`
	Amount      float64 `json:"amount"`
}
