package token

import (
    "github.com/gin-gonic/gin"
    "mit-ra-crowdsale-api/common/rest"
    "mit-ra-crowdsale-api/models"
)

func InitRoutes(router *gin.Engine) {
    mit-ra := router.Group("/mit-ra")
    {
        mit-ra.POST("/deploy", rest.SignRequired(), postDeployTokenAction)
        mit-ra.GET("/balance/:address", getMit-raBalanceAction)
        mit-ra.POST("/balances", getMit-raBalancesAction)
    }
}

func postDeployTokenAction(c *gin.Context) {
    request := &models.TokenDeployParams{}
    err := c.BindJSON(request)
    if err != nil {
        rest.NewResponder(c).ErrorValidation(err.Error())
        return
    }

    addr, tx, err := GetToken().Deploy(request.TotalSupply)
    if err != nil {
        rest.NewResponder(c).Error(err.Error())
        return
    }

    rest.NewResponder(c).Success(gin.H{
        "address": addr.String(),
        "tx":      tx.Hash().String(),
    })
}

// swagger:route GET /mit-ra/balance/:address token getMit-raBalance
//
// Get balance
//
// Get MIT-RA token balance for particular Ethereum address.
//
// Consumes:
// - application/json
// Produces:
// - application/json
// Responses:
//   200: GetBalanceSuccessResponse
//   400: RestErrorResponse
//
func getMit-raBalanceAction(c *gin.Context) {
    addr := c.Param("address")
    bal, err := GetToken().Balance(addr)
    if err != nil {
        rest.NewResponder(c).Error(err.Error())
        return
    }

    rest.NewResponder(c).Success(gin.H{
        "address": addr,
        "balance": bal.String(),
    })
}

// swagger:route POST /mit-ra/balances token getMit-raBalances
//
// Get balances
//
// Get MIT-RA token balances for list of Ethereum addresses.
//
// Consumes:
// - application/json
// Produces:
// - application/json
// Responses:
//   200: GetBalancesSuccessResponse
//   400: RestErrorResponse
//
func getMit-raBalancesAction(c *gin.Context) {
    request := &models.Addresses{}
    err := c.BindJSON(request)
    if err != nil {
        rest.NewResponder(c).ErrorValidation(err.Error())
        return
    }

    bals := map[string]string{}
    for _, addr := range request.Addresses {
        bal, err := GetToken().Balance(addr)
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
