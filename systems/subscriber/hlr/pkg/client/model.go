package clients

type SimFactoryData struct {
	Imsi           string
	Op             []byte
	Amf            []byte
	Key            []byte
	AlgoType       uint32
	UeDlAmbrBps    uint32
	UeUlAmbrBps    uint32
	Sqn            uint32
	CsgIdPrsent    bool
	CsgId          uint32
	DefaultApnName string
}
 