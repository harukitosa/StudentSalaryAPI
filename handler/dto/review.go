package dto

// ReviewBody is
type ReviewBody struct {
	CompanyName  string `json:"company_name"`
	Content      string `json:"content"`
	CreateDateJS string `json:"create_date"`
	Link         string `json:"link"`
	Reasons      string `json:"reasons"`
	Report       string `json:"report"`
	Skill        string `json:"skill"`
	UserName     string `json:"user_name"`
}

// ExportReviewBody is
// Salaryはstringで保存していたのでstringで受け取る
type ExportReviewBody struct {
	CompanyName  string `json:"company_name"`
	Content      string `json:"content"`
	CreateDateJS int    `json:"create_date"`
	Link         string `json:"link"`
	Reasons      string `json:"reasons"`
	Report       string `json:"report"`
	Skill        string `json:"skill"`
	UserName     string `json:"user_name"`
}
