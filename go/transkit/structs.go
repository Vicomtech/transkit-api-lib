//
// Transkit API libraries
// Coded by Aritz Olea <aolea@vicomtech.org>
// (c) 2022 Vicomtech - vicomtech.org
//

package transkit

//
//  Transkit ONLINE data structures
//

type OnlineTranscriptionRequest struct {
	Audio    string   `json:"audio"`
	Pipeline string   `json:"pipeline"`
	Tags     []string `json:"tags"`
}

type OnlineTranscriptionResponse struct {
	Error string       `json:"error,omitempty"`
	Text  string       `json:"text"`
	Words []OnlineWord `json:"words,omitempty"`
	ConsumedQuotaResponse
}

type OnlineWord struct {
	Confidence float64 `json:"confidence"`
	Length     float64 `json:"length"`
	Start      float64 `json:"start"`
	Word       string  `json:"word"`
}

//
// Consumed quota response
//

type ConsumedQuotaResponse struct {
	ConsumedQuota int `json:"consumedquota"`
}

//
// Transkit OFFLINE data structures
//

type OfflineTranscriptionRequest struct {
	Pipeline string      `json:"pipeline"`
	URL      string      `json:"url"`
	Config   interface{} `json:"config,omitempty"`
	Tags     []string    `json:"tags"`
}

type OfflineTranscriptionResponse struct {
	Id string `json:"id"`
	ConsumedQuotaResponse
}

type OfflineTranscriptionStatusResponse struct {
	Status  int           `json:"status"`
	Message string        `json:"message"`
	Error   *OfflineError `json:"error,omitempty"`
	URL     string        `json:"url,omitempty"`
}

type OfflineError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Module  string `json:"module"`
}
