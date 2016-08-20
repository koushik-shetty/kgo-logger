# kgo-logger
golang logger for my golang projects

* Exports an interface [`Loggable`](https://github.com/koushik-shetty/kgologger/blob/5705f9134e03b60d188d3b3f666ec0a458524635/logger.go#L17) for logging.
* Has the following methods InfoF, ErrorF, Error, FatalF, PrintF.
* Exports a type [`Logger`](https://github.com/koushik-shetty/kgologger/blob/5705f9134e03b60d188d3b3f666ec0a458524635/logger.go#L26) which does actual logging. Implements `Loggable`. Wrapper around [sirupsen/logrus](https://github.com/Sirupsen/logrus).
* Exports a [`BlankLogger`](https://github.com/koushik-shetty/kgologger/blob/5705f9134e03b60d188d3b3f666ec0a458524635/logger.go#L115) which is used when logging is not desired. Can be used in testing also.

#### __MIT License__