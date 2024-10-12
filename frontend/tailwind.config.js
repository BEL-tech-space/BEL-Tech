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
        customDarkBlue: '#0c001b',
        'shadow-dark': '#ffffff0c',
        'secondary': '#010109',
        'service-icon-bg-color': '#00000033',
        'underColor': 'rgba(255, 255, 255, 0.9)',
        'backgroundColor': 'rgba(1, 1, 9, 1)',
        'shadeWhite': 'rgba(255, 255, 255, 0.5)',
        'modalWhite': 'rgba(255, 255, 255, 0.6)',
        'modalUndercolor': 'rgba(255, 255, 255, 0.05)'


      },
      backgroundImage: {
        'custom-gradient': 'linear-gradient(to right, rgba(153, 69, 255, 0.2), rgba(18, 219, 136, 0.2))',
        'tokenCreated': "url('/assets/svg/create-new-folder.svg')",
        'footer-texture': "url('/img/footer-texture.png')",
      },
    },
  },
  plugins: [],
}


