package main

import (
    appCommon "mit-ra-crowdsale-api/common/application"
    "mit-ra-crowdsale-api/application"
    "github.com/sirupsen/logrus"
)

func init() {
    appCommon.Init()
    if err := application.Init(); err != nil {
        logrus.Fatal(err)
    }
}

func main() {
    appChannel := appCommon.GetNewChannel()

    go application.Run()

    appCommon.StartLoop(appChannel)
}
