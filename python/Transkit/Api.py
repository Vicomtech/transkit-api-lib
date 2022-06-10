# Transkit API module
# (c) 2022 Vicomtech

import requests
import json

class Transkit:
    endpoint = "https://api.transkit.vicomtech.org/"
    transcribe_online_path = "online/transcribe"
    transcribe_offline_path = "offline/transcribe"
    transcribe_offline_status_path = "offline/status/"
    quota_path = "quota"
    auth = ""
    tags = []

    def __init__(self, authcode):
        self.auth = authcode

    # Get online transcription
    def getOnlineTranscription(self, audioBase64, pipeline):
        reqJson = {'pipeline':pipeline, "audio":audioBase64}
        if len(self.tags) > 0:
            reqJson['tags'] = self.tags

        r = requests.post(url = self.endpoint+self.transcribe_online_path, data = json.dumps(reqJson), headers = {'Authorization': 'Bearer '+self.auth })

        if (r.status_code == 401):
            raise Exception('Invalid API key or expired')
        elif (r.status_code != 200):
            raise Exception('Invalid status code: '+str(r.status_code))

        return r.json()

    # Get offline transcription
    def getOfflineTranscription(self, audioURL, pipeline, config = None):
        reqJson = {'pipeline':pipeline, "url":audioURL}
        
        if len(self.tags) > 0:
            reqJson['tags'] = self.tags
        if config != None:
            reqJson['config'] = config

        r = requests.post(url = self.endpoint+self.transcribe_offline_path, data = json.dumps(reqJson), headers = {'Authorization': 'Bearer '+self.auth })

        if (r.status_code == 401):
            raise Exception('Invalid API key or expired')
        elif (r.status_code != 200):
            raise Exception('Invalid status code: '+str(r.status_code))

        return r.json()

    # Get offline transcription status
    def getOfflineTranscriptionStatus(self, jobid):
        r = requests.get(url = self.endpoint+self.transcribe_offline_status_path+jobid, headers = {'Authorization': 'Bearer '+self.auth })

        if (r.status_code == 401):
            raise Exception('Invalid API key or expired')
        elif (r.status_code != 200):
            raise Exception('Invalid status code: '+str(r.status_code))

        return r.json()

    # Get quota
    def getQuota(self):
        r = requests.get(url = self.endpoint+self.quota_path, headers = {'Authorization': 'Bearer '+self.auth })

        if (r.status_code == 401):
            raise Exception('Invalid API key or expired')
        elif (r.status_code != 200):
            raise Exception('Invalid status code: '+str(r.status_code))

        return r.json()

    # Add new tag
    def addTag(self, tag):
        self.tags.append(tag)

    # Set all tags
    def setTags(self, tags):
        self.tags = tags