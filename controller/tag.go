package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	pb "tag/grpc"
	tag "tag/module"
	r "tag/resp"

	"github.com/labstack/echo/v4"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// TagServer -
type UserServer struct {
}
type ProdectServer struct {
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

func createToModel(in *pb.UserInfo) *tag.Tag {

	return &tag.Tag{

		Name:        in.GetUsername(),
		Nationality: in.GetIsorc(),
		ID:          in.GetId(),
		PhoneNumber: in.GetPhoneNumber(),
		BirthDay:    in.GetBirthDay(),
		Mail:        in.GetMail(),
	}
}

func productToModel(in *pb.ProductInfo) *tag.Product {
	return &tag.Product{
		ProductID: in.GetProductID(),
		Brand:     in.GetBrand(),
		Name:      in.GetName(),
		Old:       in.GetOld(),
		Label:     in.GetLabel(),
		Color:     in.GetColor(),
		SalePrice: in.GetSalePrice(),
		Price:     in.GetPrice(),
		Size:      in.GetSize_(),
		Sum:       in.GetSum(),
	}
}

func ProductUpdateToModel(in *pb.ProductInfo) *tag.Product {
	return &tag.Product{
		ProductID: in.GetProductID(),
		Brand:     in.GetBrand(),
		Name:      in.GetName(),
		Old:       in.GetOld(),
		Label:     in.GetLabel(),
		Color:     in.GetColor(),
		SalePrice: in.GetSalePrice(),
		Price:     in.GetPrice(),
		Size:      in.GetSize_(),
		Sum:       in.GetSum(),
	}
}

func ProdectDeleteToModel(in *pb.ProductInfo) *tag.Product {
	return &tag.Product{
		ProductID: in.GetProductID(),
	}
}

func SearchToModel(in *pb.ProductInfo) *tag.Product {
	return &tag.Product{
		ProductID: in.GetProductID(),
	}
}

// Creat -
func (s *UserServer) Create(ctx context.Context, in *pb.UserInfo) (*pb.StatusReply, error) {
	// jsondata := ToJson(in)
	// fmt.Println(string(jsondata))
	tagModel := createToModel(in)

	if tagModel.Name == "" {
		return &pb.StatusReply{Code: 0, Msg: "空值，請輸入名稱"}, status.Error(codes.OK, "success")
	}
	// // 確認是否重複
	num, err := tag.CheckRepeat(tagModel.Name)
	if err != nil {
		fmt.Println(r.CHECKREPEATERROR)
		// return r.CHECKREPEATERROR
	}
	if num > 0 {
		fmt.Println(r.REPEATERROR)
		return &pb.StatusReply{Code: 0, Msg: "名稱重複，請更改名稱"}, status.Error(codes.OK, "success")
	}

	data, Err := tag.Creat(tagModel)
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

// Post
func (s *ProdectServer) Insert(ctx context.Context, in *pb.ProductInfo) (*pb.StatusReply, error) {
	// new(tag) same &tag
	productModel := productToModel(in)
	// Get product data

	// 確認產品序號是否重複
	num, err := tag.CheckRepeatProduct(productModel.ProductID)
	if err != nil {
		fmt.Print(r.CHECKREPEATERROR)
		// return r.CHECKREPEATERROR
	}
	if num > 0 {
		fmt.Println(r.REPEATERROR)
		// return r.REPEATERROR //目前問題是當有重複時返回 nternal Server Error, 500
	}

	resultErr := tag.Insert(productModel)
	if resultErr != nil {
		fmt.Println(r.E(err, 1))
	}
	// fmt.Println(result)
	return &pb.StatusReply{Code: 0, Msg: "success"}, status.Error(codes.OK, "success")

}

// GetOne -
func GetOne(e echo.Context) error {
	id := strings.TrimSpace(e.Param("id"))

	t, err := tag.GetOne(id)
	if err != nil {
		return e.JSON(http.StatusOK, r.E(err, 1))
	}

	return e.JSON(http.StatusOK, r.R(t))
}

// Update -
func (s *ProdectServer) Update(ctx context.Context, in *pb.ProductInfo) (*pb.StatusReply, error) {

	// pb convert model
	productModel := ProductUpdateToModel(in)

	// 確認是否重複
	num, err := tag.CheckRepeatProduct(productModel.Name)
	if err != nil {
		fmt.Println(r.CHECKREPEATERROR)
	}
	if num > 0 {
		fmt.Println(r.REPEATERROR)
	}

	err = tag.Update(productModel.ProductID, productModel)
	if err != nil {
		fmt.Println(r.E(err, 1))
		// return e.JSON(http.StatusOK, r.E(err, 1))
	}
	return &pb.StatusReply{Code: 0, Msg: "success"}, status.Error(codes.OK, "success")
}

// Delete -
func (s *ProdectServer) Delete(ctx context.Context, in *pb.ProductInfo) (*pb.StatusReply, error) {

	productdata := ProdectDeleteToModel(in)
	fmt.Println(productdata)
	// 執行刪除
	err := tag.Delete(productdata.ProductID)
	fmt.Println(productdata.ProductID)
	if err != nil {
		fmt.Print(r.E(err, 1))
	}
	return &pb.StatusReply{Code: 0, Msg: "success"}, status.Error(codes.OK, "success")
}

// Search -
func (s *ProdectServer) Search(ctx context.Context, in *pb.ProductInfo) (*pb.StatusReply, error) {
	searchdata := SearchToModel(in)
	// 相似度查詢
	t, err := tag.Search(searchdata.ProductID)
	if err != nil {
		return &pb.StatusReply{Code: 0, Msg: "success"}, status.Error(codes.OK, "success")

	}

	if len(t) > 0 {
		// 回傳使用者名稱
		return &pb.StatusReply{Code: 0, Msg: t[0].Name}, status.Error(codes.OK, "success")
	} else {
		return &pb.StatusReply{Code: 0, Msg: "查無資料"}, status.Error(codes.OK, "success")
	}

}
