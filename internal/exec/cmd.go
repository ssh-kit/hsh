package exec

type Cmd struct {
	Path string
	Args []string `in:"query=args"`
	Env  []string `in:"query=env"`
	Dir  string   `in:"query=dir"`
}

type Result struct {
	Stdout string `json:"stdout"`
	Stderr string `json:"stderr"`
	Code   int    `json:"code"`
}

const (
	CodeTimeout  = 124
	CodeIOError  = 125
	CodeLookup   = 127
	CodeOutRange = 255
)
