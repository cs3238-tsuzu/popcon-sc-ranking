package main

import (
	"sync"
)

type RankingManager struct {
	mutex *sync.RWMutex
}

func NewRankingManager() *RankingManager {
	return new(RankingManager)
}

// New 成功時に接続先ポート番号を返す
func (rm *RankingManager) New(cid int64) (int, error) {

}
