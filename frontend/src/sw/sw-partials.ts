import {matchPrecache, precacheAndRoute} from 'workbox-precaching';
import {CacheFirst} from 'workbox-strategies';

export const headerPartialStrategy = new CacheFirst({
  cacheName: 'partials-cache',
  plugins: [
    {
      handlerDidError: async ({request}) => {
        return await matchPrecache('/offline-header.php');
      }
    }
  ]
});

export const footerPartialStrategy = new CacheFirst({
  cacheName: 'partials-cache',
  plugins: [
    {
      handlerDidError: async ({request}) => {
        return await matchPrecache('/offline-footer.php');
      }
    }
  ]
});