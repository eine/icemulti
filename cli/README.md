# [DTD] Command-line interface (CLI)

```
dtd <INPUT.EXT> <OUTPUT.EXT>
dtd version
dtd serve [-v] [-p <PORT>] [[-d <DIRECTORY>] | [--nofs]]
```

Now, `go run main.go` provides the same functionality as 'dtd-slim', i.e., `api` is not included, although a dummy function is shown. This is because [spf13/cobra](https://github.com/spf13/cobra) and [spf13/viper](https://github.com/spf13/viper) are not used in `api` (yet).
