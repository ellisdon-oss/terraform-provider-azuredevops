// GatedCheckInTrigger

  pathFilters []string
  runContinuousIntegration bool
  useWorkspaceMappings bool

// BuildCompletionTrigger

  branchFilters []string
  definition DefinitionReference
  requiresSuccessfulBuild bool

// ScheduleTrigger
  schedules []Schedule

				"branch_filters" []string
				"days_to_build" string
				"schedule_job_id" string
				"schedule_only_with_changes" bool
				"start_hours" int32
				"start_minutes" int32
				"time_zone_id" string

// List
  "all",
  "batchedContinuousIntegration",
  "batchedGatedCheckIn",
  "buildCompletion",
  "continuousIntegration",
  "gatedCheckIn",
  "none",
  "pullRequest",
  "schedule",
