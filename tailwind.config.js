/** @type {import('tailwindcss').Config} */
const theme = require("tailwindcss/defaultTheme");
module.exports = {
    content: ["./ui/**/*.{html,js}"],
    theme: {
        container: {
            // you can configure the container to be centered
            center: true,

            // or have default horizontal padding
            padding: {
                DEFAULT: '1rem',
                sm: '2rem',
                lg: '4rem',
                xl: '5rem',
                '2xl': '6rem',
            },
            // default breakpoints but with 40px removed
            screens: {
                sm: '600px',
                md: '728px',
                lg: '984px',
                xl: '1240px',
                '2xl': '1496px',
                '3xl': '1700px',
            },
        },
    },
    plugins: [
        require('@tailwindcss/typography'),
        require("@tailwindcss/forms"),
        require("daisyui")
    ],
    daisyui: {
        themes: ["winter"],
    },
}
