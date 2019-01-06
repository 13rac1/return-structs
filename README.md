# Return Structs Linter

_Checks functions return structs, not interfaces._

## Why

"Accept Interfaces, Return Structs"

The output of this should be considered a warning or notification, not an error
because there are some valid reasons to return an interface.

I hadn't used the Go ast/parser libraries yet, so this was an excuse.

## Run

One file:

```bash
go build
./return-structs example/main.go
#Returned interface found on line 15: UserInterface
```

Entire project:

```
find . -name "*.go" -not -path "./vendor/*" | xargs -n1 return-structs
```

# TODO

- Support `./...` to search entire projects
- Support multiple files
- Display a help screen
- Tests
- Documentation
