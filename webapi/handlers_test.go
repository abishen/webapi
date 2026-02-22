package webapi

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

// MockAlbumStore is a mock implementation of AlbumStore for testing
type MockAlbumStore struct {
	albums map[string]*Album
}

func NewMockAlbumStore() *MockAlbumStore {
	return &MockAlbumStore{
		albums: make(map[string]*Album),
	}
}

func (m *MockAlbumStore) GetAllAlbums() ([]Album, error) {
	result := make([]Album, 0, len(m.albums))
	for _, a := range m.albums {
		result = append(result, *a)
	}
	return result, nil
}

func (m *MockAlbumStore) GetAlbumByID(id string) (*Album, error) {
	if album, exists := m.albums[id]; exists {
		return album, nil
	}
	return nil, nil
}

func (m *MockAlbumStore) CreateAlbum(album *Album) error {
	if _, exists := m.albums[album.ID]; exists {
		return nil // Simulate already exists condition
	}
	a := *album
	m.albums[album.ID] = &a
	return nil
}

func (m *MockAlbumStore) DeleteAlbum(id string) error {
	delete(m.albums, id)
	return nil
}

func (m *MockAlbumStore) UpdateAlbum(id string, album *Album) error {
	m.albums[id] = album
	return nil
}

func initTestStore() {
	store := NewMockAlbumStore()
	store.CreateAlbum(&Album{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99})
	store.CreateAlbum(&Album{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99})
	store.CreateAlbum(&Album{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99})
	SetAlbumStore(store)
}

func TestGetAlbums(t *testing.T) {
	gin.SetMode(gin.TestMode)
	initTestStore()
	router := gin.New()
	RegisterRoutes(router)

	req := httptest.NewRequest(http.MethodGet, "/album", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var got []Album
	err := json.Unmarshal(w.Body.Bytes(), &got)
	assert.NoError(t, err)
	assert.Len(t, got, 3)
}

func TestGetAlbumByID(t *testing.T) {
	gin.SetMode(gin.TestMode)
	initTestStore()
	router := gin.New()
	RegisterRoutes(router)

	req := httptest.NewRequest(http.MethodGet, "/album/1", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var got Album
	err := json.Unmarshal(w.Body.Bytes(), &got)
	assert.NoError(t, err)
	assert.Equal(t, "1", got.ID)
	assert.Equal(t, "Blue Train", got.Title)
}

func TestGetAlbumByIDNotFound(t *testing.T) {
	gin.SetMode(gin.TestMode)
	initTestStore()
	router := gin.New()
	RegisterRoutes(router)

	req := httptest.NewRequest(http.MethodGet, "/album/999", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotFound, w.Code)

	var got map[string]string
	err := json.Unmarshal(w.Body.Bytes(), &got)
	assert.NoError(t, err)
	assert.Equal(t, "album not found", got["error"])
}

func TestPostAlbum(t *testing.T) {
	gin.SetMode(gin.TestMode)
	initTestStore()
	router := gin.New()
	RegisterRoutes(router)

	req := httptest.NewRequest(http.MethodPost, "/album", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestDeleteAlbum(t *testing.T) {
	gin.SetMode(gin.TestMode)
	initTestStore()
	router := gin.New()
	RegisterRoutes(router)

	req := httptest.NewRequest(http.MethodDelete, "/album/1", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var got map[string]string
	err := json.Unmarshal(w.Body.Bytes(), &got)
	assert.NoError(t, err)
	assert.Contains(t, got["message"], "album deleted")
}

func TestDeleteAlbumNotFound(t *testing.T) {
	gin.SetMode(gin.TestMode)
	initTestStore()
	router := gin.New()
	RegisterRoutes(router)

	req := httptest.NewRequest(http.MethodDelete, "/album/999", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotFound, w.Code)

	var got map[string]string
	err := json.Unmarshal(w.Body.Bytes(), &got)
	assert.NoError(t, err)
	assert.Equal(t, "album not found", got["error"])
}
