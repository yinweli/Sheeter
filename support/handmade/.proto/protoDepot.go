package _proto

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
		filename := itor.FileName()
		data := this.loader.Load(filename)

		if data == nil || len(data) == 0 {
			continue
		}

		if err := itor.FromData(data); err != nil {
			result = false
			this.loader.Error(filename.File(), err)
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
		filename := itor.FileName()
		data := this.loader.Load(filename)

		if data == nil || len(data) == 0 {
			continue
		}

		if err := itor.MergeData(data); err != nil {
			result = false
			this.loader.Error(filename.File(), err)
		}
	}

	return result
}

func (this *Depot) Clear() {
	for _, itor := range this.readers {
		itor.Clear()
	}
}

type FileName struct {
	name string
	ext  string
}

func NewFileName(name, ext string) FileName {
	return FileName{
		name: name,
		ext:  ext,
	}
}

func (this FileName) Name() string {
	return this.name
}

func (this FileName) Ext() string {
	return this.ext
}

func (this FileName) File() string {
	return this.name + this.ext
}

type Loader interface {
	Error(name string, err error)
	Load(filename FileName) []byte
}

type Reader interface {
	FileName() FileName
	FromData(data []byte) error
	MergeData(data []byte) error
	Clear()
}
