package result

type WordArray struct {
	Words []string `json:"Words"`
}

type ImageData struct {
	ID             int       `json:"ID"`
	Url            string    `json:"Url"`
	Language       string    `json:"Language"`
	Suggestion     WordArray `json:"Suggestion"`
	MarkResult     string    `json:"MarkResult"`
	QualityResult  string    `json:"QualityResult"`
	CoordinateJson string    `json:"CoordinateJson"`
	Status         int       `json:"Status"`
}

type ImageDataResult struct {
	Total int       `json:"Total"`
	Data  ImageData `json:"Data"`
}
