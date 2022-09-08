package cache

import (
	"github.com/go-redis/redis/v8"
	"spike-frame/constant"
	"spike-frame/util"
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

func (m *Manager) Update(event ClearEvent) {
	log.Infof("clear nft cache , from : %s, to : %s", event.FromAddr, event.ToAddr)
	util.RmKeyByPrefix(event.FromAddr+constant.NFTTAG, m.RedisClient)
	util.RmKeyByPrefix(event.ToAddr+constant.NFTTAG, m.RedisClient)
}

type Observer interface {
	Update(event ClearEvent)
}

type Subject interface {
	AttachObserver(observer Observer)
	Notify(event ClearEvent)
}
