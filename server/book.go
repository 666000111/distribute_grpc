package server

import (
	"context"
	"grpc_server/dal"
	"grpc_server/model"
	"grpc_server/proto_gen/grpc_server/book"
	"sync"
)

type BookServer struct {
	book.UnimplementedBookServer
	mu sync.Mutex
}

func (h *BookServer) AddBook(c context.Context, req *book.AddBookReq) (*book.AddBookResp, error) {
	b := &model.Book{
		Name:    req.Name,
		Content: req.Content,
	}
	dal.BookDB.Select("Name", "Content").Create(b)
	resp := &book.AddBookResp{
		Book: &book.BookInfo{
			Id:      b.ID,
			Name:    b.Name,
			Content: b.Content,
		},
	}
	return resp, nil

}
func (h *BookServer) ModifyBook(c context.Context, req *book.ModifyBookReq) (*book.ModifyBookResp, error) {
	b := &model.Book{ID: req.Id}
	dal.BookDB.First(b)
	b.Content = req.Content
	b.Name = req.Name
	dal.BookDB.Save(b)
	resp := &book.ModifyBookResp{
		Book: &book.BookInfo{
			Id:      b.ID,
			Name:    b.Name,
			Content: b.Content,
		},
	}
	return resp,nil
}

func (h *BookServer) QueryBook(c context.Context, req *book.QueryBookReq) (*book.QueryBookResp, error) {
	db := dal.BookDB
	if len(req.Name) > 0 {
		db = db.Where("name in", req.Name)
	}
	var books []model.Book
	db.Find(&books, req.Id)
	resp := &book.QueryBookResp{}
	result := make([]*book.BookInfo, 0)
	for _, v := range books {
		result = append(result, &book.BookInfo{
			Id:      v.ID,
			Name:    v.Name,
			Content: v.Content,
		})
	}
	resp.Books = result
	return resp, nil
}
func (h *BookServer) DeleteBook(c context.Context, req *book.DeleteBookReq) (*book.DeleteBookResp, error) {
	dal.BookDB.Delete(&model.Book{}, req.Id)
	return &book.DeleteBookResp{},nil
}
