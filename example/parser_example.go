package example

//go:generate atgen -key=json -type=Target -output=target_at.gen.go
type Target struct {
	Name string `json:"name"`
	Age  int    `json:"age"`

	Gender string
	Other
	Temp
	Sex string `json:"gender"`

	MapString map[string]string `json:"map_string"`
}

type Other struct {
	Money int    `json:"money"`
	Sex   string `json:"gender"` // 标签优先于字段名
}
type Temp struct {
	TempStr **string `json:"temp_str"`
	TmpStr  *string
	NoneTag string
	Gender  string
	XmlTag  string `xml:"tag,attr"`
	YamlTag string `yaml:"tag"`
}
