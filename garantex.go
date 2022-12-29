package tg

type Asks []struct {
	Price  string `json:"price"`
	Volume string `json:"volume"`
	Amount string `json:"amount"`
	Factor string `json:"factor"`
	Type   string `json:"type"`
}

type Bids []struct {
	Price  string `json:"price"`
	Volume string `json:"volume"`
	Amount string `json:"amount"`
	Factor string `json:"factor"`
	Type   string `json:"type"`
}

type DepthResponse struct {
	Timestamp int `json:"timestamp"`
	Asks      `json:"asks"`
	Bids      `json:"bids"`
}
