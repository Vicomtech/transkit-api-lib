## Transkit API
## Testing file

##
## >---- INFO ----<
## Configuration variables
##
## API_KEY: 			App usage key
## ONLINE_TEST_FILE: 	Local test file to send
## OFFLINE_TEST_FILE: 	Remote test file to send
## PIPELINE: 			Pipeline name to use
## OFFLINE_JOB_ID:		Job ID to check status
##

import os
import base64

from Api import Transkit


# Transkit online testing function
def test_transkit_online():
    pipeline = os.getenv('PIPELINE')
    onlineTestFile = os.getenv('ONLINE_TEST_FILE')
    apiKey = os.getenv('API_KEY')
    transkit = Transkit(apiKey)

    with open(onlineTestFile, "rb") as f:
        encodedFile = base64.b64encode(f.read()).decode()
        transcriptionResult = transkit.getOnlineTranscription(encodedFile, pipeline)
        print("Transcription result: ", transcriptionResult)

# Transkit offline testing function
def test_transkit_offline():
    pipeline = os.getenv('PIPELINE')
    offlineVideoURL = os.getenv('OFFLINE_TEST_FILE')
    apiKey = os.getenv('API_KEY')
    transkit = Transkit(apiKey)

    transcriptionResult = transkit.getOfflineTranscription(offlineVideoURL, pipeline)
    print("Transcription result: ", transcriptionResult)
    return transcriptionResult['id']

# Transkit offline status testing function
def test_transkit_offline_status(id):
    apiKey = os.getenv('API_KEY')
    transkit = Transkit(apiKey)

    transcriptionResult = transkit.getOfflineTranscriptionStatus(id)
    print("Transcription status result: ", transcriptionResult)

if __name__ == "__main__":
    test_transkit_online()
    id = test_transkit_offline()
    test_transkit_offline_status(id)
    print("Tests OK")
