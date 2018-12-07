package nbserr

const (
	REGERR = 101
	EXE_NOHANDLE = 201
	SERVER_RESOLVE_UDP_ERR=301
	LISTEN_UDP_ERR=302
	LISTEN_UDP_RESTART=303
	UDP_MESSAGE_LENGTH_ERR=304
	DHT_PUT_ERR=401

	UDP_ADDR_PARSE=590

	ERROR_DEFAULT=99999
)


type NbsErr struct {
	Errmsg string
	ErrId uint32
}


func (ne NbsErr)Error() string{
	return ne.Errmsg
}