# Usage

```console
$ tfmigrator help
NAME:
   tfmigrator - Migrate Terraform Configuration and State. https://github.com/tfmigrator/cli

USAGE:
   tfmigrator [global options] command [command options] [arguments...]

VERSION:
   0.1.0

COMMANDS:
   run      Migrate Terraform Configuration and State
   help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h     show help (default: false)
   --version, -v  print the version (default: false)
```

```console
$ tfmigrator help run
NAME:
   tfmigrator run - Migrate Terraform Configuration and State

USAGE:
   tfmigrator run [command options] [arguments...]

OPTIONS:
   --log-level value         log level
   --config value, -c value  configuration file path
   --dry-run                 dry run (default: false)
   --state value             the output of 'terraform show -json'
```
