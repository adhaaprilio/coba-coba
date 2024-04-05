package helper

type Paginate struct {
	Page      int `json:"page"`
	PerPage   int `json:"perpage"`
	Total     int `json:"total"`
	TotalPage int `json:"totalpage"`
}

type ResponseParams struct {
	StatusCode int
	Message    string
	Paginate   *Paginate
	Data       any
}

type ResponseWithData struct {
	Message string `json:"message"`
	Data    any    `json:"data"`
}

type ResponseWithoutData struct {
	Message string `json:"message"`
}

func Response(params ResponseParams) any {
	var response any

	if params.Data != nil {
		response = &ResponseWithData{
			Message: params.Message,
			Data:    params.Data,
		}
	} else {
		response = &ResponseWithoutData{
			Message: params.Message,
		}
	}
	return response
}
