package author

// abdullin: generated from DSLv7

// Info - registered person behind the request
type Info struct {
	AccountId uint64 `json:"accountId"`
	UserId    uint64 `json:"userId"`
}

func (r *Info) validate() (err error) {
	return nil
}
func NewInfo(accountId uint64, userId uint64) *Info {
	return &Info{accountId, userId}
}
