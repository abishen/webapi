package webapi

import "gorm.io/gorm"

// AlbumStore provides database operations for albums
type AlbumStore interface {
	GetAllAlbums() ([]Album, error)
	GetAlbumByID(id string) (*Album, error)
	CreateAlbum(album *Album) error
	DeleteAlbum(id string) error
	UpdateAlbum(id string, album *Album) error
}

// GormAlbumStore implements AlbumStore using GORM
type GormAlbumStore struct {
	db *gorm.DB
}

// NewGormAlbumStore creates a new GORM-based album store
func NewGormAlbumStore(db *gorm.DB) AlbumStore {
	return &GormAlbumStore{db: db}
}

// GetAllAlbums retrieves all albums from the database
func (s *GormAlbumStore) GetAllAlbums() ([]Album, error) {
	var albums []Album
	result := s.db.Find(&albums)
	return albums, result.Error
}

// GetAlbumByID retrieves a specific album by ID
func (s *GormAlbumStore) GetAlbumByID(id string) (*Album, error) {
	var album Album
	result := s.db.Where("id = ?", id).First(&album)
	if result.Error == gorm.ErrRecordNotFound {
		return nil, nil
	}
	return &album, result.Error
}

// CreateAlbum creates a new album in the database
func (s *GormAlbumStore) CreateAlbum(album *Album) error {
	return s.db.Create(album).Error
}

// DeleteAlbum deletes an album by ID from the database
func (s *GormAlbumStore) DeleteAlbum(id string) error {
	return s.db.Delete(&Album{}, "id = ?", id).Error
}

// UpdateAlbum updates an existing album
func (s *GormAlbumStore) UpdateAlbum(id string, album *Album) error {
	return s.db.Model(&Album{}).Where("id = ?", id).Updates(album).Error
}
