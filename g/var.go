package g

import (
	"fmt"
	"sync"
	"errors"
	"container/list"
)

var (
	MaxCount = 10

	InputHosts []string

	HostIpMap = &SafeHostIpMap{
		M: make(map[string]string)}

	LastRTT = &SafeRttMap{
		M: make(map[string]int64)}

	HistoryRttMap = &SafeHistoryRttMap{
		M: make(map[string]*SafeLinkedList )}
)

type SafeRttMap struct {
	sync.RWMutex
	M map[string]int64
}

func (this *SafeRttMap) Set(k string, v int64) {
	this.Lock()
	defer this.Unlock()
	this.M[ k ] = v
}

func (this *SafeRttMap) Get(k string) (int64, bool) {
	this.RLock()
	defer this.RUnlock()
	t, ok := this.M[ k ]
	return t, ok
}

func (this *SafeRttMap) GetAll() map[string]int64 {
	this.RLock()
	defer this.RUnlock()
	return this.M
}

type SafeHostIpMap struct {
	sync.RWMutex
	M map[string]string
}

func (this *SafeHostIpMap) Set(k, v string) {
	this.Lock()
	defer this.Unlock()
	this.M[ k ] = v
}

func (this *SafeHostIpMap) Get(k string) (string, bool) {
	this.RLock()
	defer this.RUnlock()
	t, ok := this.M[ k ]
	return t, ok
}

func (this *SafeHostIpMap) GetAll() map[string]string {
	this.RLock()
	defer this.RUnlock()
	return this.M
}

type SafeHistoryRttMap struct {
	sync.RWMutex
	M map[string]*SafeLinkedList
}

func (this *SafeHistoryRttMap) Set(k string, v *SafeLinkedList) {
	this.Lock()
	defer this.Unlock()
	this.M[ k ] = v
}

func (this *SafeHistoryRttMap) Get(k string) (*SafeLinkedList, bool) {
	this.RLock()
	defer this.RUnlock()
	v, ok := this.M[ k ]
	return v, ok
}

func (this *SafeHistoryRttMap) PushFrontAndMaintain(k string, v int64){
	if linkedList, exist := this.Get(k); exist {

		linkedList.PushFrontAndMaintain(v)
	}else{

		nl := list.New()
		nl.PushFront(v)
		safeList := &SafeLinkedList{L: nl}
		this.Set(k, safeList)
	}
}

type PingStats struct{
	Max int64 `json:"max"`
	Min int64 `json:"min"`
	Avg int64 `json:""avg"`
	Ploss string `json:"ploss"`
}

type SafeLinkedList struct {
	sync.RWMutex
	L *list.List
}

func (this *SafeLinkedList) PushFrontAndMaintain(v int64){

	this.Lock()
	defer this.Unlock()

	sz := this.L.Len()

	this.L.PushFront(v)

	sz++
	if sz <= MaxCount {
		return
	}

	del := sz - MaxCount
	for i := 0; i < del; i++ {

		this.L.Remove(this.L.Back())
	}
}

func (this *SafeLinkedList) GetSummary()(PingStats, error){

	this.RLock()
	defer this.RUnlock()

	e := this.L.Front()
	if e == nil {
		return PingStats{}, errors.New("Not data in list")
	}

	var max, min, avg, sum, success, loss, total int64

	total = 1

	v := e.Value.(int64)
	if v >= 0 {
		max = v
		min = v
		sum = v
		success = 1
	}else{
		loss = 1
	}

	e = e.Next()

	for ; e != nil; e = e.Next() {
		v := e.Value.(int64)

		total ++

		if v < 0 {
			loss ++
			continue
		}

		success ++ 
		sum += v

		if v > max {
			max = v
		}

		if v < min {
			min = v
		}
	}

	if success > 0 {
		avg = sum / success
	}

	if loss == total {
		max = -1
		min = -1
		avg = -1
	}

	ploss := fmt.Sprintf("%.2f%%", float64(loss * 100) / float64(total))

	stats := PingStats{
		Max: max,
		Min: min,
		Avg: avg,
		Ploss: ploss }

	return stats, nil
}
