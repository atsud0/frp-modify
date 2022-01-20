// Copyright 2016 fatedier, fatedier@gmail.com
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package msg

import "net"

const (
	TypeLogin                 = 'o'
	TypeLoginResp             = '1'
	TypeNewProxy              = 'p'
	TypeNewProxyResp          = '2'
	TypeCloseProxy            = 'c'
	TypeNewWorkConn           = 'w'
	TypeReqWorkConn           = 'r'
	TypeStartWorkConn         = 's'
	TypeNewVisitorConn        = 'v'
	TypeNewVisitorConnResp    = '3'
	TypePing                  = 'h'
	TypePong                  = '4'
	TypeUDPPacket             = 'u'
	TypeNatHoleVisitor        = 'i'
	TypeNatHoleClient         = 'n'
	TypeNatHoleResp           = 'm'
	TypeNatHoleClientDetectOK = 'd'
	TypeNatHoleSid            = '5'
)

var (
	msgTypeMap = map[byte]interface{}{
		TypeLogin:                 Login{},
		TypeLoginResp:             LoginResp{},
		TypeNewProxy:              NewProxy{},
		TypeNewProxyResp:          NewProxyResp{},
		TypeCloseProxy:            CloseProxy{},
		TypeNewWorkConn:           NewWorkConn{},
		TypeReqWorkConn:           ReqWorkConn{},
		TypeStartWorkConn:         StartWorkConn{},
		TypeNewVisitorConn:        NewVisitorConn{},
		TypeNewVisitorConnResp:    NewVisitorConnResp{},
		TypePing:                  Ping{},
		TypePong:                  Pong{},
		TypeUDPPacket:             UDPPacket{},
		TypeNatHoleVisitor:        NatHoleVisitor{},
		TypeNatHoleClient:         NatHoleClient{},
		TypeNatHoleResp:           NatHoleResp{},
		TypeNatHoleClientDetectOK: NatHoleClientDetectOK{},
		TypeNatHoleSid:            NatHoleSid{},
	}
)

// When frpc start, client send this message to login to server.
type Login struct {
	Version      string            `json:"21hDceCd"`
	Hostname     string            `json:"HDFv4o7N"`
	Os           string            `json:"lvYjOhf7"`
	Arch         string            `json:"Tixb4ctH"`
	User         string            `json:"DZwSVL8b"`
	PrivilegeKey string            `json:"DUNzvi35"`
	Timestamp    int64             `json:"TRx8zOBn"`
	RunID        string            `json:"hGUH2mb7"`
	Metas        map[string]string `json:"Hg3virfp"`

	// Some global configures.
	PoolCount int `json:"3z8WmsHt"`
}

type LoginResp struct {
	Version       string `json:"nP6Ggpyl"`
	RunID         string `json:"Fj9dVzlw"`
	ServerUDPPort int    `json:"jaECf64x"`
	Error         string `json:"pcoMqTgE"`
}

// When frpc login success, send this message to frps for running a new proxy.
type NewProxy struct {
	ProxyName      string            `json:"dCycL2NA"`
	ProxyType      string            `json:"HjtgCNkX"`
	UseEncryption  bool              `json:"E401kyeh"`
	UseCompression bool              `json:"0l8j3gq4"`
	Group          string            `json:"G1IypxSU"`
	GroupKey       string            `json:"ZF205SpL"`
	Metas          map[string]string `json:"GdqlhXA0"`

	// tcp and udp only
	RemotePort int `json:"ifue9yDY"`

	// http and https only
	CustomDomains     []string          `json:"CHpI8ozs"`
	SubDomain         string            `json:"JgmATi7P"`
	Locations         []string          `json:"zklJHhXv"`
	HTTPUser          string            `json:"tFXxvnuE"`
	HTTPPwd           string            `json:"dFDEBz1f"`
	HostHeaderRewrite string            `json:"n1SxY9yN"`
	Headers           map[string]string `json:"aZEkOLq9"`

	// stcp
	Sk string `json:"68hVtL3S"`

	// tcpmux
	Multiplexer string `json:"Y6Mj57Ai"`
}

type NewProxyResp struct {
	ProxyName  string `json:"Idw2MztU"`
	RemoteAddr string `json:"WK92xi0f"`
	Error      string `json:"S0EBnhAu"`
}

type CloseProxy struct {
	ProxyName string `json:"Z3dE8isF"`
}

type NewWorkConn struct {
	RunID        string `json:"0FAzfVUi"`
	PrivilegeKey string `json:"b8T5jWxY"`
	Timestamp    int64  `json:"Ucg37V1O"`
}

type ReqWorkConn struct {
}

type StartWorkConn struct {
	ProxyName string `json:"8mKAF0UV"`
	SrcAddr   string `json:"LTAWORtf"`
	DstAddr   string `json:"HjIery3J"`
	SrcPort   uint16 `json:"xfAjs0Ve"`
	DstPort   uint16 `json:"xvTGAEau"`
	Error     string `json:"x6dMEqI9"`
}

type NewVisitorConn struct {
	ProxyName      string `json:"TrYIaoQU"`
	SignKey        string `json:"ESUruxIW"`
	Timestamp      int64  `json:"XBlmj3Z0"`
	UseEncryption  bool   `json:"aX5UErvb"`
	UseCompression bool   `json:"MGmlf6nN"`
}

type NewVisitorConnResp struct {
	ProxyName string `json:"lUIRkJKX"`
	Error     string `json:"gzAwSrIe"`
}

type Ping struct {
	PrivilegeKey string `json:"c4LRJIk8"`
	Timestamp    int64  `json:"JjXqw3Di"`
}

type Pong struct {
	Error string `json:"fp8CsJiF"`
}

type UDPPacket struct {
	Content    string       `json:"dwWSHxvq"`
	LocalAddr  *net.UDPAddr `json:"CDw1PcyT"`
	RemoteAddr *net.UDPAddr `json:"1HDI6uxw"`
}

type NatHoleVisitor struct {
	ProxyName string `json:"cp7RywO6"`
	SignKey   string `json:"gVtuo43P"`
	Timestamp int64  `json:"ZlF8d6fO"`
}

type NatHoleClient struct {
	ProxyName string `json:"TcqKQO3s"`
	Sid       string `json:"i6pjgPnk"`
}

type NatHoleResp struct {
	Sid         string `json:"BbdgX2nE"`
	VisitorAddr string `json:"9MgdYisy"`
	ClientAddr  string `json:"qf42xT3N"`
	Error       string `json:"QPd7kmBs"`
}

type NatHoleClientDetectOK struct {
}

type NatHoleSid struct {
	Sid string `json:"eJb9736r"`
}
