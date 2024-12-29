import {
  CacheFirst,
  NetworkFirst,
  StaleWhileRevalidate,
} from "workbox-strategies";
import { ExpirationPlugin } from "workbox-expiration";
import { strategy as composeStrategies } from "workbox-streams";

export const contentCache = "contents-cache";
export const partialsCache = "partials-cache";
export const imagesCache = "images-cache";

export const cacheFirstStrategy = new CacheFirst({
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

export const networkFirstStrategy = new NetworkFirst({
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
export const staleWhileRevalidateStrategy = new StaleWhileRevalidate({
  cacheName: contentCache,
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

export const dashboardPageStream = composeStrategies(
  [
    async ({ event }) =>
      cacheFirstStrategy.handle({
        request: new Request("/account/dashboard/openingHtmlTag-Partial.html"),
        event,
      }),
    async ({ event }) =>
      cacheFirstStrategy.handle({
        request: new Request("/account/dashboard/header-Partial.html"),
        event,
      }),
    async ({ event }) =>
      cacheFirstStrategy.handle({
        request: new Request("/account/dashboard/mainOpeningTag-Partial.html"),
        event,
      }),
    async ({ event }) =>
      staleWhileRevalidateStrategy.handle({
        request: new Request("/account/dashboard/swiperOne"),
        event,
      }),
    async ({ event }) =>
      staleWhileRevalidateStrategy.handle({
        request: new Request("/account/dashboard/echarts?month=0"),
        event,
      }),
    async ({ event }) =>
      staleWhileRevalidateStrategy.handle({
        request: new Request("/account/dashboard/swiperTwo"),
        event,
      }),
    // async ({ event }) =>
    //   networkFirstStrategy.handle({
    //     request: new Request("/account/dashboard", {
    //       headers: [["X-Content-Mode", "partials"]],
    //     }),
    //     event,
    //   }),
    async ({ event }) =>
      cacheFirstStrategy.handle({
        request: new Request("/account/dashboard/footer-Partial.html"),
        event,
      }),
    async ({ event }) =>
      cacheFirstStrategy.handle({
        request: new Request("/account/dashboard/closingHtmlTag-Partial.html"),
        event,
      }),
  ],
  {},
);

export const locationPageStream = composeStrategies(
  [
    async ({ event }) =>
      cacheFirstStrategy.handle({
        request: new Request("/account/location/openingHtmlTag-Partial.html"),
        event,
      }),
    async ({ event }) =>
      cacheFirstStrategy.handle({
        request: new Request("/account/location/header-Partial.html"),
        event,
      }),
    async ({ event }) =>
      networkFirstStrategy.handle({
        request: new Request("/account/location/main", {
          headers: [["X-Content-Mode", "partials"]],
        }),
        event,
      }),
    async ({ event }) =>
      cacheFirstStrategy.handle({
        request: new Request("/account/location/footer-Partial.html"),
        event,
      }),
    async ({ event }) =>
      cacheFirstStrategy.handle({
        request: new Request("/account/location/closingHtmlTag-Partial.html"),
        event,
      }),
  ],
  {},
);

export const walletPageStream = composeStrategies(
  [
    async ({ event }) =>
      cacheFirstStrategy.handle({
        request: new Request("/account/wallet/openingHtmlTag-Partial.html"),
        event,
      }),
    async ({ event }) =>
      cacheFirstStrategy.handle({
        request: new Request("/account/wallet/header-Partial.html"),
        event,
      }),
    async ({ event }) =>
      cacheFirstStrategy.handle({
        request: new Request("/account/wallet/mainOpeningTag-Partial.html"),
        event,
      }),
    async ({ event }) =>
      staleWhileRevalidateStrategy.handle({
        request: new Request("/account/wallet/balance"),
        event,
      }),
    async ({ event }) =>
      cacheFirstStrategy.handle({
        request: new Request("/account/wallet/options"),
        event,
      }),
    async ({ event }) =>
      networkFirstStrategy.handle({
        request: new Request("/account/wallet/coinData", {}),
        event,
      }),
    // async ({ event }) =>
    //   networkFirstStrategy.handle({
    //     request: new Request("/account/wallet", {
    //       headers: [["X-Content-Mode", "partials"]],
    //     }),
    //     event,
    //   }),
    async ({ event }) =>
      staleWhileRevalidateStrategy.handle({
        request: new Request("/account/wallet/swiper"),
        event,
      }),
    async ({ event }) =>
      cacheFirstStrategy.handle({
        request: new Request("/account/wallet/footer-Partial.html"),
        event,
      }),
    async ({ event }) =>
      cacheFirstStrategy.handle({
        request: new Request("/account/wallet/closingHtmlTag-Partial.html"),
        event,
      }),
  ],
  {},
);

export const spendPageStream = composeStrategies(
  [
    async ({ event }) =>
      cacheFirstStrategy.handle({
        request: new Request("/account/spend/openingHtmlTag-Partial.html"),
        event,
      }),
    async ({ event }) =>
      cacheFirstStrategy.handle({
        request: new Request("/account/spend/header-Partial.html"),
        event,
      }),
    async ({ event }) =>
      cacheFirstStrategy.handle({
        request: new Request("/account/spend/mainOpeningTag-Partial.html"),
        event,
      }),
    async ({ event }) =>
      networkFirstStrategy.handle({
        request: new Request("/account/spend/balance"),
        event,
      }),
    async ({ event }) =>
      cacheFirstStrategy.handle({
        request: new Request("/account/spend/otherhalf"),
        event,
      }),
    // async ({ event }) =>
    //   networkFirstStrategy.handle({
    //     request: new Request("/account/spend", {
    //       headers: [["X-Content-Mode", "partials"]],
    //     }),
    //     event,
    //   }),
    async ({ event }) =>
      cacheFirstStrategy.handle({
        request: new Request("/account/spend/footer-Partial.html"),
        event,
      }),
    async ({ event }) =>
      cacheFirstStrategy.handle({
        request: new Request("/account/spend/closingHtmlTag-Partial.html"),
        event,
      }),
  ],
  {},
);
