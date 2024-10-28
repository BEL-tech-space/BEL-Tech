/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["./pages/account/*.templ","./pages/agent/*.templ" , "*.html"],
  theme: {
    extend: {
      spacing: {
      },
       fontFamily: {
          montserrat: [ 'Montserrat', 'sans-serif' ],
          poppins: [ 'Poppins', 'sans-serif' ],
        },
      gridTemplateColumns: {
      },
      gridTemplateRows: {
        'agentCollectPage': 'auto 1fr minmax(2.563rem, auto)',
        'agentDashboardPage': '1fr minmax(2.563rem, auto)',
      },
      backgroundImage: {
      },
    },
  },
  plugins: [],
}


