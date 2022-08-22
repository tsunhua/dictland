# Dictland

[中文版戳我](readme-zh.md)

**Dictland** is an OPEN and PUBLIC Blockchain Project for ALL Dictionaries, built using Cosmos SDK and Tendermint and created with [Ignite CLI](https://ignite.com/cli).

## Get started

```
ignite chain serve
```

`serve` command installs dependencies, builds, initializes, and starts Dictland in development.

### Configure

Dictland in development can be configured with `config.yml`. To learn more, see the [Ignite CLI docs](https://docs.ignite.com).

## Release

To release a new version of Dictland, create and push a new tag with `v` prefix. A new draft release with the configured targets will be created.

```
git tag v0.1
git push origin v0.1
```

After a draft release is created, make your final changes from the release page and publish it.

### Install

To install the latest version of Dictland node's binary, execute the following command on your machine:

```
curl https://get.ignite.com/tsunhua/dictland@latest! | sudo bash
```

## Learn more

- [Ignite CLI docs](https://docs.ignite.com)
- [Cosmos SDK docs](https://docs.cosmos.network)
