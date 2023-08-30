package dto

type DashboardResponseBody struct {
	TotalUser    int64 `json:"total_user"`
	TotalProduct int64 `json:"total_product"`
	TotalOrder   int64 `json:"total_order"`
	TotalAdmin   int64 `json:"total_admin"`
}
