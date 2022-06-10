# Transkit API libraries

In this repository, you can find example ready-to-use libraries. For more technical information, please visit https://api.transkit.vicomtech.org/doc/


## Go

### Functions in package transkit

- TranscribeOnline(audio, pipeline string, tags ...string) (OnlineTranscriptionResponse, error)
- TranscribeOffline(audioURL, pipeline string, tags ...string) (OfflineTranscriptionResponse, error)
- GetQuota(auth string) (QuotaResponse, error)
- TranscribeOfflineStatus(id string) (OfflineTranscriptionStatusResponse, error)
- GetQuota() (ConsumedQuotaResponse, error)
- SetConfig(config interface{})
- SetEndpoint(apiEndpoint string)

### Example
Import module:
```
import("github.com/Vicomtech/transkit-api-lib/go/transkit")
```

Initialize object:
```
transkit := transkit.Transkit{}
transkit.Init("api key")
```

Send an online transcription (fileString is a base64-encoded binary file)
```
qResp, err := transkit.TranscribeOnline(fileString, "PipelineName")
if err != nil {
    panic(err)
}

fmt.Println("Transcription:", qResp.Text)
```

Get used quota:
```
if responseQuota, err := transkit.GetQuota(); err != nil {
    fmt.Println(responseQuota.ConsumedQuota)
}
```

Send an offline transcription:
```
qResp, err := transkit.TranscribeOffline("https://url.com/to/my/file", "PipelineName")
if err != nil {
    panic(err)
}

fmt.Println("Transcription job ID:", qResp.Id)
```

Get offline transcription status:
```
qResp, err := transkit.TranscribeOfflineStatus("job id")
if err != nil {
	panic(err)
}

fmt.Println("Transcription status data object:", qResp)
```


## Python 3
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
