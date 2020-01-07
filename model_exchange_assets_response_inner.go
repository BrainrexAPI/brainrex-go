/*
 * BrainRex API
 *
 * The Brainrex API is a collection of analytics tools and data integrations made for blockchain developers. In particular we offer Natural Language Processing and Anomaly detection algorithms that have been fine tune to understand text data and time series in the domain speficic setting of cryptocurrency and blockchain technology. This technology is intended to be use as building blocks to bigger applications, we offer examples on how to use them for Trading Backtesting and Smart Contract anomaly monitoring.
 *
 * API version: 1.0.2
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package brainrex

type ExchangeAssetsResponseInner struct {
	// Highest price of the time frame with two decimal points
	Name string `json:"name,omitempty"`
	// Percetange change in the last 24 hours
	Id float32 `json:"id,omitempty"`
	// Volume of currency exchanged in the time frame with two decimal points
	TradinSym string `json:"tradinSym,omitempty"`
	// Volume of currency exchanged in the time frame with two decimal points
	Symbol string `json:"symbol,omitempty"`
}
