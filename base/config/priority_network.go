package config

type priorityNetwork struct {
	Networks []string `yaml:"priority-network"`
}

var PriorityNetwork = &priorityNetwork{
	Networks: make([]string, 0),
}
