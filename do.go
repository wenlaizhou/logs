package logs

const (
	DEBUG = 0
	INFO  = 1
	ERROR = 2
)

func record(level int, pattern string, args ...interface{}) {

	logIdle <- logRecord{
		Args:    args,
		Level:   level,
		Pattern: pattern,
	}

}
