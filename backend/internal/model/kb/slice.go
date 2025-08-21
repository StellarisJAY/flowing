package kb

type Slice struct {
	SliceId         string
	DocumentId      int64
	KnowledgeBaseId int64
	SliceContent    string
	Embedding       []float32
	Index           int64
	Meta            map[string]string
}

func (s *Slice) Content() string {
	return s.SliceContent
}

func (s *Slice) DocId() int64 {
	return s.DocumentId
}

func (s *Slice) Id() string {
	return s.SliceId
}

func (s *Slice) DenseVector() []float32 {
	return s.Embedding
}

func (s *Slice) Metadata() map[string]string {
	return s.Meta
}
