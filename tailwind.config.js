/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["./templates/**/*.html"],
  theme: {
    extend: {
      backgroundImage:{
        "bgone":"url('/static/images/motph-test.png')",
        "mastercard":"url('/static/images/mastercard.png')"
      }
    },
  },
  plugins: [],
}

