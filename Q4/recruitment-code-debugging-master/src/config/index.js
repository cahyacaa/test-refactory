require("dotenv").config();

process.env.NODE_ENV = process.env.NODE_ENV || "development";
var config = {
  port: process.env.PORT,
  clientId: process.env.CLIENT_ID,
  clientSecret: process.env.CLIENT_SECRET,
  oauthUrl: process.env.OAUTH_URL,
  apiUrl: process.env.API_URL,
  redUrl: process.env.REDIRECT_URL,
}


if (config.error) {
  throw new Error("⚠️  Couldn't find .env file  ⚠️");
}

module.exports = {
  config
};