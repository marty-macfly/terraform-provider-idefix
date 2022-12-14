---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "idefix_ci Resource - terraform-provider-idefix"
subcategory: ""
description: |-
  Manages CI.
---

# idefix_ci (Resource)

Manages CI.

## Example Usage

```terraform
resource "idefix_ci" "example" {
  name        = "myci"
  company_id  = 1234
  project_ids = [1, 2]
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `company_id` (Number) The company ID associated to the CI.
- `name` (String) The name of this CI.
- `project_ids` (List of Number) The projects associated to the CI.

### Optional

- `comment` (String) Comment.
- `is_owner_lbn` (Boolean) The owner of the CI.
- `key_dates` (Block Set) Use And Key Date. (see [below for nested schema](#nestedblock--key_dates))
- `outsourcing_name` (String) The Outsourcing level name.
- `service_at` (Block Set) Services AT. (see [below for nested schema](#nestedblock--service_at))
- `service_cloud` (Block Set) Service Cloud. (see [below for nested schema](#nestedblock--service_cloud))
- `service_level_id` (Number) The Level of the service.
- `team` (String) The team in charge.
- `type_id` (Number) The type of the CI.

### Read-Only

- `id` (String) The ID of this resource.

<a id="nestedblock--key_dates"></a>
### Nested Schema for `key_dates`

Required:

- `environment_ids` (List of Number) Environments of the CI.
- `function_ids` (List of Number) Functions of the CI.


<a id="nestedblock--service_at"></a>
### Nested Schema for `service_at`

Required:

- `monitoring_tool` (List of Number) Monitoring Tool IDs.
- `required_services` (List of Number) Required Services IDs.


<a id="nestedblock--service_cloud"></a>
### Nested Schema for `service_cloud`

Required:

- `product_id` (Number) The Product ID of the CI.
- `region_id` (Number) The Region ID of the CI.
- `subscription_id` (Number) The Subscription ID of the CI.

## Import

Import is supported using the following syntax:

```shell
terraform import idefix_ci.example 1234
```
