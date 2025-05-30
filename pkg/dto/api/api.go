package api

import "encoding/json"

func (p *Params) JSONRawMessage() (json.RawMessage, error) {
	marshal, err := json.Marshal(p)
	if err != nil {
		return nil, err
	}
	return marshal, nil
}

type Params struct {

	//form,table,echarts,bi,3D .....
	RenderType string       `json:"render_type"`
	Children   []*ParamInfo `json:"children"`
}

type ParamInfo struct {
	//英文标识
	Code string `json:"code"`
	//中文名称
	Name string `json:"name"`
	//中文介绍
	Desc string `json:"desc"`
	//是否必填
	Required bool `json:"required"`

	Callbacks    string      `json:"callbacks"`
	Validates    string      `json:"validates"`
	WidgetConfig interface{} `json:"widget_config"`
	WidgetType   string      `json:"widget_type"`
	ValueType    string      `json:"value_type"`
	Example      string      `json:"example"`
}

type Info struct {
	Router      string   `json:"router"`
	Method      string   `json:"method"`
	User        string   `json:"user"`
	Runner      string   `json:"runner"`
	ApiDesc     string   `json:"api_desc"`
	Labels      []string `json:"labels"`
	ChineseName string   `json:"chinese_name"`
	EnglishName string   `json:"english_name"`
	Classify    string   `json:"classify"`
	//输入参数
	ParamsIn *Params `json:"params_in"`
	//输出参数
	ParamsOut *Params  `json:"params_out"`
	UseTables []string `json:"use_tables"`
	UseDB     []string `json:"use_db"`
	Callbacks []string `json:"callbacks"`
}
