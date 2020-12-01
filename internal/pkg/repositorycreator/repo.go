package repositorycreator

// Repo struct is the base type for a Repo
type Repo struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	HomePage    string `json:"homepage"`
	Private     bool   `json:"private"`
}
