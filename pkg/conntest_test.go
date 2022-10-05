package pkg

import (
	"encoding/json"
	"reflect"
	"strings"
	"testing"
)

func TestQueryFor(t *testing.T) {
	res := queryFor("postgres")
	if !strings.Contains(res, "information_schema") {
		t.Fail()
	}
}

func TestMarshall(t *testing.T) {
	event := NewEvent(NewResult("lorem", nil, nil, map[string]string{"lorem": "ipsum"}, 1))
	var unmarshaled Event
	marshaled, err := json.Marshal(event)
	res := json.Unmarshal(marshaled, &unmarshaled)

	if err != nil {
		t.Fail()
	}

	if reflect.DeepEqual(event, unmarshaled) {
		t.Log(unmarshaled, res)
		t.Fail()
	}
}

func TestRegisterScheme(t *testing.T) {

	t.Log("Register Databricks scheme")
	RegisterDatabricks()

	t.Log("Get protocols for database scheme")
	protocols := GetProtocols("databricks")
	t.Log(protocols)

	if isElementExist(protocols, "snowplow") != false {
		t.Fail()
	}
	if isElementExist(protocols, "databricks") != true {
		t.Fail()
	}

	t.Log("Get scheme driver and aliases for scheme")
	t.Log(SchemeDriverAndAliases("databricks"))
}

func isElementExist(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}
	return false
}

// func TestParseTags(t *testing.T) {
// 	expected := map[string]string{"dolor": "sit-amet", "lorem": "ipsum"}
// 	sample := "lorem=ipsum;dolor=sit-amet"
// 	actuals := []map[string]string{
// 		ParseTags(sample),
// 		ParseTags(";" + sample + ";"),
// 	}

// 	for i, actual := range actuals {
// 		if !reflect.DeepEqual(actual, expected) {
// 			t.Log("error at", i, ":"8, actual, expected)
// 			t.Fail()
// 		}
// 	}
// }
