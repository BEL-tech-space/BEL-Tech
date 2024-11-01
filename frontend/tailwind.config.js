/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["./pages/account/*.templ","./pages/agent/*.templ" , "*.html"],
  theme: {
    extend: {
      spacing: {
      },
       fontFamily: {
         montserrat: [ 'Montserrat', 'sans-serif' ],
         montserratSemibold: [ 'Montserrat-semibold', 'sans-serif' ],
         poppins: [ 'Poppins', 'sans-serif' ],
        },
      gridTemplateColumns: {
        'profileButton': '1fr auto',
      },
      gridTemplateRows: {
        'agentCollectPage': 'auto 1fr minmax(2.563rem, auto)',
        'agentDashboardPage': '1fr minmax(2.563rem, auto)',
        'loginPage': '1fr minmax(2.563rem, auto)',
        
      },
      backgroundImage: {
      },
    },
  },
  plugins: [],
}


