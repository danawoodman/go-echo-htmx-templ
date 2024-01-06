/** @type {import('tailwindcss').Config} */
module.exports = {
	content: ["./**/*.templ"],
	// Disable Tailwind colors:
	theme: { colors: {} },
	plugins: [require("daisyui")],
};
