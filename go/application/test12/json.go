package test12

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"reflect"
	"strings"
	"study/prototype/study/proto/editorapi"
	"study/prototype/study/proto/oneof"
	"study/prototype/study/proto/smartsheetapi"
	"study/prototype/study/proto/user_info"

	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/proto"
)

func test1() {
	jsonFilePath := "/root/work_code/programming_learning/go/study/json/user_info.json"
	pbFilePath := "/root/work_code/programming_learning/go/user_info.proto"

	buf, err := ioutil.ReadFile(jsonFilePath)
	if err != nil {
		fmt.Println("Read file err: ", err)
		os.Exit(0)
	}

	userInfo := &user_info.UserInfo{}

	if err = jsonpb.UnmarshalString(string(buf), userInfo); err != nil {
		fmt.Println("jsonpb UnmarshalString fail: ", err)
		os.Exit(0)
	}
	fmt.Println("user info pb: ", userInfo.String())

	var buf1 bytes.Buffer
	if err := (&jsonpb.Marshaler{EmitDefaults: true, EnumsAsInts: true}).Marshal(&buf1, userInfo); err != nil {
		fmt.Println("jsonpb UnmarshalString fail: ", userInfo.String())
	}

	fmt.Println("test: ", string(buf1.Bytes()))
	//fmt.Println("user info pb: ", userInfo.String())
	//fmt.Println(proto.MarshalTextString(userInfo))

	data, err := proto.Marshal(userInfo)
	if err != nil {
		fmt.Println("proto Marshal fail: ", err)
		os.Exit(0)
	}

	if err = ioutil.WriteFile(pbFilePath, data, os.ModePerm); err != nil {
		fmt.Println("Write file err: ", err)
	}
}

func jsonpbMarshal() {
	// 构造数据
	view := &editorapi.View{
		ViewID:   "yzyyzy",
		ViewName: "看板视图",
		ViewType: "1",
	}
	addViewResp := &editorapi.AddViewResponse{
		View: view,
	}
	// 序列化数据，模拟后台
	b, err := json.Marshal(addViewResp)
	if err != nil {
		fmt.Println("json Marshal fail: ", err)
		os.Exit(0)
	}
	c := &smartsheetapi.EditSmartsheetResourceRsp{
		Data: b,
	}
	// 反序列化，模拟开平
	s := &editorapi.AddViewResponse{}
	result := make(map[string]interface{})
	if err := json.Unmarshal(c.GetData(), &result); err != nil {
		fmt.Println("json Unmarshal fail: ", err)
		os.Exit(0)
	}
	fmt.Println("result", result)
	setStructField(s, result)
	fmt.Println("SmartsheetResourceResponse", s.String())
}

func setStructField(ptr interface{}, result map[string]interface{}) {
	v := reflect.ValueOf(ptr).Elem()
	for i := 0; i < v.NumField(); i++ {
		fieldInfo := v.Type().Field(i)
		tag := fieldInfo.Tag
		name := tag.Get("json")
		if name == "" {
			name = strings.ToLower(fieldInfo.Name)
		}
		name = strings.Split(name, ",")[0]
		//fmt.Println("name: ", name)
		fmt.Println("fieldInfo.Name:", fieldInfo.Name)
		if value, ok := result[name]; ok {
			fmt.Println("类型1：", reflect.ValueOf(value).Type(), "类型2：", v.FieldByName(fieldInfo.Name).Type())
			if reflect.ValueOf(value).Type() == v.FieldByName(fieldInfo.Name).Type() {
				v.FieldByName(fieldInfo.Name).Set(reflect.ValueOf(value))
			}
		}
	}
}

func jsonTest() {
	src := `{
        "deleteFields": {
            "fieldIDs": [
                "ajs78U",
                "uys6L1"
            ]
        },
        "addRecords": {
            "records": [{
                "values": {
                    "星球": "地球",
                    "大洲": "亚洲",
                    "国家": "中国",
                    "首都": "北京",
                    "省/自治区/直辖市/特别行政区": "河南",
                    "省会/首府": "郑州"
                }
            }]
        },
        "getRecords": {
            "viewID": "sdj3hy",
            "offset": 0,
            "limit": 2
        }
    }`
	/*var params interface{}
	  err := json.Unmarshal([]byte(src), &params)
	  if err != nil {
	      fmt.Printf("%v\n", err)
	      return
	  }
	  p, _ := json.Marshal(params)
	  fmt.Printf("11%v, 22%s\n", params, p)*/

	req := editorapi.SmartsheetResourceRequest{}
	if err := jsonpb.UnmarshalString(src, &req); err != nil {
		fmt.Println("jsonpb UnmarshalString fail: ", err)
		os.Exit(0)
	}
	fmt.Println("user info pb: ", req.String())

	var params interface{}
	err := json.Unmarshal([]byte(src), &params)
	if err != nil {
		fmt.Println(params)
	}
	//p, _ := json.Marshal(params)
	fmt.Printf("22%s\n", params)

	var buf bytes.Buffer
	if err := (&jsonpb.Marshaler{EmitDefaults: true, EnumsAsInts: true}).Marshal(&buf, &req); err != nil {
		fmt.Println("jsonpb Marshal fail: ", err)
		os.Exit(0)
	}
	result := make(map[string]interface{})
	if err := json.Unmarshal(buf.Bytes(), &result); err != nil {
		fmt.Println("json Unmarshal fail: ", err)
		os.Exit(0)
	}
}

func oneofTest() {
	ny := &oneof.NiuYa{
		Foo: 1,
		Oy:  "Yong Zhaoyu",
	}
	ots := &oneof.StatusTest{
		Jk: "string",
		Show: &oneof.StatusTest_Hh{
			Hh: 123,
		},
	}
	fmt.Println("NiuYa pb: ", ny.String())
	fmt.Println("StatusTest pb: ", ots.String())
	otl := &oneof.StatusTest{
		Show: &oneof.StatusTest_Test{
			Test: "zhe shi sha",
		},
	}
	fmt.Println("otl pb: ", otl.String())
	ots = &oneof.StatusTest{
		Show: &oneof.StatusTest_Test{
			Test: "Zai kankan",
		},
	}
	fmt.Println("StatusTest pb1: ", ots.String())
	tt := &oneof.Test{
		Bar: 789,
		St:  otl,
	}
	fmt.Println("StatusTest pb1: ", tt.String())
}
