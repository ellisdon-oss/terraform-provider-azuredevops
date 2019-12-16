package helper

import "github.com/hashicorp/terraform-plugin-sdk/helper/schema"

func EnvironmentResourceSchema() *schema.Schema {
	return &schema.Schema{
		Type:     schema.TypeList,
		Optional: true,
		Elem: &schema.Resource{
			Schema: EnvironmentSchema(),
		},
	}
}

func EnvironmentSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"name": &schema.Schema{
			Type:     schema.TypeString,
			Required: true,
			ForceNew: true,
		},
		"rank": &schema.Schema{
			Type:     schema.TypeInt,
			Required: true,
		},
		"pre_deploy_approval": preDeployApprovalSchema(),
		"variable":            ReleaseVariableSchema(),
		"variable_groups": &schema.Schema{
			Type:     schema.TypeList,
			Optional: true,
			Elem:     &schema.Schema{Type: schema.TypeInt},
		},
		"condition":    conditionSchema(),
		"deploy_phase": deployPhaseSchema(),
	}
}

func EnvironmentSingleSchema() *schema.Schema {
	return &schema.Schema{
		Type:     schema.TypeList,
		MinItems: 1,
		Required: true,
		Elem: &schema.Resource{
			Schema: EnvironmentSchema(),
		},
	}
}

func preDeployApprovalSchema() *schema.Schema {
	return &schema.Schema{
		Type:     schema.TypeList,
		MaxItems: 1,
		Optional: true,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"options":   optionsSchema(),
				"approvals": approvalsSchema(),
			},
		},
	}
}

func approvalsSchema() *schema.Schema {
	return &schema.Schema{
		Type:     schema.TypeList,
		Optional: true,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"approver_id": &schema.Schema{
					Type:     schema.TypeString,
					Optional: true,
				},
				"rank": &schema.Schema{
					Type:     schema.TypeInt,
					Optional: true,
				},
				"is_automated": &schema.Schema{
					Type:     schema.TypeBool,
					Optional: true,
					Default:  true,
				},
				"is_notification_on": &schema.Schema{
					Type:     schema.TypeBool,
					Optional: true,
					Default:  false,
				},
			},
		},
	}
}

func optionsSchema() *schema.Schema {
	return &schema.Schema{
		Type:     schema.TypeList,
		Optional: true,
		MaxItems: 1,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"required_approver_count": &schema.Schema{
					Type:     schema.TypeInt,
					Optional: true,
					Default:  0,
				},
				"execution_order": &schema.Schema{
					Type:     schema.TypeString,
					Optional: true,
					Default:  "beforeGates",
				},
				"timeout_in_minutes": &schema.Schema{
					Type:     schema.TypeInt,
					Required: true,
				},
				"release_creator_can_be_approver": &schema.Schema{
					Type:     schema.TypeBool,
					Optional: true,
					Default:  true,
				},
			},
		},
	}
}

func conditionSchema() *schema.Schema {
	return &schema.Schema{
		Type:     schema.TypeList,
		Optional: true,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"condition_type": &schema.Schema{
					Type:     schema.TypeString,
					Required: true,
				},
				"name": &schema.Schema{
					Type:     schema.TypeString,
					Required: true,
				},
				"value": &schema.Schema{
					Type:     schema.TypeString,
					Optional: true,
				},
			},
		},
	}
}

func deployPhaseSchema() *schema.Schema {
	return &schema.Schema{
		Type:     schema.TypeList,
		Required: true,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"workflow_task": WorkflowTaskSchema(),
				"phase_type": &schema.Schema{
					Type:     schema.TypeString,
					Required: true,
				},
				"name": &schema.Schema{
					Type:     schema.TypeString,
					Required: true,
				},
				"rank": &schema.Schema{
					Type:     schema.TypeInt,
					Required: true,
				},
				"deployment_input": &schema.Schema{
					Type:     schema.TypeString,
					Required: true,
				},
			},
		},
	}
}

func WorkflowTaskSingleSchema() *schema.Schema {
	return &schema.Schema{
		Type:     schema.TypeList,
		Required: true,
		MaxItems: 1,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"name": &schema.Schema{
					Type:     schema.TypeString,
					Required: true,
				},
				"definition_type": &schema.Schema{
					Type:     schema.TypeString,
					Optional: true,
					Default:  "task",
				},
				"version": &schema.Schema{
					Type:     schema.TypeString,
					Optional: true,
					Default:  "0.*",
				},
				"task_id": &schema.Schema{
					Type:     schema.TypeString,
					Required: true,
				},
				"enabled": &schema.Schema{
					Type:     schema.TypeBool,
					Optional: true,
					Default:  true,
				},
				"always_run": &schema.Schema{
					Type:     schema.TypeBool,
					Optional: true,
					Default:  false,
				},
				"ref_name": &schema.Schema{
					Type:     schema.TypeString,
					Optional: true,
				},
				"continue_on_error": &schema.Schema{
					Type:     schema.TypeBool,
					Optional: true,
					Default:  false,
				},
				"condition": &schema.Schema{
					Type:     schema.TypeString,
					Optional: true,
					Default:  "succeeded()",
				},
				"environment": &schema.Schema{
					Type:     schema.TypeMap,
					Optional: true,
					Elem: &schema.Schema{
						Type: schema.TypeString,
					},
				},
				"inputs": &schema.Schema{
					Type:     schema.TypeMap,
					Required: true,
					Elem: &schema.Schema{
						Type: schema.TypeString,
					},
				},
			},
		},
	}
}

func WorkflowTaskSchema() *schema.Schema {
	return &schema.Schema{
		Type:     schema.TypeList,
		Required: true,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"name": &schema.Schema{
					Type:     schema.TypeString,
					Required: true,
				},
				"definition_type": &schema.Schema{
					Type:     schema.TypeString,
					Optional: true,
					Default:  "task",
				},
				"version": &schema.Schema{
					Type:     schema.TypeString,
					Optional: true,
					Default:  "0.*",
				},
				"task_id": &schema.Schema{
					Type:     schema.TypeString,
					Required: true,
				},
				"enabled": &schema.Schema{
					Type:     schema.TypeBool,
					Optional: true,
					Default:  true,
				},
				"ref_name": &schema.Schema{
					Type:     schema.TypeString,
					Optional: true,
				},
				"always_run": &schema.Schema{
					Type:     schema.TypeBool,
					Optional: true,
					Default:  false,
				},
				"continue_on_error": &schema.Schema{
					Type:     schema.TypeBool,
					Optional: true,
					Default:  false,
				},
				"condition": &schema.Schema{
					Type:     schema.TypeString,
					Optional: true,
					Default:  "succeeded()",
				},
				"environment": &schema.Schema{
					Type:     schema.TypeMap,
					Optional: true,
					Elem: &schema.Schema{
						Type: schema.TypeString,
					},
				},
				"inputs": &schema.Schema{
					Type:     schema.TypeMap,
					Required: true,
					Elem: &schema.Schema{
						Type: schema.TypeString,
					},
				},
			},
		},
	}
}

func ReleaseVariableSchema() *schema.Schema {
	return &schema.Schema{
		Type:     schema.TypeSet,
		Optional: true,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"is_secret": &schema.Schema{
					Type:     schema.TypeBool,
					Optional: true,
					Default:  false,
				},
				"name": &schema.Schema{
					Type:     schema.TypeString,
					Required: true,
				},
				"value": &schema.Schema{
					Type:     schema.TypeString,
					Required: true,
				},
			},
		},
	}
}

func ArtifactSchema() *schema.Schema {
	return &schema.Schema{
		Type:     schema.TypeSet,
		Optional: true,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"alias": &schema.Schema{
					Type:     schema.TypeString,
					Required: true,
				},
				"source_id": &schema.Schema{
					Type:     schema.TypeString,
					Required: true,
				},
				"type": &schema.Schema{
					Type:     schema.TypeString,
					Required: true,
				},
				"definition_reference": &schema.Schema{
					Type:     schema.TypeString,
					Required: true,
				},
			},
		},
	}
}

func TriggerSchema() *schema.Schema {
	return &schema.Schema{
		Type:     schema.TypeSet,
		Optional: true,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"alias": &schema.Schema{
					Type:     schema.TypeString,
					Required: true,
				},
				"trigger_type": &schema.Schema{
					Type:     schema.TypeString,
					Required: true,
				},
				"branch_filters": &schema.Schema{
					Type:     schema.TypeList,
					Optional: true,
					Elem:     &schema.Schema{Type: schema.TypeString},
				},
			},
		},
	}
}
