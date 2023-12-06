package logging

type AppLogger interface {
	LogInfo(message string)
	LogError(message string)
}
