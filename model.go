package logs

type logRecord struct {
	Level   int
	Pattern string
	Args    []interface{}
}
