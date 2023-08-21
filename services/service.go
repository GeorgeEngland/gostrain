package services

import "github.com/GeorgeEngland/gostrain/common"

type requestHandler func(common.Message) error

type Service interface {
	SetRequestHandler(requestHandler)
}

type defaultNode struct {
	requestHandler
}

func (n *defaultNode) SetRequestHandler(r requestHandler) {
	n.requestHandler = r
}

func NewDefaultService(requestHandler requestHandler) (Service, error) {
	n := defaultNode{}
	n.SetRequestHandler(requestHandler)
	return &n, nil
}

// Node should on startup:
/*

* Connect to Sequencer
* Optionally Replay
* Start listening to live updates
	* Call request handler on each message
* Start listening to clock updates (heartbeats)
	* Update clock time with time from clock
* Maintain accounting information
	* On Startup, get latest CommandSent Index
	* Messages received index (should inc by 1, otherwise replay)
	* Commands sent index (don't send any commands until passed the index)

*/
