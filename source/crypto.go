package source

type Crypto struct {
	buffer [64]byte
	offset int
}

func NewCrypto() *Crypto { _ = "STUB: not implemented"; return nil }

func (s *Crypto) refillBuffer() { _ = "STUB: not implemented"; return }

func (s *Crypto) Uint64() uint64 { _ = "STUB: not implemented"; return 0 }
