# Overview  
A simple logging package for [Go](https://go.dev/).
- Outputs logs to clean, queryable JSON format.
- Performs quickly under medium loads (~100,000 logs). Messages stored by the Logger are encoded to JSON concurrently in goroutines once the user flushes them.
- Provides the following log levels, with ascending severity. User's may set the minimum log level threshold for a message to actually be logged when instantiating a Logger.
    1. Debug (0)
    2. Info (10)
    3. Warn (20)
    4. Error (30)
    5. Crit (40)

# Usage
To begin, instantiate a Logger.
```go
log := jlogger.NewLogger(logLevel, out)
```

To add a message to the Logger, use any of the following methods on your Logger.
```go
log.Debug("This is a Debug level message.")

log.Info("This is an Info level message.")

log.Warn("This is a Warn level message.")

log.Error("This is an Error level message.")

log.Crit("This is a Crit level message.")
```

Finally, all of a Logger's stored messages may be encoded to JSON and output via,
```go
log.FlushAll()
```

This produces the following output,
```json
{"prefix":"WARN","level":20,"message":"This is a Warn level message.","time":"2024-03-08 13:51:07.507257008 -0500 EST m=+0.001634967"}
{"prefix":"CRIT","level":40,"message":"This is a Crit level message.","time":"2024-03-08 13:51:07.507259452 -0500 EST m=+0.001637421"}
{"prefix":"ERROR","level":30,"message":"This is an Error level message.","time":"2024-03-08 13:51:07.50725828 -0500 EST m=+0.001636249"}
{"prefix":"DEBUG","level":0,"message":"This is a Debug level message.","time":"2024-03-08 13:51:07.507071973 -0500 EST m=+0.001449962"}
{"prefix":"INFO","level":10,"message":"This is an Info level message.","time":"2024-03-08 13:51:07.507254583 -0500 EST m=+0.001632552"}
```

# License
This project is freely availalble under the MIT license.

