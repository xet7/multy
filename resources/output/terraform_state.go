package output

import (
	"encoding/json"
	"fmt"
)

type TfResource struct {
	Address string                 `json:"address"`
	Values  map[string]interface{} `json:"values"`
}

type TfState struct {
	Values struct {
		RootModule struct {
			Resources []TfResource `json:"resources"`
		} `json:"root_module"`
	} `json:"values"`
}

func (t *TfState) GetValues(resourceType any, resourceId string) (map[string]interface{}, error) {
	address := fmt.Sprintf("%s.%s", GetResourceName(resourceType), resourceId)
	return t.Get(address)
}
func (t *TfState) Get(resourceRef string) (map[string]interface{}, error) {
	if o, exists := t.MaybeGet(resourceRef); exists {
		return o, nil
	}

	return nil, fmt.Errorf("unable to parse resource %s as it doesn't exist on state", resourceRef)
}

func (t *TfState) MaybeGet(resourceRef string) (map[string]interface{}, bool) {
	for _, r := range t.Values.RootModule.Resources {
		if r.Address == resourceRef {
			return r.Values, true
		}
	}

	return nil, false
}

func GetParsed[T any](state *TfState, resourceRef string) (*T, error) {
	rawResourceState, err := state.Get(resourceRef)
	if err != nil {
		return nil, err
	}

	jsonResourceState, err := json.Marshal(rawResourceState)
	if err != nil {
		return nil, err
	}

	stateResource := new(T)
	err = json.Unmarshal(jsonResourceState, stateResource)
	if err != nil {
		return nil, err
	}

	return stateResource, nil
}

func MaybeGetParsed[T any](state *TfState, resourceRef string) (*T, bool, error) {
	rawResourceState, exists := state.MaybeGet(resourceRef)
	if !exists {
		return nil, exists, nil
	}

	jsonResourceState, err := json.Marshal(rawResourceState)
	if err != nil {
		return nil, exists, err
	}

	stateResource := new(T)
	err = json.Unmarshal(jsonResourceState, stateResource)
	if err != nil {
		return nil, exists, err
	}

	return stateResource, exists, nil
}

func GetParsedById[T any](state *TfState, resourceId string) (*T, error) {
	resourceRef := fmt.Sprintf("%s.%s", GetResourceName(*new(T)), resourceId)
	return GetParsed[T](state, resourceRef)
}

func MaybeGetParsedById[T any](state *TfState, resourceId string) (*T, bool, error) {
	resourceRef := fmt.Sprintf("%s.%s", GetResourceName(*new(T)), resourceId)
	return MaybeGetParsed[T](state, resourceRef)
}
