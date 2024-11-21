package usecases

import (
	"encoding/json"
	"fmt"

	"example.com/m/v2/interfaces"
	"example.com/m/v2/models"
)

type LocalCommic struct {
	*models.Commics
}

type PostCard struct {
	IsPostCard bool
}

func (b *LocalCommic) SetTitle(title string) {
	b.Title = title
}

func (b *LocalCommic) SetAuthor(author string) {
	b.Author = author
}

func (b *LocalCommic) SetPage(page uint) {
	b.Page = page
}

func (b *LocalCommic) SetPrice(price float64) {
	b.Price = price
}

func (b *LocalCommic) GetTitle() string {
	return b.Title
}

func (b *LocalCommic) GetAuthor() string {
	return b.Author
}

func (b *LocalCommic) GetPage() uint {
	return b.Page
}

func (b *LocalCommic) GetPrice() float64 {
	return b.Price
}

func (b *LocalCommic) Print() {
	obj, _ := json.MarshalIndent(b, "", "  ")
	fmt.Println(string(obj))
}
func NewCommic() interfaces.IBook {
	return &LocalCommic{
		Commics: &models.Commics{
			Book: models.Book{
				Title:  "commic",
				Author: "john doe",
				Page:   10,
				Price:  100,
			},
			PostCard: &models.PostCard{
				IsPostCard: true,
			},
		},
	}
}
