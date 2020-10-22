package tag

import (
	r "tag/resp"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	D = "local"
	C = "user"
	P = "product"
)

type TagListBinder struct {
	ID   primitive.ObjectID `json:"_id" bson:"_id"`
	Name string             `json:"name" bson:"name"`
	Page string
	Size string
	Skip string
}

// Tag -
type Tag struct {
	Name        string `json:"name"`
	Nationality int32  `json:"nationality"` //isROC
	ID          string `json:"id"`
	PhoneNumber string `json:"phoneNumber"`
	BirthDay    string `json:"birthDay"`
	Mail        string `json:"mail"`
}

// Slide - 主要結構
type Slide struct {
	Data []*Tag `json:"data"`
}

// User - 賣家資訊
type User struct {
	Name        string `json:"name"`
	Nationality int32  `json:"nationality"` //isROC
	ID          string `json:"id"`
	PhoneNumber string `json:"phoneNumber"`
	BirthDay    string `json:"birthDay"`
	Mail        string `json:"mail"`
}

// Product - 商品資訊
type Product struct {
	ProductID string  `json:"productid" validate:"required"`
	Brand     string  `json:"brand" validate:"required"`
	Name      string  `json:"name" validate:"required"`
	Old       string  `json:"old" validate:"required"`
	Label     string  `json:"label"`
	Color     string  `json:"color"`
	SalePrice float64 `json:"salePrice" validate:"required"`
	Price     float64 `json:"price" validate:"required"`
	Size      string
	Sum       int32
}

// checkID -
func checkID(id string) (primitive.ObjectID, error) {
	i, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return i, r.FORMATERROR
	}
	return i, nil
}
