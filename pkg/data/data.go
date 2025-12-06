package data

import "time"

type Cache struct {
	Status string `json:"status" yaml:"status"`
}

//
// Payloads
//

type AddBookmarkPayload struct {
	CollectionId int64    `json:"collectionId,omitempty"`
	Excerpt      string   `json:"excerpt,omitempty"`
	Highlights   []string `json:"highlights,omitempty"`
	Important    bool     `json:"important,omitempty"`
	Link         string   `json:"link,omitempty"`
	Note         string   `json:"note,omitempty"`
	Tags         []string `json:"tags,omitempty"`
	Title        string   `json:"title,omitempty"`
}

type AddCollectionPayload struct {
	Title  string `json:"title,omitempty"`
	Parent int64  `json:"parent,omitempty"`
	Public bool   `json:"public,omitempty"`
	View   string `json:"view,omitempty"`
}

type SortCollectionPayload struct {
	Sort string `json:"sort,omitempty"`
}

type RemoveHighlightItem struct {
	Id   string `json:"_id,omitempty"`
	Text string `json:"text,omitempty"`
}

type RemoveHighlightsPayload struct {
	Highlights []RemoveHighlightItem `json:"highlights,omitempty"`
}

type RemoveTagsPayload struct {
	CollectionId int64    `json:"collectionId,omitempty"`
	Tags         []string `json:"tags,omitempty"`
}
type RenameTagPayload struct {
	CollectionId int64    `json:"collectionId,omitempty"`
	NewName      string   `json:"replace,omitempty"`
	OldName      []string `json:"tags,omitempty"`
}

type UpdateBookmarkPayload struct {
	CollectionId int64    `json:"collectionId,omitempty"`
	Excerpt      string   `json:"excerpt,omitempty"`
	Highlights   []string `json:"highlights,omitempty"`
	Important    bool     `json:"important,omitempty"`
	Link         string   `json:"link,omitempty"`
	Note         string   `json:"note,omitempty"`
	Tags         []string `json:"tags,omitempty"`
	Title        string   `json:"title,omitempty"`
}

type UpdateCollectionPayload struct {
	Title  string `json:"title,omitempty"`
	Parent int64  `json:"parent,omitempty"`
	Public bool   `json:"public,omitempty"`
	View   string `json:"view,omitempty"`
}

//
// Results
//

// Login
type LoginResult struct {
	Result       bool   `json:"result" yaml:"result"`
	ErrorMessage string `json:"errorMessage" yaml:"errorMessage"`
}

// Bookmarks
type AddBookmarkResult struct {
	Result       bool     `json:"result" yaml:"result"`
	Item         Bookmark `json:"item" yaml:"item"`
	ErrorMessage string   `json:"errorMessage" yaml:"errorMessage"`
}

type ListBookmarksResult struct {
	Result       bool        `json:"result" yaml:"result"`
	Items        []*Bookmark `json:"items" yaml:"items"`
	Count        int         `json:"count" yaml:"count"`
	ErrorMessage string      `json:"errorMessage" yaml:"errorMessage"`
}
type RemoveBookmarkResult struct {
	Result       bool   `json:"result" yaml:"result"`
	ErrorMessage string `json:"errorMessage" yaml:"errorMessage"`
}

type UpdateBookmarkResult struct {
	Result       bool     `json:"result" yaml:"result"`
	Item         Bookmark `json:"item" yaml:"item"`
	ErrorMessage string   `json:"errorMessage" yaml:"errorMessage"`
}

// Collections
type AddCollectionResult struct {
	Result       bool       `json:"result" yaml:"result"`
	Item         Collection `json:"item" yaml:"item"`
	ErrorMessage string     `json:"errorMessage" yaml:"errorMessage"`
}

type ListCollectionsResult struct {
	Result       bool          `json:"result" yaml:"result"`
	Items        []*Collection `json:"items" yaml:"items"`
	Count        int           `json:"count" yaml:"count"`
	ErrorMessage string        `json:"errorMessage" yaml:"errorMessage"`
}

type ListCollectionsChildrenResult struct {
	Result       bool          `json:"result" yaml:"result"`
	Items        []*Collection `json:"items" yaml:"items"`
	Count        int           `json:"count" yaml:"count"`
	ErrorMessage string        `json:"errorMessage" yaml:"errorMessage"`
}

type RemoveCollectionResult struct {
	Result       bool   `json:"result" yaml:"result"`
	ErrorMessage string `json:"errorMessage" yaml:"errorMessage"`
}

type SortCollectionsResult struct {
	Result       bool   `json:"result" yaml:"result"`
	ErrorMessage string `json:"errorMessage" yaml:"errorMessage"`
}

type UpdateCollectionResult struct {
	Result       bool       `json:"result" yaml:"result"`
	Item         Collection `json:"item" yaml:"item"`
	ErrorMessage string     `json:"errorMessage" yaml:"errorMessage"`
}

// Highlights
type ListHighlightsResult struct {
	Result       bool         `json:"result" yaml:"result"`
	Items        []*Highlight `json:"items" yaml:"items"`
	ErrorMessage string       `json:"errorMessage" yaml:"errorMessage"`
}

type RemoveHighlightResult struct {
	Result       bool                 `json:"result" yaml:"result"`
	Item         RemovedHighlightItem `json:"item" yaml:"item"`
	ErrorMessage string               `json:"errorMessage" yaml:"errorMessage"`
}

// Tags
type ListTagsResult struct {
	Result       bool   `json:"result" yaml:"result"`
	Items        []*Tag `json:"items" yaml:"items"`
	ErrorMessage string `json:"errorMessage" yaml:"errorMessage"`
}

type RemoveTagsResult struct {
	Result       bool   `json:"result" yaml:"result"`
	ErrorMessage string `json:"errorMessage" yaml:"errorMessage"`
}

type RenameTagResult struct {
	Result       bool   `json:"result" yaml:"result"`
	ErrorMessage string `json:"errorMessage" yaml:"errorMessage"`
}

//
// Refs
//

type CollectionRef struct {
	Ref  string `json:"$ref" yaml:"$ref"`
	Id   int64  `json:"$id" yaml:"$id"`
	Oid  int64  `json:"oid" yaml:"oid"`
	Name string `json:"name" yaml:"name"`
}

type CreatorRef struct {
	Id     int64  `json:"_id" yaml:"_id"`
	Avatar string `json:"avatar" yaml:"avatar"`
	Name   string `json:"name" yaml:"name"`
	Email  string `json:"email" yaml:"email"`
}

type LinkRef struct {
	Link string `json:"link" yaml:"link"`
	Type string `json:"type" yaml:"type"`
}

type ReminderRef struct {
	Date time.Time `json:"date" yaml:"date"`
}

type UserRef struct {
	Ref string `json:"$ref" yaml:"$ref"`
	Id  int64  `json:"$id" yaml:"$id"`
}

//
// Raindrop
//

type Bookmark struct {
	Id           uint64              `json:"_id" yaml:"_id"`
	Link         string              `json:"link" yaml:"link"`
	Title        string              `json:"title" yaml:"title"`
	Excerpt      string              `json:"excerpt" yaml:"excerpt"`
	Note         string              `json:"note" yaml:"note"`
	Type         string              `json:"type" yaml:"type"`
	User         UserRef             `json:"user" yaml:"user"`
	Cover        string              `json:"cover" yaml:"cover"`
	Media        []LinkRef           `json:"media" yaml:"media"`
	Tags         []string            `json:"tags" yaml:"tags"`
	Important    bool                `json:"important" yaml:"important"`
	Reminder     ReminderRef         `json:"reminder" yaml:"reminder"`
	Removed      bool                `json:"removed" yaml:"removed"`
	Created      time.Time           `json:"created" yaml:"created"`
	Collection   CollectionRef       `json:"collection" yaml:"collection"`
	Highlights   []BookmarkHighlight `json:"highlights" yaml:"highlights"`
	LastUpdate   time.Time           `json:"lastUpdate" yaml:"lastUpdate"`
	Domain       string              `json:"domain" yaml:"domain"`
	CreatorRef   CreatorRef          `json:"creatorRef" yaml:"creatorRef"`
	Sort         int64               `json:"sort" yaml:"sort"`
	CollectionId int64               `json:"collectionId" yaml:"collectionId"`
	Cache        Cache               `json:"cache" yaml:"cache"`
}

type BookmarkHighlight struct {
	Text       string    `json:"text" yaml:"text"`
	Note       string    `json:"note" yaml:"note"`
	Color      string    `json:"color" yaml:"color"`
	Position   uint64    `json:"position" yaml:"position"`
	Created    time.Time `json:"created" yaml:"created"`
	LastUpdate time.Time `json:"lastUpdate" yaml:"lastUpdate"`
	CreatorRef uint64    `json:"creatorRef" yaml:"creatorRef"`
	Id         string    `json:"_id" yaml:"_id"`
}

//
// Collection
//

type CollectionAccess struct {
	For       uint64 `json:"for" yaml:"for"`
	Level     int    `json:"level" yaml:"level"`
	Root      bool   `json:"root" yaml:"root"`
	Draggable bool   `json:"draggable" yaml:"draggable"`
}

type ParentRef struct {
	Ref string `json:"$ref" yaml:"$ref"`
	Id  uint64 `json:"$id" yaml:"$id"`
}

type Collection struct {
	Access        CollectionAccess `json:"access" yaml:"access"`
	Collaborators UserRef          `json:"collaborators" yaml:"collaborators"`
	Color         string           `json:"color" yaml:"color"`
	Count         uint64           `json:"count" yaml:"count"`
	Cover         []string         `json:"cover" yaml:"cover"`
	Created       time.Time        `json:"created" yaml:"created"`
	CreatorRef    CreatorRef       `json:"creatorRef" yaml:"creatorRef"`
	Description   string           `json:"description" yaml:"description"`
	Expanded      bool             `json:"expanded" yaml:"expanded"`
	Id            uint64           `json:"_id" yaml:"_id"`
	LastAction    time.Time        `json:"lastAction" yaml:"lastAction"`
	LastUpdate    time.Time        `json:"lastUpdate" yaml:"lastUpdate"`
	Parent        ParentRef        `json:"parent" yaml:"parent"`
	Public        bool             `json:"public" yaml:"public"`
	Slug          string           `json:"slug" yaml:"slug"`
	Sort          int64            `json:"sort" yaml:"sort"`
	Title         string           `json:"title" yaml:"title"`
	User          UserRef          `json:"user" yaml:"user"`
	View          string           `json:"view" yaml:"view"`

	Author bool `json:"author" yaml:"autor"`
}

//
// Highlight
//

type Highlight struct {
	Id      string    `json:"_id" yaml:"_id"`
	Text    string    `json:"text" yaml:"text"`
	Title   string    `json:"title" yaml:"title"`
	Color   string    `json:"color" yaml:"color"`
	Note    string    `json:"note" yaml:"note"`
	Created time.Time `json:"created" yaml:"created"`
	Tags    []string  `json:"tags" yaml:"tags"`
	Link    string    `json:"link" yaml:"link"`
}

type RemovedHighlightItem struct {
	Id         uint64      `json:"_id" yaml:"_id"`
	Highlights []Highlight `json:"highlights" yaml:"highlights"`
}

//
// Tag
//

type Tag struct {
	Id    string `json:"_id" yaml:"_id"`
	Count uint64 `json:"count" yaml:"count"`
}
