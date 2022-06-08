# Waterlog

A small collection of log adapters for [Watermill].

The `LoggerAdapter` affords Watermill applications to bring their own loggers
with their applications. This package provides lightly opinionated logger
implementations that satisfy the `LoggerAdapter` interface so your logs are
uniform across all components in your system.

All implementations will follow these conventions:

- logger structs are exported, named `Logger` and the underlying logger is
  exported as `Logger.L`
- new loggers use JSON formatting and use a constructor that matches
  `func(output io.Writer, debug, trace bool) watermill.LoggerAdapter`
- debug and trace are just boolean gates; they aren't separate loggers (unlike
  `watermill.StdLoggerAdapter`)

To Do

- [x] [go-kit/log]
- [x] [rs/zerolog]
- [ ] [sirupsen/logrus]
- [ ] [uber-go/zap]

[Watermill]: https://github.com/ThreeDotsLabs/watermill
[go-kit/log]: https://github.com/go-kit/log
[rs/zerolog]: https://github.com/rs/zerolog
[sirupsen/logrus]: https://github.com/Sirupsen/logrus
[uber-go/zap]: https://github.com/uber-go/zap
