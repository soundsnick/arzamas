package docs

var PostIndex = HandlerDocs{
	Route:       "/posts",
	Description: "Get all posts",
	Method:      "GET",
}

var PostSearch = HandlerDocs{
	Route:       "/posts/search",
	Description: "Search in posts by query",
	Method:      "GET",
	Params: []HandlerDocsParam{
		{
			Name:        "query",
			Type:        "string",
			Description: "Query string for search",
			Required:    true,
		},
	},
}

var PostUser = HandlerDocs{
	Route:       "/posts/user/:user_id",
	Description: "Get posts of user",
	Method:      "GET",
	Params: []HandlerDocsParam{
		{
			Name:     "user_id",
			Type:     "integer",
			Required: true,
		},
	},
}

var PostComments = HandlerDocs{
	Route:       "/posts/comments/:id",
	Description: "Get comments of post",
	Method:      "GET",
	Params: []HandlerDocsParam{
		{
			Name:     "id",
			Type:     "integer",
			Required: true,
		},
	},
}

var PostCreate = HandlerDocs{
	Route:       "/post/create",
	Description: "Creates post",
	Method:      "POST",
	Params: []HandlerDocsParam{
		{
			Name:        "title",
			Type:        "string",
			Required:    true,
			Description: "Post title",
		},
		{
			Name:        "content",
			Type:        "string",
			Required:    true,
			Description: "Post content",
		},
		{
			Name:        "cover",
			Type:        "string",
			Required:    true,
			Description: "Post cover absolute url",
		},
		{
			Name:        "token",
			Type:        "string",
			Required:    true,
			Description: "User authentication token",
		},
	},
}

var PostGet = HandlerDocs{
	Route:       "/post/read/:id",
	Description: "Get post by id",
	Method:      "GET",
	Params: []HandlerDocsParam{
		{
			Name:     "id",
			Type:     "integer",
			Required: true,
		},
	},
}

var PostUpdate = HandlerDocs{
	Route:       "/post/update/:id",
	Description: "Update post (pass param you want to change)",
	Method:      "PUT",
	Params: []HandlerDocsParam{
		{
			Name:        "title",
			Type:        "string",
			Description: "Post title",
		},
		{
			Name:        "content",
			Type:        "string",
			Description: "Post content",
		},
		{
			Name:        "cover",
			Type:        "string",
			Description: "Post cover absolute url",
		},
		{
			Name:        "token",
			Type:        "string",
			Required:    true,
			Description: "User authentication token",
		},
	},
}

var PostDelete = HandlerDocs{
	Route:       "/post/delete/:id",
	Description: "Delete post by id",
	Method:      "DELETE",
	Params: []HandlerDocsParam{
		{
			Name:     "id",
			Type:     "integer",
			Required: true,
		},
	},
}
