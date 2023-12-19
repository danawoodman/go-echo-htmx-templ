/** @type {import('tailwindcss').Config} */
module.exports = {
	content: ["./views/html/**/*.templ"],
	theme: {
		colors: {},
		// extend: {},
	},
	plugins: [require("daisyui")],
};
