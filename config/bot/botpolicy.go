package bot

type Botpolicy struct {
	Builtin     interface{} `json:"builtin,omitempty"`
	Comment     string      `json:"comment,omitempty"`
	Feature     string      `json:"feature,omitempty"`
	Hits        int         `json:"hits,omitempty"`
	Logaction   string      `json:"logaction,omitempty"`
	Name        string      `json:"name,omitempty"`
	Newname     string      `json:"newname,omitempty"`
	Profilename string      `json:"profilename,omitempty"`
	Rule        string      `json:"rule,omitempty"`
	Undefaction string      `json:"undefaction,omitempty"`
	Undefhits   int         `json:"undefhits,omitempty"`
}
