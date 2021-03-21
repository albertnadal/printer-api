package models

type WorkerStats struct {
	TotalProcessedOps uint64 `json:"totalProcessedOps"`
	ProcessingSpeed int `json:"processingSpeed"` //operations processed per minute
}

type AggregationAPIStats struct {
	TotalReceivedOps uint64 `json:"totalReceivedOps"` //total api calls(operations) received
	TotalProcessedOps uint64 `json:"totalProcessedOps"` //total operations processed
	TotalOpsAwaiting uint64 `json:"totalOpsAwaiting"` //total operations waiting to be processed
	ProcessingSpeed string `json:"processingSpeed"` //current operations processed per minute
	ProcessingTimeLeft string `json:"processingTimeLeft"` //time left to process all awaiting operations
}
