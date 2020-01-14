package object // import "github.com/idcooldi/vksdk/5.92/object"

// DatabaseCity struct
type DatabaseCity struct {
	ID        int    `json:"id"`    // City ID
	Title     string `json:"title"` // City title
	Area      string `json:"area"`
	Region    string `json:"region"`
	Important int    `json:"important"`
}

// DatabaseMetroStation  struct
type DatabaseMetroStation struct {
	ID    int    `json:"id"`    // Metro station ID
	Name  string `json:"name"`  // Metro station name
	Color string `json:"color"` // Metro station color
}

// DatabaseFaculty struct
type DatabaseFaculty struct {
	ID    int    `json:"id"`    // Faculty ID
	Title string `json:"title"` // Faculty title
}

// DatabaseRegion struct
type DatabaseRegion struct {
	ID    int    `json:"id"`    // Region ID
	Title string `json:"title"` // Region title
}

// DatabaseSchool struct
type DatabaseSchool struct {
	ID    int    `json:"id"`    // School ID
	Title string `json:"title"` // School title
}

// DatabaseStation struct
type DatabaseStation struct {
	CityID int    `json:"city_id"` // City ID
	Color  string `json:"color"`   // Hex color code without #
	ID     int    `json:"id"`      // Station ID
	Name   string `json:"name"`    // Station name
}

// DatabaseUniversity struct
type DatabaseUniversity struct {
	ID    int    `json:"id"`    // University ID
	Title string `json:"title"` // University title
}
