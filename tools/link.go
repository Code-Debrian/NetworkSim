package tools

type LinkJSON struct {
	PropogationRate float32 `json:"prop_rate"`
	Length          int     `json:"length"`
	Source          int     `json:"source"`
	Target          int     `json:"target"`
}
