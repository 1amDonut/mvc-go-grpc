package controller

import (
	"context"
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
type TagServer struct {
}

func createToModel(in *pb.UserInfo) {

}

// Creat -
func (s *TagServer) Create(ctx context.Context, in *pb.UserInfo) (*pb.StatusReply, error) {
	// Get user and member from the query string
	// u := new(tag.Tag)
	// // Bind 屬於url form-data 方式
	// if err := e.Bind(u); err != nil {
	// 	return r.FORMATERROR
	// }
	// // 確認是否重複
	// num, err := tag.CheckRepeat(u.Name)
	// if err != nil {
	// 	return r.CHECKREPEATERROR
	// }
	// if num > 0 {
	// 	return r.REPEATERROR //	目前問題是當有重複時會返回 Internal Server Error, 500
	// }

	// err = tag.Creat(u)
	// if err != nil {
	// 	return e.JSON(http.StatusOK, r.E(err, 1))
	// }
	fmt.Print(ctx)
	// return e.JSON(http.StatusOK, r.R(err))
	return &pb.StatusReply{Code: 0, Msg: "Succes"}, status.Error(codes.OK, "success")
}

// Post
func Insert(e echo.Context) error {

	// Get product data
	u := new(tag.Product)

	if err := e.Bind(u); err != nil {
		return r.FORMATERROR
	}
	// 驗證有無空值
	if e := e.Validate(u); e != nil {
		return r.FORMATERROR
	}
	// 確認產品序號是否重複
	num, err := tag.CheckRepeatProduct(u.ProductID)
	if err != nil {
		return r.CHECKREPEATERROR
	}
	if num > 0 {
		return r.REPEATERROR //目前問題是當有重複時返回 nternal Server Error, 500
	}

	err = tag.Insert(u)
	if err != nil {
		return e.JSON(http.StatusOK, r.E(err, 1))
	}

	return e.JSON(http.StatusOK, r.R(err))
}

// Get -
func Get(e echo.Context) error {
	l := new(tag.TagListBinder)
	if err := e.Bind(l); err != nil {
		return r.FORMATERROR
	}
	// l.Skip = l.Size * (l.Page - 1)
	t, err := tag.Get(l.Page, l.Size)
	if err != nil {
		return e.JSON(http.StatusOK, r.E(err, 1))
	}

	return e.JSON(http.StatusOK, r.R(t))
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
func Update(e echo.Context) error {
	id := strings.TrimSpace(e.Param("id"))

	u := new(tag.Tag)
	if err := e.Bind(u); err != nil {
		return r.FORMATERROR
	}

	// 確認是否重複
	num, err := tag.CheckRepeat(u.Name)
	if err != nil {
		return r.CHECKREPEATERROR
	}
	if num > 0 {
		return r.REPEATERROR //	目前問題是當有重複時會返回 Internal Server Error, 500
	}

	err = tag.Update(id, u)
	if err != nil {
		return e.JSON(http.StatusOK, r.E(err, 1))
	}

	return e.JSON(http.StatusOK, "update OK!")
}

// Delete -
func Delete(e echo.Context) error {
	id := strings.TrimSpace(e.Param("id"))

	err := tag.Delete(id)
	if err != nil {
		return e.JSON(http.StatusOK, r.E(err, 1))
	}

	return e.JSON(http.StatusOK, "delete OK!")
}

// Search -
func Search(e echo.Context) error {
	code := strings.TrimSpace(e.FormValue("code"))
	t, err := tag.Search(code)
	if err != nil {
		return e.JSON(http.StatusOK, r.E(err, 1))
	}

	return e.JSON(http.StatusOK, r.R(t))
}
