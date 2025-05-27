package jsonctx

import (
	"reflect"
	"testing"
)

func TestGetJsonContextInNs(t *testing.T) {
	tests := []struct {
		name     string
		key      string
		ctx      RootJsonData
		sep      string
		loopEnd  int
		expected GetJsonContextResult
	}{
		{
			name:    "Simple object path",
			key:     "a.b.c",
			ctx:     RootJsonData{},
			sep:     ".",
			loopEnd: -1,
			expected: GetJsonContextResult{
				Context: map[string]interface{}{},
				LastNS:  "c",
				Root:    RootJsonData{"a": map[string]interface{}{"b": map[string]interface{}{}}},
			},
		},
		{
			name:    "Array notation",
			key:     "names[0].firstName",
			ctx:     RootJsonData{},
			sep:     ".",
			loopEnd: -1,
			expected: GetJsonContextResult{
				Context: []interface{}{},
				LastNS:  "firstName",
				Root:    RootJsonData{"names": []interface{}{}},
			},
		},
		{
			name:    "Mixed object and array",
			key:     "users[0].addresses[1].street",
			ctx:     RootJsonData{},
			sep:     ".",
			loopEnd: -1,
			expected: GetJsonContextResult{
				Context: []interface{}{},
				LastNS:  "street",
				Root:    RootJsonData{"users": []interface{}{}},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := GetJsonContextInNs(tt.key, tt.ctx, tt.sep, tt.loopEnd)
			
			// Compare LastNS
			if result.LastNS != tt.expected.LastNS {
				t.Errorf("LastNS = %v, want %v", result.LastNS, tt.expected.LastNS)
			}

			// Compare Context type
			if reflect.TypeOf(result.Context) != reflect.TypeOf(tt.expected.Context) {
				t.Errorf("Context type = %T, want %T", result.Context, tt.expected.Context)
			}

			// 在TestGetJsonContextInNs的t.Run中添加：
			if !reflect.DeepEqual(result.Root, tt.expected.Root) {
			    t.Errorf("Root structure = %v, want %v", result.Root, tt.expected.Root)
			}
		})
	}
}

func TestInitializeContextValue(t *testing.T) {
	tests := []struct {
		name     string
		ctx      map[string]interface{}
		current  string
		next     string
		expected interface{}
	}{
		{
			name:     "Array initialization with bracket notation",
			ctx:      map[string]interface{}{},
			current:  "names",
			next:     "[0]",
			expected: []interface{}{},
		},
		{
			name:     "Object initialization",
			ctx:      map[string]interface{}{},
			current:  "user",
			next:     "details",
			expected: map[string]interface{}{},
		},
		{
			name:     "Array initialization with numeric next",
			ctx:      map[string]interface{}{},
			current:  "items",
			next:     "0",
			expected: []interface{}{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			initializeContextValue(tt.ctx, tt.current, tt.next)
			
			result, exists := tt.ctx[tt.current]
			if !exists {
				t.Errorf("Key %s was not initialized", tt.current)
				return
			}

			if reflect.TypeOf(result) != reflect.TypeOf(tt.expected) {
				t.Errorf("Type = %T, want %T", result, tt.expected)
			}
		})
	}
}

func TestInitializeKey(t *testing.T) {
	tests := []struct {
		name       string
		ctx        map[string]interface{}
		key        string
		defaultVal interface{}
		expected   interface{}
	}{
		{
			name:       "Initialize new key",
			ctx:        map[string]interface{}{},
			key:        "test",
			defaultVal: map[string]interface{}{},
			expected:   map[string]interface{}{},
		},
		{
			name:       "Existing key not modified",
			ctx:        map[string]interface{}{"test": "value"},
			key:        "test",
			defaultVal: "new-value",
			expected:   "value",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := InitializeKey(tt.ctx, tt.key, tt.defaultVal)
			
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("Result = %v, want %v", result, tt.expected)
			}
		})
	}
}