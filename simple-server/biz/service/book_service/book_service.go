package book_service

import (
	"log/slog"

	"simple-server/dao/book_dao"

	"github.com/go-spring/spring-core/gs"
)

func init() {
	gs.Object(&BookService{})
}

type BookService struct {
	BookDao *book_dao.BookDao `autowire:""`
	Logger  *slog.Logger      `autowire:"biz"`
}

func (s *BookService) ListBooks() ([]book_dao.Book, error) {
	return s.BookDao.ListBooks()
}

func (s *BookService) GetBook(sn string) (book_dao.Book, error) {
	return s.BookDao.GetBook(sn)
}

func (s *BookService) SaveBook(book book_dao.Book) error {
	return s.BookDao.SaveBook(book)
}

func (s *BookService) DeleteBook(sn string) error {
	return s.BookDao.DeleteBook(sn)
}
