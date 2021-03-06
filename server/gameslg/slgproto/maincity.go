package slgproto

import (
	"github.com/llr104/LiFrame/proto"
)

const (
	Building_Dwelling = iota
	Building_Minefield
	Building_Farmland
	Building_Lumberyard
	Building_Barrack
)

type QryBuildingQeq struct {
	BuildType    int8	 `json:"type"`
}

type QryBuildingAck struct {
	proto.BaseAck
	BuildType    int8	 `json:"type"`
	Buildings    string  `json:"buildings"`
}
