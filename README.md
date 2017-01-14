# pushbullet

```go
package main

import (
    "github.com/mute1997/pushbullet-go"
)

func main(){
    request := pushbullet.Request{Push_type:"note",Title:"title",Body:"body"}
    request_link := pushbullet.Request{Push_type:"link",Title:"link",Uri:"http://google.com"}
    request_file := pushbullet.Request{Push_type:"file",Title:"file",File_name:"マグロ",File_type:"image/jpeg",File_url:"https://pbs.twimg.com/media/C2I8DjPUcAAr6ba.jpg"}

    pb := pushbullet.NewPushbullet("api_key")
    pb.Push_note(request)
    pb.Push_note(request_link)
    pb.Push_note(request_file)
}
```
