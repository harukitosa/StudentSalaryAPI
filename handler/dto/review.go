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
