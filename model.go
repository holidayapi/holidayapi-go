package holidayapi

type Holiday struct {
	Name     string `json:"name"`
	Date     string `json:"date"`
	Observed string `json:"observed"`
	Public   string `json:"public"`
}

type ResponseHoliday struct {
	Status int `json:"status"`
	Error string `json:"error"`
	Holidays map[string][]Holiday `json:"holidays"`
}
