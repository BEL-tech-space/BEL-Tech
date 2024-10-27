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
       colors: {
      },
      gridTemplateRows: {
        'agentCollectPage': 'auto 1fr auto',
      },
      backgroundImage: {
      },
    },
  },
  plugins: [],
}


