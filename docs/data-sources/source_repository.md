---
page_title: "azuredevops_source_repository Data Source - terraform-provider-azuredevops"
subcategory: ""
description: |-
  
---

# Data Source `azuredevops_source_repository`





## Schema

### Required

- **org_name** (String, Required)
- **project_id** (String, Required)
- **repo_name** (String, Required)
- **service_endpoint_id** (String, Required)
- **type** (String, Required)

### Optional

- **id** (String, Optional) The ID of this resource.

### Read-only

- **default_branch** (String, Read-only)
- **properties** (Map of String, Read-only)
- **url** (String, Read-only)


