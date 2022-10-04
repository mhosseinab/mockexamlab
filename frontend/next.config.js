/** @type {import('next').NextConfig} */
const nextConfig = {
  reactStrictMode: true,
  swcMinify: true,
};

const path = require("path");
const withSass = require("@zeit/next-sass");
module.exports = withSass({
  /* bydefault config  option Read For More Optios
  here https://github.com/vercel/next-plugins/tree/master/packages/next-sass
  */
  cssModules: true,
});
module.exports = {
  env: {
    API_KEY: "",
    APP_ID: "",
    AUTH_DOMAIN: "",
  },
  /* Add Your Scss File Folder Path Here */
  sassOptions: {
    includePaths: [path.join(__dirname, "styles")],
  },
};

module.exports = nextConfig;
