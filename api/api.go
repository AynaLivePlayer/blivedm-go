package api

import (
	log "github.com/sirupsen/logrus"
)

type IApi interface {
	GetDanmuInfo(roomID int) (int, *DanmuInfo, error)
	GetRoomInfo(roomID int) (*RoomInfo, error)
}

type defaultClient struct {
	cookie string
}

func (d *defaultClient) GetDanmuInfo(roomID int) (int, *DanmuInfo, error) {
	var uid int = 0
	var err error
	if d.cookie != "" {
		uid, err = GetUid(d.cookie)
		if err != nil {
			uid = 0
			log.Error(err)
		}
	}
	result, err := GetDanmuInfo(roomID, d.cookie)
	return uid, result, err
}

func (d *defaultClient) GetRoomInfo(roomID int) (*RoomInfo, error) {
	return GetRoomInfo(roomID)
}

func NewDefaultClient(cookie string) IApi {
	return &defaultClient{cookie: cookie}
}
