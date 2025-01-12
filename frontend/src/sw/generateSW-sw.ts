import { generateSW } from "workbox-build";
async function buildSW() {
  const { count, size, warnings } = await generateSW({
    globDirectory: "public",
    globPatterns: ["**/*.{js,css,ttf}"],
    swDest: "public/sw.js",
    maximumFileSizeToCacheInBytes: 2000000000,
    skipWaiting: true,
    disableDevLogs: false,
    cacheId: "local",
    cleanupOutdatedCaches: true,
    clientsClaim: true,
    mode: "development",
    runtimeCaching: [
      {
        urlPattern: /.*\.(?:png|jpg|jpeg|svg)$/,
        handler: "CacheFirst",
        options: {
          cacheName: "images-cache",
          expiration: {
            maxEntries: 300,
            maxAgeSeconds: 60 * 60 * 24 * 30 * 3, // about 3 month
          },
        },
      },
      {
        urlPattern: /.*partial.*/,
        handler: "CacheFirst",
        options: {
          cacheName: "partials-cache",
          expiration: {
            maxEntries: 100,
            maxAgeSeconds: 60 * 60 * 24 * 30 * 3, // about 3 month
          },
        },
      },
    ],
  });

  console.log(
    `Generated a service worker, which will precache ${count} files, totaling ${size} bytes and \n warning: \n ${warnings}.`
  );
}

buildSW();
