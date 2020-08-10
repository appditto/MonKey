module.exports = {
  purge: {
    enabled: process.env.NODE_ENV === 'development' ? false : true,
    content: ['./src/**/*.html', './src/**/*.svelte'],
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
        grayLight: "#9b9ba1"
      },
      spacing: {
        '36': '9rem'
      },
      borderRadius: {
        'xl': '0.65rem'
      },
      maxWidth: {
        'xxs': '16rem'
      },
      width: {
        '28': '7rem'
      },
      margin: {
        '-14': '-3.5rem'
      },
      scale: {
        '200': '2',
        '300': '3'
      },
      inset: {
        '1/2': '50%'
      },
      borderWidth: {
        '3': '3px'
      },
      maxHeight: {
        xs: '20rem',
        sm: '24rem',
        md: '28rem',
        lg: '32rem',
      }
    },
  },
  variants: {},
  plugins: [],
}