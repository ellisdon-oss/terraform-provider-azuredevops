---
page_title: "azuredevops_release_tasks Data Source - terraform-provider-azuredevops"
subcategory: ""
description: |-
  
---

# Data Source `azuredevops_release_tasks`





## Schema

### Required

- **definition_id** (String, Required)
- **job_rank** (Number, Required)
- **project_id** (String, Required)
- **stage_name** (String, Required)

### Optional

- **id** (String, Optional) The ID of this resource.

### Read-only

- **tasks** (List of Object, Read-only) (see [below for nested schema](#nestedatt--tasks))

<a id="nestedatt--tasks"></a>
### Nested Schema for `tasks`

- **always_run** (Boolean)
- **condition** (String)
- **continue_on_error** (Boolean)
- **definition_type** (String)
- **enabled** (Boolean)
- **environment** (Map of String)
- **inputs** (Map of String)
- **name** (String)
- **ref_name** (String)
- **task_id** (String)
- **version** (String)


