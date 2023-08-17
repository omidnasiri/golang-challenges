package keyboard

import "strings"

type Keyboard struct {
	Durability int
	History    map[string]int
}

func NewKeyboard(dure int) *Keyboard {
	return &Keyboard{
		Durability: dure,
		History:    make(map[string]int),
	}
}

func (keyboard *Keyboard) Enter(inp string) string {
	var res string
	for _, v := range inp {
		s := string(v)
		lower := strings.ToLower(s)
		if s != lower {
			i, ok := keyboard.History["shift"]
			if !ok || i < keyboard.Durability {
				keyboard.History["shift"]++
			} else {
				s = lower
			}
		}
		i, ok := keyboard.History[lower]
		if !ok || i < keyboard.Durability {
			keyboard.History[lower]++
			res += s
		} else {
			continue
		}
	}
	return res
}
