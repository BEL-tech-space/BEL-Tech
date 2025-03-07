// sw.js
import { StaleWhileRevalidate, CacheFirst } from "workbox-strategies";
import { registerRoute } from "workbox-routing";
import { precacheAndRoute } from "workbox-precaching";
import { ExpirationPlugin } from "workbox-expiration";
import { partialsCache } from "./sw-strategies";

declare const self: ServiceWorkerGlobalScope & { __WB_MANIFEST: any[] };

self.addEventListener("message", (event) => {
  if (event.data && event.data.type === "SKIP_WAITING") {
    (self as any).skipWaiting();
  }
});

(self as any).__WB_DISABLE_DEV_LOGS = true;

registerRoute(
  ({ url }) => /\.(png|jpe?g|gif|webp|svg)$/i.test(url.pathname),
  new CacheFirst({
    cacheName: "image-cache",
    plugins: [
      new ExpirationPlugin({
        maxEntries: 400, // Adjust based on your needs
        maxAgeSeconds: 60 * 60 * 24 * 30, // Cache for 30 days
      }),
    ],
  })
);

// account/auth
registerRoute(
  ({ url }) =>
    url.pathname === "/account/login.html" ||
    url.pathname === "/account/signup.html",
  new StaleWhileRevalidate({
    cacheName: partialsCache,
    plugins: [
      new ExpirationPlugin({
        maxEntries: 2,
        maxAgeSeconds: 60 * 60 * 24 * 10, // about 10 days
      }),
    ],
  })
);

// agent/auth
registerRoute(
  ({ url }) =>
    url.pathname === "/agent/login.html" ||
    url.pathname === "/agent/signout.html",
  new StaleWhileRevalidate({
    cacheName: partialsCache,
    plugins: [
      new ExpirationPlugin({
        maxEntries: 2,
        maxAgeSeconds: 60 * 60 * 24 * 10, // about 10 days
      }),
    ],
  })
);

// account/scan
registerRoute(
  ({ url }) => url.pathname === "/account/scan",
  new StaleWhileRevalidate({
    cacheName: partialsCache,
    plugins: [
      new ExpirationPlugin({
        maxEntries: 2,
        maxAgeSeconds: 60 * 60 * 24 * 10, // about 10 days
      }),
    ],
  })
);

// agent/collect
registerRoute(
  ({ url }) => url.pathname === "/agent/collect",
  new StaleWhileRevalidate({
    cacheName: partialsCache,
    plugins: [
      new ExpirationPlugin({
        maxEntries: 2,
        maxAgeSeconds: 60 * 60 * 24 * 10, // about 10 days
      }),
    ],
  })
);

precacheAndRoute(self.__WB_MANIFEST || []);
