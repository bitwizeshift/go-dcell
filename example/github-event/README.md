# `github-event` example

This small example demonstrates how the `dcell` library can be used to
emulate GitHub Actions event expressions.

## Usage

This program takes only 1 argument: the expression to evaluate.

```bash
go run ./example/github-event <expression>
```

### Example

```bash
$ go run ./example/github-event "github.event_name == 'push' && github.actor == 'octocat'"
true
```
