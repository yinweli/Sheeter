package _json

type Depot struct {
	Reward  RewardReader
	loader  Loader
	readers []Reader
}

func NewDepot(loader Loader) *Depot {
	depot := &Depot{}
	depot.loader = loader
	depot.readers = append(
		depot.readers,
		&depot.Reward,
	)
	return depot
}

func (this *Depot) FromData() bool {
	if this.loader == nil {
		return false
	}

	result := true

	for _, itor := range this.readers {
		data := this.loader.Load(itor.DataName(), itor.DataExt(), itor.DataFile())

		if data == nil || len(data) == 0 {
			continue
		}

		if err := itor.FromData(data); err != nil {
			result = false
			this.loader.Error(itor.DataName(), err)
		}
	}

	return result
}

func (this *Depot) MergeData() bool {
	if this.loader == nil {
		return false
	}

	result := true

	for _, itor := range this.readers {
		data := this.loader.Load(itor.DataName(), itor.DataExt(), itor.DataFile())

		if data == nil || len(data) == 0 {
			continue
		}

		if err := itor.MergeData(data); err != nil {
			result = false
			this.loader.Error(itor.DataName(), err)
		}
	}

	return result
}

type Loader interface {
	Error(name string, err error)
	Load(name, ext, fullname string) []byte
}

type Reader interface {
	DataName() string
	DataExt() string
	DataFile() string
	FromData(data []byte) error
	MergeData(data []byte) error
}
