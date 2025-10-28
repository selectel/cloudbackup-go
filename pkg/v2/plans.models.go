package cloudbackup

type (
	Plan struct {
		BackupMode        string          `json:"backup_mode"`
		CreatedAt         string          `json:"created_at,omitempty"`
		ID                string          `json:"id,omitempty"`
		FullBackupsAmount int             `json:"full_backups_amount"`
		Name              string          `json:"name"`
		Resources         []*PlanResource `json:"resources"`
		SchedulePattern   string          `json:"schedule_pattern"`
		ScheduleType      string          `json:"schedule_type"`
		Status            string          `json:"status,omitempty"`
	}

	PlanResource struct {
		ID   string `json:"id"`
		Name string `json:"name"`
		Type string `json:"type"`
	}

	PlansResponse struct {
		Plans []*Plan `json:"plans"`
		Total int     `json:"total"`
	}
)

const (
	PlanStatusStarted   = "started"
	PlanStatusSuspended = "suspended"
)

type PlanUpdateRequest struct {
	FullBackupsAmount int             `json:"full_backups_amount"`
	Name              string          `json:"name"`
	Resources         []*PlanResource `json:"resources"`
	SchedulePattern   string          `json:"schedule_pattern"`
	ScheduleType      string          `json:"schedule_type"`
}
