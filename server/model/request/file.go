package request

type FileInfo struct {
	Token   string `json:"token"`
	Login   string `json:"login"`
	Path    string `json:"path"`
	Content string `json:"content"`
	Sha     string `json:"sha"`
}
