package main

import (
        "log"
        "strconv"

        "./bot"
        _ "./modules"
        "github.com/jimlawless/cfg"
)

func main() {
        config := make(map[string]string)
        err := cfg.Load("config.cfg", config)
        if err != nil {
                log.Fatal(err)
        }

        var ssl bool
        ssl, err = strconv.ParseBool(config["ssl"])
        if ssl == true {
                bot.UseTLS()
        }

        err = bot.Connect(config["network"])
        if err != nil {
                log.Fatal(err)
                return
        }
        bot.Nick(config["nick"])
        bot.Join(config["channel"])
        bot.Serve(config["server-port"])
        bot.Loop()
}
