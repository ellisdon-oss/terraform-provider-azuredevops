---
page_title: "azuredevops_project Resource - terraform-provider-azuredevops"
subcategory: ""
description: |-
  
---

# Resource `azuredevops_project`





## Schema

### Required

- **name** (String, Required)

### Optional

- **capabilities** (Block List, Max: 1) (see [below for nested schema](#nestedblock--capabilities))
- **description** (String, Optional)
- **id** (String, Optional) The ID of this resource.
- **visibility** (String, Optional)

<a id="nestedblock--capabilities"></a>
### Nested Schema for `capabilities`

Required:

- **process_template** (Block List, Min: 1, Max: 1) (see [below for nested schema](#nestedblock--capabilities--process_template))
- **version_control** (Block List, Min: 1, Max: 1) (see [below for nested schema](#nestedblock--capabilities--version_control))

<a id="nestedblock--capabilities--process_template"></a>
### Nested Schema for `capabilities.process_template`

Optional:

- **template_type_id** (String, Optional)


<a id="nestedblock--capabilities--version_control"></a>
### Nested Schema for `capabilities.version_control`

Optional:

- **source_control_type** (String, Optional)


