package source

type JSF struct {
	a, b, c, d uint32
}

func NewJSF(seed uint64) *JSF { _ = "STUB: not implemented"; return nil }

func (jsf *JSF) Seed(seed uint64) { _ = "STUB: not implemented"; return }

func (jsf *JSF) Uint64() uint64 { _ = "STUB: not implemented"; return 0 }
