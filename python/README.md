# Transkit API libraries

In this repository, you can find example ready-to-use libraries. For more technical information, please visit https://api.transkit.vicomtech.org/doc/

### Functions in "Transkit" package

- getOnlineTranscription(audioBase64, pipeline). Returns a JSON
- getOfflineTranscription(audioURL, pipeline, config = None). Returns a JSON
- getOfflineTranscriptionStatus(jobid). Returns a JSON
- getQuota(). Returns a JSON
- addTag(tag).
- setTags(tags).

An exception will be thrown if any error detected, so we recommend to use try/except statements.

### Example

Install the package

> pip install Transkit

Import:
`from Transkit import Transkit`

Initialize with your api key:

`transkitapi = Transkit("apikey")`

Make an online transcription:
```
with open("file.mp3", "rb") as f:
    encodedFile = base64.b64encode(f.read()).decode()
    transcriptionResult = transkit.getOnlineTranscription(encodedFile, "myPipeline")
    print("Transcription result: ", transcriptionResult)
```

Make an offline transcription:
```
transcriptionResult = transkit.getOfflineTranscription("https://myvideo.url/path/here", "myPipeline")
```

Make an offline transcription status request:
```
transcriptionResult = transkit.getOfflineTranscriptionStatus("my-job-id")
```
