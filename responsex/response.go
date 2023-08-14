package responsex

// Response http 响应
type Response interface {
	Fail(error)
	Ok(any)
	OkPage(any, int64)
}
