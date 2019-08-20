package serializer

// import "go-test/model"

// User 用户序列化器
type User struct {
	ID uint `json:"id"`
	// post           json提交 自动注入
	UserName  string `json:"user_name" xml:"user_name" form:"user_name"`
	Nickname  string `json:"nickname"`
	Status    string `json:"status"`
	Avatar    string `json:"avatar"`
	CreatedAt int64  `json:"created_at"`
}
