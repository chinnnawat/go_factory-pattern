package models

type Commics struct {
	Book
	PostCard IPostCard
}

type IPostCard interface {
	setPostCard(isPostCard bool)
	getPostCard() bool
}

type PostCard struct {
	IsPostCard bool
}

// Implement IPostCard methods for PostCard
func (p *PostCard) setPostCard(isPostCard bool) {
	p.IsPostCard = isPostCard
}

func (p *PostCard) getPostCard() bool {
	return p.IsPostCard
}
