package dto

// OUTPUT DATAFLOW
type ResponseError struct {
	Errors []string `json:"errors"`
}

type ResponseDetail struct {
	Detail string `json:"detail"`
}
