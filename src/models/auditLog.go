package models

import "time"

type AuditLog struct {
	ID        string          `json:"id" bson:"_id"`
	Type      string          `json:"type"`
	User      UserSummary     `json:"user"`
	Token     string          `json:"token"`
	Timestamp time.Time       `json:"timestamp"`
	Context   AuditLogContext `json:"context"`
}

type AuditLogContext struct {
	User    *UserSummary    `json:"user"`
	Project *ProjectSummary `json:"project"`
	Action  *ActionSummary  `json:"action"`
	Task    *TaskSummary    `json:"task"`
	Request interface{}     `json:"request"`
}
