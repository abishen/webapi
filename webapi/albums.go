package webapi

import "sync"

var (
	// albums stores the album data with ID as key for O(1) lookups
	albums = map[string]album{
		"1": {ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
		"2": {ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
		"3": {ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
	}
	// albumsMutex protects concurrent access to albums
	albumsMutex sync.RWMutex
)
