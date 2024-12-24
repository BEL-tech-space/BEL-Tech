// sw.js
import * as navigationPreload from "workbox-navigation-preload";
import { CacheFirst, NetworkFirst } from "workbox-strategies";
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
} from "./sw-strategies";

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

// account/location
registerRoute(
  ({ url }) => url.pathname == "/account/location",
  locationPageStream
);

// account/wallet
registerRoute(({ url }) => url.pathname == "/account/wallet", walletPageStream);

// account/wallet
registerRoute(({ url }) => url.pathname == "/account/spend", spendPageStream);

// account/profile/displayName
registerRoute(
  ({ url }) => url.pathname == "/account/profile/displayName",
  new NetworkFirst({
    cacheName: partialsCache,
    plugins: [
      new ExpirationPlugin({
        maxEntries: 30,
        maxAgeSeconds: 60 * 60 * 24 * 30, // about a month
      }),
    ],
  })
);

precacheAndRoute(self.__WB_MANIFEST || []);
