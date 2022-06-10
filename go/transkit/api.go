//
// Transkit API libraries
// Coded by Aritz Olea <aolea@vicomtech.org>
// (c) 2022 Vicomtech - vicomtech.org
//

// Package transkit provides utilities to communicate with Vicomtech Transkit API service
// on online or offline mode.
package transkit

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"path"
	"strings"
)

// Predefined constant paths
const (
	URL_API                           string = "https://api.transkit.vicomtech.org/"
	PATH_ONLINE_TRANSCRIPTION                = "online/transcribe"
	PATH_OFFLINE_TRANSCRIPTION               = "offline/transcribe"
	PATH_OFFLINE_TRANSCRIPTION_STATUS        = "offline/status"
	PATH_GET_QUOTA                           = "quota"
)

// Transkit Main transkit structure for data initialization
type Transkit struct {
	key         string
	apiURL      string
	extraConfig interface{}
}

// Init Initializes transkit object using provided user key
func (tr *Transkit) Init(key string) {
	(*tr).key = key
	(*tr).apiURL = URL_API
}

// Used error list in this package
var (
	errJSONMarshalObject error = errors.New("error creating request bytes")
	errInvalidAPIKey     error = errors.New("invalid API key or expired")
	errInvalidStatusCode error = errors.New("invalid status code")
	errInvalidParameters error = errors.New("invalid parameters")
	errNoAuthData        error = errors.New("No auth string provided")
)

// Request type definitions
var (
	reqTypeGET  string = "GET"
	reqTypePOST string = "POST"
)

// TranscribeOnline
// Makes a sync online transcription request to API.
// If success, a OnlineTranscriptionResponse object will be returned.
// If any error, an error will be returned.
func (tr *Transkit) TranscribeOnline(audio, pipeline string, tags ...string) (OnlineTranscriptionResponse, error) {

	// Basic error check
	if len(audio) <= 0 || len(pipeline) <= 0 {
		return OnlineTranscriptionResponse{}, errInvalidParameters
	}

	// Make new request object
	requestObj := OnlineTranscriptionRequest{Audio: audio, Pipeline: pipeline}
	if len(tags) > 0 {
		requestObj.Tags = append(requestObj.Tags, tags...)
	}

	// Get path
	pathReq, err := (*tr).joinURL(PATH_ONLINE_TRANSCRIPTION)
	if err != nil {
		return OnlineTranscriptionResponse{}, err
	}

	// Make HTTP request
	var responseObj OnlineTranscriptionResponse
	return responseObj, (*tr).makeHTTPRequest(reqTypePOST, pathReq, &responseObj, requestObj)
}

// TranscribeOffline
// Makes a async offline transcription request to API.
// If success, a OfflineTranscriptionResponse object will be returned that includes the job ID.
// If any error, an error will be returned.
func (tr *Transkit) TranscribeOffline(audioURL, pipeline string, tags ...string) (OfflineTranscriptionResponse, error) {

	// Basic error check
	if len(audioURL) <= 0 || len(pipeline) <= 0 {
		return OfflineTranscriptionResponse{}, errInvalidParameters
	}

	// Make new request object
	requestObj := OfflineTranscriptionRequest{URL: audioURL, Pipeline: pipeline}
	if len(tags) > 0 {
		requestObj.Tags = append(requestObj.Tags, tags...)
	}

	if (*tr).extraConfig != nil {
		requestObj.Config = (*tr).extraConfig
	}

	// Get path
	pathReq, err := (*tr).joinURL(PATH_OFFLINE_TRANSCRIPTION)
	if err != nil {
		return OfflineTranscriptionResponse{}, err
	}

	// Make HTTP request
	var responseObj OfflineTranscriptionResponse
	return responseObj, (*tr).makeHTTPRequest(reqTypePOST, pathReq, &responseObj, requestObj)
}

// TranscribeOfflineStatus
// Checks job status.
// If valid request, OfflineTranscriptionStatusResponse object will be returned.
// if not, an error will be thrown.
func (tr *Transkit) TranscribeOfflineStatus(id string) (OfflineTranscriptionStatusResponse, error) {
	var responseObj OfflineTranscriptionStatusResponse

	// Basic check
	if len((*tr).key) <= 0 {
		return OfflineTranscriptionStatusResponse{}, errNoAuthData
	}

	pathQuota, err := (*tr).joinURL(PATH_OFFLINE_TRANSCRIPTION_STATUS, id)
	if err != nil {
		return responseObj, err
	}

	// Return object
	return responseObj, (*tr).makeHTTPRequest(reqTypeGET, pathQuota, &responseObj)
}

// GetQuota
// Checks user quota.
// If valid, QuotaResponse object will be returned.
// if not, an error will be thrown.
func (tr *Transkit) GetQuota() (ConsumedQuotaResponse, error) {
	var responseObj ConsumedQuotaResponse

	// Basic check
	if len((*tr).key) <= 0 {
		return ConsumedQuotaResponse{}, errNoAuthData
	}

	pathQuota, err := (*tr).joinURL(PATH_GET_QUOTA)
	if err != nil {
		return responseObj, err
	}

	// Return object
	return responseObj, (*tr).makeHTTPRequest(reqTypeGET, pathQuota, &responseObj)
}

// SetConfig
// Set configuration to provide to transcription services.
func (tr *Transkit) SetConfig(config interface{}) {
	(*tr).extraConfig = config
}

// SetEndpoint
// Change API endpoint to provided apiEndpoint.
func (tr *Transkit) SetEndpoint(apiEndpoint string) {
	(*tr).apiURL = apiEndpoint
}

// makeHTTPRequest
// Makes a HTTP reqType request type to a given reqPath.
// Response object address must be provided as first reqresObject value.
// Request object must be provided as second reqresObject value, if any.
func (tr *Transkit) makeHTTPRequest(reqType, reqPath string, reqresObject ...interface{}) error {
	var responseObj interface{}
	var req *http.Request
	var err error

	// Define response object
	responseObj = reqresObject[0]

	if reqType == reqTypePOST {
		reqObjBytes, err := json.Marshal(reqresObject[1])
		if err != nil {
			return errJSONMarshalObject
		}
		req, err = http.NewRequest(reqType, reqPath, strings.NewReader(string(reqObjBytes)))
	} else {
		req, err = http.NewRequest(reqType, reqPath, nil)
	}

	// Create a new request using http
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+(*tr).key)

	// Make request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	// Check status code 401
	if resp.StatusCode == 401 {
		return errInvalidAPIKey
	}

	// Check status code not 200
	if resp.StatusCode != 200 {
		return errInvalidStatusCode
	}

	// Read body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	return json.Unmarshal(body, responseObj)
}

// joinURL
// Joins a API url with a list of paths provided on pathList variable.
// Result path will be ordered as provided arguments.
func (tr *Transkit) joinURL(pathList ...string) (string, error) {
	u, err := url.Parse((*tr).apiURL)
	if err != nil {
		return "", err
	}

	for _, pathStr := range pathList {
		u.Path = path.Join(u.Path, pathStr)
	}

	return u.String(), nil
}
