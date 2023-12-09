package entity

type SearchResults struct {
	Posts []*Post
	Users []*UserSearch
}

func NewSearchResults(posts []*Post, users []*UserSearch) *SearchResults {
	return &SearchResults{
		Posts: posts,
		Users: users,
	}
}
