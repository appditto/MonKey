const tailwindcss = require("tailwindcss");

const production = process.env.NODE_ENV === "production";

module.exports = {
  plugins: [
    tailwindcss("./tailwind.config.js"),
    production &&
      require("cssnano")({
        preset: "default",
      }),
  ],
};
