self.addEventListener('fetch', (event) => {
    console.log(`SW intercepted: ${event}`);
});	

console.log('script was run.');