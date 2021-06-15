package model

type Peer struct {
	ID      uint   `json:"-" gorm:"primarykey" `
	Name    string `json:"peerAlais"`
	IP      string `json:"peerIP"`
	NodeID  string `json:"peerID"`
	Version string `json:"peerVersion"`
	Type    int    `json:"peerType"`
	Status  int    `json:"peerStatus"`
	Time    string `json:"peerStartIme"`
}
