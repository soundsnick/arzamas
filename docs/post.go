package docs

var PostIndex = HandlerDocs{
	Route: "/posts",
	Description: "Get all posts",
	Method: "GET",
	Params: []HandlerDocsParam{
		HandlerDocsParam{ Name: "", Type: "" },
	},
}