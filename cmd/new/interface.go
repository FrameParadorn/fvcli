package new

import "github.com/spf13/cobra"

type Interface interface {
	CreateCommand() *cobra.Command
}
