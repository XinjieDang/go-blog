package response

type ArticleDetailResp struct {
	//文章标题
	Title string `json:"title"`
	//作者
	Author string `json:"author"`
	//分类名称
	Category string `json:"category"`
}
