package domain

// WorkReview is student review
type Review struct {
	ID           int
	CompanyName  string
	Content      string
	CreateDateJS string
	Link         string
	Reasons      string
	Report       string
	Skill        string
	UserName     string
}

type ReviewRepository interface {
	Insert(review Review) (id int, err error)
	SelectByID(id int) (Review, error)
	SelectByName(name string) ([]Review, error)
	SelectAll() ([]Review, error)
}
