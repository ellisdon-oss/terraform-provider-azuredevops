---
page_title: "azuredevops_variable_group Resource - terraform-provider-azuredevops"
subcategory: ""
description: |-
  
---

# Resource `azuredevops_variable_group`





## Schema

### Required

- **name** (String, Required)
- **project_id** (String, Required)
- **variable** (Block Set, Min: 1) (see [below for nested schema](#nestedblock--variable))

### Optional

- **id** (String, Optional) The ID of this resource.

### Read-only

- **group_id** (Number, Read-only)

<a id="nestedblock--variable"></a>
### Nested Schema for `variable`

Required:

- **name** (String, Required)
- **value** (String, Required)

Optional:

- **is_secret** (Boolean, Optional)


