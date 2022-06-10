package build

import (
	"fmt"
	"os/exec"

	"Sheeter/internal/command/build/core"

	"github.com/spf13/cobra"
)

// NewCommand 建立命令物件
func NewCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "build [config]",
		Short: "build sheet",
		Long:  "build all the sheet in the configuration",
		Args:  cobra.ExactArgs(1),
		Run:   execute,
	}

	return cmd
}

// execute 執行命令
func execute(cmd *cobra.Command, args []string) {
	err := exec.Command("go").Run()

	if err != nil {
		cmd.Println(fmt.Errorf("not install go, download from https://go.dev/dl/"))
		return
	} // if

	err = exec.Command("protoc").Run()

	if err != nil {
		cmd.Println(fmt.Errorf("not install protoc, download from https://github.com/protocolbuffers/protobuf/releases"))
		return
	} // if

	config, err := core.ReadConfig(args[0])

	if err != nil {
		cmd.Println(err)
		return
	} // if

	count := len(config.Elements)
	signaler := make(chan error, 1)

	for _, itor := range config.Elements {
		go taskRoutine(config.Global, itor, signaler)
	} // for

	for i := 0; i < count; i++ {
		err := <-signaler

		if err != nil {
			cmd.Println(err)
		} // if
	} // for
}

// taskRoutine 多執行緒包裝函式
func taskRoutine(global core.Global, element core.Element, signaler chan<- error) {
	signaler <- core.NewTask(&global, &element).Execute()
}
