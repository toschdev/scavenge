# scavenge
**scavenge** is a blockchain built using Cosmos SDK and Tendermint and created with [Ignite CLI](https://ignite.com/cli).

## Get started

```
ignite chain serve
```

`serve` command installs dependencies, builds, initializes, and starts your blockchain in development.

### Configure

Your blockchain in development can be configured with `config.yml`. To learn more, see the [Ignite CLI docs](https://docs.ignite.com).

## Scavenge Hunt

Example walkthrough

```bash
# Create New Scavenge
scavenged tx scavenge create-question "I have cities, but no houses. I have mountains, but no trees. I have water, but no fish. I have roads, but no cars. What am I?" "60be9861750facbfad8758254a2f76c0cfe78d54459a3bc187d49b1401fcd8e8" 100 --from alice --chain-id scavenge


# Queries
scavenged q scavenge show-question 0

scavenged q scavenge list-questions

# Commit Answer
scavenged tx scavenge commit-answer 0 --from bob --chain-id scavenge

# Queries
scavenged q scavenge show-commit 0

scavenged q scavenge list-commits

# Reveal Answer
scavenged tx scavenge reveal-answer 0 map --from bob --chain-id scavenge
```

### Web Frontend

Additionally, Ignite CLI offers both Vue and React options for frontend scaffolding:

For a Vue frontend, use: `ignite scaffold vue`
For a React frontend, use: `ignite scaffold react`
These commands can be run within your scaffolded blockchain project. 


For more information see the [monorepo for Ignite front-end development](https://github.com/ignite/web).

## Release
To release a new version of your blockchain, create and push a new tag with `v` prefix. A new draft release with the configured targets will be created.

```
git tag v0.1
git push origin v0.1
```

After a draft release is created, make your final changes from the release page and publish it.

### Install
To install the latest version of your blockchain node's binary, execute the following command on your machine:

```
curl https://get.ignite.com/username/scavenge@latest! | sudo bash
```
`username/scavenge` should match the `username` and `repo_name` of the Github repository to which the source code was pushed. Learn more about [the install process](https://github.com/allinbits/starport-installer).

## Learn more

- [Ignite CLI](https://ignite.com/cli)
- [Tutorials](https://docs.ignite.com/guide)
- [Ignite CLI docs](https://docs.ignite.com)
- [Cosmos SDK docs](https://docs.cosmos.network)
- [Developer Chat](https://discord.gg/ignite)
