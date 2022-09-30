include "global" {
  path           = "${get_terragrunt_dir()}/../../../../_base_modules/{{global}}.hcl"
  expose         = true
  merge_strategy = "deep"
}