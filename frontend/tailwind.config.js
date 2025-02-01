/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["./pages/**/*.templ", "./src/**/*.ts"],
  theme: {
    extend: {
      spacing: {},
      fontFamily: {
        montserrat: ["Montserrat", "sans-serif"],
        montserratSemibold: ["Montserrat-semibold", "sans-serif"],
        poppins: ["Poppins", "sans-serif"],
        tt_chocolate: ["TT Chocolates", "sans-serif"],
      },
      gridTemplateColumns: {
        profileButton: "1fr auto",
      },
      gridTemplateRows: {
        agentCollectPage: "auto 1fr auto",
        agentDashboardPage: "1fr minmax(2.563rem, auto)",
        loginPage: "1fr minmax(2.563rem, auto)",
        spendPage: "repeat(4, auto) 1fr auto",
        wastePage: "repeat(3, auto) 1fr",
      },
      backgroundImage: {},
    },
  },
  plugins: [],
};
