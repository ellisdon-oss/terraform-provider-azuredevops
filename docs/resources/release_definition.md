---
page_title: "azuredevops_release_definition Resource - terraform-provider-azuredevops"
subcategory: ""
description: |-
  
---

# Resource `azuredevops_release_definition`





## Schema

### Required

- **name** (String, Required)
- **project_id** (String, Required)

### Optional

- **artifact** (Block Set) (see [below for nested schema](#nestedblock--artifact))
- **environment** (Block List) (see [below for nested schema](#nestedblock--environment))
- **id** (String, Optional) The ID of this resource.
- **path** (String, Optional)
- **release_variable** (Block Set) (see [below for nested schema](#nestedblock--release_variable))
- **release_variable_groups** (List of Number, Optional)
- **trigger** (Block Set) (see [below for nested schema](#nestedblock--trigger))

### Read-only

- **revision** (Number, Read-only)

<a id="nestedblock--artifact"></a>
### Nested Schema for `artifact`

Required:

- **alias** (String, Required)
- **definition_reference** (String, Required)
- **source_id** (String, Required)
- **type** (String, Required)


<a id="nestedblock--environment"></a>
### Nested Schema for `environment`

Required:

- **deploy_phase** (Block List, Min: 1) (see [below for nested schema](#nestedblock--environment--deploy_phase))
- **name** (String, Required)
- **rank** (Number, Required)

Optional:

- **condition** (Block List) (see [below for nested schema](#nestedblock--environment--condition))
- **pre_deploy_approval** (Block List, Max: 1) (see [below for nested schema](#nestedblock--environment--pre_deploy_approval))
- **variable** (Block Set) (see [below for nested schema](#nestedblock--environment--variable))
- **variable_groups** (List of Number, Optional)

<a id="nestedblock--environment--deploy_phase"></a>
### Nested Schema for `environment.deploy_phase`

Required:

- **deployment_input** (String, Required)
- **name** (String, Required)
- **phase_type** (String, Required)
- **rank** (Number, Required)
- **workflow_task** (Block List, Min: 1) (see [below for nested schema](#nestedblock--environment--deploy_phase--workflow_task))

<a id="nestedblock--environment--deploy_phase--workflow_task"></a>
### Nested Schema for `environment.deploy_phase.workflow_task`

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



<a id="nestedblock--environment--condition"></a>
### Nested Schema for `environment.condition`

Required:

- **condition_type** (String, Required)
- **name** (String, Required)

Optional:

- **value** (String, Optional)


<a id="nestedblock--environment--pre_deploy_approval"></a>
### Nested Schema for `environment.pre_deploy_approval`

Optional:

- **approvals** (Block List) (see [below for nested schema](#nestedblock--environment--pre_deploy_approval--approvals))
- **options** (Block List, Max: 1) (see [below for nested schema](#nestedblock--environment--pre_deploy_approval--options))

<a id="nestedblock--environment--pre_deploy_approval--approvals"></a>
### Nested Schema for `environment.pre_deploy_approval.approvals`

Optional:

- **approver_id** (String, Optional)
- **is_automated** (Boolean, Optional)
- **is_notification_on** (Boolean, Optional)
- **rank** (Number, Optional)


<a id="nestedblock--environment--pre_deploy_approval--options"></a>
### Nested Schema for `environment.pre_deploy_approval.options`

Required:

- **timeout_in_minutes** (Number, Required)

Optional:

- **execution_order** (String, Optional)
- **release_creator_can_be_approver** (Boolean, Optional)
- **required_approver_count** (Number, Optional)



<a id="nestedblock--environment--variable"></a>
### Nested Schema for `environment.variable`

Required:

- **name** (String, Required)
- **value** (String, Required)

Optional:

- **is_secret** (Boolean, Optional)



<a id="nestedblock--release_variable"></a>
### Nested Schema for `release_variable`

Required:

- **name** (String, Required)
- **value** (String, Required)

Optional:

- **is_secret** (Boolean, Optional)


<a id="nestedblock--trigger"></a>
### Nested Schema for `trigger`

Required:

- **alias** (String, Required)
- **trigger_type** (String, Required)

Optional:

- **branch_filters** (List of String, Optional)


