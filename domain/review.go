package domain

import (
	"fmt"
	"net/url"
)

// WorkReview is student review
type Review struct {
	id           reviewID
	companyName  companyName
	content      content
	createDateJS createDate
	link         link
	reasons      reasons
	report       report
	skill        skill
	userName     userName
}

func (r *Review) GetID() reviewID {
	return r.id
}

func (r *Review) GetCompanyName() companyName {
	return r.companyName
}

func (r *Review) GetContent() content {
	return r.content
}

func (r *Review) GetCreateDate() createDate {
	return r.createDateJS
}

func (r *Review) GetLink() link {
	return r.link
}

func (r *Review) GetReasons() reasons {
	return r.reasons
}

func (r *Review) GetReport() report {
	return r.report
}

func (r *Review) GetSkill() skill {
	return r.skill
}

func (r *Review) GetUserName() userName {
	return r.userName
}

// VO: レビューid
type reviewID int

func newreviewID(value *int) (reviewID, error) {
	if value == nil {
		return reviewID(0), fmt.Errorf("reviewID is empty")
	}
	return reviewID(*value), nil
}

func (r *reviewID) Int() int {
	return int(*r)
}

// VO: コンテンツ
type content string

func newcontent(value *string) content {
	if value == nil || *value == "" {
		return content("無記入")
	}
	return content(*value)
}

func (c content) String() string {
	return string(c)
}

// VO: リンク
type link string

func newlink(value *string) link {
	if value == nil || *value == "" {
		return link("")
	}
	if isUrl(*value) {
		return link(*value)
	}
	// invalid
	return link("")
}

func isUrl(str string) bool {
	u, err := url.Parse(str)
	return err == nil && u.Scheme != "" && u.Host != ""
}

func (l link) String() string {
	return string(l)
}

// VO: 応募した理由
type reasons string

func newreasons(value *string) reasons {
	if value == nil || *value == "" {
		return reasons("無記入")
	}
	return reasons(*value)
}

func (c reasons) String() string {
	return string(c)
}

// VO: 応募した理由
type report string

func newreport(value *string) report {
	if value == nil || *value == "" {
		return report("無記入")
	}
	return report(*value)
}

func (r report) String() string {
	return string(r)
}

// VO: スキル
type skill string

func newskill(value *string) skill {
	if value == nil || *value == "" {
		return skill("無記入")
	}
	return skill(*value)
}

func (r skill) String() string {
	return string(r)
}

// VO: ユーザーネーム
type userName string

func newuserName(value *string) userName {
	if value == nil || *value == "" {
		return userName("名無しの天才エンジニア")
	}
	return userName(*value)
}

func (r userName) String() string {
	return string(r)
}

func NewReview(
	id *int,
	CompanyName *string,
	Content *string,
	CreateDataJs *string,
	Link *string,
	Reasons *string,
	Report *string,
	Skill *string,
	UserName *string,
) (Review, error) {
	rid, err := newreviewID(id)
	if err != nil {
		return Review{}, err
	}
	companyName, err := newcompanyName(CompanyName)
	if err != nil {
		return Review{}, err
	}
	content := newcontent(Content)
	createDate := newcreateDate(CreateDataJs)
	link := newlink(Link)
	reasons := newreasons(Reasons)
	report := newreport(Report)
	skill := newskill(Skill)
	userName := newuserName(UserName)
	return Review{
		id:           rid,
		companyName:  companyName,
		content:      content,
		createDateJS: createDate,
		link:         link,
		reasons:      reasons,
		report:       report,
		skill:        skill,
		userName:     userName,
	}, nil
}
