package utils

import (
	"encoding/json"
	"io"
	"net/http"
)

func ParseBody(r *http.Request, x interface{}) {
	// 读取整个请求主体
	if body, err := io.ReadAll(r.Body); err == nil {
		// 尝试将JSON主题反序列化到提供的接口中
		if err := json.Unmarshal([]byte(body), x); err != nil {
			return
		}
	}
}
