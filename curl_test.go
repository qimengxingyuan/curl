package curl

import "testing"

func TestParse(t *testing.T) {
	cases := []string{
		`curl 'https://dolphin.boe.bytedance.net/api/v2/draft/new?draft_id=1&type=' \
  -H 'accept: */*' \
  -H 'accept-language: zh-CN,zh;q=0.9' \
  -H 'content-type: application/json' \
  -H 'origin: http://localhost:8441' \
  -H 'priority: u=1, i' \
  -H 'referer: http://localhost:8441/' \
  -H 'sec-ch-ua: "Google Chrome";v="125", "Chromium";v="125", "Not.A/Brand";v="24"' \
  -H 'sec-ch-ua-mobile: ?0' \
  -H 'sec-ch-ua-platform: "macOS"' \
  -H 'sec-fetch-dest: empty' \
  -H 'sec-fetch-mode: cors' \
  -H 'sec-fetch-site: cross-site' \
  -H 'user-agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/125.0.0.0 Safari/537.36' \
  -H 'x-jwt-token: eyJhbGciOiJSUzI1NiIsImtpZCI6IiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJwYWFzLnBhc3Nwb3J0LmF1dGgiLCJleHAiOjE3MTkyMzYyOTIsImlhdCI6MTcxOTIzMjYzMiwidXNlcm5hbWUiOiJkaW5nd2Vpd3UudG9wIiwidHlwZSI6InBlcnNvbl9hY2NvdW50IiwicmVnaW9uIjoiY24iLCJ0cnVzdGVkIjp0cnVlLCJ1dWlkIjoiZDQ3ZWYxOTktMzNiNy00OWNiLThkMzMtYWIzYjY0ODljNTcwIiwic2l0ZSI6ImJvZSIsInNjb3BlIjoiYnl0ZWRhbmNlIiwic2VxdWVuY2UiOiJSRCIsIm9yZ2FuaXphdGlvbiI6IkRhdGEt6aOO5o6nLemjjuaOp-eglOWPkS3liY3nq68iLCJ3b3JrX2NvdW50cnkiOiJDSE4iLCJsb2NhdGlvbiI6IkNOIiwiYXZhdGFyX3VybCI6Imh0dHBzOi8vczEtaW1maWxlLmZlaXNodWNkbi5jb20vc3RhdGljLXJlc291cmNlL3YxL3YzXzAwNmlfNmMxNzA3Y2UtOTlhOS00NjExLWIyODYtM2EzNTJmNmI4YzZnfj9pbWFnZV9zaXplPW5vb3BcdTAwMjZjdXRfdHlwZT1cdTAwMjZxdWFsaXR5PVx1MDAyNmZvcm1hdD1wbmdcdTAwMjZzdGlja2VyX2Zvcm1hdD0ud2VicCIsImVtYWlsIjoiZGluZ3dlaXd1LnRvcEBieXRlZGFuY2UuY29tIiwiZW1wbG95ZWVfaWQiOjMzMTE3ODV9.cdSfkahIg8_H3gnXieEImVr3SyGJh7Bc_3vkCFZyI1EB9H6_43O4CEd9X-gmOK_8pMDpcNbsoNkEvzk708sZnviH2p2bIu70ROy5MQs0TKXdYrTARDImJT6vLF2stvyz8UR1I2XK-xhIXExKCROYLhaJRYUI81dw1NE9dUFxjBU' \
  -H 'x-tt-env: boe_zyh_dev' \
  --data-raw '{"bizline_id":120,"group_id":306071,"event_id":32444,"content":"{\"origin_version\":10,\"rules\":[{\"decision\":{\"operation\":\"M_DECISION\",\"config\":\"{\\\"kv_v2\\\":[{\\\"op_type\\\":\\\"SET\\\",\\\"op_key\\\":\\\"1212312312123123123123123123123123123123123123\\\",\\\"op_value\\\":\\\"\\\\\\\"23123123123123123123123123123123177\\\\\\\"\\\"},{\\\"op_type\\\":\\\"RETURN\\\",\\\"op_key\\\":\\\"\\\",\\\"op_value\\\":\\\"\\\"},{\\\"op_type\\\":\\\"ACTION\\\",\\\"op_key\\\":\\\"testDataBus\\\",\\\"op_value\\\":\\\"\\\"}],\\\"kv\\\":{\\\"SET\\\":{\\\"1212312312123123123123123123123123123123123123\\\":\\\"\\\\\\\"23123123123123123123123123123123177\\\\\\\"\\\"},\\\"RETURN\\\":{},\\\"ACTION\\\":{\\\"testDataBus\\\":\\\"\\\"}}}\",\"id\":304660,\"visual_decisions\":null},\"id\":252522,\"name\":\"New Rule 1\",\"expression\":\"true\",\"disabled\":false,\"visual_info\":null}]}"}'`,
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
