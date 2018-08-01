// Package classification MIT-RA Crowdsale Ethereum API.
//
// Simple HTTP API to the Ethereum blockchain for MIT-RA Crowdsale.
//
//     Schemes: http, https
//     BasePath: /
//     Version: 0.3.12
//     License: MIT http://opensource.org/licenses/MIT
//     Contact: Dan Gartman<drjohn.dobbin@gmail.com> https://github.com/mit-ra-industries
//
//     Consumes:
//     - application/json
//
//     Produces:
//     - application/json
//
//     Security:
//     - hash
//     - jwt
//
//     SecurityDefinitions:
//     hash:
//          type: apiKey
//          name: X-Authorization
//          in: header
//
// swagger:meta
package application

import (
    httpCommon "mit-ra-crowdsale-api/common/rest"
    "github.com/gin-gonic/gin"
    "sync"
    "mit-ra-crowdsale-api/common/rest"
    "mit-ra-crowdsale-api/wallet"
    "mit-ra-crowdsale-api/token"
    "mit-ra-crowdsale-api/crowdsale"
)

var o sync.Once

func Run() error {
    return httpCommon.Run(initRoutes)
}

func initRoutes(router *gin.Engine) {
    o.Do(func() {
        router.Use(rest.ExecutionTime())

        wallet.InitRoutes(router)
        token.InitRoutes(router)
        crowdsale.InitRoutes(router)
    })
}
