package interfaces

type IBook interface {
	SetTitle(string)
	SetAuthor(string)
	SetPage(uint)
	SetPrice(float64)
	GetTitle() string
	GetAuthor() string
	GetPage() uint
	GetPrice() float64
	Print()
}
