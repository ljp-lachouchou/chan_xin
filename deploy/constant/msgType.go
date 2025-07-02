package constant

type MType int

const (
	TestType MType = iota
	ImageType
	FileType
)

type ChatType int

const (
	SingleChat ChatType = iota
	GroupChat
)
