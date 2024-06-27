package curl

import "testing"

func TestGetArgsToken(t *testing.T) {
	cases := []string{
		//"curl -X POST -H 'Content-Type: application/json' -d '{\"key\":\"value\"}' https:www.example.com/api/v1/test",
		"curl --location 'https://dolphin.bccc.net/access/v2/mbasic_judge/test/test' \\\n--header 'Content-Type: application/json' \\\n--header 'Authorization: Basic emhhb3lvbmc6c2FjamRh' \\\n--data '{\n    \"1\":\"{}\"\n}'",
	}

	for _, c := range cases {
		args, err := GetArgsToken(c)
		if err != nil {
			t.Errorf("GetArgsToken error: %v", err)
		}
		t.Logf("args: %+v", args)
	}
}
