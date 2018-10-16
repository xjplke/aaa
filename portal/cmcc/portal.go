package cmcc

import (
	"log"
	"net"

	"github.com/xjplke/aaa"
)

const (
	max_package_size = 4096
)

const (
	_             = iota
	REQ_CHALLENGE = iota
	ACK_CHALLENGE = iota
	REQ_AUTH      = iota
	ACK_AUTH      = iota
	REQ_LOGOUT    = iota
	ACK_LOGOUT    = iota
	AFF_ACK_AUTH  = iota
	NTF_LOGOUT    = iota
	REQ_INFO      = iota
	ACK_INFO      = iota
)

const (
	ATTR_USERNAME  = iota
	ATTR_PASSWORD  = iota
	ATTR_CHALLENGE = iota
	ATTR_CHAPPWD   = iota
	ATTR_USERMAC   = 0xE0
)

//Attribute Attribute
type Attribute interface {
	Type() byte
	Length() byte
	Byte() []byte
}

//CMCCOpts CMCCOpts
type CMCCOpts struct {
	Address string
	Port    string
}

//Message Message
type Message interface {
	Bytes() []byte
	Type() byte
	ReqId() uint16
	SerialId() uint16
	UserIp() net.IP
	CheckFor(Message, string) error
	AttributeLen() int
	Attribute(int) Attribute
}

type Version interface {
	Unmarshall([]byte) Message
	IsResponse(Message) bool
	NewChallenge(net.IP, string) Message
	NewAuth(net.IP, string, []byte, []byte, uint16, []byte) Message
	NewAffAckAuth(net.IP, string, uint16, uint16) Message
	NewLogout(net.IP, string) Message
	NewReqInfo(net.IP, string) Message
}

//CMCCPortal CMCCPortal
type CMCCPortal struct {
	authenticator aaa.Authenticator
	accounter     aaa.Accounter
	authorizer    aaa.Authorizer
	opts          *CMCCOpts
	ver           Version
}

//NewCMCCPortal NewCMCCPortal
func NewCMCCPortal(opts *CMCCOpts, authenticator aaa.Authenticator,
	accounter aaa.Accounter, authorizer aaa.Authorizer, ver Version) *CMCCPortal {
	return &CMCCPortal{
		authenticator: authenticator,
		accounter:     accounter,
		authorizer:    authorizer,
		opts:          opts,
		ver:           ver,
	}
}

//Start Start
func (p *CMCCPortal) Start() error {
	addr, err := net.ResolveUDPAddr("udp", p.opts.Address+":"+p.opts.Port)
	if err != nil {
		log.Fatal(err)
	}
	listener, err2 := net.ListenUDP("udp", addr)
	if err2 != nil {
		log.Fatal(err2)
	}

	log.Println("listen on addr :", addr)
	data := make([]byte, max_package_size)
	for {
		log.Println("before read")
		n, remoteAddr, err := listener.ReadFrom(data)
		if err != nil {
			log.Printf("error during read: %s", err)
		}
		log.Println("after read")

		log.Printf("<%s> %s", remoteAddr, data[:n])
		message := p.ver.Unmarshall(data)
		t := message.Type()
		switch t {
		case REQ_CHALLENGE: //
		case REQ_AUTH:

		case REQ_LOGOUT:
		case AFF_ACK_AUTH:
		}
	}
}
