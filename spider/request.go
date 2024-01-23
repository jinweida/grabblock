package spider

type Request struct {
	Url       string
	Height    int64
	Hash      string
	BlockHash string
	FailCount int64
}
