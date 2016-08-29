package tasks

import "github.com/SierraSoftworks/chieftan-server/models"

type CreateUserRequest struct {
	Name        string   `json:"name"`
	Email       string   `json:"email"`
	Permissions []string `json:"permissions,omitempty"`
}

func CreateUser(req *CreateUserRequest) (*models.User, *models.AuditLogContext, error) {
	user := models.User{
		Name:        req.Name,
		Email:       req.Email,
		Permissions: req.Permissions,
		Tokens:      []string{},
	}

	if user.Permissions == nil {
		user.Permissions = []string{}
	}

	user.UpdateID()

	if err := models.DB().Users().Insert(&user); err != nil {
		return nil, nil, formatError(err)
	}

	return &user, &models.AuditLogContext{
		User:    user.Summary(),
		Request: req,
	}, nil
}
