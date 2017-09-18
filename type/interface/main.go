package main

type S struct {
	Name  string
	Money int
}

func main() {
	st := []S{
		{"kevin.gao", 1000},
		{"aqua", 1000},
	}

	i := make([]interface{}, len(st))
	for k, v := range st {
		i[k] = v
	}
}
