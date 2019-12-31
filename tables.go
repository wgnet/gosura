package gosura

type TableArgs struct {
	Schema string `json:"schema"`
	Name   string `json:"name"`
}

func DefaultTableSchemaArgs(name string) TableArgs {
	return TableArgs{
		Name:   name,
		Schema: "public",
	}
}
