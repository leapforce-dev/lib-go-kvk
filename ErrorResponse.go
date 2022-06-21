package kvk

type ErrorResponse struct {
	Fout []struct {
		Code         string `json:"code"`
		Omschrijving string `json:"omschrijving"`
	} `json:"fout"`
}
