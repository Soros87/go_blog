# blog

**blog** is a blockchain built using Cosmos SDK and Tendermint and created with [Ignite CLI](https://ignite.com/cli).

## Get started

```
ignite chain serve
```

`serve` command installs dependencies, builds, initializes, and starts your blockchain in development.

### Configure

Your blockchain in development can be configured with `config.yml`. To learn more, see the [Ignite CLI docs](https://docs.ignite.com).

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
curl https://get.ignite.com/username/blog@latest! | sudo bash
```

`username/blog` should match the `username` and `repo_name` of the Github repository to which the source code was pushed. Learn more about [the install process](https://github.com/allinbits/starport-installer).

## Learn more

- [Ignite CLI](https://ignite.com/cli)
- [Tutorials](https://docs.ignite.com/guide)
- [Ignite CLI docs](https://docs.ignite.com)
- [Cosmos SDK docs](https://docs.cosmos.network)
- [Developer Chat](https://discord.gg/ignite)

## Consensus-Breaking Change

This branch introduces a consensus-breaking change that affects how blocks are validated in the network. The change is [describe the change].
Eg.This branch introduces a consensus-breaking change by switching the block hash algorithm from SHA-256 to SHA-3.

### Implications

- [List the implications and potential risks of the change]
  Example
- All nodes in the network must upgrade their software to support the new block hash algorithm.
- Old and new nodes will not be able to reach consensus due to the different hash algorithms.

### Upgrade Instructions

- [Provide instructions for users or node operators on how to upgrade their software]
- Node operators must update their software to the latest version that supports the SHA-3 block hash algorithm.

### Feedback and Support

If you have questions or need support, please [provide contact information or link to support channels].
