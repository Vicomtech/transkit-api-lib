//
// Transkit API libraries
// Coded by Aritz Olea <aolea@vicomtech.org>
// (c) 2022 Vicomtech - vicomtech.org
//

//
// >---- INFO ----<
// Configuration variables
//
// API_KEY: 			App usage key
// ONLINE_TEST_FILE: 	Local test file to send
// OFFLINE_TEST_FILE: 	Remote test file to send
// PIPELINE: 			Pipeline name to use
// OFFLINE_JOB_ID:		Job ID to check status
//

package transkit

import (
	"encoding/base64"
	"encoding/json"
	"os"
	"testing"
)

var transkit Transkit

// initTest
// initializes test objects
func initTest() {
	transkit.Init(os.Getenv("API_KEY"))
}

// TestQuota
// Makes a request to quota method and prints the response as log
func TestQuota(t *testing.T) {
	initTest()

	qResp, err := transkit.GetQuota()
	if err != nil {
		t.Fatal(err)
	}

	t.Log("User quota:", qResp.ConsumedQuota)
}

// TestTranscribeOnline
// Reads an audio file, encodes into base64 and
// makes a request to online transcription service and prints the response as log
func TestTranscribeOnline(t *testing.T) {
	initTest()

	// Read test file
	fileString, err := readFileBase64(os.Getenv("ONLINE_TEST_FILE"))
	if err != nil {
		t.Fatal(err)
	}

	// Transcribe file
	qResp, err := transkit.TranscribeOnline(fileString, os.Getenv("PIPELINE"))
	if err != nil {
		t.Fatal(err)
	}

	t.Log("Transcription:", qResp.Text)
}

// TestTranscribeOffline
// Gets an URL and makes a request to offline transcription service
// and prints the response ID as log
func TestTranscribeOffline(t *testing.T) {
	initTest()

	// Get test file
	URL := os.Getenv("OFFLINE_TEST_FILE")

	// Transcribe file
	qResp, err := transkit.TranscribeOffline(URL, os.Getenv("PIPELINE"))
	if err != nil {
		t.Fatal(err)
	}

	t.Log("Transcription request ID:", qResp.Id)
}

// TestTranscribeOfflineStatus
// Makes a status request for a job. Then, prints as log the result.
func TestTranscribeOfflineStatus(t *testing.T) {
	initTest()

	// Get test file
	jobID := os.Getenv("OFFLINE_JOB_ID")

	// Transcribe file
	qResp, err := transkit.TranscribeOfflineStatus(jobID)
	if err != nil {
		t.Fatal(err)
	}

	qRespJSON, err := json.Marshal(qResp)
	if err != nil {
		t.Fatal(err)
	}

	t.Log("Transcription status JSON:", string(qRespJSON))
}

// readFileBase64
// Reads an audio file and encodes into base64
// throws an error if any
func readFileBase64(filename string) (string, error) {
	var (
		readedFile string
		retErr     error
	)

	originalFile, retErr := os.ReadFile(filename)
	if retErr != nil {
		return readedFile, retErr
	}

	return base64.StdEncoding.EncodeToString(originalFile), retErr
}
