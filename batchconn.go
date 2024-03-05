package kcp

const (
	batchSize = 16
)

type batchConn interface {
	WriteBatch(ms []Message, flags int) (int, error)
	ReadBatch(ms []Message, flags int) (int, error)
}
