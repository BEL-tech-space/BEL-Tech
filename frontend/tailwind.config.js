/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["./pages/*.templ", "*.html"],
  theme: {
    extend: {
      pacing: {
        'minus-menu': 'calc(100vh - 5rem)',
      },
    },
  },
  plugins: [],
}

