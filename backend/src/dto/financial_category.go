package dto

type CreateFinancialCategoryRequest struct {
	Name          string `json:"name"`
	OperationType string `json:"operation_type"`
}

type UpdateFinancialCategoryRequest struct {
	Name          string `json:"name"`
	OperationType string `json:"operation_type"`
}

type FinancialCategoryResponse struct {
	ID            int64  `json:"id"`
	Name          string `json:"name"`
	OperationType string `json:"operation_type"`
}

type FinancialCategoryListResponse struct {
	Items    []FinancialCategoryResponse `json:"items"`
	Page     int                         `json:"page"`
	PageSize int                         `json:"page_size"`
	Total    int64                       `json:"total"`
}
