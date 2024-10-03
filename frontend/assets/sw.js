const VERSION = "v1.0.0";

const CACHE_NAME = `pwa-${VERSION}`;

const APP_STATIC_RESOURCES = [
  "/",
  "/images/*",
  "/svg/*",
  "/libriaries/*",
  "/styles/output.css",
];

self.addEventListener("install", (e) => {
  console.log("[Service Worker] Install");
  e.waitUntil(
    (async () => {
      const cache = await caches.open(CACHE_NAME);
      console.log("[Service Worker] Caching all: app shell and content");
      await cache.addAll(APP_STATIC_RESOURCESe);
    })(),
  );
});


self.addEventListener("activate", (event) => {
  event.waitUntil(
    (async () => {
      const names = await caches.keys();
      await Promise.all(
        names.map((name) => {
          if (name !== CACHE_NAME) {
            return caches.delete(name);
          }
        }),
      );
      await clients.claim();
    })(),
  );
});

