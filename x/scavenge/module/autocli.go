package scavenge

import (
	autocliv1 "cosmossdk.io/api/cosmos/autocli/v1"

	modulev1 "scavenge/api/scavenge/scavenge"
)

// AutoCLIOptions implements the autocli.HasAutoCLIConfig interface.
func (am AppModule) AutoCLIOptions() *autocliv1.ModuleOptions {
	return &autocliv1.ModuleOptions{
		Query: &autocliv1.ServiceCommandDescriptor{
			Service: modulev1.Query_ServiceDesc.ServiceName,
			RpcCommandOptions: []*autocliv1.RpcCommandOptions{
				{
					RpcMethod: "Params",
					Use:       "params",
					Short:     "Shows the parameters of the module",
				},
				// this line is used by ignite scaffolding # autocli/query
			},
		},
		Tx: &autocliv1.ServiceCommandDescriptor{
			Service:              modulev1.Msg_ServiceDesc.ServiceName,
			EnhanceCustomCommand: true, // only required if you want to use the custom command
			RpcCommandOptions: []*autocliv1.RpcCommandOptions{
				{
					RpcMethod: "UpdateParams",
					Skip:      true, // skipped because authority gated
				},
				{
					RpcMethod:      "CreateQuestion",
					Use:            "create-question [question] [encrypted-answer] [bounty]",
					Short:          "Send a create-question tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "question"}, {ProtoField: "encryptedAnswer"}, {ProtoField: "bounty"}},
				},
				{
					RpcMethod:      "CommitAnswer",
					Use:            "commit-answer [question-id] [hash-answer]",
					Short:          "Send a commit-answer tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "questionId"}, {ProtoField: "hashAnswer"}},
				},
				// this line is used by ignite scaffolding # autocli/tx
			},
		},
	}
}
