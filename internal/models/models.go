package models

// Version is checksum calculated from content and filename
type Version struct {
	Checksum string `json:"sha256sum"`
}

// Source holds content and filename
type Source struct {
	Content  string `json:"content"`
	FileName string `json:"filename"`
}

// Metadata is merely key-value pair for extra information
type Metadata struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}
