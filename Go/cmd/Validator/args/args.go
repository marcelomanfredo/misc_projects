package args

import (
	"fmt"

	"marcelomanfredo/misc/Go/cmd/Validator/internal/generate"
	"marcelomanfredo/misc/Go/internal/validate"
	"marcelomanfredo/misc/Go/logger"

	"github.com/spf13/cobra"
)

const (
	version        = "0.0.1"
	discard string = "\n  If -n is present, any value passed will be ignored"
)

// Flags
var (
	verbose bool // THIS WILL BE -C, SO -V CAN PRINT THE VERSION NUMBER
	create  bool
	modType int
	value   string
)

// Commands
var (
	rootCmd = &cobra.Command{
		Use:     "val",
		Version: version,
		Short:   "A validator for common calculations",
		Long: `Validator is used to validate common Brazilian calculations, such as CPF numbers or MOD10/MOD11 check digits.
        All generated values are random and follow no specific pattern.
        `,
	}

	cmdCPF = &cobra.Command{
		Use:   "cpf",
		Short: "Validate or create a new CPF string",
		RunE: func(cmd *cobra.Command, args []string) error {
			logger.Init(verbose)
			if create {
				if err := generate.GenerateCpf(); err != nil {
					return err
				}
				return nil
			}

			var n int = len(args)
			if n != 1 {
				return fmt.Errorf("expected 1 value, but got %d", n)
			}
			if err := validate.ValidateCpf(args[0]); err != nil {
				return err
			}
			return nil
		},
	}

	cmdMOD = &cobra.Command{
		Use:   "mod",
		Short: "Validate or create new MOD<TYPE> check number",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Running CPF command")
		},
	}
)

func init() {
	rootCmd.SilenceErrors = false
}

func addCommonFlags(cmd *cobra.Command) {
	cmd.Flags().BoolVarP(&verbose, "calculations", "c", false, "show details of the calculation")
	cmd.Flags().BoolVarP(&create, "new", "n", false, "generate new value instead of validating one - ignores any value passed")
}

func Execute() error {
	// Disable script generation > Not needed in this project
	rootCmd.CompletionOptions.DisableDefaultCmd = true

	// CPF flags
	addCommonFlags(cmdCPF)

	// MOD flags
	addCommonFlags(cmdMOD)
	cmdMOD.Flags().IntVarP(&modType, "type", "t", 0, "select the MOD type to use")
	cmdMOD.MarkFlagRequired("type")

	rootCmd.AddCommand(cmdCPF, cmdMOD)

	return rootCmd.Execute()
}
