---
page_title: "azuredevops_variable_group Data Source - terraform-provider-azuredevops"
subcategory: ""
description: |-
  
---

# Data Source `azuredevops_variable_group`





## Schema

### Required

- **project_id** (String, Required)

### Optional

- **group_id** (Number, Optional)
- **id** (String, Optional) The ID of this resource.
- **name** (String, Optional)

### Read-only

- **variables** (Set of Object, Read-only) (see [below for nested schema](#nestedatt--variables))

<a id="nestedatt--variables"></a>
### Nested Schema for `variables`

- **is_secret** (Boolean)
- **name** (String)
- **value** (String)


