package models

import "time"

type AuditLog struct {
	ID        string          `json:"id" bson:"_id"`
	Type      string          `json:"type"`
	User      UserSummary     `json:"user,omitempty"`
	Token     string          `json:"token"`
	Timestamp time.Time       `json:"timestamp"`
	Context   AuditLogContext `json:"context"`
}

type AuditLogContext struct {
	User    *UserSummary    `json:"user,omitempty"`
	Project *ProjectSummary `json:"project,omitempty"`
	Action  *ActionSummary  `json:"action,omitempty"`
	Task    *TaskSummary    `json:"task,omitempty"`
	Request interface{}     `json:"request,omitempty"`
}
