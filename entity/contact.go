package entity

type Contact struct {
	ID          interface{} `json:"id,omitempty" bson:"_id,omitempty"`
	Fullname    string      `json:"fullname"`
	Address     string      `json:"address"`
	PhoneNumber string      `json:"phonenumber"`
	Email       string      `json:"email"`
}
