package constants

var Status = map[string]string{
	"pending":  "PENDING",
	"approved": "APPROVED",
	"rejected": "REJECTED",
}

var Topics = map[string]string{
	"pending":  "payment.pending",
	"approved": "payment.approved",
	"rejected": "payment.rejected",
}