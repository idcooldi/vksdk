package object // import "github.com/idcooldi/vksdk/5.92/object"

type adsAccesses struct {
	ClientID string `json:"client_id"`
	Role     string `json:"role"`
}

// AdsAccount struct
type AdsAccount struct {
	AccessRole    string `json:"access_role"`
	AccountID     int    `json:"account_id"`     // Account ID
	AccountStatus int    `json:"account_status"` // Information whether account is active
	AccountType   string `json:"account_type"`
}

// AdsAdLayout struct
type AdsAdLayout struct {
	AdFormat    int    `json:"ad_format"`   // Ad format
	CampaignID  int    `json:"campaign_id"` // Campaign ID
	CostType    int    `json:"cost_type"`
	Description string `json:"description"`  // Ad description
	ID          int    `json:"id"`           // Ad ID
	ImageSrc    string `json:"image_src"`    // Image URL
	ImageSrc2x  string `json:"image_src_2x"` // URL of the preview image in double size
	LinkDomain  string `json:"link_domain"`  // Domain of advertised object
	LinkURL     string `json:"link_url"`     // URL of advertised object
	PreviewLink string `json:"preview_link"` // TODO: check it link to preview an ad as it is shown on the website
	Title       string `json:"title"`        // Ad title
	Video       int    `json:"video"`        // Information whether the ad is a video
}

// AdsCampaign struct
type AdsCampaign struct {
	AllLimit  string `json:"all_limit"`  // Campaign's total limit, rubles
	DayLimit  string `json:"day_limit"`  // Campaign's day limit, rubles
	ID        int    `json:"id"`         // Campaign ID
	Name      string `json:"name"`       // Campaign title
	StartTime int    `json:"start_time"` // Campaign start time, as Unixtime
	Status    int    `json:"status"`
	StopTime  int    `json:"stop_time"` // Campaign stop time, as Unixtime
	Type      string `json:"type"`
}

// AdsCategory struct
type AdsCategory struct {
	ID            int                  `json:"id"`   // Category ID
	Name          string               `json:"name"` // Category name
	Subcategories []baseObjectWithName `json:"subcategories"`
}

// AdsClient struct
type AdsClient struct {
	AllLimit string `json:"all_limit"` // Client's total limit, rubles
	DayLimit string `json:"day_limit"` // Client's day limit, rubles
	ID       int    `json:"id"`        // Client ID
	Name     string `json:"name"`      // Client name
}

// AdsCriteria struct
type AdsCriteria struct {
	AgeFrom              int    `json:"age_from"`               // Age from
	AgeTo                int    `json:"age_to"`                 // Age to
	Apps                 string `json:"apps"`                   // Apps IDs
	AppsNot              string `json:"apps_not"`               // Apps IDs to except
	Birthday             int    `json:"birthday"`               // Days to birthday
	Cities               string `json:"cities"`                 // Cities IDs
	CitiesNot            string `json:"cities_not"`             // Cities IDs to except
	Country              int    `json:"country"`                // Country ID
	Districts            string `json:"districts"`              // Districts IDs
	Groups               string `json:"groups"`                 // Communities IDs
	InterestCategories   string `json:"interest_categories"`    // Interests categories IDs
	Interests            string `json:"interests"`              // Interests
	Paying               int    `json:"paying"`                 // Information whether the user has proceeded VK payments before
	Positions            string `json:"positions"`              // Positions IDs
	Religions            string `json:"religions"`              // Religions IDs
	RetargetingGroups    string `json:"retargeting_groups"`     // Retargeting groups IDs
	RetargetingGroupsNot string `json:"retargeting_groups_not"` // Retargeting groups IDs to except
	SchoolFrom           int    `json:"school_from"`            // School graduation year from
	SchoolTo             int    `json:"school_to"`              // School graduation year to
	Schools              string `json:"schools"`                // Schools IDs
	Sex                  int    `json:"sex"`
	Stations             string `json:"stations"`      // Stations IDs
	Statuses             string `json:"statuses"`      // Relationship statuses
	Streets              string `json:"streets"`       // Streets IDs
	Travellers           int    `json:"travellers"`    // Travellers only
	UniFrom              int    `json:"uni_from"`      // University graduation year from
	UniTo                int    `json:"uni_to"`        // University graduation year to
	UserBrowsers         string `json:"user_browsers"` // Browsers
	UserDevices          string `json:"user_devices"`  // Devices
	UserOs               string `json:"user_os"`       // Operating systems
}

// AdsDemoStats struct
type AdsDemoStats struct {
	ID    int                `json:"id"` // Object ID
	Stats adsDemostatsFormat `json:"stats"`
	Type  string             `json:"type"`
}

type adsDemostatsFormat struct {
	Age     []adsStatsAge    `json:"age"`
	Cities  []adsStatsCities `json:"cities"`
	Day     string           `json:"day"`     // Day as YYYY-MM-DD
	Month   string           `json:"month"`   // Month as YYYY-MM
	Overall int              `json:"overall"` // 1 if period=overall
	Sex     []adsStatsSex    `json:"sex"`
	SexAge  []adsStatsSexAge `json:"sex_age"`
}

// AdsFloodStats struct
type AdsFloodStats struct {
	Left    int `json:"left"`    // Requests left
	Refresh int `json:"refresh"` // Time to refresh in seconds
}

// AdsLinkStatus struct
type AdsLinkStatus struct {
	Description string `json:"description"`  // Reject reason
	RedirectURL string `json:"redirect_url"` // URL
	Status      string `json:"status"`       // Link status
}

type adsParagraphs struct {
	Paragraph string `json:"paragraph"` // Rules paragraph
}

// AdsRejectReason struct
type AdsRejectReason struct {
	Comment string     `json:"comment"` // Comment text
	Rules   []adsRules `json:"rules"`
}

type adsRules struct {
	Paragraphs []adsParagraphs `json:"paragraphs"`
	Title      string          `json:"title"` // Comment
}

// AdsStats struct
type AdsStats struct {
	ID    int            `json:"id"` // Object ID
	Stats adsStatsFormat `json:"stats"`
	Type  string         `json:"type"`
}

type adsStatsAge struct {
	ClicksRate      float64 `json:"clicks_rate"`      // Clicks rate
	ImpressionsRate float64 `json:"impressions_rate"` // Impressions rate
	Value           string  `json:"value"`            // Age interval
}

type adsStatsCities struct {
	ClicksRate      float64 `json:"clicks_rate"`      // Clicks rate
	ImpressionsRate float64 `json:"impressions_rate"` // Impressions rate
	Name            string  `json:"name"`             // City name
	Value           int     `json:"value"`            // City ID
}

type adsStatsFormat struct {
	Clicks          int    `json:"clicks"`            // Clicks number
	Day             string `json:"day"`               // Day as YYYY-MM-DD
	Impressions     int    `json:"impressions"`       // Impressions number
	JoinRate        int    `json:"join_rate"`         // Events number
	Month           string `json:"month"`             // Month as YYYY-MM
	Overall         int    `json:"overall"`           // 1 if period=overall
	Reach           int    `json:"reach"`             // Reach
	Spent           int    `json:"spent"`             // Spent funds
	VideoClicksSite int    `json:"video_clicks_site"` // Clickthoughs to the advertised site
	VideoViews      int    `json:"video_views"`       // Video views number
	VideoViewsFull  int    `json:"video_views_full"`  // Video views (full video)
	VideoViewsHalf  int    `json:"video_views_half"`  // Video views (half of video)
}

type adsStatsSex struct {
	ClicksRate      float64 `json:"clicks_rate"`      // Clicks rate
	ImpressionsRate float64 `json:"impressions_rate"` // Impressions rate
	Value           string  `json:"value"`
}

type adsStatsSexAge struct {
	ClicksRate      float64 `json:"clicks_rate"`      // Clicks rate
	ImpressionsRate float64 `json:"impressions_rate"` // Impressions rate
	Value           string  `json:"value"`            // Sex and age interval
}

// AdsTargSettings struct
type AdsTargSettings struct {
}

// AdsTargStats struct
type AdsTargStats struct {
	AudienceCount  int     `json:"audience_count"`  // Audience
	RecommendedCpc float64 `json:"recommended_cpc"` // Recommended CPC value
	RecommendedCpm float64 `json:"recommended_cpm"` // Recommended CPM value
}

// AdsTargSuggestions struct
type AdsTargSuggestions struct {
	ID   int    `json:"id"`   // Object ID
	Name string `json:"name"` // Object name
}

// AdsTargSuggestionsCities struct
type AdsTargSuggestionsCities struct {
	ID     int    `json:"id"`     // Object ID
	Name   string `json:"name"`   // Object name
	Parent string `json:"parent"` // Parent object
}

// AdsTargSuggestionsRegions struct
type AdsTargSuggestionsRegions struct {
	ID   int    `json:"id"`   // Object ID
	Name string `json:"name"` // Object name
	Type string `json:"type"` // Object type
}

// AdsTargSuggestionsSchools struct
type AdsTargSuggestionsSchools struct {
	Desc   string `json:"desc"`   // Full school title
	ID     int    `json:"id"`     // School ID
	Name   string `json:"name"`   // School title
	Parent string `json:"parent"` // City name
	Type   string `json:"type"`
}

// AdsTargetGroup struct
type AdsTargetGroup struct {
	AudienceCount int    `json:"audience_count"` // Audience
	Domain        string `json:"domain"`         // Site domain
	ID            int    `json:"id"`             // Group ID
	Lifetime      int    `json:"lifetime"`       // Number of days for user to be in group
	Name          string `json:"name"`           // Group name
	Pixel         string `json:"pixel"`          // Pixel code
}

// AdsUsers struct
type AdsUsers struct {
	Accesses []adsAccesses `json:"accesses"`
	UserID   int           `json:"user_id"` // User ID
}

// AdsAd struct
type AdsAd struct {
	AdFormat              int         `json:"ad_format"`   // Ad format
	AdPlatform            interface{} `json:"ad_platform"` // Ad platform
	AllLimit              int         `json:"all_limit"`   // Total limit
	Approved              int         `json:"approved"`
	CampaignID            int         `json:"campaign_id"`  // Campaign ID
	Category1ID           int         `json:"category1_id"` // Category ID
	Category2ID           int         `json:"category2_id"` // Additional category ID
	CostType              int         `json:"cost_type"`
	Cpc                   int         `json:"cpc"`                    // Cost of a click, kopecks
	Cpm                   int         `json:"cpm"`                    // Cost of 1000 impressions, kopecks
	DisclaimerMedical     int         `json:"disclaimer_medical"`     // Information whether disclaimer is enabled
	DisclaimerSpecialist  int         `json:"disclaimer_specialist"`  // Information whether disclaimer is enabled
	DisclaimerSupplements int         `json:"disclaimer_supplements"` // Information whether disclaimer is enabled
	ID                    int         `json:"id"`                     // Ad ID
	ImpressionsLimit      int         `json:"impressions_limit"`      // Impressions limit
	ImpressionsLimited    int         `json:"impressions_limited"`    // Information whether impressions are limited
	Name                  string      `json:"name"`                   // Ad title
	Status                int         `json:"status"`
	Video                 int         `json:"video"` // Information whether the ad is a video
}

// AdsPromotedPostReach struct
type AdsPromotedPostReach struct {
	Hide             int `json:"hide"`              // Hides amount
	ID               int `json:"id"`                // Object ID from 'ids' parameter
	JoinGroup        int `json:"join_group"`        // Community joins
	Links            int `json:"links"`             // Link clicks
	ReachSubscribers int `json:"reach_subscribers"` // Subscribers reach
	ReachTotal       int `json:"reach_total"`       // Total reach
	Report           int `json:"report"`            // Reports amount
	ToGroup          int `json:"to_group"`          // Community clicks
	Unsubscribe      int `json:"unsubscribe"`       // 'Unsubscribe' events amount
	VideoViews100p   int `json:"video_views_100p"`  // Video views for 100 percent
	VideoViews25p    int `json:"video_views_25p"`   // Video views for 25 percent
	VideoViews3s     int `json:"video_views_3s"`    // Video views for 3 seconds
	VideoViews50p    int `json:"video_views_50p"`   // Video views for 50 percent
	VideoViews75p    int `json:"video_views_75p"`   // Video views for 75 percent
	VideoViewsStart  int `json:"video_views_start"` // Video starts
}
