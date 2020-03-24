package model

type Answer struct {
	Status   string `json:"status"`
	Customer `json:"customer"`
}

type Customer struct {
	ProcessingStatus string `json:"processingStatus"`
	Ids              `json:"ids"`
}

type Ids struct {
	MindboxId        int64
	StoreezWebsiteId string
}
