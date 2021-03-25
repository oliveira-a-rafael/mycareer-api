package domains

var StrongFoots = map[string]string{
	"R": "Right",
	"L": "Left",
	"B": "Both",
}

func validateStrongFoot(key string) bool {
	_, ok := StrongFoots[key]
	return ok
}
