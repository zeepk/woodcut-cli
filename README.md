# Woodcut CLI Tool

## Installation

### Homebrew

Run the following series of commands

```
brew tap zeepk/woodcut-cli https://github.com/zeepk/woodcut-cli
brew install woodcut-cli
```

### Manual

Install manually from the [releases page](https://github.com/zeepk/woodcut-cli/releases)

## Commands

> All commands begin with `woodcut` with flags appended - e.g. `woodcut ge Lobster` or `woodcut vos -h`

- `ge [item]` - returns the current Grand Exchange price for that item
- `stats [username]` - displays that users current skills
- `vos` - returns current Voice of Seren clans
  - `-h` - shows the last 10 VoS clan hours

## Dev Docs

Project created with [Cobra](https://cobra.dev/) using the [Cobra CLI](https://github.com/spf13/cobra-cli/blob/main/README.md)

### Release instructions

1. Create tag with `git tag -a v0.1.0 -m "commit message"`
2. Push tag with `git push origin v0.1.0`
3. Create release on Github using [this link](https://github.com/zeepk/woodcut-cli/releases/new)
4. Push build files to the release with `goreleaser release --rm-dist`
5. Update version in `package.json`
6. Publish to NPM with `npm publish --access public`
