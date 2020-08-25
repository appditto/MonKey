module.exports = {
  purge: {
    enabled: process.env.NODE_ENV === "development" ? false : true,
    content: ["./src/**/*.html", "./src/**/*.svelte"],
  },
  theme: {
    extend: {
      colors: {
        primary: "#404040",
        danger: "#BF1323",
        dangerDark: "#97000E",
        brown: "#6c4725",
        brownLight: "#cd9e6c",
        gray: "#404040",
        grayLight: "#9b9ba1",
        offWhite: "#E4E5E9",
        orangeLight: "#FFCD98",
        purpleLight: "#C6B1FF",
        cyanLight: "#A6FFEA",
        white: "#ffffff",
      },
      spacing: {
        "11": "2.75rem",
        "36": "9rem",
      },
      borderRadius: {
        xl: "0.65rem",
      },
      maxWidth: {
        xxs: "16rem",
      },
      width: {
        "28": "7rem",
      },
      margin: {
        "-14": "-3.5rem",
        "1/8": "8%",
        "1/10": "10%",
        "1/20": "20%",
        "1/30": "30%",
        "1/35": "35%",
      },
      scale: {
        "103": "1.03",
        "200": "2",
        "300": "3",
      },
      inset: {
        "1/2": "50%",
      },
      borderWidth: {
        "3": "3px",
      },
      maxHeight: {
        xs: "20rem",
        sm: "24rem",
        md: "28rem",
        lg: "32rem",
      },
      transitionDuration: {
        "250": "250ms",
        "350": "350ms",
        "400": "400ms",
      },
    },
  },
  variants: {},
  plugins: [],
};
