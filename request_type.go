package pushbullet

type Request struct {
    Push_type string
    Title string
    Body string
    Uri string
    File_name string
    File_type string
    File_url string
    Source_device_iden string
    Device_iden string
    Client_iden string
    Channel_tag string
    Email string
    Guid string
}

func (request Request) To_json() string{
    push_type := request.Push_type

    if push_type == "note" {
        //note
        return `{"body":"` + request.Body + 
        `","title":"` + request.Title + 
        `","type":"` + request.Push_type + `"}`
    }else if push_type == "file" {
        //file
        return `{","title":"` + request.Title + 
        `","file_name":"` + request.File_name + 
        `","file_type":"` + request.File_type + 
        `","file_url":"` + request.File_url + 
        `","type":"` + request.Push_type + `"}`
    }else{
        //link
        return `{","title":"` + request.Title + 
        `","url":"` + request.Uri + 
        `","type":"` + request.Push_type + `"}`
    }
}
