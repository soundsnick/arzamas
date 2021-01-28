package docs

var CommentCreate = HandlerDocs{
	Route: "/comment/create/:id",
	Description: "Create comment",
	Method: "POST",
	Params: []HandlerDocsParam{
		{
			Name: "id",
			Type: "integer",
			Description: "Post id",
			Required: true,
		},
		{
			Name: "token",
			Type: "string",
			Description: "User authentication token",
			Required: true,
		},
		{
			Name: "content",
			Type: "string",
			Description: "Comment text",
			Required: true,
		},
	},
}

var CommentRead = HandlerDocs{
	Route: "/comment/read/:id",
	Description: "Read comment by id",
	Method: "GET",
	Params: []HandlerDocsParam{
		{
			Name: "id",
			Type: "integer",
			Description: "Comment id",
			Required: true,
		},
	},
}

var CommentUpdate = HandlerDocs{
	Route: "/comment/update/:id",
	Description: "Update comment",
	Method: "PUT",
	Params: []HandlerDocsParam{
		{
			Name: "id",
			Type: "integer",
			Description: "Comment id",
			Required: true,
		},
		{
			Name: "token",
			Type: "string",
			Description: "User authentication token",
			Required: true,
		},
		{
			Name: "content",
			Type: "string",
			Description: "Comment text",
			Required: true,
		},
	},
}

var CommentDelete = HandlerDocs{
	Route: "/comment/delete/:id",
	Description: "Delete comment",
	Method: "DELETE",
	Params: []HandlerDocsParam{
		{
			Name: "id",
			Type: "integer",
			Description: "Comment id",
			Required: true,
		},
		{
			Name: "token",
			Type: "string",
			Description: "User authentication token",
			Required: true,
		},
	},
}