import {generateSW} from 'workbox-build';
import {} from 'workbox-routing';

export async function buildSW() {

  const {count, size, warnings}  = await generateSW({
    globDirectory: 'public',
    globPatterns: ['**/*.{js,css,svg,img,ttf}'],
    swDest: 'public/sw.js',
    maximumFileSizeToCacheInBytes: 2000000000,
    disableDevLogs: false,
    runtimeCaching: [{
      urlPattern: ({url}) => {
        console.log(url);
        return /\.(?:png|jpg|jpeg|svg|gif|webp)$/.test(url.pathname);
      },
      handler: 'CacheFirst',
      options: {
        cacheName: 'images',
        expiration: {
          maxEntries: 10,
          maxAgeSeconds: 60 * 60 * 24 * 30, // about a month
        },
      },
    }, {
  urlPattern: ({request}) => {
    console.log(request, request.destination);
    return request.destination === 'image';
  },
  handler: 'CacheFirst',
  options: {
    cacheName: 'image-cache',
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


