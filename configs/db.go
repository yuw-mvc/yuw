package configs

type (
	MPoT struct {
		Types string `json:"types"`// Todo: "ONE"->"FetchOne" "ALL"->"FetchAll"
		Table string `json:"table"`
		Field string `json:"field"`
		Joins [][]interface{} `json:"joins"`
		Limit int `json:"limit"`
		Start []int `json:"start"`
		Query interface{} `json:"query"`
		QueryArgs []interface{} `json:"query_args"`
		Columns []string `json:"columns"`
		OrderType string `json:"order_type"`
		OrderArgs []string `json:"order_args"`
	}
)
