package bot

import (
        "crypto/tls"
        "net/http"

        "github.com/gorilla/mux"
        "github.com/thoj/go-ircevent"
)

var _bot *irc.Connection
var _r *mux.Router

func init() {
        _bot = irc.IRC("nick", "user")
        _r = mux.NewRouter()
}

// irc-event functions
func UseTLS() {
        _bot.UseTLS = true
        _bot.TLSConfig = &tls.Config{InsecureSkipVerify: true}
}

func Connect(network string) error {
        return _bot.Connect(network)
}

func Join(channel string) {
        _bot.Join(channel)
}

func Nick(nick string) {
        _bot.Nick(nick)
}

func AddCallback(eventcode string, callback func(*irc.Event)) string {
        return _bot.AddCallback(eventcode, callback)
}

func Privmsg(target, format string, a ...interface{}) {
        _bot.Privmsg(target, format)
}

func Loop() {
        _bot.Loop()
}

// gorilla mux functions

func HandleFunc(path string, f func(http.ResponseWriter,
        *http.Request)) *mux.Route {
        return _r.HandleFunc(path, f)
}

func Vars(r *http.Request) map[string]string {
        return mux.Vars(r)
}

func Serve(port string) {
        go http.ListenAndServe(port, _r)
}
