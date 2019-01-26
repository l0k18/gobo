package args

// ItemS is the node in the parsing graph. Groups means item names with which this item causes a conflict with
type ItemS struct {
	Name    string
	Short   string
	Handler func([]interface{}) error
	Groups  []string
}

// Item is a short way of declaring an item
func Item(
	name string,
	short string,
	handler func([]interface{}) error,
	groups ...string,
) ItemS {
	return ItemS{name, short, handler, groups}
}

// CmdS is a mandatory option that performs some action. The parser will flag an error if it finds no commands in the parameters
type CmdS struct {
	ItemS
	Parameters []interface{}
}

// Cmd is a short way of declaring an Cmd item
func Cmd(
	name string,
	short string,
	handler func([]interface{}) error,
	groups []string,
	params ...interface{},
) CmdS {
	return CmdS{ItemS{name, short, handler, groups}, params}
}

// TriggerS causes an action to be performed that may be terminal, but all triggers will be processed unless they are in the same group
type TriggerS struct {
	ItemS
}

// Trigger is a short way of declaring a trigger
func Trigger(
	name string,
	short string,
	handler func([]interface{}) error,
	groups []string,
) TriggerS {
	return TriggerS{
		ItemS{name, short, handler, groups},
	}
}

// VarS is the combination of a label and a value, it is singular and it is an error to put a list into a Var
type VarS struct {
	ItemS
	Value interface{}
}

// Var is a short way of declaring an Var item
func Var(
	item ItemS,
	value interface{},
) VarS {
	return VarS{item, value}
}

// ListS is a label with a collection of items following it, terminated by anything not in its list. This is distinct from the previous to simplify the parsing
type ListS struct {
	ItemS
	Values []interface{}
}

// List is a short way of declaring an List item
func List(
	item ItemS,
	values ...interface{},
) VarS {
	return VarS{item, values}
}
