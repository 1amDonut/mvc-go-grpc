package tag

import (
	"context"
	"fmt"
	"strconv"
	r "tag/resp"

	"github.com/axolotlteam/thunder/db/mongov2"
	"go.mongodb.org/mongo-driver/bson"
)

// CheckRepeat -
func CheckRepeat(name string) (int, error) {
	m, err := mongov2.M()
	if err != nil {
		return 0, r.MONGOERROR
	}

	c := m.Database(D).Collection(C)
	filter := bson.M{
		"name": name,
	}
	num, err := c.CountDocuments(context.Background(), filter)
	if err != nil {
		return 0, r.MONGOERROR
	}

	count1 := strconv.FormatInt(num, 10) // int64 to string
	count2, err := strconv.Atoi(count1)  // string to int
	if err != nil {
		return 0, r.STRCONVERROR
	}

	return count2, nil
}

// Creat -
func Creat(tag *Tag) (string, error) {
	m, err := mongov2.M()
	if err != nil {
		return "", r.MONGOERROR
	}

	c := m.Database(D).Collection(C)
	fmt.Println("insert data")
	// 插入多條数据
	data, Err := c.InsertOne(context.Background(), tag)
	if Err != nil {
		return "", r.INSERTERROR
	}

	// data.InsertedID.(string)
	// msg := data.InsertedID.(string)
	fmt.Println(data)
	return tag.Name, nil
}
