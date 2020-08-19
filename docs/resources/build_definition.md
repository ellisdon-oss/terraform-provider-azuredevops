---
page_title: "azuredevops_build_definition Resource - terraform-provider-azuredevops"
subcategory: ""
description: |-
  
---

# Resource `azuredevops_build_definition`





## Schema

### Required

- **name** (String, Required)
- **project_id** (String, Required)
- **repository** (Block List, Min: 1, Max: 1) (see [below for nested schema](#nestedblock--repository))

### Optional

- **build_variable** (Block Set) (see [below for nested schema](#nestedblock--build_variable))
- **demand** (Block Set) (see [below for nested schema](#nestedblock--demand))
- **id** (String, Optional) The ID of this resource.
- **process** (Block List, Max: 1) (see [below for nested schema](#nestedblock--process))
- **queue** (Block List, Max: 1) (see [below for nested schema](#nestedblock--queue))
- **triggers** (Block List, Max: 1) (see [below for nested schema](#nestedblock--triggers))

### Read-only

- **revision** (Number, Read-only)

<a id="nestedblock--repository"></a>
### Nested Schema for `repository`

Required:

- **name** (String, Required)
- **properties** (Map of String, Required)
- **url** (String, Required)

Optional:

- **checkout_submodules** (Boolean, Optional)
- **clean** (String, Optional)
- **default_branch** (String, Optional)
- **type** (String, Optional)


<a id="nestedblock--build_variable"></a>
### Nested Schema for `build_variable`

Required:

- **name** (String, Required)
- **value** (String, Required)

Optional:

- **is_secret** (Boolean, Optional)


<a id="nestedblock--demand"></a>
### Nested Schema for `demand`

Required:

- **name** (String, Required)
- **value** (String, Required)


<a id="nestedblock--process"></a>
### Nested Schema for `process`

Optional:

- **phase** (Block List) (see [below for nested schema](#nestedblock--process--phase))
- **target** (Block List, Max: 1) (see [below for nested schema](#nestedblock--process--target))
- **type** (Number, Optional)
- **yaml_file_name** (String, Optional)

<a id="nestedblock--process--phase"></a>
### Nested Schema for `process.phase`

Required:

- **name** (String, Required)
- **step** (Block List, Min: 1) (see [below for nested schema](#nestedblock--process--phase--step))

Optional:

- **condition** (String, Optional)
- **job_authorization_scope** (String, Optional)
- **target** (Block List, Max: 1) (see [below for nested schema](#nestedblock--process--phase--target))

<a id="nestedblock--process--phase--step"></a>
### Nested Schema for `process.phase.step`

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


<a id="nestedblock--process--phase--target"></a>
### Nested Schema for `process.phase.target`

Optional:

- **agent_specification** (String, Optional)
- **allow_scripts_auth_access_option** (Boolean, Optional)
- **demands** (List of String, Optional)
- **queue** (Number, Optional)
- **type** (Number, Optional)



<a id="nestedblock--process--target"></a>
### Nested Schema for `process.target`

Optional:

- **agent_specification** (String, Optional)



<a id="nestedblock--queue"></a>
### Nested Schema for `queue`

Optional:

- **agent_name** (String, Optional)


<a id="nestedblock--triggers"></a>
### Nested Schema for `triggers`

Optional:

- **continuous_integration** (Block List, Max: 1) (see [below for nested schema](#nestedblock--triggers--continuous_integration))
- **pull_request** (Block List, Max: 1) (see [below for nested schema](#nestedblock--triggers--pull_request))

<a id="nestedblock--triggers--continuous_integration"></a>
### Nested Schema for `triggers.continuous_integration`

Required:

- **branch_filters** (List of String, Required)

Optional:

- **batch_changes** (Boolean, Optional)
- **max_concurrent_builds_per_branch** (Number, Optional)
- **path_filters** (List of String, Optional)
- **polling_interval** (Number, Optional)
- **polling_job_id** (String, Optional)
- **settings_source_type** (Number, Optional)


<a id="nestedblock--triggers--pull_request"></a>
### Nested Schema for `triggers.pull_request`

Required:

- **branch_filters** (List of String, Required)

Optional:

- **auto_cancel** (Boolean, Optional)
- **forks** (Block List, Max: 1) (see [below for nested schema](#nestedblock--triggers--pull_request--forks))
- **is_comment_required_for_pull_request** (Boolean, Optional)
- **path_filters** (List of String, Optional)
- **settings_source_type** (Number, Optional)

<a id="nestedblock--triggers--pull_request--forks"></a>
### Nested Schema for `triggers.pull_request.forks`

Required:

- **enabled** (Boolean, Required)

Optional:

- **allow_secrets** (Boolean, Optional)


