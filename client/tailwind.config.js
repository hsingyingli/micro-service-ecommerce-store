const defaultTheme = require('tailwindcss/defaultTheme')
/** @type {import('tailwindcss').Config} */
module.exports = {
  content: [
    "./app/**/*.{js,ts,jsx,tsx}",
    "./pages/**/*.{js,ts,jsx,tsx}",
    "./components/**/*.{js,ts,jsx,tsx}",
 
    // Or if using `src` directory:
    "./src/**/*.{js,ts,jsx,tsx}",
  ],
  theme: {
    extend: {
      fontFamily: {
        'sans': ['"M PLUS Rounded 1c"', ...defaultTheme.fontFamily.sans]
      },
      colors: {
        'secondary': {
          50: '#e9e8e8',
          100: '#d6d4d4',
          200: '#c3c0c0',
          300: '#b0acab',
          400: '#9d9897',
          500: '#8a8483',
          600: '#6c6766',
          700: '#625d5c',
          800: '#585453',
          900: '#4e4a49',
          1000: '#444140'
        },
        'primary': {
          50: '#fcfbfa',
          100: '#f3f2ef',
          200: '#ebe9e4',
          300: '#e3e1d8',
          400: '#dad8cd',
          500: '#d2cfc2',
          600: '#cac6b7',
          700: '#c1bdab',
          800: '#b9b4a0',
          900: '#b1ab95',
          1000: '#a8a289'
        }
      }
    },
  },
  plugins: [],
}
