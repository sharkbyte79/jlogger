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

# License
This project is freely available under the MIT license.

