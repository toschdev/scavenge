package cli

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strconv"

	"github.com/spf13/cobra"

	"scavenge/x/scavenge/types"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
)

// GetTxCmd returns the transaction commands for the scavenge module
func GetTxCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("%s transactions subcommands", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(
		CmdCreateQuestion(),
		CmdCommitSolution(),
		CmdRevealSolution(),
	)

	return cmd
}

func CmdCommitSolution() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "commit-answer [question-id] [solution]",
		Short: "Commit a answer to a scavenge question",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			questionID, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return err
			}

			// Get the solution from args
			solution := args[1]

			// Get creator address
			creator := clientCtx.GetFromAddress().String()

			plainTextToHash := []byte(solution)
			plainTextSha := sha256.Sum256(plainTextToHash)
			encodedPlainText := hex.EncodeToString(plainTextSha[:])

			// Create the hash from solution and creator address
			commitBytes := []byte(encodedPlainText + creator)
			hash := sha256.Sum256(commitBytes)
			hashString := hex.EncodeToString(hash[:])

			msg := types.NewMsgCommitAnswer(
				creator,
				questionID,
				hashString,
			)

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)
	return cmd
}

func CmdCreateQuestion() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-question [question] [answer] [bounty]",
		Short: "Create a new scavenge question with a bounty",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			question := args[0]
			answer := args[1]

			// Hash the answer
			answerHash := sha256.Sum256([]byte(answer))
			answerHashString := hex.EncodeToString(answerHash[:])

			bounty, err := strconv.ParseUint(args[2], 10, 64)
			if err != nil {
				return err
			}

			msg := types.NewMsgCreateQuestion(
				clientCtx.GetFromAddress().String(),
				question,
				answerHashString,
				bounty,
			)

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)
	return cmd
}

func CmdRevealSolution() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "reveal-answer [question-id] [solution]",
		Short: "Reveal the answer for a committed answer",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			questionID, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return err
			}

			answer := args[1]

			// Hash the answer
			answerHash := sha256.Sum256([]byte(answer))
			answerHashString := hex.EncodeToString(answerHash[:])

			msg := types.NewMsgRevealAnswer(
				clientCtx.GetFromAddress().String(),
				questionID,
				answerHashString,
				args[1],
			)

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)
	return cmd
}
