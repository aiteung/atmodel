package atmodel

type MediaType int

const (
	_              = iota
	MediaTypeImage = MediaType(iota)
	MediaTypeVideo
	MediaTypeAudio
	MediaTypeDocument
)

var (
	AppInfo = map[MediaType]string{
		MediaTypeImage:    "WhatsApp Image Keys",
		MediaTypeVideo:    "WhatsApp Video Keys",
		MediaTypeAudio:    "WhatsApp Audio Keys",
		MediaTypeDocument: "WhatsApp Document Keys",
	}
)

type IteungMessage struct {
	Phone_number string  `json:"phone_number"`
	Group_name   string  `json:"group_name"`
	Alias_name   string  `json:"alias_name"`
	Messages     string  `json:"messages"`
	Is_group     string  `json:"is_group"`
	Filename     string  `json:"filename"`
	Filedata     string  `json:"filedata"`
	Latitude     float64 `json:"latitude"`
	Longitude    float64 `json:"longitude"`
	Api_key      string  `json:"api_key"`
}

type IteungRespon struct {
	Message string `json:"message"`
}

type QRScan struct {
	QR      string `json:"qr"`
	Status  string `json:"status"`
	Message string `json:"message"`
}

type Chat struct {
	Phone_number string `json:"phone_number"`
	Messages     string `json:"messages"`
}

type Notif struct {
	User     string `json:"user"`
	Server   string `json:"server"`
	Messages string `json:"messages"`
}

type NotifPool struct {
	Pool         string `json:"pool"`
	NotifMessage Notif  `json:"message"`
}
