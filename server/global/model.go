package global

// {"access_token":"5e558bfc91105abaf3705a5c3bdc30b8","token_type":"bearer","expires_in":86400,"refresh_token":"001e520ed86ba1472fb411d4b9f0ebe38030131a03e71de9629d203d25bddb2e","scope":"user_info projects","created_at":1624282839}
type AuthInfo struct {
	AccessToken  string `json:"access_token"`
	TokenType    string `json:"token_type"`
	ExpiresIn    int64  `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
	Scope        string `json:"scope"`
	CreatedAt    int64  `json:"created_at"`
	RepoType     string `json:"repo_type"`
}

const (
	APP_NAME       = "shownote"
	REPO_NAME      = "sn-repo"
	SHARE_NAME     = "sn-share"
	WORKSPACE_PATH = ".shownote/workspace.json"
	COMMIT_MSG     = "shownote commit"
)
