package docs

type HandlerDocsParam struct {
	Name        string
	Type        string
	Description string `json:"Description,omitempty"`
	Required    bool
}

type HandlerDocs struct {
	Route       string
	Description string
	Method      string
	Params      []HandlerDocsParam `json:"Params,omitempty"`
}
