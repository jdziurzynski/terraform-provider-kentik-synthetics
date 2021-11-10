package synthetics

import (
	"fmt"
	"time"
)

func formatTime(t *time.Time) string {
	if t == nil {
		return ""
	}
	return t.Format(time.RFC3339Nano)
}

func getObjectFromNestedResourceData(data interface{}) (map[string]interface{}, error) {
	dataSlice, ok := data.([]interface{})
	if !ok {
		return nil, fmt.Errorf("invalid data type, got: %T, want: []interface{}", data)
	}

	if len(dataSlice) == 0 {
		return nil, nil
	}

	if dataSlice[0] == nil {
		return nil, nil
	}

	m, ok := dataSlice[0].(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf(
			"invalid dataSlice[0] type, got: %T, want: map[string]interface{}",
			dataSlice[0],
		)
	}

	return m, nil
}

func ifSliceToStringSlice(s []interface{}) []string {
	if s == nil {
		return nil
	}

	result := make([]string, 0, len(s))
	for _, v := range s {
		result = append(result, v.(string))
	}
	return result
}

func ifSliceToInt64Slice(s []interface{}) []int64 {
	if s == nil {
		return nil
	}

	result := make([]int64, 0, len(s))
	for _, v := range s {
		result = append(result, int64(v.(int)))
	}
	return result
}

func stringInterfaceMapToStringMap(m map[string]interface{}) (map[string]string, error) {
	if m == nil {
		return nil, nil
	}

	result := make(map[string]string, len(m))
	for k, v := range m {
		vv, ok := v.(string)
		if !ok {
			return nil, fmt.Errorf("stringInterfaceMapToStringMap: m[%v] should be string, got: %#v", k, v)
		}
		result[k] = vv
	}
	return result, nil
}

func stringMapPtrToStringMap(p *map[string]string) map[string]string {
	if p == nil {
		return nil
	}
	return *p
}
