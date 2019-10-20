package models

type Artist struct {
	// gorm.Model // Extend our own model with the properties given at the gorm basic model
	ArtistId int `gorm:"primary_key"`
	Name     string
	Albums   []Album `gorm:"foreignkey:ArtistId"`
}

type Album struct {
	AlbumId  int `gorm:"primary_key"`
	Title    string
	ArtistId int
	Artist   Artist `gorm:"foreignkey:ArtistId"`
}

type Customer struct {
	CustomerId int `gorm:"primary_key"`
	Email      string
	Invoices   []Invoice `gorm:"foreignkey:CustomerId"`
}

type Invoice struct {
	InvoiceId    int      `gorm:"primary_key"`
	Customer     Customer `gorm:"foreignkey:CustomerId"`
	CustomerId   int
	Total        float32
	InvoiceLines []InvoiceLine `gorm:"foreignkey:InvoiceId"`
}

type InvoiceLine struct {
	InvoiceLineId int     `gorm:"primary_key"`
	Invoice       Invoice `gorm:"foreignkey:InvoiceId"`
	InvoiceId     int
	TrackId       int
	UnitPrice     float32
	Quantity      int
}
