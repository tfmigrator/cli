# Example 1

## Requirements

* Terraform
* tfmigrator CLI

## Setup

Run `terraform apply` to create `terraform.tfstate`.

```
$ terraform init
$ terraform apply
```

## Dry Run

```console
$ tfmigrator run -dry-run main.tf
2021/07/18 18:28:18 [INFO] [DRYRUN] + terraform state mv null_resource.foo null_resource.bar
migrated_resources:
- source_address: null_resource.foo
  source_tf_file_path: main.tf
  new_address: null_resource.bar
removed_resources: []
not_migrated_resources: []
```

## Migrate

```console
$ tfmigrator run main.tf
2021/07/18 18:31:24 [INFO] + terraform state mv null_resource.foo null_resource.bar
migrated_resources:
- source_address: null_resource.foo
  source_tf_file_path: main.tf
  new_address: null_resource.bar
removed_resources: []
not_migrated_resources: []
```

Confirm that main.tf is updated without losing code comment.

```diff
$ git diff main.tf
diff --git a/examples/example1/main.tf b/examples/example1/main.tf
index b9aaa38..43f6c1a 100644
--- a/examples/example1/main.tf
+++ b/examples/example1/main.tf
@@ -1,5 +1,5 @@
 # comment
-resource "null_resource" "foo" {}
+resource "null_resource" "bar" {}
 
 locals {
   foo = "foo"
```
