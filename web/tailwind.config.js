/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ['./src/**/*.{html,js,ts}'],
  theme: {
    colors: {
      'white': '#FFF',
      'black': '#000',
      'blue': 'blue',
      'neutral': '#1E3E62',
      'primary-600': 'rgb(230, 91, 0)',
      'primary-500': 'rgb(255, 116, 26)',
      'primary': 'rgb(255, 116, 26)',
      'primary-300': 'rgb(255, 178, 128)',
      'primary-200': 'rgb(255, 209, 179)',
      'secondary-800': 'rgb(30, 62, 98)',
      'secondary-700': 'rgb(42, 86, 137)',
      'secondary-300': 'rgb(157, 189, 225)',
      'dark-900': '#111827',
      'dark-800': '#1f2937',
      'dark-700': '#3741',
      'dark-600': '#4b5563',
    },
    extend: {},
  },
  plugins: [
  ],
  daisyui: {
  }
}

