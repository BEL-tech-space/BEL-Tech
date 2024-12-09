// sw.js
import * as navigationPreload from 'workbox-navigation-preload';
import {CacheFirst, NetworkFirst} from 'workbox-strategies';
import {registerRoute} from 'workbox-routing';
import {matchPrecache, precacheAndRoute} from 'workbox-precaching';
import {ExpirationPlugin} from 'workbox-expiration';
import {strategy as composeStrategies} from 'workbox-streams';
import {headerPartialStrategy, footerPartialStrategy} from './sw-partials';



// Enable navigation preload
// navigationPreload.enable();

// RuntimeCaching
registerRoute(RegExp(/.*\.(?:png|jpg|jpeg|svg)$/), 
new CacheFirst(
  {
    cacheName: 'images-cache',
    plugins: [
      new ExpirationPlugin({
        maxEntries: 300,
        maxAgeSeconds: 60 * 60 * 24 * 30, // about a month
      }),
    ],
  }
))

registerRoute(RegExp(/.*partial.*/), 
new CacheFirst(
  {
    cacheName: 'partials-cache',
    plugins: [
      new ExpirationPlugin({
        maxEntries: 300,
        maxAgeSeconds: 60 * 60 * 24 * 30, // about a month
      }),
    ],
  }
))

declare const self: ServiceWorkerGlobalScope & { __WB_MANIFEST: any[] };
precacheAndRoute(self.__WB_MANIFEST || []);


// const accountDashboardContentStrategy = new NetworkFirst({
//   cacheName: 'contents-cache',
//   networkTimeoutSeconds : 5,
//   plugins: [
//     {
//       requestWillFetch: async ({request, event}) => {
//         const headers = new Headers();
//         headers.append('X-Content-Mode', 'partial');
//         return Promise.resolve(new Request(request.url, {method: 'GET', headers }));
//       },
//       handlerDidError: async ({request}) => {
//         return await matchPrecache('/offline.php');
//       }
//     }
//   ]
// });


// const accountDashboardNavigation = composeStrategies([
//   async ({event}) => headerPartialStrategy.handle({request: new Request('/partial-header.php'), event}),
//   async ({event}) => accountDashboardContentStrategy.handle(event),
//   async ({event}) => footerPartialStrategy.handle({request: new Request('/partial-header.php'), event}),
// ],
//  {}
// );


// registerRoute(({request}) => request.mode === 'navigate', accountDashboardNavigation);
