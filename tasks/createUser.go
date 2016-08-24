package tasks

import "github.com/SierraSoftworks/chieftan-server/models"

type CreateUserRequest struct {
	Name        string
	Email       string
	Permissions []string
}

func CreateUser(req *CreateUserRequest) (*models.User, *models.AuditLogContext, error) {
	user := models.User{
		Name:        req.Name,
		Email:       req.Email,
		Permissions: req.Permissions,
		Tokens:      []string{},
	}

	user.UpdateID()

	if err := models.DB().Users().Insert(&user); err != nil {
		return nil, nil, formatError(err)
	}

	return &user, &models.AuditLogContext{
		User: user.Summary(),
	}, nil
}
