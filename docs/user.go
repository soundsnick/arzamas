package docs

var UserByEmail = HandlerDocs{
	Route:       "/users/email/:email",
	Description: "Get user by email",
	Method:      "GET",
	Params: []HandlerDocsParam{
		{
			Name:        "email",
			Type:        "string",
			Required:    true,
			Description: "Email string",
		},
	},
}

var UsersByName = HandlerDocs{
	Route:       "/users/name/:name",
	Description: "Get users by name",
	Method:      "GET",
	Params: []HandlerDocsParam{
		{
			Name:        "name",
			Type:        "string",
			Required:    true,
			Description: "Name of user",
		},
	},
}

var UsersByLastName = HandlerDocs{
	Route:       "/users/last_name/:name",
	Description: "Get users by last name",
	Method:      "GET",
	Params: []HandlerDocsParam{
		{
			Name:        "name",
			Type:        "string",
			Required:    true,
			Description: "Last name of user",
		},
	},
}

var UserRead = HandlerDocs{
	Route:       "/user/read/:id",
	Description: "Get user by id",
	Method:      "GET",
	Params: []HandlerDocsParam{
		{
			Name:        "id",
			Type:        "integer",
			Required:    true,
			Description: "id of user",
		},
	},
}

var UserUpdate = HandlerDocs{
	Route:       "/user/update",
	Description: "Update user (pass field you want to change)",
	Method:      "PUT",
	Params: []HandlerDocsParam{
		{
			Name:        "token",
			Type:        "string",
			Required:    true,
			Description: "User authentication token",
		},
		{
			Name:        "name",
			Type:        "string",
			Required:    false,
			Description: "Name of user",
		},
		{
			Name:        "last_name",
			Type:        "string",
			Required:    false,
			Description: "Last name of user",
		},
		{
			Name:        "avatar",
			Type:        "string",
			Required:    false,
			Description: "Avatar absolute url",
		},
	},
}

var UserDelete = HandlerDocs{
	Route:       "/user/delete",
	Description: "Delete user by token",
	Method:      "DELETE",
	Params: []HandlerDocsParam{
		{
			Name:        "token",
			Type:        "string",
			Required:    true,
			Description: "User authentication token",
		},
	},
}

var UserAuthenticate = HandlerDocs{
	Route:       "/user/auth",
	Description: "Authenticate user",
	Method:      "POST",
	Params: []HandlerDocsParam{
		{
			Name:     "email",
			Type:     "string",
			Required: true,
		},
		{
			Name:     "password",
			Type:     "string",
			Required: true,
		},
	},
}

var UserRegister = HandlerDocs{
	Route:       "/user/register",
	Description: "Register user",
	Method:      "POST",
	Params: []HandlerDocsParam{
		{
			Name:     "email",
			Type:     "string",
			Required: true,
		},
		{
			Name:     "name",
			Type:     "string",
			Required: true,
		},
		{
			Name:     "last_name",
			Type:     "string",
			Required: true,
		},
		{
			Name:     "password",
			Type:     "string",
			Required: true,
		},
		{
			Name:     "password_confirmation",
			Type:     "string",
			Required: true,
		},
	},
}
