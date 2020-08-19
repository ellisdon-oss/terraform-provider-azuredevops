---
page_title: "azuredevops_release_stage_variables Data Source - terraform-provider-azuredevops"
subcategory: ""
description: |-
  
---

# Data Source `azuredevops_release_stage_variables`





## Schema

### Required

- **definition_id** (String, Required)
- **project_id** (String, Required)
- **stage_name** (String, Required)

### Optional

- **id** (String, Optional) The ID of this resource.

### Read-only

- **variables** (Set of Object, Read-only) (see [below for nested schema](#nestedatt--variables))

<a id="nestedatt--variables"></a>
### Nested Schema for `variables`

- **is_secret** (Boolean)
- **name** (String)
- **value** (String)


