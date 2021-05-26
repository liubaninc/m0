package keeper

import (
	"container/list"
	"time"
)

// 临时锁定对象
type LockedItem struct {
	timestamp int64 // 过期时间
	holder    *list.Element
}

// 是否处于临时锁定状态
func (k Keeper) isLocked(key string) bool {
	k.mutexMem.Lock()
	defer k.mutexMem.Unlock()
	_, exist := k.lockKeys[key]
	return exist
}

// 试图临时锁定, 返回是否锁定成功
func (k Keeper) tryLockKey(key string, expire int64) bool {
	k.mutexMem.Lock()
	defer k.mutexMem.Unlock()
	timestamp := time.Now().Unix() + expire
	if _, exist := k.lockKeys[key]; !exist {
		elem := k.lockKeyList.Back()
		for ; elem != nil; elem = elem.Prev() {
			key := elem.Value.(string)
			if item := k.lockKeys[key]; item != nil {
				if item.timestamp <= timestamp {
					break
				}
			}
		}
		if elem != nil {
			holder := k.lockKeyList.InsertAfter(key, elem)
			k.lockKeys[key] = &LockedItem{timestamp: timestamp, holder: holder}
		} else {
			holder := k.lockKeyList.PushBack(key)
			k.lockKeys[key] = &LockedItem{timestamp: timestamp, holder: holder}
		}
		return true
	}
	return false
}

// 解锁对象
func (k Keeper) unlockKey(key string) {
	k.mutexMem.Lock()
	defer k.mutexMem.Unlock()
	lockItem := k.lockKeys[key]
	if lockItem != nil {
		k.lockKeyList.Remove(lockItem.holder)
		delete(k.lockKeys, key)
	}
}

// 清理过期的锁定对象
func (k Keeper) clearExpiredLocks() {
	k.mutexMem.Lock()
	defer k.mutexMem.Unlock()
	now := time.Now().Unix()
	for {
		topItem := k.lockKeyList.Front()
		if topItem == nil {
			break
		}
		key := topItem.Value.(string)
		lockItem, exist := k.lockKeys[key]
		if !exist {
			k.lockKeyList.Remove(topItem)
		} else if lockItem.timestamp <= now {
			k.lockKeyList.Remove(topItem)
			delete(k.lockKeys, key)
		} else {
			break
		}
	}
}
