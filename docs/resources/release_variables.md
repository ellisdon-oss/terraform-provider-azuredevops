---
page_title: "azuredevops_release_variables Resource - terraform-provider-azuredevops"
subcategory: ""
description: |-
  
---

# Resource `azuredevops_release_variables`





## Schema

### Required

- **definition_id** (Number, Required)
- **project_id** (String, Required)

### Optional

- **id** (String, Optional) The ID of this resource.
- **variable** (Block Set) (see [below for nested schema](#nestedblock--variable))

<a id="nestedblock--variable"></a>
### Nested Schema for `variable`

Required:

- **name** (String, Required)
- **value** (String, Required)

Optional:

- **is_secret** (Boolean, Optional)
- **stage_name** (String, Optional)


