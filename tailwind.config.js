/** @type {import('tailwindcss').Config} */
module.exports = {
    content: ["./ui/**/*.{html,js}"],
    theme: {
        extend: {},
    },
    plugins: [require("daisyui")],
}
