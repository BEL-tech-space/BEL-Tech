import { injectManifest } from "workbox-build";

async function injectSW() {
  const { count, size, warnings } = await injectManifest({
    globDirectory: "public",
    globPatterns: ["**/*.{js,css,ttf,otf,json,png,jpg,jpeg,svg,webp}"],
    swSrc: "public/sw.js",
    swDest: "public/sw.js",
    maximumFileSizeToCacheInBytes: 20000000,
  });

  console.log(
    `Generated a service worker, which will precache ${count} files, totaling ${size} bytes and \n warning: \n ${warnings}.`
  );
}

injectSW();
