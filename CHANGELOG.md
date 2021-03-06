# 0.3.0 (July 17, 2020)

BREAKING CHANGES:
  - Stop exporting `exec.Cmd` versions of methods ([#25](https://github.com/hashicorp/terraform-plugin-sdk/issues/25))
  - Require `address` and `id` arguments in `Import()` ([#24](https://github.com/hashicorp/terraform-plugin-sdk/issues/24))
  - Rename `StateShow()` to `Show()` ([#30](https://github.com/hashicorp/terraform-plugin-sdk/issues/30))
  
BUG FIXES:
  - Fix bug in `Import()` config argument ([#28](https://github.com/hashicorp/terraform-plugin-sdk/issues/28))

# 0.2.2 (July 13, 2020)

BUG FIXES:
  - Version number is now correctly reported by the tfinstall package. Please note that `tfinstall.Version` was incorrect between versions 0.1.1 and 0.2.1 inclusive.

# 0.2.1 (July 10, 2020)

BUG FIXES:
  - Minor code changes to allow for compilation in Go 1.12 ([#21](https://github.com/hashicorp/terraform-exec/pull/21))

# 0.2.0 (July 8, 2020)

NEW FEATURES:
  - add `Import()` function ([#20](https://github.com/hashicorp/terraform-exec/pull/20))

# 0.1.1 (July 7, 2020)

BUG FIXES:
 - Downgrade `github.com/hashicorp/go-getter` dependency, which added a requirement for Go 1.13.

# 0.1.0 (July 3, 2020)

Initial release. 

This Go module contains two packages, `github.com/hashicorp/terraform-exec/tfexec`, and `github.com/hashicorp/terraform-exec/tfinstall`, which share the same version.
