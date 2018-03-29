# [icemulti] Command-line interface (CLI)

```
./cli -v 1 list ../test/ -r -o db.json
```

At now, `go run main.go` provides the same functionality as 'icemulti-slim', i.e., `api` is not included, although a dummy function is shown. This is because [spf13/cobra](https://github.com/spf13/cobra) and [spf13/viper](https://github.com/spf13/viper) are not used in `api` (yet).

---

To format the codebase either run `yarn fmt-cli` or `yarn fmt-all` from the root of the repo, or run `gofmt -s -w .` from subdir `cli`.
