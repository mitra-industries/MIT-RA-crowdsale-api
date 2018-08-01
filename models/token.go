package models

type TokenDeployParams struct {
    TotalSupply  string `json:"totalSupply"`
}

// swagger:parameters getMit-raBalance
type GetMit-raBalanceParams struct {
    // Ethereum address
    // example: 0xFdb3Ae550c4f6a8FD170C3019c97D4F152b65373
    // in: query
    Address string `json:"address"`
}

// swagger:parameters getMit-raBalances
type Mit-raAddressesParams struct {
    // in: body
    Body Addresses `json:"body"`
}
