import {generateSW} from 'workbox-build';
async function buildSW() {
  const {count, size, warnings}  = await generateSW({
    globDirectory: 'public',
    globPatterns: ['**/*.{js,css,ttf}'],
    maximumFileSizeToCacheInBytes: 2000000000,
    skipWaiting: true,
    disableDevLogs: false,
    cacheId: "local",
    cleanupOutdatedCaches: true,
    clientsClaim: true,
    mode: "development",
    runtimeCaching: [{
      urlPattern: '/\.(?:png|jpg|jpeg|svg)$/',
      handler: 'CacheFirst',
      options: {
        cacheName: 'image-cache',
        expiration: {
          maxEntries: 10,
          maxAgeSeconds: 60 * 60 * 24 * 30, // about a month
        },
      },
    }],
    swDest: 'public/sw.js',
  })
  
  console.log(`Generated a service worker, which will precache ${count} files, totaling ${size} bytes and \n warning: \n ${warnings}.`);  
}


buildSW();


