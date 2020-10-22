package controller

import (
	"context"
	"encoding/json"
	"fmt"
	pb "tag/grpc/user"
	tag "tag/module"
	r "tag/resp"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// TagServer -
type UserServer struct {
}

func exampleToModel(in *pb.UserInfo) []byte {
	type Person struct {
		Id   int    `json:"id"`
		Name string `json:name`
	}

	data := []byte(`{"id" : 1 , "name" : "josh" }`)
	var person Person
	json.Unmarshal(data, &person)
	jsondata, _ := json.Marshal(person)
	return jsondata
}

func ToJson(in *pb.UserInfo) []byte {

	type Person struct {
		Id   int
		Name string
	}
	var person = []Person{
		{Id: 1, Name: in.Username},
	}

	// sturct 轉換json字串
	data, _ := json.Marshal(person)

	return data

}

func createToModel(in *pb.Slide) *tag.Slide {

	return &tag.Slide{
		Data: datasToModel(in.GetData()),
	}
}

func datasToModel(in []*pb.UserInfo) []*tag.Tag {
	m := make([]*tag.Tag, len(in))

	for i, v := range in {
		m[i] = dataToModel(v)
	}

	return m
}

func dataToModel(in *pb.UserInfo) *tag.Tag {
	return &tag.Tag{
		Name:        in.GetUsername(),
		Nationality: in.GetIsorc(),
		ID:          in.GetId(),
		PhoneNumber: in.GetPhoneNumber(),
		BirthDay:    in.GetBirthDay(),
		Mail:        in.GetMail(),
	}
}

// Creat -
func (s *UserServer) Create(ctx context.Context, in *pb.Slide) (*pb.StatusReply, error) {
	// jsondata := ToJson(in)
	// fmt.Println(string(jsondata))
	tagModel := createToModel(in)

	// // 確認是否重複
	num, err := tag.CheckRepeat(tagModel.Data[0].Name)
	if err != nil {
		fmt.Println(r.CHECKREPEATERROR)
		// return r.CHECKREPEATERROR
	}
	if num > 0 {
		fmt.Println(r.REPEATERROR)
		return &pb.StatusReply{Code: 0, Msg: "名稱重複，請更改名稱"}, status.Error(codes.OK, "success")
	}

	data, Err := tag.Creat(tagModel.Data[0])
	fmt.Println(data)
	if Err != nil {
		fmt.Println(r.E(err, 1))
	}

	// if err != nil {
	// 	return &pb.StatusReply{Code: 0, Msg: "Error"}, status.Error(codes.OK, "success")

	// }
	return &pb.StatusReply{Code: 0, Msg: data}, status.Error(codes.OK, "success")

	// return e.JSON(http.StatusOK, r.R(err))
}
