package cfgParser

import (
    "log"
    "os"
    "github.com/pelletier/go-toml/v2"

    "srv/util"
)

type cfgStruct struct {
    Port string
    Template string
}

func Parse() (cfgStruct){
    tomlData, err := os.ReadFile("config.toml")
    util.CheckErr(err)
    log.Println("Cfg file reading complete.")

    var cfgFields cfgStruct
    err = toml.Unmarshal(tomlData, &cfgFields)
    util.CheckErr(err)
    log.Println("Cfg file unmarshalling complete.")
    
    log.Println("Port:", cfgFields.Port)
    log.Println("HTML Template File:", cfgFields.Template)

    return cfgFields
}
