// sw.js
import * as navigationPreload from "workbox-navigation-preload";
import {
  CacheFirst,
  NetworkFirst,
  StaleWhileRevalidate,
} from "workbox-strategies";
import { registerRoute } from "workbox-routing";
import { precacheAndRoute } from "workbox-precaching";
import { ExpirationPlugin } from "workbox-expiration";
import { clientsClaim } from "workbox-core";
import {
  imagesCache,
  contentCache,
  partialsCache,
  dashboardPageStream,
  locationPageStream,
  walletPageStream,
  spendPageStream,
  staleWhileRevalidateStrategy,
  profilePageStream,
  airtimePageStream,
  beneficiariesPageStream,
  wastePageStream,
  cacheFirstStrategy,
} from "./sw-strategies";

declare const self: ServiceWorkerGlobalScope & { __WB_MANIFEST: any[] };

clientsClaim();

// Enable navigation preload
navigationPreload.enable();

// RuntimeCaching
// registerRoute(
//   RegExp(/.*\.(?:png|jpg|jpeg|svg|webp)$/),
//   new CacheFirst({
//     cacheName: imagesCache,
//     plugins: [
//       new ExpirationPlugin({
//         maxEntries: 300,
//         maxAgeSeconds: 60 * 60 * 24 * 30, // about a month
//       }),
//     ],
//   }),
// );

// account/auth
// registerRoute(
//   ({ url }) =>
//     url.pathname === "/account/signin-Partial.html" ||
//     url.pathname === "/account/signup-Partial.html",
//   new CacheFirst({
//     cacheName: partialsCache,
//     plugins: [
//       new ExpirationPlugin({
//         maxEntries: 30,
//         maxAgeSeconds: 60 * 60 * 24 * 30, // about a month
//       }),
//     ],
//   }),
// );

registerRoute(
  ({ url }) =>
    url.pathname === "/account/signin-Partial.html" ||
    url.pathname === "/account/signup-Partial.html",
  async ({ request, event }) => {
    const cache = await caches.open(partialsCache);
    let response = await cache.match(request);

    // If there's no cached response, fetch it and cache it
    if (!response) {
      response = await fetch(request);
      await cache.put(request, response.clone());
    }

    // Modify the response headers
    const modifiedResponse = new Response(response.body, {
      status: response.status,
      statusText: response.statusText,
      headers: {
        ...Object.fromEntries(response.headers.entries()),
        "HX-Reswap": "none",
      },
    });

    return modifiedResponse;
  },
);

// account/dashboard
registerRoute(
  new RegExp("^.*/account/dashboard/chart\\?month=[1-4]$"),
  new StaleWhileRevalidate({
    cacheName: contentCache,
    plugins: [
      new ExpirationPlugin({
        maxEntries: 30,
        maxAgeSeconds: 60 * 60 * 24 * 30, // about a month
      }),
    ],
  }),
);

registerRoute(
  new RegExp("^.*/account/dashboard/chart\\?month=0$"),
  new NetworkFirst({
    networkTimeoutSeconds: 3,
    cacheName: contentCache,
    plugins: [
      new ExpirationPlugin({
        maxEntries: 30,
        maxAgeSeconds: 60 * 60 * 24 * 30, // about a month
      }),
    ],
  }),
);

registerRoute(
  ({ url }) => url.pathname === "/account/dashboard",
  dashboardPageStream,
);

// account/scan
registerRoute(
  ({ url }) => url.pathname === "/account/scan",
  new CacheFirst({
    cacheName: contentCache,
    plugins: [
      new ExpirationPlugin({
        maxEntries: 30,
        maxAgeSeconds: 60 * 60 * 24 * 30, // about a month
      }),
    ],
  }),
);

// account/location
registerRoute(
  ({ url }) => url.pathname === "/account/location",
  locationPageStream,
);

// account/wallet
registerRoute(
  ({ url }) => url.pathname === "/account/wallet",
  walletPageStream,
);

// account/wallet
registerRoute(({ url }) => url.pathname === "/account/spend", spendPageStream);

// account/profile
registerRoute(
  ({ url }) => url.pathname === "/account/profile/changeDisplayName",
  async ({ request }) => {
    const response = await fetch(request);

    if (!response.ok) {
      return response;
    }

    const cache = await caches.open(partialsCache);
    const allCacheKeys = await cache.keys();

    const matchingResponses = allCacheKeys.filter((request) =>
      request.url.endsWith("commonHeader-Partial.html"),
    );

    for (const req of matchingResponses) {
      await cache.put(req, response.clone());
    }

    return response;
  },
  "POST",
);

registerRoute(
  ({ url }) => url.pathname === "/account/profile",
  profilePageStream,
);

// account/airtime
registerRoute(
  ({ url }) => url.pathname === "/account/airtime",
  airtimePageStream,
  "POST",
);

// account/beneficiaries
registerRoute(
  ({ url }) => url.pathname === "/account/beneficiaries",
  beneficiariesPageStream,
);

// account/waste
registerRoute(({ url }) => {
  const wastePathPattern = /.*\/account\/waste\/(paper|glass|plastics|metal)$/;
  return wastePathPattern.test(url.href);
}, wastePageStream);

// account/tips
registerRoute(
  ({ url }) => {
    const tipsPathPattern = /.*\/account\/tips\/(paper|glass|plastics|metal)$/;
    return tipsPathPattern.test(url.href);
  },
  new CacheFirst({
    cacheName: contentCache,
    plugins: [
      new ExpirationPlugin({
        maxEntries: 30,
        maxAgeSeconds: 60 * 60 * 24 * 30, // about a month
      }),
    ],
  }),
);

precacheAndRoute(self.__WB_MANIFEST || []);
