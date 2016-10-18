package supplier

// abdullin: generated from DSLv7

// Info - Information about product supplier
type Info struct {
	SupplierId        uint64 `json:"supplierId"`
	SupplierName      string `json:"supplierName"`
	SupplierPreferred bool   `json:"supplierPreferred"`
}

func (r *Info) validate() (err error) {
	return nil
}
func NewInfo(supplierId uint64, supplierName string, supplierPreferred bool) *Info {
	return &Info{supplierId, supplierName, supplierPreferred}
}
