---
page_title: "azuredevops_task_group Resource - terraform-provider-azuredevops"
subcategory: ""
description: |-
  
---

# Resource `azuredevops_task_group`





## Schema

### Required

- **name** (String, Required)
- **project_id** (String, Required)
- **task** (Block List, Min: 1) (see [below for nested schema](#nestedblock--task))

### Optional

- **category** (String, Optional)
- **id** (String, Optional) The ID of this resource.
- **input** (Block List) (see [below for nested schema](#nestedblock--input))
- **runs_on** (List of String, Optional)
- **version** (Block List, Max: 1) (see [below for nested schema](#nestedblock--version))

### Read-only

- **group_id** (String, Read-only)
- **revision** (Number, Read-only)

<a id="nestedblock--task"></a>
### Nested Schema for `task`

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
- **version** (String, Optional)


<a id="nestedblock--input"></a>
### Nested Schema for `input`

Required:

- **name** (String, Required)
- **type** (String, Required)

Optional:

- **default** (String, Optional)
- **help_text** (String, Optional)
- **label** (String, Optional)
- **required** (Boolean, Optional)


<a id="nestedblock--version"></a>
### Nested Schema for `version`

Required:

- **major** (Number, Required)
- **minor** (Number, Required)
- **patch** (Number, Required)

Optional:

- **is_test** (Boolean, Optional)


