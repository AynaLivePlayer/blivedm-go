package api

type IApi interface {
	GetUid() (int, error)
	GetDanmuInfo(roomID int) (*DanmuInfo, error)
	GetRoomInfo(roomID int) (*RoomInfo, error)
}

type defaultClient struct {
	cookie string
}

func (d *defaultClient) GetUid() (int, error) {
	return GetUid(d.cookie)
}

func (d *defaultClient) GetDanmuInfo(roomID int) (*DanmuInfo, error) {
	return GetDanmuInfo(roomID, d.cookie)
}

func (d *defaultClient) GetRoomInfo(roomID int) (*RoomInfo, error) {
	return GetRoomInfo(roomID)
}

func NewDefaultClient(cookie string) IApi {
	return &defaultClient{cookie: cookie}
}
