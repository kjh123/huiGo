package log

import (
    "fmt"
    "github.com/RichardKnop/logging"
    "github.com/urfave/negroni"
    "log"
    "net/http"
    "os"
    "time"
)

var (
    logger = logging.New(nil, nil, new(logging.ColouredFormatter))

    // INFO ...
    INFO = logger[logging.INFO]
    // WARNING ...
    WARNING = logger[logging.WARNING]
    // ERROR ...
    ERROR = logger[logging.ERROR]
    // FATAL ...
    FATAL = logger[logging.FATAL]
)

type Logger struct {
    *log.Logger
}

// return new logger instance
func NewLogger() *Logger {
    return &Logger{log.New(os.Stdout, "[huiGo] ", 0)}
}

func (l *Logger) ServeHTTP(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
    start := time.Now()
    ip := r.RemoteAddr
    if xff := r.Header.Get("X-Forwarded-For"); xff != "" {
        ip = xff
    }

    INFO.Printf("Started %s %s for %s", r.Method, r.URL.Path, ip)

    next(rw, r)

    res := rw.(negroni.ResponseWriter)

    msg := fmt.Sprintf("Finished %s %s : %v %s in %v", r.Method, r.URL.Path, res.Status(), http.StatusText(res.Status()), time.Since(start))

    switch {
    case res.Status() < 400:
        INFO.Print(msg)
    case res.Status() < 500:
        WARNING.Print(msg)
    default:
        ERROR.Print(msg)
    }
}

// Set sets a custom logger
func Set(l logging.LoggerInterface) {
    INFO = l
    WARNING = l
    ERROR = l
    FATAL = l
}
