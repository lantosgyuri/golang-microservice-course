package auth

// Error is sent when an error accours
type Error struct {
	StatusCode int64
	Message    string
}
