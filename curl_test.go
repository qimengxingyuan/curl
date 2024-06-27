package curl

import "testing"

func TestParse(t *testing.T) {
	cases := []string{
		`curl -X GET www.baidu.com`,
	}
	for _, c := range cases {
		curl, err := Parse(c)
		if err != nil {
			t.Errorf("Parse(%q): %v", c, err)
		} else {
			t.Logf("curl: %+v", curl)
		}

		body := map[string]interface{}{}

		err = curl.Body.UnmarshalParse(&body)
		if err != nil {
			t.Errorf("body.UnmarshalParse: %v", err)
		} else {
			t.Logf("body: %+v", body)
		}
	}
}
