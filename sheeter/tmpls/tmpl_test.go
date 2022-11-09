package tmpls

import (
	"path/filepath"
	"strconv"
	"testing"

	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/yinweli/Sheeter/sheeter"
	"github.com/yinweli/Sheeter/sheeter/utils"
	"github.com/yinweli/Sheeter/testdata"
)

func TestTmpl(t *testing.T) {
	suite.Run(t, new(SuiteTmpl))
}

type SuiteTmpl struct {
	suite.Suite
	testdata.TestEnv
	name  string
	tmpl1 string
	tmpl2 string
}

func (this *SuiteTmpl) SetupSuite() {
	this.Change("test-tmpl")
	this.name = "tmpl.txt"
	this.tmpl1 = "tmpl1"
	this.tmpl2 = "tmpl2"
}

func (this *SuiteTmpl) TearDownSuite() {
	this.Restore()
}

func (this *SuiteTmpl) target() *Tmpl {
	target := &Tmpl{
		Name: this.name,
		Data: this.tmpl1,
	}
	return target
}

func (this *SuiteTmpl) TestInitialize() {
	cmd := SetFlags(&cobra.Command{})
	assert.Nil(this.T(), Initialize(cmd))
	testdata.CompareFile(this.T(), JsonReaderCs.path(), []byte(JsonReaderCs.Data))
	testdata.CompareFile(this.T(), JsonReaderGo.path(), []byte(JsonReaderGo.Data))

	cmd = SetFlags(&cobra.Command{})
	assert.Nil(this.T(), cmd.Flags().Set(flagClean, strconv.FormatBool(true)))
	assert.Nil(this.T(), Initialize(cmd))
	testdata.CompareFile(this.T(), JsonReaderCs.path(), []byte(JsonReaderCs.Data))
	testdata.CompareFile(this.T(), JsonReaderGo.path(), []byte(JsonReaderGo.Data))

	cmd = SetFlags(&cobra.Command{})
	assert.Nil(this.T(), utils.WriteFile(JsonReaderCs.path(), []byte(this.tmpl1)))
	assert.Nil(this.T(), utils.WriteFile(JsonReaderGo.path(), []byte(this.tmpl2)))
	assert.Nil(this.T(), Initialize(cmd))
	assert.Equal(this.T(), this.tmpl1, JsonReaderCs.Data)
	assert.Equal(this.T(), this.tmpl2, JsonReaderGo.Data)
	testdata.CompareFile(this.T(), JsonReaderCs.path(), []byte(JsonReaderCs.Data))
	testdata.CompareFile(this.T(), JsonReaderGo.path(), []byte(JsonReaderGo.Data))
}

func (this *SuiteTmpl) TestLoad() {
	target := this.target()
	assert.Nil(this.T(), target.load())
	assert.Equal(this.T(), this.tmpl1, target.Data)

	target = this.target()
	assert.Nil(this.T(), utils.WriteFile(target.path(), []byte(this.tmpl2)))
	assert.Nil(this.T(), target.load())
	assert.Equal(this.T(), this.tmpl2, target.Data)
}

func (this *SuiteTmpl) TestSave() {
	target := this.target()
	assert.Nil(this.T(), target.save())
	testdata.CompareFile(this.T(), target.path(), []byte(target.Data))
}

func (this *SuiteTmpl) TestPath() {
	target := this.target()
	assert.Equal(this.T(), filepath.Join(sheeter.TmplPath, target.Name), target.path())
}
