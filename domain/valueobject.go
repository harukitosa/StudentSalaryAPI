package domain

// VO: 投稿日時
// 色々あってJSのdateが入っていたりいなかったり
type createDate string

func newcreateDate(value *string) createDate {
	if value == nil {
		return createDate("")
	}
	return createDate(*value)
}

func (c createDate) String() string {
	return string(c)
}

// VO: 詳細内容
type workdetail string

func newworkdetail(value *string) workdetail {
	if value == nil || *value == "" {
		return workdetail("なし")
	}
	return workdetail(*value)
}

func (c workdetail) String() string {
	return string(c)
}
