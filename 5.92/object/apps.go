package object // import "github.com/idcooldi/vksdk/5.92/object"

// AppsApp struct
type AppsApp struct {
	AuthorGroup     int           `json:"author_group"`     // Official community's ID
	AuthorID        int           `json:"author_id"`        // Application author's ID
	AuthorURL       string        `json:"author_url"`       // Application author's URL
	Banner1120      string        `json:"banner_1120"`      // URL of the app banner with 1120 px in width
	Banner560       string        `json:"banner_560"`       // URL of the app banner with 560 px in width
	CatalogPosition int           `json:"catalog_position"` // Catalog position
	Description     string        `json:"description"`      // Application description
	Friends         []int         `json:"friends"`
	Genre           string        `json:"genre"`         // Genre name
	GenreID         int           `json:"genre_id"`      // Genre ID
	Icon139         string        `json:"icon_139"`      // URL of the app icon with 139 px in width
	Icon150         string        `json:"icon_150"`      // URL of the app icon with 150 px in width
	Icon278         string        `json:"icon_278"`      // URL of the app icon with 279 px in width
	Icon75          string        `json:"icon_75"`       // URL of the app icon with 75 px in width
	ID              int           `json:"id"`            // Application ID
	International   int           `json:"international"` // Information whether the application is multilanguage
	IsInCatalog     int           `json:"is_in_catalog"` // Information whether application is in mobile catalog
	LeaderboardType int           `json:"leaderboard_type"`
	MembersCount    int           `json:"members_count"`  // Members number
	PlatformID      int           `json:"platform_id"`    // Application ID in store
	PublishedDate   int           `json:"published_date"` // Date when the application has been published in Unixtime
	ScreenName      string        `json:"screen_name"`    // Screen name
	Screenshots     []PhotosPhoto `json:"screenshots"`
	Section         string        `json:"section"` // Application section name
	Title           string        `json:"title"`   // Application title
	Type            string        `json:"type"`
	Icon16          string        `json:"icon_16"`
	PushEnabled     int           `json:"push_enabled"`
}

// AppsLeaderboard struct
type AppsLeaderboard struct {
	Level  int `json:"level"`   // Level
	Points int `json:"points"`  // Points number
	Score  int `json:"score"`   // Score number
	UserID int `json:"user_id"` // User ID
}

// AppsScope Scope description
type AppsScope struct {
	Name  string `json:"name"`  // Scope name
	Title string `json:"title"` // Scope title
}
