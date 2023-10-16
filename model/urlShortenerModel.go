package model

type Shortener struct {
	Id       int64  `json:"id" gorm:"primay_key"`
	BaseUrl  string `json:"baseUrl" gorm:"not null;unique"`
	Redirect string `json:"redirect" gorm:"not null"`
	Clicked  int64  `json:"clicked"`
	User     User   `gorm:"references:Id"`
	UserId   int64  `json:"userId" gorm:"not null"`
}
