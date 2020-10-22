package tag

import (
	"context"
	"fmt"
	"log"
	"strconv"
	r "tag/resp"

	"github.com/axolotlteam/thunder/db/mongov2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func CheckRepeatProduct(name string) (int, error) {
	m, err := mongov2.M()
	if err != nil {
		return 0, r.MONGOERROR
	}

	c := m.Database(D).Collection(P)
	filter := bson.M{
		"productid": name,
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

// Insert -
func Insert(product *Product) error {
	m, err := mongov2.M()
	if err != nil {
		return r.MONGOERROR
	}

	c := m.Database(D).Collection("product")
	fmt.Println(product)
	// 插入多条数据
	_, err = c.InsertOne(context.Background(), product)
	if err != nil {
		return r.INSERTERROR
	}
	return nil
}

// Get -
func Get(page, size string) ([]*TagListBinder, error) {
	m, err := mongov2.M()
	if err != nil {
		return nil, r.MONGOERROR
	}

	c := m.Database(D).Collection(C)
	p, err := strconv.ParseInt(page, 10, 64) //string to int64
	if err != nil {
		return nil, err
	}

	s, err := strconv.ParseInt(size, 10, 64) //string to int64
	if err != nil {
		return nil, err
	}

	skip := s * (p - 1)

	findOptions := options.Find()
	findOptions.SetLimit(s)   //顯示筆數
	findOptions.SetSkip(skip) //頁碼

	cur, err := c.Find(context.Background(), bson.M{}, findOptions)
	if err != nil {
		return nil, r.FINDERROR
	}

	defer cur.Close(context.Background())
	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	var t []*TagListBinder

	for cur.Next(context.Background()) {

		var k TagListBinder
		err := cur.Decode(&k)
		if err != nil {
			// log.Fatal(err)
			return nil, r.FINDDECODEERROR
		}

		t = append(t, &k)
	}

	return t, nil
}

// GetOne -
func GetOne(id string) (*TagListBinder, error) {
	i, err := checkID(id)
	if err != nil {
		return nil, r.OBJECTIDERROR
	}

	m, err := mongov2.M()
	if err != nil {
		return nil, r.MONGOERROR
	}

	c := m.Database(D).Collection(C)

	t := &TagListBinder{}
	err = c.FindOne(context.Background(), bson.M{"_id": i}).Decode(t)
	if err != nil {
		return nil, r.FINDONEERROR
	}

	return t, nil
}

// Update -
func Update(id string, product *Product) error {
	// i, err := checkID(id)
	// if err != nil {
	// 	return r.OBJECTIDERROR
	// }

	m, err := mongov2.M()
	if err != nil {
		return r.MONGOERROR
	}

	c := m.Database(D).Collection(P)

	filter := bson.M{
		"productid": id,
	}

	update := bson.M{
		"$set": bson.M{
			"name": product.Name,
		},
	}

	_, err = c.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return r.UPDATEEERROR
	}

	return nil
}

// Delete -
func Delete(id string) error {
	fmt.Println(id)
	// i, err := checkID(id)
	// fmt.Println(i)
	// if err != nil {
	// 	return r.OBJECTIDERROR
	// }

	m, err := mongov2.M()
	if err != nil {
		return r.MONGOERROR
	}

	c := m.Database(D).Collection(P)

	_, err = c.DeleteOne(context.Background(), bson.M{"productid": id})
	if err != nil {
		return r.DELETEERROR
	}

	return nil
}

// Search -
func Search(code string) ([]*TagListBinder, error) {
	m, err := mongov2.M()
	if err != nil {
		return nil, r.MONGOERROR
	}

	c := m.Database(D).Collection(P)
	// 修改常數
	filter := bson.M{"productid": primitive.Regex{Pattern: code}} // 模糊查詢

	cur, err := c.Find(context.Background(), filter)
	if err != nil {
		return nil, r.SEARCHERROR
	}

	defer cur.Close(context.Background())
	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	var t []*TagListBinder

	for cur.Next(context.Background()) {

		var k TagListBinder
		err := cur.Decode(&k)
		fmt.Print(k)
		if err != nil {
			// log.Fatal(err)
			return nil, r.FINDDECODEERROR
		}

		t = append(t, &k)
	}

	return t, nil
}
