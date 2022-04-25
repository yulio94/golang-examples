package cli

import (
	"fmt"
	"github.com/spf13/cobra"
)

var stores = map[string]string{
	"01DC9ZAPGKEQJS4P4A48EG3P43": "Mercadona",
	"01DC9ZB23EW0J0ARAER09SJDKC": "Carrefour",
	"01DC9ZB89V1PQD977ZE6QXSQHH": "Alcampo",
}

func InitStoresCmd() (storesCmd *cobra.Command) {
	storesCmd = &cobra.Command{
		Use:   "stores",
		Short: "Print data about beer stores",
		Run:   runStoresFn(),
	}

	storesCmd.Flags().StringP(IdFlag, "i", "", "Id of the store")
	return
}

func runStoresFn() CobraFn {
	return func(cmd *cobra.Command, args []string) {
		id, _ := cmd.Flags().GetString(IdFlag)

		if id != "" {
			fmt.Println(stores[id])
		} else {
			fmt.Println(stores)
		}
	}
}
