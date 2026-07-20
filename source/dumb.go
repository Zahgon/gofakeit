package source

type Dumb struct {
	state uint64
}

func NewDumb(seed uint64) *Dumb { _ = "STUB: not implemented"; return nil }

func (d *Dumb) Seed(seed uint64) { _ = "STUB: not implemented"; return }

func (d *Dumb) Uint64() uint64 { _ = "STUB: not implemented"; return 0 }
