"use strict";
var __defProp = Object.defineProperty;
var __getOwnPropDesc = Object.getOwnPropertyDescriptor;
var __getOwnPropNames = Object.getOwnPropertyNames;
var __hasOwnProp = Object.prototype.hasOwnProperty;
var __export = (target, all) => {
  for (var name in all)
    __defProp(target, name, { get: all[name], enumerable: true });
};
var __copyProps = (to, from, except, desc) => {
  if (from && typeof from === "object" || typeof from === "function") {
    for (let key of __getOwnPropNames(from))
      if (!__hasOwnProp.call(to, key) && key !== except)
        __defProp(to, key, { get: () => from[key], enumerable: !(desc = __getOwnPropDesc(from, key)) || desc.enumerable });
  }
  return to;
};
var __toCommonJS = (mod) => __copyProps(__defProp({}, "__esModule", { value: true }), mod);

// src/build-sw.ts
var build_sw_exports = {};
__export(build_sw_exports, {
  buildSW: () => buildSW
});
module.exports = __toCommonJS(build_sw_exports);
var import_workbox_build = require("workbox-build");
async function buildSW() {
  const { count, size, warnings } = await (0, import_workbox_build.generateSW)({
    globDirectory: "public/assets",
    globPatterns: ["**/*.{js,css}"],
    swDest: "public/swThree.js",
    maximumFileSizeToCacheInBytes: 2e9,
    runtimeCaching: [{
      urlPattern: new RegExp("/.(?:png|jpg|jpeg|svg)$/,"),
      handler: "CacheFirst",
      options: {
        cacheName: "moreAssets",
        expiration: {
          maxEntries: 10,
          maxAgeSeconds: 60 * 60 * 24 * 30
          // about a month
        }
      }
    }, {
      urlPattern: new RegExp("/.(?:xml|html)$/,"),
      handler: "StaleWhileRevalidate",
      options: {
        cacheName: "html-cache",
        expiration: {
          maxEntries: 10,
          maxAgeSeconds: 60 * 60 * 24 * 30
          // about a month
        }
      }
    }]
  });
  console.log(`Generated a service worker, which will precache ${count} files, totaling ${size} bytes and 
 warning: 
 ${warnings}.`);
}
buildSW();
// Annotate the CommonJS export names for ESM import in node:
0 && (module.exports = {
  buildSW
});
//# sourceMappingURL=workbox-config.js.map
