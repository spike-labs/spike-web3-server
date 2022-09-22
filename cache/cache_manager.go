package cache

import (
	"github.com/go-redis/redis/v8"
	"github.com/spike-engine/spike-web3-server/constant"
	"github.com/spike-engine/spike-web3-server/util"
)

type ClearEvent struct {
	FromAddr string
	ToAddr   string
}

type Manager struct {
	RedisClient *redis.Client
}

func NewManager(client *redis.Client) *Manager {
	return &Manager{
		client,
	}
}

func (m *Manager) Update(event interface{}) {
	if e, ok := event.(ClearEvent); ok {
		log.Infof("clear nft cache , from : %s, to : %s", e.FromAddr, e.ToAddr)
		util.RmKeyByPrefix(e.FromAddr+constant.NFTTAG, m.RedisClient)
		util.RmKeyByPrefix(e.ToAddr+constant.NFTTAG, m.RedisClient)
	}
}

type Observer interface {
	Update(event interface{})
}

type Subject interface {
	AttachObserver(observer Observer)
	Notify(event interface{})
}
