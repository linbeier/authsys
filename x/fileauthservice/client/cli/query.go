package cli

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/codec"

	"github.com/linbeier/authsys/x/fileauthservice/internal/types"
)

// GetQueryCmd returns the cli query commands for this module
func GetQueryCmd(queryRoute string, cdc *codec.Codec) *cobra.Command {
	// Group fileauthservice queries under a subcommand
	fileauthserviceQueryCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("Querying commands for the %s module", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	fileauthserviceQueryCmd.AddCommand(
		flags.GetCommands(
			// TODO: Add query Cmds
			GetCmdGetAccounts(queryRoute, cdc),
			GetCmdGetFiles(queryRoute, cdc),
			GetCmdGetAuth(queryRoute, cdc),
			GetCmdGetRecords(queryRoute, cdc),
			GetCmdGetTrace(queryRoute, cdc),
		)...,
	)

	return fileauthserviceQueryCmd
}

// TODO: Add Query Commands
func GetCmdGetAccounts(queryRoute string, cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "accounts",
		Short: "return accounts",
		//Args:	cobra.ExactArgs()
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)

			res, _, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/accounts", queryRoute), nil)
			if err != nil {
				fmt.Printf("could not get query accounts\n %s\n", err.Error())
				return nil
			}

			var out types.QueryResAccounts
			cdc.MustUnmarshalJSON(res, &out)
			return cliCtx.PrintOutput(out)
		},
	}
}

func GetCmdGetFiles(queryRoute string, cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "files [account]",
		Short: "account string as arg, return files",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)
			account := args[0]
			res, _, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/files/%s", queryRoute, account), nil)
			if err != nil {
				fmt.Printf("could not get this accounts' files\n %s\n", err.Error())
				return nil
			}

			var out []string
			cdc.MustUnmarshalJSON(res, &out)
			return cliCtx.PrintOutput(out)
		},
	}
}

func GetCmdGetAuth(queryRoute string, cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "authority [account] [filename] [filehash]",
		Short: "account string, file name and file hash needed, return the authority that the account has",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)
			account := args[0]
			filename := args[1]
			filehash := args[2]

			res, _, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/authority/%s/%s/%s", queryRoute, account, filename, filehash), nil)
			if err != nil {
				fmt.Printf("could not get this file's authority\n %s\n", err.Error())
				return nil
			}

			var out string
			cdc.MustUnmarshalJSON(res, &out)
			return cliCtx.PrintOutput(out)
		},
	}
}

func GetCmdGetRecords(queryRoute string, cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "records [account]",
		Short: "Get records about an account",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)
			account := args[0]

			res, _, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/records/%s", queryRoute, account), nil)
			if err != nil {
				fmt.Printf("could not get records\n %s\n", err.Error())
				return nil
			}

			var out []types.Filerecord
			cdc.MustUnmarshalJSON(res, &out)
			return cliCtx.PrintOutput(out)

		},
	}
}

func GetCmdGetTrace(queryRoute string, cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "trace [account] [filename]",
		Short: "trace account's file, return its tracing route",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)
			account := args[0]
			filename := args[1]
			fp, err0 := os.OpenFile(fmt.Sprintf("/home/lin/go/src/github.com/linbeier/authsys/data.json"), os.O_CREATE|os.O_RDWR, 0755)
			if err0 != nil {
				fmt.Printf("%s", err0)
			}
			defer fp.Close()

			var keymap map[string][]byte
			keymap = make(map[string][]byte)
			keymapbyte, err0 := ioutil.ReadAll(fp)

			if err0 != nil {
				fmt.Printf("error with read\n")
			}
			if !bytes.Equal(nil, keymapbyte) {
				codec.Cdc.MustUnmarshalJSON(keymapbyte, &keymap)
			}

			res, _, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/trace/%s/%s", queryRoute, account, filename), nil)
			if err != nil {
				fmt.Printf("could not get records\n %s\n", err.Error())
				return nil
			}

			var out types.Filerecord
			cdc.MustUnmarshalJSON(res, &out)
			cliCtx.PrintOutput(out)

			for !out.From.Equals(out.Origin) {
				account = ""
				for key, val := range keymap {
					if bytes.Equal(val, out.From) {
						account = key
					}
				}

				if account != "" {
					res, _, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/trace/%s/%s", queryRoute, account, filename), nil)
					if err != nil {
						fmt.Printf("could not get more records\n %s\n", err.Error())
						return nil
					}

					cdc.MustUnmarshalJSON(res, &out)
					cliCtx.PrintOutput(out)
				} else {
					return errors.New("could not find this account")
				}
				fmt.Printf("from: %s\n", out.From)
				fmt.Printf("origin: %s\n", out.Origin)
				fmt.Printf("from与origin是否相等？: %t\n", out.From.Equals(out.Origin))
			}

			return nil
		},
	}
}
