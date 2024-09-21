package utils

type Bitwise int

func (set Bitwise) Has(flag Bitwise) bool {
    return set&flag == flag
}
func (set Bitwise) Remove(flag Bitwise) Bitwise {
    return set &^ flag
}

func (set Bitwise) Add(flag Bitwise) Bitwise {
    return set | flag
}
