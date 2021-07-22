package response

type FileInfo struct {
	Uid         string `json:"uid"`
	Name        string `json:"name"`
	Path        string `json:"path"`
	Content     string `json:"content"`
	Sha         string `json:"sha"`
	Size        int    `json:"size"`
	Url         string `json:"url"`
	FileType    string `json:"type"`
	HtmlUrl     string `json:"html_url"`
	DownloadUrl string `json:"download_url"`
	Html        string `json:"html"`
}
