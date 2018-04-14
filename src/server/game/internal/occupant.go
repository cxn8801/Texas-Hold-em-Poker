package internal

import (
	"server/model"
	"github.com/dolotech/leaf/gate"
	"server/algorithm"
)

type Occupant struct {
	*model.User
	gate.Agent
	room  *Room
	cards algorithm.Cards
	//Actions chan *ws.Message
	//timer *time.Timer // action timer

	Status uint8 // 1为离线状态
}

func NewOccupant(data *model.User, conn gate.Agent) *Occupant {
	o := &Occupant{
		User:  data,
		Agent: conn,
	}
	return o
}