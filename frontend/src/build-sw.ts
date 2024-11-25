import {generateSW} from 'workbox-build';

export async function buildSW() {

  const {count, size, warnings}  = await generateSW({
    globDirectory: 'public/assets',
    globPatterns: ['**/*.{js,css}'],
    swDest: 'public/swThree.js',
    maximumFileSizeToCacheInBytes: 2000000000,
    runtimeCaching: [{
      urlPattern: new RegExp('/\.(?:png|jpg|jpeg|svg)$/,'),
      handler: 'CacheFirst',
      options: {
        cacheName: 'moreAssets',
        expiration: {
          maxEntries: 10,
          maxAgeSeconds: 60 * 60 * 24 * 30, // about a month
        },
      },
    }, {
      urlPattern: new RegExp('/\.(?:xml|html)$/,'),
      handler: 'StaleWhileRevalidate',
      options: {
        cacheName: 'html-cache',
        expiration: {
          maxEntries: 10,
          maxAgeSeconds: 60 * 60 * 24 * 30, // about a month
        },
      },
    }],
  })

  console.log(`Generated a service worker, which will precache ${count} files, totaling ${size} bytes and \n warning: \n ${warnings}.`);

}

buildSW();


