package geocoder

// Response відображає кореневу структуру JSON відповіді
type Response struct {
	Results []*Result `json:"results"`
	Status  *Status   `json:"status"`
}

type Result struct {
	Formatted *string   `json:"formatted"`
	Geometry  *Geometry `json:"geometry"`
}

type Geometry struct {
	Lat *float64 `json:"lat"`
	Lng *float64 `json:"lng"`
}

type Status struct {
	Code    *int    `json:"code"`
	Message *string `json:"message"`
}
