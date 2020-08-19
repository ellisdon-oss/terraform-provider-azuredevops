---
page_title: "azuredevops_service_hook Resource - terraform-provider-azuredevops"
subcategory: ""
description: |-
  
---

# Resource `azuredevops_service_hook`





## Schema

### Required

- **consumer** (Block List, Min: 1, Max: 1) (see [below for nested schema](#nestedblock--consumer))
- **event_type** (String, Required)
- **publisher** (Block List, Min: 1, Max: 1) (see [below for nested schema](#nestedblock--publisher))

### Optional

- **custom_path** (String, Optional)
- **id** (String, Optional) The ID of this resource.

<a id="nestedblock--consumer"></a>
### Nested Schema for `consumer`

Required:

- **action_id** (String, Required)
- **id** (String, Required) The ID of this resource.
- **inputs** (Map of String, Required)


<a id="nestedblock--publisher"></a>
### Nested Schema for `publisher`

Required:

- **id** (String, Required) The ID of this resource.
- **inputs** (Map of String, Required)


