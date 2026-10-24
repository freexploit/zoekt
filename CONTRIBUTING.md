# Contributing

We welcome contributions to the project! To propose a change, please fork the repository, make your changes, then submit
a pull request. If the change is significant or potentially controversial, please open an issue first to discuss it.
Zoekt does not require a CLA to contribute.

Development requires Go 1.26.10 or newer, including the standard-library security
fixes in that release. `go.mod`, `.tool-versions`, and the container builder use
this minimum version. The Nix shell provides Go from nixpkgs; with the default
`GOTOOLCHAIN=auto`, Go downloads the required toolchain if nixpkgs is older.

Before opening a pull request, make sure that you have run the tests locally:
```sh
go test ./...
```

It's also good to run a local smoke test for the relevant component. For example, if you've made changes in
`zoekt-git-index`, you can try indexing a repository locally:
```sh
go run ./cmd/zoekt-git-index /path/to/repo
```

After updating dependencies, check for known vulnerabilities as well:
```sh
go run golang.org/x/vuln/cmd/govulncheck@latest ./...
```

