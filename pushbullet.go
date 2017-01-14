package pushbullet

type Pushbullet struct {
    Access_token string
}

func NewPushbullet(access_token string) *Pushbullet {
    return &Pushbullet{Access_token: access_token}
}

func (pushbullet *Pushbullet) Push_note(request Request) error{
    uri := "https://api.pushbullet.com/v2/pushes"

    err := Post(uri, pushbullet.Access_token, request.To_json())
    if err != nil {
        return err
    }
    return nil
}

