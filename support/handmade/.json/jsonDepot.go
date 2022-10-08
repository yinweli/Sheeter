package _json

type Depot struct {
	Reward  RewardReader
	readers []ReaderInterface
}

func (this *Depot) FromData(load DepotLoad, error DepotError) bool {
	this.build()
	result := true

	for _, itor := range this.readers {
		data := load(itor.DataName(), itor.DataExt())

		if data == nil || len(data) == 0 {
			continue
		}

		if err := itor.FromData(data); err != nil {
			result = false
			error(itor.DataName(), err)
		}
	}

	return result
}

func (this *Depot) MergeData(load DepotLoad, error DepotError) bool {
	this.build()
	result := true

	for _, itor := range this.readers {
		data := load(itor.DataName(), itor.DataExt())

		if data == nil || len(data) == 0 {
			continue
		}

		if err := itor.MergeData(data); err != nil {
			result = false
			error(itor.DataName(), err)
		}
	}

	return result
}

func (this *Depot) build() {
	if len(this.readers) == 0 {
		this.readers = append(
			this.readers,
			&this.Reward,
		)
	}
}

type DepotError func(name string, err error)
type DepotLoad func(name, ext string) []byte

type ReaderInterface interface {
	DataName() string
	DataExt() string
	DataFile() string
	FromData(data []byte) error
	MergeData(data []byte) error
}
