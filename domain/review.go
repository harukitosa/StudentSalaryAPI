package domain

// WorkReview is student review
type Review struct {
	ID           int
	CompanyName  string
	Detail       string
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
	GetNewReview() ([]Review, error)
}

func NewReview(
	CompanyName *string,
	Detail *string,
	Content *string,
	CreateDataJs *string,
	Link *string,
	Reasons *string,
	Report *string,
	Skill *string,
	UserName *string,
) Review {
	return Review{
		CompanyName:  convertNilString(CompanyName),
		Detail:       convertNilString(Detail),
		Content:      convertNilString(Content),
		CreateDateJS: convertNilString(CreateDataJs),
		Link:         convertNilString(Link),
		Reasons:      convertNilString(Reasons),
		Report:       convertNilString(Report),
		Skill:        convertNilString(Skill),
		UserName:     convertNilString(UserName),
	}
}
