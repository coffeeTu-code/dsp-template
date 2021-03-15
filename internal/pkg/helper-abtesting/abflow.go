package abtesting

type AbFlow struct {
	AllAbflags map[string]string
	UsedAbflag map[string]string
}

func (flow *AbFlow) GetAbflag(key string) (string, bool) {
	value, fined := flow.AllAbflags[key]
	if !fined {
		return "", false
	}
	flow.UsedAbflag[key] = value
	return value, true
}
