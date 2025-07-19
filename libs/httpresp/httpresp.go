package httpresp

type Response struct {
	Data  any    `json:"data,omitempty"`
	Error string `json:"error,omitempty"`
}

func OK(data any) Response {
	return Response{
		Data: data,
	}
}

func Err(err error) Response {
	return Response{
		Error: err.Error(),
	}
}
