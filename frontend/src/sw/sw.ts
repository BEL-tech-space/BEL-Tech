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
import {
  contentCache,
  partialsCache,
  dashboardPageStream,
  locationPageStream,
  walletPageStream,
  spendPageStream,
  profilePageStream,
  airtimePageStream,
  beneficiariesPageStream,
  wastePageStream,
  tipsPageStream,
} from "./sw-strategies";

declare const self: ServiceWorkerGlobalScope & { __WB_MANIFEST: any[] };

self.addEventListener("message", (event) => {
  if (event.data && event.data.type === "SKIP_WAITING") {
    (self as any).skipWaiting();
  }
});

(self as any).__WB_DISABLE_DEV_LOGS = true;

navigationPreload.enable();

// account/auth
registerRoute(
  ({ url }) =>
    url.pathname === "/account/login.html" ||
    url.pathname === "/account/signup.html",
  new StaleWhileRevalidate({
    cacheName: partialsCache,
    plugins: [
      new ExpirationPlugin({
        maxEntries: 25,
        maxAgeSeconds: 60 * 60 * 24 * 30 * 3, // about 3 month
      }),
    ],
  })
);

// account/dashboard/chart
registerRoute(
  new RegExp("^.*/account/dashboard/chart\\?month=[1-4]$"),
  new StaleWhileRevalidate({
    cacheName: contentCache,
    plugins: [
      new ExpirationPlugin({
        maxEntries: 25,
        maxAgeSeconds: 60 * 60 * 24 * 30 * 3, // about 3 month
      }),
    ],
  })
);

registerRoute(
  new RegExp("^.*/account/dashboard/chart\\?month=0$"),
  new NetworkFirst({
    networkTimeoutSeconds: 3,
    cacheName: contentCache,
    plugins: [
      new ExpirationPlugin({
        maxEntries: 25,
        maxAgeSeconds: 60 * 60 * 24 * 30 * 3, // about 3 month
      }),
    ],
  })
);

// registerRoute(
//   ({ url }) => url.pathname === "/account/dashboard",
//   dashboardPageStream
// );

// account/scan
registerRoute(
  ({ url }) => url.pathname === "/account/scan",
  new StaleWhileRevalidate({
    cacheName: partialsCache,
    plugins: [
      new ExpirationPlugin({
        maxEntries: 25,
        maxAgeSeconds: 60 * 60 * 24 * 30 * 3, // about 3 month
      }),
    ],
  })
);

// account/location
// registerRoute(
//   ({ url }) => url.pathname === "/account/location",
//   locationPageStream
// );

// account/wallet
// registerRoute(
//   ({ url }) => url.pathname === "/account/wallet",
//   walletPageStream
// );

// account/wallet
// registerRoute(({ url }) => url.pathname === "/account/spend", spendPageStream);

// account/profile/changeDisplayName
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
      request.url.endsWith("commonHeader-Partial.html")
    );

    for (const req of matchingResponses) {
      await cache.put(req, response.clone());
    }

    return response;
  },
  "POST"
);

registerRoute(
  ({ url }) => url.pathname === "/account/profile",
  profilePageStream
);

// account/airtime
// registerRoute(
//   ({ url }) => url.pathname === "/account/airtime",
//   airtimePageStream,
//   "POST"
// );

// account/beneficiaries
// registerRoute(
//   ({ url }) => url.pathname === "/account/beneficiaries",
//   beneficiariesPageStream
// );

// account/waste
// registerRoute(({ url }) => {
//   const wastePathPattern = /.*\/account\/waste\/(paper|glass|plastics|metal)$/;
//   return wastePathPattern.test(url.href);
// }, wastePageStream);

// account/tips
// registerRoute(({ url }) => {
//   const tipsPathPattern = /.*\/account\/tips\/(paper|glass|plastics|metal)$/;
//   return tipsPathPattern.test(url.href);
// }, tipsPageStream);

precacheAndRoute(self.__WB_MANIFEST || []);

//
