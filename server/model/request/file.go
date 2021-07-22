package request

type FileInfo struct {
	Token   string `json:"token"`
	Login   string `json:"login"`
	Uid     string `json:"uid"`
	Content string `json:"content"`
	Sha     string `json:"sha"`
}
