package peanuts

type Image struct {
	Link      string `json:"link"`
	IsDefault bool   `json:"is_default"`
	Width     int    `json:"width"`
	Height    int    `json:"height"`
}

type ContentOfUser struct {
	Text        string   `json:"text"`
	Html        string   `json:"html"`
	Entities    Entities `json:"entities"`
	AvatarImage Image    `json:"avatar_image"`
	CoverImage  Image    `json:"cover_image"`
}

type CountsOfUser struct {
	Bookmarks int `json:"bookmarks"`
	Clients   int `json:"clients"`
	Followers int `json:"followers"`
	Following int `json:"following"`
	Posts     int `json:"posts"`
	Users     int `json:"users"`
}

type Verified struct {
	Domain string `json:"domain"`
	Link   string `json:"link"`
}

type User struct {
	CreatedAt    string        `json:"created_at"`
	Guid         string        `json:"guid"`
	Id           string        `json:"id"`
	Locale       string        `json:"locale"`
	Timezone     string        `json:"timezon"`
	Type         string        `json:"type"`
	Username     string        `json:"username"`
	Name         string        `json:"name"`
	Content      ContentOfUser `json:"content"`
	Counts       CountsOfUser  `json:"counts"`
	FollowsYou   bool          `json:"follows_you"`
	YouBlocked   bool          `json:"you_blocked"`
	YouFollow    bool          `json:"you_follow"`
	YouMuted     bool          `json:"you_muted"`
	YouCanFollow bool          `json:"you_can_follow"`
	Verified     Verified      `json:"verified"`
}

type Source struct {
	Name string `json:"name"`
	Link string `json:"link"`
	Id   string `json:"id"`
}

type CountsOfPost struct {
	Bookmarks int `json:"bookmarks"`
	Replies   int `json:"replies"`
	Reposts   int `json:"reposts"`
	Threads   int `json:"threads"`
}

type ContentOfPost struct {
	Text           string   `json:"text"`
	Html           string   `json:"html"`
	Entities       Entities `json:"entities"`
	LinksNotParsed bool     `json:"links_not_parsed"`
}

type Post struct {
	CreatedAt     string        `json:"created_at"`
	Guid          string        `json:"guid"`
	Id            string        `json:"id"`
	IsDeleted     bool          `json:"is_deleted"`
	Source        Source        `json:"source"`
	User          User          `json:"user"`
	ThreadId      string        `json:"thread_id"`
	IsRevised     bool          `json:"is_revised"`
	Revision      string        `json:"revision"`
	ReplyTo       string        `json:"reply_to"`
	RepostOf      *Post         `json:"repost_of"`
	Counts        CountsOfPost  `json:"counts"`
	Content       ContentOfPost `json:"content"`
	YouBookmarked bool          `json:"you_bookmarked"`
	YouReposted   bool          `json:"you_reposted"`
	PaginationId  string        `json:"pagination_id"`
}

type Action struct {
	PaginationId string `json:"pagination_id"`
	EventDate    string `json:"event_date"`
	Action       string `json:"action"`
	Users        []User `json:"users"`
}
