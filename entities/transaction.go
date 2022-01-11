package entities

type TXStatus string
const (
	Pending  TXStatus = "PENDING"
	Success = "SUCCESS"
	Failed = "FAILED"
)

type Transaction struct{
	TransactionId string `json:"tx_id"`
	TransactionURL string `json:"tx_url"`
	CointType string `json:"coin"`
	Status TXStatus `json:"status"`
}


// Inmemory repository for the awaiting transactions
var (
	Transactions = map[int]*Transaction{}
	seq   = 1
 )