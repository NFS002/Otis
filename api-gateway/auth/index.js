const jwt = require("express-jwt")
const jwksRsa = require("jwks-rsa")

const AUTH0_DOMAIN = process.env.OTIS_AUTH0_DOMAIN
const AUTH0_API_IDENTIFIER = process.env.OTIS_AUTH0_API_IDENTIFIER

const basicSecret = jwksRsa.expressJwtSecret({
	cache: true,
	rateLimit: true,
	jwksRequestsPerMinute: 5,
	jwksUri: `https://${AUTH0_DOMAIN}/.well-known/jwks.json`
})

const basicJwtCheck = jwt({
	secret: basicSecret,
	audience: AUTH0_API_IDENTIFIER,
	issuer: `https://${AUTH0_DOMAIN}/`,
	algorithms: ["RS256"]
})

module.exports = {
	basicSecret,
	basicJwtCheck
}
