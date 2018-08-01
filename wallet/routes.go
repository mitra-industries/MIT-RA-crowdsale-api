package wallet

import (
    "github.com/gin-gonic/gin"
    "mit-ra-crowdsale-api/common/rest"
    "mit-ra-crowdsale-api/common/ethereum"
    "mit-ra-crowdsale-api/models"
)

func InitRoutes(router *gin.Engine) {
    eth := router.Group("/wallet")
    {
        eth.GET("/balance/:address", getEthBalanceAction)
        eth.POST("/balances", getEthBalancesAction)
    }
}

// swagger:route GET /wallet/balance/:address ethereum getEthBalance
//
// Get balance
//
// Get particular Ethereum wallet balance.
//
// Consumes:
// - application/json
// Produces:
// - application/json
// Responses:
//   200: GetBalanceSuccessResponse
//   400: RestErrorResponse
//
func getEthBalanceAction(c *gin.Context) {
    addr := c.Param("address")
    bal, err := ethereum.GetWallet().BalanceAt(addr)
    if err != nil {
        rest.NewResponder(c).Error(err.Error())
        return
    }

    rest.NewResponder(c).Success(gin.H{
        "address": addr,
        "balance": bal.String(),
    })
}

// swagger:route POST /wallet/balances ethereum getEthBalances
//
// Get balances
//
// Get balances for list of Ethereum addresses.
//
// Consumes:
// - application/json
// Produces:
// - application/json
// Responses:
//   200: GetBalancesSuccessResponse
//   400: RestErrorResponse
//
func getEthBalancesAction(c *gin.Context) {
    request := &models.Addresses{}
    err := c.BindJSON(request)
    if err != nil {
        rest.NewResponder(c).ErrorValidation(err.Error())
        return
    }

    bals := map[string]string{}
    for _, addr := range request.Addresses {
        bal, err := ethereum.GetWallet().BalanceAt(addr)
        if err != nil {
            rest.NewResponder(c).Error(err.Error())
            return
        }
        bals[addr] = bal.String()
    }

    rest.NewResponder(c).Success(gin.H{
        "balances": bals,
    })
}
