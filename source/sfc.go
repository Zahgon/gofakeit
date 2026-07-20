package source

type SFC struct {
	a, b, c, counter uint64
}

func NewSFC(seed uint64) *SFC { _ = "STUB: not implemented"; return nil }

func (s *SFC) Seed(seed uint64) { _ = "STUB: not implemented"; return }

func (s *SFC) Uint64() uint64 { _ = "STUB: not implemented"; return 0 }
