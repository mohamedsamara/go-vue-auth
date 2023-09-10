/** @type {import('tailwindcss').Config} */
export default {
  content: ['./index.html', './src/**/*.{vue,js,ts,jsx,tsx}'],
  theme: {
    extend: {
      fontFamily: {
        poppins: 'Poppins',
        archivo: 'Archivo'
      },
      colors: {
        primary: '#4863f4',
        secondary: '#000000',
        'white-lilacapprox': '#f2f2fA'
      }
    }
  },
  plugins: []
}
