package customer

type ReqInitCustomer struct {
	CustomerXid string `json:"customer_xid" binding:"required"`
}
