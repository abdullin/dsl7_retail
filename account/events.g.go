package account

import "github.com/abdullin/dsl7_retail/lang/author"

// abdullin: generated from DSLv7

// Registered -
type Registered struct {
	AuthorInfo  author.Info `json:"authorInfo"`
	AccountName string      `json:"accountName"`
}

func (r *Registered) validate() (err error) {
	return nil
}
func NewRegistered(authorInfo author.Info, accountName string) *Registered {
	return &Registered{authorInfo, accountName}
}
