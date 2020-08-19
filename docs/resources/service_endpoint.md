---
page_title: "azuredevops_service_endpoint Resource - terraform-provider-azuredevops"
subcategory: ""
description: |-
  
---

# Resource `azuredevops_service_endpoint`





## Schema

### Required

- **authorization** (Block List, Min: 1, Max: 1) (see [below for nested schema](#nestedblock--authorization))
- **name** (String, Required)
- **owner** (String, Required)
- **project_id** (String, Required)
- **type** (String, Required)
- **url** (String, Required)

### Optional

- **allow_all_pipelines** (Boolean, Optional)
- **data** (Map of String, Optional)
- **id** (String, Optional) The ID of this resource.

<a id="nestedblock--authorization"></a>
### Nested Schema for `authorization`

Required:

- **parameters** (Map of String, Required)
- **scheme** (String, Required)


