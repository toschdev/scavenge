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
				{
					RpcMethod:      "ListQuestions",
					Use:            "list-questions",
					Short:          "Query list-questions",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{},
				},

				{
					RpcMethod:      "ShowQuestion",
					Use:            "show-question [question-id]",
					Short:          "Query show-question",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "questionId"}},
				},

				{
					RpcMethod:      "ListCommits",
					Use:            "list-commits",
					Short:          "Query list-commits",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{},
				},

				{
					RpcMethod:      "ShowCommit",
					Use:            "show-commit [question-id] [creator]",
					Short:          "Query show-commit",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "questionId"}, {ProtoField: "creator"}},
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
				{
					RpcMethod:      "RevealAnswer",
					Use:            "reveal-answer [question-id] [answer] [plain-text]",
					Short:          "Send a reveal-answer tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "questionId"}, {ProtoField: "answer"}, {ProtoField: "plainText"}},
				},
				{
					RpcMethod:      "CompleteQuestion",
					Use:            "complete-question [question-id] [winner]",
					Short:          "Send a complete-question tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "questionId"}, {ProtoField: "winner"}},
				},
				// this line is used by ignite scaffolding # autocli/tx
			},
		},
	}
}
