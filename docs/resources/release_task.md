---
page_title: "azuredevops_release_task Resource - terraform-provider-azuredevops"
subcategory: ""
description: |-
  
---

# Resource `azuredevops_release_task`





## Schema

### Required

- **definition_id** (Number, Required)
- **project_id** (String, Required)
- **stage_name** (String, Required)
- **task_info** (Block List, Min: 1, Max: 1) (see [below for nested schema](#nestedblock--task_info))

### Optional

- **after** (String, Optional)
- **before** (String, Optional)
- **id** (String, Optional) The ID of this resource.
- **job_name** (String, Optional)
- **job_rank** (Number, Optional)
- **rank** (Number, Optional)
- **replace** (Boolean, Optional)

<a id="nestedblock--task_info"></a>
### Nested Schema for `task_info`

Required:

- **inputs** (Map of String, Required)
- **name** (String, Required)
- **task_id** (String, Required)

Optional:

- **always_run** (Boolean, Optional)
- **condition** (String, Optional)
- **continue_on_error** (Boolean, Optional)
- **definition_type** (String, Optional)
- **enabled** (Boolean, Optional)
- **environment** (Map of String, Optional)
- **ref_name** (String, Optional)
- **version** (String, Optional)


