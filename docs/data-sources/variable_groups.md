---
page_title: "azuredevops_variable_groups Data Source - terraform-provider-azuredevops"
subcategory: ""
description: |-
  
---

# Data Source `azuredevops_variable_groups`





## Schema

### Required

- **project_id** (String, Required)

### Optional

- **id** (String, Optional) The ID of this resource.

### Read-only

- **groups** (List of Object, Read-only) (see [below for nested schema](#nestedatt--groups))

<a id="nestedatt--groups"></a>
### Nested Schema for `groups`

- **group_id** (Number)
- **name** (String)
- **variables** (List of Object) (see [below for nested schema](#nestedobjatt--groups--variables))

<a id="nestedobjatt--groups--variables"></a>
### Nested Schema for `groups.variables`

- **is_secret** (Boolean)
- **name** (String)
- **value** (String)


