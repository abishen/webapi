package webapi

// Album represents an album in the database
type Album struct {
	ID     string  `gorm:"primaryKey" json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

// TableName specifies the table name for Album
func (Album) TableName() string {
	return "albums"
}
