package pushbullet

import (
    "net/http"
    "bytes"
)

func Post(uri, access_token, body string) error {
    req, err := http.NewRequest(
        "POST",
        uri,
        bytes.NewBuffer([]byte(body)),
    )
    if err != nil {
        return err
    }

    req.Header.Set("Access-Token", access_token)
    req.Header.Set("Content-Type", "application/json")

    client := &http.Client{}
    resp, err := client.Do(req)

    if err != nil {
        return err
    }
    defer resp.Body.Close()

    return nil
}
