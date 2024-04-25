package config

// Inspired by: https://github.com/edulinq/autograder-server/tree/main/config

var options = make(map[string]*option)

type option struct {
	Key         string
	Value       any
	Description string
}

type StringOption struct{ *option }
type BoolOption struct{ *option }
type IntOption struct{ *option }

func newOption(key string, value any, description string) *option {
	_, ok := options[key]
	if ok {
		return nil
	}

	opt := &option{
		Key:         key,
		Value:       value,
		Description: description,
	}

	options[key] = opt
	return opt
}

func NewStringOption(key string, value string, description string) *StringOption {
	return &StringOption{newOption(key, value, description)}
}

func NewBoolOption(key string, value bool, description string) *BoolOption {
	return &BoolOption{newOption(key, value, description)}
}

func NewIntOption(key string, value int, description string) *IntOption {
	return &IntOption{newOption(key, value, description)}
}

func (this *StringOption) Get() string {
	return options[this.Key].Value.(string)
}

func (this *BoolOption) Get() bool {
	return options[this.Key].Value.(bool)
}

func (this *IntOption) Get() int {
	return options[this.Key].Value.(int)
}

func (this *option) Set(value any) {
	options[this.Key].Value = value
}
