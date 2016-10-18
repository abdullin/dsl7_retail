package user

import "github.com/abdullin/dsl7/lang/author"

// abdullin: generated from DSLv7

// Added -
type Added struct {
	AuthorInfo author.Info `json:"authorInfo"`
	UserId     uint64      `json:"userId"`
	UserName   string      `json:"userName"`
	UserEmail  string      `json:"userEmail"`
}

func (r *Added) validate() (err error) {
	return nil
}

// Renamed -
type Renamed struct {
	AuthorInfo author.Info `json:"authorInfo"`
	UserName   string      `json:"userName"`
}

func (r *Renamed) validate() (err error) {
	return nil
}

// Deleted -
type Deleted struct {
	AuthorInfo author.Info `json:"authorInfo"`
	UserName   string      `json:"userName"`
	UserEmail  string      `json:"userEmail"`
}

func (r *Deleted) validate() (err error) {
	return nil
}
func NewAdded(authorInfo author.Info, userId uint64, userName string, userEmail string) *Added {
	return &Added{authorInfo, userId, userName, userEmail}
}
func NewRenamed(authorInfo author.Info, userName string) *Renamed {
	return &Renamed{authorInfo, userName}
}
func NewDeleted(authorInfo author.Info, userName string, userEmail string) *Deleted {
	return &Deleted{authorInfo, userName, userEmail}
}
