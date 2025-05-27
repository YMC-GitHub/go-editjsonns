package jsonctx

import (
	"go-editjsonns/pkg/jsonns"
	"regexp"
)

// RootJsonData represents a JSON object type
type RootJsonData map[string]interface{}

// GetJsonContextResult holds the result of getting JSON context
type GetJsonContextResult struct {
	Context interface{}  // 当前路径下的JSON上下文内容
	LastNS  string       // 最后一个命名空间部分(路径的最后一段)
	Root    RootJsonData // 完整的根JSON数据结构
}

// GetJsonContextInNs gets JSON context using namespace notation
func GetJsonContextInNs(key string, ctx RootJsonData, sep string, loopEnd int) GetJsonContextResult {
	if sep == "" {
		sep = "."
	}
	keys := jsonns.NSStd(key, sep, nil)
	root := make(RootJsonData)
	if len(ctx) > 0 {
		// Copy the original context to avoid modifying it directly
		for k, v := range ctx {
			root[k] = v
		}
	}

	end := getLastIndex(len(keys), loopEnd)

	currentCtx := interface{}(root)

	for i := 0; i < end; i++ {
		cur := keys[i]
		var next string
		if i+1 < len(keys) {
			next = keys[i+1]
		}
		name := jsonns.NSPureName(cur)

		switch ctx := currentCtx.(type) {
		case map[string]interface{}:
			initializeContextValue(ctx, name, next)
			currentCtx = ctx[name]
			// Update the root context with the new structure
			if i == 0 {
				root[name] = ctx[name]
			} else {
				updateRootContext(root, keys[:i+1], ctx[name])
			}
		case []interface{}:
			// For array context, just return the slice itself
			// Don't initialize with empty map as first element
			return GetJsonContextResult{
				Context: ctx,
				LastNS:  jsonns.NSPureName(keys[end]),
				Root:    root,
			}
		default:
			newCtx := make(map[string]interface{})
			initializeContextValue(newCtx, name, next)
			currentCtx = newCtx[name]
			// Update the root context with the new structure
			if i == 0 {
				root[name] = newCtx[name]
			} else {
				updateRootContext(root, keys[:i+1], newCtx[name])
			}
		}
	}

	last := jsonns.NSPureName(keys[end])
	return GetJsonContextResult{
		Context: currentCtx,
		LastNS:  last,
		Root:    root,
	}
}

// getLastIndex calculates the final index for iteration based on length and loopEnd parameters
func getLastIndex(length int, loopEnd int) int {
	if loopEnd > 0 {
		return length - loopEnd
	}
	return length + loopEnd
}

// initializeContextValue initializes a value in the context based on the next namespace part
func initializeContextValue(ctx map[string]interface{}, current string, next string) {
	if _, exists := ctx[current]; !exists {
		pureName := jsonns.NSPureName(next)

		// Check if next part indicates an array
		isArrayIndex := regexp.MustCompile(`\[\d+\]`).MatchString(next) ||
			regexp.MustCompile(`^\d+$`).MatchString(pureName)

		if isArrayIndex {
			ctx[current] = make([]interface{}, 0)
		} else {
			ctx[current] = make(map[string]interface{})
		}
	}
}

// InitializeKey initializes a key in the context with a default value
func InitializeKey(ctx map[string]interface{}, key string, defaultVal interface{}) interface{} {
	if val, exists := ctx[key]; exists {
		return val
	}
	ctx[key] = defaultVal
	return defaultVal
}

// updateRootContext updates the root context with the new structure
func updateRootContext(root RootJsonData, keys []string, value interface{}) {
	current := root
	for i, key := range keys {
		name := jsonns.NSPureName(key)
		if i == len(keys)-1 {
			current[name] = value
		} else {
			if _, ok := current[name]; !ok {
				current[name] = make(map[string]interface{})
			}
			current = current[name].(map[string]interface{})
		}
	}
}
