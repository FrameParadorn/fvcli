package new

import "github.com/spf13/cobra"

type New struct{}

func NewNew() Interface {
	return &New{}
}

func (*New) CreateCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "new",
		Short: "New project",
		Run: func(cmd *cobra.Command, args []string) {
		},
	}
}
