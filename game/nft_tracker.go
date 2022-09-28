package game

import "github.com/spike-engine/spike-web3-server/model"

type NftOwnerTracker interface {
	QueryNftList(ownerAddr string, contractAddr string) ([]model.NftOwner, error)
	AddNftOwner(no model.NftOwner) error
	UpdateNftOwner(ownerAddr string, contractAddr string, tokenId int64, updateTime int64) error
	QueryNftOwner(tokenId int64, contractAddr string) ([]model.NftOwner, error)
}
