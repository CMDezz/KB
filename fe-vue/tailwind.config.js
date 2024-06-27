/** @type {import('tailwindcss').Config} */
export default {
  content: [],
  purge: ['./index.html', './src/**/*.{vue,js,ts,jsx,tsx}'],
  theme: {
    colors: {
      primary: '#fdcf41',
      text: '#222',
      neutral: '#f0f0f0',
      neutralText: '#9f9e9e',
      dark: '#9f9e9e',
      white: '#fff'
    },

    fontFamily: {
      sans: ['Arimo', 'sans-serif'],
      serif: ['Arimo', 'serif'],
    },
    extend: {},
  },
  plugins: [],
}

