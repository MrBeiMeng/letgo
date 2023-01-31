package logger

type LoggerI interface {
	Warn(msg string, asg ...interface{})
	Break(msg string, asg ...interface{})
	Success(msg string, asg ...interface{})
	Info(msg string, asg ...interface{})
}

var Logger LoggerI = &LoggerImpl{}
