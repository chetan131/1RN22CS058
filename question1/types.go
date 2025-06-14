package main

type NumberResponse struct {
	Numbers []int `json:"numbers"`
}

type OutputResponse struct {
	WindowPrevState []int   `json:"windowPrevState"`
	WindowCurrState []int   `json:"windowCurrState"`
	Numbers         []int   `json:"numbers"`
	Avg             float64 `json:"avg"`
}
