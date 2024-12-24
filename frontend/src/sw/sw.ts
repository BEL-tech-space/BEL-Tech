// sw.js
import * as navigationPreload from "workbox-navigation-preload";
import { CacheFirst, NetworkFirst } from "workbox-strategies";
import { registerRoute } from "workbox-routing";
import { matchPrecache, precacheAndRoute } from "workbox-precaching";
import { ExpirationPlugin } from "workbox-expiration";
import { strategy as composeStrategies } from "workbox-streams";
import { clientsClaim } from "workbox-core";

const contentCache = "contents-cache";
const partialsCache = "partials-cache";
const imagesCache = "images-cache";

declare const self: ServiceWorkerGlobalScope & { __WB_MANIFEST: any[] };

clientsClaim();

// Enable navigation preload
navigationPreload.enable();

// RuntimeCaching
registerRoute(
  RegExp(/.*\.(?:png|jpg|jpeg|svg|webp)$/),
  new CacheFirst({
    cacheName: imagesCache,
    plugins: [
      new ExpirationPlugin({
        maxEntries: 300,
        maxAgeSeconds: 60 * 60 * 24 * 30, // about a month
      }),
    ],
  })
);

const cacheFirstPartialStrategy = new CacheFirst({
  cacheName: partialsCache,
  plugins: [
    new ExpirationPlugin({
      maxEntries: 30,
      maxAgeSeconds: 60 * 60 * 24 * 30, // about a month
    }),
    // {
    //   handlerDidError: async ({request}) => {
    //     return await matchPrecache('/offline-header.php');
    //   }
    // }
  ],
});

const networkFirstContentStrategy = new NetworkFirst({
  cacheName: contentCache,
  networkTimeoutSeconds: 5,
  // plugins: [
  //   {
  //     requestWillFetch: async ({request, event}) => {
  //       const headers = new Headers();
  //       headers.append('X-Content-Mode', 'partial');
  //       console.log('requestWillFetch', request.url);
  //       return Promise.resolve(new Request(request.url, {method: 'GET', headers }));
  //     },
  //     handlerDidError: async ({request}) => {
  //       return await matchPrecache('/offline.html');
  //     }
  //   }
  // ]
});

const dashboardPageStream = composeStrategies(
  [
    async ({ event }) =>
      cacheFirstPartialStrategy.handle({
        request: new Request("/account/dashboard/openingHtmlTag-Partial.html"),
        event,
      }),
    async ({ event }) =>
      cacheFirstPartialStrategy.handle({
        request: new Request("/account/dashboard/header-Partial.html"),
        event,
      }),
    async ({ event }) =>
      networkFirstContentStrategy.handle({
        request: new Request("/account/dashboard", {
          headers: [["X-Content-Mode", "partials"]],
        }),
        event,
      }),
    async ({ event }) =>
      cacheFirstPartialStrategy.handle({
        request: new Request("/account/dashboard/footer-Partial.html"),
        event,
      }),
    async ({ event }) =>
      cacheFirstPartialStrategy.handle({
        request: new Request("/account/dashboard/closingHtmlTag-Partial.html"),
        event,
      }),
  ],
  {}
);

// account/dashboard
registerRoute(
  ({ url }) => url.pathname == "/account/dashboard",
  dashboardPageStream
);

// account/scan
registerRoute(
  ({ url }) => url.pathname == "/account/scan",
  new CacheFirst({
    cacheName: contentCache,
    plugins: [
      new ExpirationPlugin({
        maxEntries: 30,
        maxAgeSeconds: 60 * 60 * 24 * 30, // about a month
      }),
    ],
  })
);

const locationPageStream = composeStrategies(
  [
    async ({ event }) =>
      cacheFirstPartialStrategy.handle({
        request: new Request("/account/location/openingHtmlTag-Partial.html"),
        event,
      }),
    async ({ event }) =>
      cacheFirstPartialStrategy.handle({
        request: new Request("/account/location/header-Partial.html"),
        event,
      }),
    async ({ event }) =>
      networkFirstContentStrategy.handle({
        request: new Request("/account/location", {
          headers: [["X-Content-Mode", "partials"]],
        }),
        event,
      }),
    async ({ event }) =>
      cacheFirstPartialStrategy.handle({
        request: new Request("/account/location/footer-Partial.html"),
        event,
      }),
    async ({ event }) =>
      cacheFirstPartialStrategy.handle({
        request: new Request("/account/location/closingHtmlTag-Partial.html"),
        event,
      }),
  ],
  {}
);

// account/location
registerRoute(
  ({ url }) => url.pathname == "/account/location",
  locationPageStream
);

const walletPageStream = composeStrategies(
  [
    async ({ event }) =>
      cacheFirstPartialStrategy.handle({
        request: new Request("/account/wallet/openingHtmlTag-Partial.html"),
        event,
      }),
    async ({ event }) =>
      cacheFirstPartialStrategy.handle({
        request: new Request("/account/wallet/header-Partial.html"),
        event,
      }),
    async ({ event }) =>
      networkFirstContentStrategy.handle({
        request: new Request("/account/wallet", {
          headers: [["X-Content-Mode", "partials"]],
        }),
        event,
      }),
    async ({ event }) =>
      cacheFirstPartialStrategy.handle({
        request: new Request("/account/wallet/footer-Partial.html"),
        event,
      }),
    async ({ event }) =>
      cacheFirstPartialStrategy.handle({
        request: new Request("/account/wallet/closingHtmlTag-Partial.html"),
        event,
      }),
  ],
  {}
);

// account/wallet
registerRoute(({ url }) => url.pathname == "/account/wallet", walletPageStream);

const spendPageStream = composeStrategies(
  [
    async ({ event }) =>
      cacheFirstPartialStrategy.handle({
        request: new Request("/account/spend/openingHtmlTag-Partial.html"),
        event,
      }),
    async ({ event }) =>
      cacheFirstPartialStrategy.handle({
        request: new Request("/account/spend/header-Partial.html"),
        event,
      }),
    async ({ event }) =>
      networkFirstContentStrategy.handle({
        request: new Request("/account/spend", {
          headers: [["X-Content-Mode", "partials"]],
        }),
        event,
      }),
    async ({ event }) =>
      cacheFirstPartialStrategy.handle({
        request: new Request("/account/spend/footer-Partial.html"),
        event,
      }),
    async ({ event }) =>
      cacheFirstPartialStrategy.handle({
        request: new Request("/account/spend/closingHtmlTag-Partial.html"),
        event,
      }),
  ],
  {}
);

// account/wallet
registerRoute(({ url }) => url.pathname == "/account/spend", spendPageStream);

precacheAndRoute(self.__WB_MANIFEST || []);
