module.exports = {
	globDirectory: 'public/',
	globPatterns: [
		'**/*.{ttf,png,jpg,jpeg,js,css,svg}'
	],
	swDest: 'public/sw.js',
	ignoreURLParametersMatching: [
		/^utm_/,
		/^fbclid$/
	]
};