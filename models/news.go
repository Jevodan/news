package models

type Post struct {
	Title       string `xml:"title" json:"Title"`
	Description string `xml:"description" json:"Content"`
	PubDate     int64  `xml:"pubDate" json:"PubTime"`
	Link        string `xml:"link" json:"Link"`
}
