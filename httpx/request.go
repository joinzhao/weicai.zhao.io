package httpx

//import (
//	"context"
//	"net/http"
//	"time"
//	"weicai.zhao.io/tools"
//)
//
//func DoJsonGet(url string, body any) (*http.Response, error) {
//	return Do(http.MethodGet, url, body, WithJsonHeader())
//}
//
//func DoJsonPost(url string, body any) (*http.Response, error) {
//	return Do(http.MethodPost, url, body, WithJsonHeader())
//}
//
//// Do do http request
//func Do(method string, url string, body any, handles ...RequestHandle) (*http.Response, error) {
//	req, err := newRequest(method, url, body)
//	if err != nil {
//		return nil, err
//	}
//
//	for _, handle := range handles {
//		handle(req)
//	}
//
//	return http.DefaultClient.Do(req)
//}
//
//type RequestHandle func(*http.Request)
//
//func newRequest(method string, url string, body any) (*http.Request, error) {
//	data, err := tools.AnyReader(body)
//	if err != nil {
//		return nil, err
//	}
//
//	return http.NewRequest(method, url, data)
//}
//
//func WithTimeOut(t time.Duration) RequestHandle {
//	ctx, _ := context.WithTimeout(context.Background(), t)
//	return WithContext(ctx)
//}
//
//func WithContext(ctx context.Context) RequestHandle {
//	return func(req *http.Request) {
//		req.WithContext(ctx)
//	}
//}
//
//func WithHeaderMap(header map[string]string) RequestHandle {
//	h := http.Header{}
//	if header != nil {
//		for key, val := range header {
//			h.Add(key, val)
//		}
//	}
//	return WithHeader(h)
//}
//
//func WithHeader(header http.Header) RequestHandle {
//	return func(req *http.Request) {
//		req.Header = header
//	}
//}
//
//func WithJsonHeader() RequestHandle {
//	h := http.Header{}
//	h.Add("Content-Type", "application/json")
//	return WithHeader(h)
//}
