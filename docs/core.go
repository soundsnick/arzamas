package docs 

type HandlerDocsParam struct {
	Name string 
	Type string
}

type HandlerDocs struct {
	Route  string	
	Description string
	Method string	
	Params []HandlerDocsParam
}