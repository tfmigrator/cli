# Configuration

path | type | required | default | description
--- | --- | --- | --- | ---
.rules | rule | true | | 

## type: rule

path | type | required | default | example | description
--- | --- | --- | --- | --- | ---
if | bool expression | true | | `Resource.Type == "null_resource"` | If the result is `true`, the resource is proceeded by the rule
address | template | false | no change | `{{.Resource.Type}}.{{.Resource.Name \| replace "-" "_"}}` |
dirname | template | false | no change | `foo` |
hcl_file_basename | template | false | no change | `{{.Resource.Type}}.tf` |
state_basename | template | false | terraform.tfstate | `foo.tfstate` |
removed | bool | false | false | | If this is true, resources which match the rule are removed 
ignored | bool | false | false | | If this is true, resources which match the rule aren't migrated by tfmigrator 
skip_hcl_migration | bool | false | false | | If this is true, Terraform Configuration isn't changed
skip_state_migration | bool | false | false | | If this is true, Terraform State isn't changed

## type: bool expression

[expr](https://github.com/antonmedv/expr/blob/master/docs/Language-Definition.md) expression.
The expression must be returnes boolean (true or false).

## type: template

Go's [text/template](https://golang.org/pkg/text/template/)

[sprig](http://masterminds.github.io/sprig/) function can be used.

## expression and template parameter

[*tfmigrator.Source](https://pkg.go.dev/github.com/tfmigrator/tfmigrator@v0.5.1/tfmigrator#Source) is passed.
