---
page_title: "azuredevops_release_environment Data Source - terraform-provider-azuredevops"
subcategory: ""
description: |-
  
---

# Data Source `azuredevops_release_environment`





## Schema

### Required

- **definition_id** (String, Required)
- **project_id** (String, Required)
- **stage_name** (String, Required)

### Optional

- **id** (String, Optional) The ID of this resource.

### Read-only

- **environment** (List of Object, Read-only) (see [below for nested schema](#nestedatt--environment))

<a id="nestedatt--environment"></a>
### Nested Schema for `environment`

- **condition** (List of Object) (see [below for nested schema](#nestedobjatt--environment--condition))
- **deploy_phase** (List of Object) (see [below for nested schema](#nestedobjatt--environment--deploy_phase))
- **name** (String)
- **pre_deploy_approval** (List of Object) (see [below for nested schema](#nestedobjatt--environment--pre_deploy_approval))
- **rank** (Number)
- **variable** (Set of Object) (see [below for nested schema](#nestedobjatt--environment--variable))
- **variable_groups** (List of Number)

<a id="nestedobjatt--environment--condition"></a>
### Nested Schema for `environment.condition`

- **condition_type** (String)
- **name** (String)
- **value** (String)


<a id="nestedobjatt--environment--deploy_phase"></a>
### Nested Schema for `environment.deploy_phase`

- **deployment_input** (String)
- **name** (String)
- **phase_type** (String)
- **rank** (Number)
- **workflow_task** (List of Object) (see [below for nested schema](#nestedobjatt--environment--deploy_phase--workflow_task))

<a id="nestedobjatt--environment--deploy_phase--workflow_task"></a>
### Nested Schema for `environment.deploy_phase.workflow_task`

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



<a id="nestedobjatt--environment--pre_deploy_approval"></a>
### Nested Schema for `environment.pre_deploy_approval`

- **approvals** (List of Object) (see [below for nested schema](#nestedobjatt--environment--pre_deploy_approval--approvals))
- **options** (List of Object) (see [below for nested schema](#nestedobjatt--environment--pre_deploy_approval--options))

<a id="nestedobjatt--environment--pre_deploy_approval--approvals"></a>
### Nested Schema for `environment.pre_deploy_approval.approvals`

- **approver_id** (String)
- **is_automated** (Boolean)
- **is_notification_on** (Boolean)
- **rank** (Number)


<a id="nestedobjatt--environment--pre_deploy_approval--options"></a>
### Nested Schema for `environment.pre_deploy_approval.options`

- **execution_order** (String)
- **release_creator_can_be_approver** (Boolean)
- **required_approver_count** (Number)
- **timeout_in_minutes** (Number)



<a id="nestedobjatt--environment--variable"></a>
### Nested Schema for `environment.variable`

- **is_secret** (Boolean)
- **name** (String)
- **value** (String)


