package httpresp

type Response struct {
	Data  any   `json:"data"`
	Error error `json:"error,omitempty"`
}

func OK(data any) Response {
	return Response{
		Data:  data,
		Error: nil,
	}
}

func Err(err error) Response {
	return Response{
		Data:  nil,
		Error: err,
	}
}
