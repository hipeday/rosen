package response

// GeneralPageResponse 通用页面响应
type GeneralPageResponse struct {
	Lang  string `json:"lang"`  // 页面语言
	Title string `json:"title"` // 页面标题
}
