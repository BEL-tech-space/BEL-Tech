/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["./pages/*.templ", "*.html"],
  theme: {
    extend: {
      spacing: {
        'minus-menu': 'calc(100vh - 8rem)',
      },
    },
  },
  plugins: [],
}

