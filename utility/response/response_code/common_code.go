// 返回数据code码定义：公共code码
// 创建人： 黄翠刚
// 创建时间： 2022.10.09

package response_code

type code struct {
	Code    int    // 响应码
	Message string // 错误信息
}

var (
	Ok    = code{200, "请求成功！"}
	Error = code{-1, "请求失败！"}
)
