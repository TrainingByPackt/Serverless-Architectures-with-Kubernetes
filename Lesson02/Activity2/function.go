package p

import (
    "bytes"
    "net/http"
)

func Reminder(http.ResponseWriter, *http.Request) {
    url := "https://hooks.slack.com/services/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"
    
    var jsonStr = []byte(`{"text": "Time for a stand-up meeting!"}`)
    req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
    
    client := &http.Client{}
    _, err = client.Do(req)
    if err != nil {
        panic(err)
    }
}
