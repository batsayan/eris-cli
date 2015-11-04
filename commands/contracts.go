package commands

import (
	"os"

	"github.com/eris-ltd/eris-cli/contracts"

	. "github.com/eris-ltd/eris-cli/Godeps/_workspace/src/github.com/eris-ltd/common/go/common"
	"github.com/eris-ltd/eris-cli/Godeps/_workspace/src/github.com/spf13/cobra"
)

// Primary Contracts Sub-Command
var Contracts = &cobra.Command{
	Use:   "contracts",
	Short: "Deploy, Test, and Manage Your Smart Contracts.",
	Long: `The contracts subcommand is used to test and deploy
smart contracts for use by your application.`,
	Run: func(cmd *cobra.Command, args []string) { cmd.Help() },
}

// build the contracts subcommand
func buildContractsCommand() {
	// TODO: finish when the PR which is blocking
	//   eris files put --dir is integrated into
	//   ipfs
	// Contracts.AddCommand(contractsImport)
	// Contracts.AddCommand(contractsExport)
	Contracts.AddCommand(contractsTest)
	Contracts.AddCommand(contractsDeploy)
	addContractsFlags()
}

var contractsImport = &cobra.Command{
	Use:   "import [hash] [packageName]",
	Short: "Pull a package of smart contracts from IPFS.",
	Long: `Pull a package of smart contracts from IPFS
via its hash and save it locally.`,
	Run: ContractsImport,
}

var contractsExport = &cobra.Command{
	Use:   "export",
	Short: "Post a package of smart contracts to IPFS.",
	Long:  `Post a package of smart contracts to IPFS.`,
	Run:   ContractsExport,
}

var contractsTest = &cobra.Command{
	Use:   "test",
	Short: "Test a package of smart contracts.",
	Long: `Test a package of smart contracts.

Tests can be structured using three different
test types.

* epm -- epm apps can be tested against tendermint
style blockchains.
* embark -- embark apps can be tested against
ethereum style blockchains.
* truffle -- HELP WANTED!
* solUnit -- pure solidity smart contract packages
may be tested via solUnit test framework.
* manual -- a simple gulp task can be given to the
test environment.`,
	Run: ContractsTest,
}

var contractsDeploy = &cobra.Command{
	Use:   "deploy",
	Short: "Deploy a package of smart contracts to a chain.",
	Long: `Deploy a package of smart contracts to a chain.

Deployments can be structured using three different
deploy types.

* epm -- epm apps can be deployed to tendermint style
blockchains simply.
* embark -- embark apps can be deployed to an
ethereum style blockchain simply.
* truffle -- HELP WANTED!
* pyepm -- IF THIS IS STILL A THING, HELP WANTED!
* manual -- a simple gulp task can be given to the
deployer.
`,
	Run: ContractsDeploy,
}

//----------------------------------------------------

func addContractsFlags() {
	contractsTest.Flags().StringVarP(&do.ChainName, "chain", "c", "", "chain to be used for testing")
	contractsTest.Flags().StringSliceVarP(&do.ServicesSlice, "services", "s", []string{}, "comma separated list of services to start")
	contractsTest.Flags().StringVarP(&do.Type, "type", "t", "mint", "app type paradigm to be used for testing (overrides package.json)")
	contractsTest.Flags().StringVarP(&do.Task, "task", "k", "", "gulp task to be ran (overrides package.json; forces --type manual)")
	contractsTest.Flags().StringVarP(&do.Path, "dir", "i", "", "root directory of app (will use $pwd by default)")
	contractsTest.Flags().StringVarP(&do.NewName, "dest", "e", "", "working directory to be used for testing")
	contractsTest.Flags().BoolVarP(&do.Rm, "rm", "r", true, "remove containers after stopping")

	contractsDeploy.Flags().StringVarP(&do.ChainName, "chain", "c", "", "chain to be used for deployment")
	contractsDeploy.Flags().StringSliceVarP(&do.ServicesSlice, "services", "s", []string{}, "comma separated list of services to start")
	contractsDeploy.Flags().StringVarP(&do.Type, "type", "t", "mint", "app type paradigm to be used for deployment (overrides package.json)")
	contractsDeploy.Flags().StringVarP(&do.Task, "task", "k", "", "gulp task to be ran (overrides package.json; forces --type manual)")
	contractsDeploy.Flags().StringVarP(&do.Path, "dir", "i", "", "root directory of app (will use $pwd by default)")
	contractsDeploy.Flags().StringVarP(&do.NewName, "dest", "e", "", "working directory to be used for deployment")
	contractsDeploy.Flags().StringVarP(&do.ConfigFile, "yaml", "y", "", "yaml file for deployment. epm apps require this; other apps ignore")
	contractsDeploy.Flags().BoolVarP(&do.Rm, "rm", "r", true, "remove containers after stopping")
}

//----------------------------------------------------

func ContractsImport(cmd *cobra.Command, args []string) {
	IfExit(ArgCheck(2, "eq", cmd, args))
	do.Name = args[0]
	do.Path = args[1]
	IfExit(contracts.GetPackage(do))
}

func ContractsExport(cmd *cobra.Command, args []string) {
	IfExit(ArgCheck(1, "eq", cmd, args))
	do.Name = args[0]
	IfExit(contracts.PutPackage(do))
	logger.Println(do.Result)
}

func ContractsTest(cmd *cobra.Command, args []string) {
	IfExit(ArgCheck(0, "eq", cmd, args))
	if do.Path == "" {
		do.Path, _ = os.Getwd() // we aren't catching this error, but revisit later if it becomes a problem
	}
	do.Name = "test"
	IfExit(contracts.RunPackage(do))
}

func ContractsDeploy(cmd *cobra.Command, args []string) {
	IfExit(ArgCheck(0, "eq", cmd, args))
	if do.Path == "" {
		do.Path, _ = os.Getwd() // we aren't catching this error, but revisit later if it becomes a problem
	}
	do.Name = "deploy"
	IfExit(contracts.RunPackage(do))
}
