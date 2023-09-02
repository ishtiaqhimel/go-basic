## Building CLI Applications
Command line interfaces (CLIs), unlike graphical user interfaces (GUIs), are text-only. Cloud and infrastructure applications are primarily CLI-based due to their easy automation and remote capabilities.

## CLI Library in Go
- [Cobra](https://github.com/spf13/cobra)
- [Urfave CLI](https://github.com/urfave/cli)

## Cobra
Cobra is a library for creating powerful modern CLI applications.

Cobra is built on a structure of commands, arguments & flags.

Commands represent actions, Args are things and Flags are modifiers for those actions.

Using Cobra is easy. First, use `go get` to install the latest version of the library.
```
go get -u github.com/spf13/cobra@latest
```

Next, include Cobra in your application:
```
import "github.com/spf13/cobra"
```

`cobra-cli` is a command line program to generate cobra applications and command files. It will bootstrap your application scaffolding to rapidly develop a Cobra-based application. It is the easiest way to incorporate Cobra into your application.

It can be installed by running:
```
go install github.com/spf13/cobra-cli@latest
```

To learn more see [cobra generator](https://github.com/spf13/cobra-cli/blob/main/README.md) and [user guide](https://github.com/spf13/cobra/blob/main/site/content/user_guide.md)


