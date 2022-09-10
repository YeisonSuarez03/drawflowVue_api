package drawflow

type Program struct {
	Uid         string   `json:"uid,omitempty"`
	DgraphType  []string `json:"dgraph.type,omitempty"`
	Name        string   `json:"name,omitempty"`
	Description string   `json:"description,omitempty"`
	Drawflow    string   `json:"drawflow,omitempty"`
}

type JsonPrograms struct {
	Programs []Program `json:"programs"`
}
type JsonProgram struct {
	Programs []Program `json:"program"`
}
type JsonError struct {
	Error       string `json:"error"`
	Description string `json:"description"`
	Status      int    `json:"status"`
}