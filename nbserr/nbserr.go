package nbserr

const (

	REGERR = 101
	EXE_NOHANDLE = 201
	SERVER_RESOLVE_UDP_ERR=301
	LISTEN_UDP_ERR=302
	LISTEN_UDP_RESTART=303
	UDP_MESSAGE_LENGTH_ERR=304
	DHT_PUT_ERR=401

	UDP_BAD_CONN=440
	UDP_CONN_NOTREADY = 441
	UDP_CONN_NODATA = 442
	UDP_BUFFOVERFLOW = 443
	UDP_CONN_LISTEN = 444
	UDP_DIAL_ERR = 445

	UDP_ADDR_PARSE=590
	UDP_ADDR_TOSTRING_ERR = 591
	UDP_SND_DEFAULT_ERR    = 592

	UDP_SND_MTU_ERR = 593
	UDP_SND_READER_IO_ERR = 594
	UDP_SND_WRITER_IO_ERR = 595
	UDP_SND_TIMEOUT_ERR=596
	UDP_SND_RUNNING  = 597
	UDP_SND_CLOSED = 598


	UDP_RCV_DEFAULT_ERR = 601
	UDP_RCV_WRITER_IO_ERR = 602
	UDP_RCV_RUNNING=603

	UDP_TRANSINFO_ERR = 1024
	UDP_CONN_EXISTS = 1025


	PEER_RUNNING = 1100
	PEER_DEAD = 1101



	ERROR_DEFAULT=99999
)


type NbsErr struct {
	Errmsg string
	ErrId uint32
}


func (ne NbsErr)Error() string{
	return ne.Errmsg
}