package web

type ApiResult struct {
	Success bool
	Msg     string
	Data    interface{}
}

func DefaultApiResult() ApiResult {
	return ApiResult{
		false,
		"",
		nil,
	}
}
