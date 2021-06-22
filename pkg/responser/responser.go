package responser

import "encoding/json"

type Body struct {
	Code int64 `json:"code"`
	Message string `json:"message"`
}

func GenerateJSON(b *Body) string {
	jsons, err := json.Marshal(b)
	if err != nil {
		panic(err)
	}
	return string(jsons)
}