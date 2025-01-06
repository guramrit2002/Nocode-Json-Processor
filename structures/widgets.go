package structures

type Widget struct {
	ID       string            `json:"id"`
	Type     string            `json:"type"`
	Style    map[string]string `json:"style"`
	Content  string            `json:"content,omitempty"`
	Children []Widget          `json:"children,omitempty"`
}
