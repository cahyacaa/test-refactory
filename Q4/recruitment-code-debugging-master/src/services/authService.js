const {
  config
} = require("../config");

function redirectUri() {
  return `${config.oauthUrl}/authorize?client_id=${config.clientId}&redirect_uri=http://localhost:8080/${config.redUrl}`;
}

module.exports = {
  redirectUri:redirectUri
}