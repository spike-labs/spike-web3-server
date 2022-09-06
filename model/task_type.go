package model

const (
	NftQuery = iota
	NativeTxRecordQuery
	Erc20TxRecordQuery
)

var taskTypeShortNames = map[TaskType]string{
	NftQuery:            "nftQuery",
	NativeTxRecordQuery: "nativeTxRecordQuery",
	Erc20TxRecordQuery:  "erc20TxRecordQuery",
}

func (t TaskType) String() string {
	n, ok := taskTypeShortNames[t]
	if !ok {
		return "unknown"
	}
	return n
}

type TaskType int
