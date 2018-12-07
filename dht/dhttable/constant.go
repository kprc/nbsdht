package dhttable


const(
	COUNT_PER_BUCKET=8
	COUNT_PER_CANDBUCKET=16
	COUNT_PING_CHANNEL=2048
	COUNT_REQUEST_A=3
)


var Bitspos = [...]int{
	0,
	1,
	2,2,
	3,3,3,3,
	4,4,4,4,4,4,4,4,
}