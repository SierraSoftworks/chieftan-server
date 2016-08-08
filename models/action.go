package models

type Action struct {
	ID             string                `json:"id" bson:"_id"`
	Name           string                `json:"name"`
	Description    string                `json:"description"`
	Project        ProjectSummary        `json:"project"`
	Variables      map[string]string     `json:"vars"`
	Configurations []ActionConfiguration `json:"configurations"`
	HTTP           *Request              `json:"http"`
}

type ActionSummary struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type ActionConfiguration struct {
	Name      string            `json:"name"`
	Variables map[string]string `json:"vars"`
}

func (a *Action) Summary() ActionSummary {
	return ActionSummary{
		ID:          a.ID,
		Name:        a.Name,
		Description: a.Description,
	}
}
