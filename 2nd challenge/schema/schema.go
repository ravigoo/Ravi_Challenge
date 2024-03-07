package schema

type MyJsonName struct {
	List2 struct {
		L string `json:"L"`
	} `json:"list_2"`
	List3 struct {
		L []string `json:"L"`
	} `json:"list_3"`
	Map1 struct {
		M struct {
			Bool1 struct {
				Bool string `json:"BOOL"`
			} `json:"bool_1"`
			List1 struct {
				L []struct {
					Bool string `json:"BOOL"`
					N    string `json:"N"`
					Null string `json:"NULL"`
					S    string `json:"S"`
				} `json:"L"`
			} `json:"list_1"`
			Null1 struct {
				NULL string `json:"NULL "`
			} `json:"null_1"`
		} `json:"M"`
	} `json:"map_1"`
	Number1 struct {
		N string `json:"N"`
	} `json:"number_1"`
	String1 struct {
		S string `json:"S"`
	} `json:"string_1"`
	String2 struct {
		S string `json:"S"`
	} `json:"string_2"`
	EmptyField struct {
		S string `json:"S"`
	} `json:""`
}

type InputFile struct {
	Number1    string  `json:"number_1,omitempty"`
	String1    string  `json:"string_1,omitempty"`
	String2    string  `json:"string_2,omitempty"`
	Map1       Map1    `json:"map_1,omitempty"`
	List2      GetList `json:"list_2,omitempty"`
	List3      GetList `json:"list_3,omitempty"`
	EmptyField string  `json:""`
}
type Map1 struct {
	M struct {
		Bool1 string `json:"bool_1,omitempty"`
		Null1 string `json:"null_1,omitempty"`
		List1 List1  `json:"list_1,omitempty"`
	} `json:"M"`
}
type List1 struct {
	L []L `json:"L"`
}
type GetList struct {
	L []string `json:"L"`
}
type L struct {
	S    string `json:"S,omitempty"`
	N    string `json:"N,omitempty"`
	Bool string `json:"BOOL,omitempty"`
	Null string `json:"NULL,omitempty"`
}
type OutputFile []struct {
	Map1 struct {
		List1 []interface{} `json:"list_1"`
		Null1 interface{}   `json:"null_1"`
	} `json:"map_1"`
	Number1 float64 `json:"number_1"`
	String1 string  `json:"string_1"`
	String2 int64   `json:"string_2"`
}
