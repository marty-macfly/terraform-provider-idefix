---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "idefix_project Resource - terraform-provider-idefix"
subcategory: ""
description: |-
  Manages project.
---

# idefix_project (Resource)

Manages project.

## Example Usage

```terraform
resource "idefix_project" "example" {
  name            = "example"
  company_id      = 1234
  contract_number = "1234"
  wbs_france      = "MYWBS"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `company_id` (Number) The company ID associated to the CI.
- `contract_number` (String) Contract number
- `name` (String) The name of project (must be unique).

### Optional

- `parent_id` (Number) The ID of the parent project.
- `wbs_belgique` (String) The WBS of this project
- `wbs_canada` (String) The WBS of this project
- `wbs_chine` (String) The WBS of this project
- `wbs_france` (String) The WBS of this project
- `wbs_hong_kong` (String) The WBS of this project
- `wbs_luxembourg` (String) The WBS of this project
- `wbs_maurice` (String) The WBS of this project
- `wbs_singapour` (String) The WBS of this project
- `wbs_vietnam` (String) The WBS of this project

### Read-Only

- `id` (String) The id of the project.

## Import

Import is supported using the following syntax:

```shell
terraform import idefix_project.example 1234
```
