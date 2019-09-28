/*
 * Brainrex API Explorer
 *
 * Welcome to the Brainrex API explorer, we make analytics tools for crypto and blockchain. Our currently propiertary models offer sentiment analysis, market making, blockchain monitoring and face-id verification. This AI models can be consumed from this API. We also offer integrations to open data and propietary data providers, as well as free test data we collect. There is a collection of data transformation tools. Join our Telegram group to get the latest news and ask questions [https://t.me/brainrex, #brainrex](https://t.me/brainrex). More about Brainrex at [https://brainrex.com](http://brainrex.com). Full Documentation can be found at [https://brainrexapi.github.io/docs](https://brainrexapi.github.io/docs)
 *
 * API version: 0.1.1
 * Contact: support@brainrex.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type Request1 struct {
	// Name of the exchange
	Exchange string `json:"exchange,omitempty"`
	// Name of the currency pair
	Market string `json:"market,omitempty"`
	// Name of the data format availables (standard)
	DataFormat string `json:"data_format,omitempty"`
	// Name of the market
	Orient string `json:"orient,omitempty"`
	// Start date in YYYY/MM/DD
	StartDate string `json:"start_date,omitempty"`
	// End date YYYY/MM/DD
	EndDate string `json:"end_date,omitempty"`
}
