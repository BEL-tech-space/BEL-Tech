"use strict";
var __createBinding =
  (this && this.__createBinding) ||
  (Object.create
    ? function (o, m, k, k2) {
        if (k2 === undefined) k2 = k;
        var desc = Object.getOwnPropertyDescriptor(m, k);
        if (
          !desc ||
          ("get" in desc ? !m.__esModule : desc.writable || desc.configurable)
        ) {
          desc = {
            enumerable: true,
            get: function () {
              return m[k];
            },
          };
        }
        Object.defineProperty(o, k2, desc);
      }
    : function (o, m, k, k2) {
        if (k2 === undefined) k2 = k;
        o[k2] = m[k];
      });
var __setModuleDefault =
  (this && this.__setModuleDefault) ||
  (Object.create
    ? function (o, v) {
        Object.defineProperty(o, "default", { enumerable: true, value: v });
      }
    : function (o, v) {
        o["default"] = v;
      });
var __importStar =
  (this && this.__importStar) ||
  (function () {
    var ownKeys = function (o) {
      ownKeys =
        Object.getOwnPropertyNames ||
        function (o) {
          var ar = [];
          for (var k in o)
            if (Object.prototype.hasOwnProperty.call(o, k)) ar[ar.length] = k;
          return ar;
        };
      return ownKeys(o);
    };
    return function (mod) {
      if (mod && mod.__esModule) return mod;
      var result = {};
      if (mod != null)
        for (var k = ownKeys(mod), i = 0; i < k.length; i++)
          if (k[i] !== "default") __createBinding(result, mod, k[i]);
      __setModuleDefault(result, mod);
      return result;
    };
  })();
var __awaiter =
  (this && this.__awaiter) ||
  function (thisArg, _arguments, P, generator) {
    function adopt(value) {
      return value instanceof P
        ? value
        : new P(function (resolve) {
            resolve(value);
          });
    }
    return new (P || (P = Promise))(function (resolve, reject) {
      function fulfilled(value) {
        try {
          step(generator.next(value));
        } catch (e) {
          reject(e);
        }
      }
      function rejected(value) {
        try {
          step(generator["throw"](value));
        } catch (e) {
          reject(e);
        }
      }
      function step(result) {
        result.done
          ? resolve(result.value)
          : adopt(result.value).then(fulfilled, rejected);
      }
      step((generator = generator.apply(thisArg, _arguments || [])).next());
    });
  };
Object.defineProperty(exports, "__esModule", { value: true });
// sw.js
const navigationPreload = __importStar(require("workbox-navigation-preload"));
const workbox_strategies_1 = require("workbox-strategies");
const workbox_routing_1 = require("workbox-routing");
const workbox_precaching_1 = require("workbox-precaching");
const workbox_expiration_1 = require("workbox-expiration");
const sw_strategies_1 = require("./sw-strategies");
self.addEventListener("message", (event) => {
  if (event.data && event.data.type === "SKIP_WAITING") {
    self.skipWaiting();
  }
});
self.__WB_DISABLE_DEV_LOGS = true;
navigationPreload.enable();
// account/auth
(0, workbox_routing_1.registerRoute)(
  ({ url }) =>
    url.pathname === "/account/login.html" ||
    url.pathname === "/account/signup.html",
  new workbox_strategies_1.CacheFirst({
    cacheName: sw_strategies_1.partialsCache,
    plugins: [
      new workbox_expiration_1.ExpirationPlugin({
        maxEntries: 25,
        maxAgeSeconds: 60 * 60 * 24 * 30 * 3, // about 3 month
      }),
    ],
  })
);
// account/dashboard
(0, workbox_routing_1.registerRoute)(
  new RegExp("^.*/account/dashboard/chart\\?month=[1-4]$"),
  new workbox_strategies_1.StaleWhileRevalidate({
    cacheName: sw_strategies_1.contentCache,
    plugins: [
      new workbox_expiration_1.ExpirationPlugin({
        maxEntries: 25,
        maxAgeSeconds: 60 * 60 * 24 * 30 * 3, // about 3 month
      }),
    ],
  })
);
(0, workbox_routing_1.registerRoute)(
  new RegExp("^.*/account/dashboard/chart\\?month=0$"),
  new workbox_strategies_1.NetworkFirst({
    networkTimeoutSeconds: 3,
    cacheName: sw_strategies_1.contentCache,
    plugins: [
      new workbox_expiration_1.ExpirationPlugin({
        maxEntries: 25,
        maxAgeSeconds: 60 * 60 * 24 * 30 * 3, // about 3 month
      }),
    ],
  })
);
(0, workbox_routing_1.registerRoute)(
  ({ url }) => url.pathname === "/account/dashboard",
  sw_strategies_1.dashboardPageStream
);
// account/scan
(0, workbox_routing_1.registerRoute)(
  ({ url }) => url.pathname === "/account/scan",
  new workbox_strategies_1.CacheFirst({
    cacheName: sw_strategies_1.contentCache,
    plugins: [
      new workbox_expiration_1.ExpirationPlugin({
        maxEntries: 25,
        maxAgeSeconds: 60 * 60 * 24 * 30 * 3, // about 3 month
      }),
    ],
  })
);
// account/location
(0, workbox_routing_1.registerRoute)(
  ({ url }) => url.pathname === "/account/location",
  sw_strategies_1.locationPageStream
);
// account/wallet
(0, workbox_routing_1.registerRoute)(
  ({ url }) => url.pathname === "/account/wallet",
  sw_strategies_1.walletPageStream
);
// account/wallet
(0, workbox_routing_1.registerRoute)(
  ({ url }) => url.pathname === "/account/spend",
  sw_strategies_1.spendPageStream
);
// account/profile/changeDisplayName
(0, workbox_routing_1.registerRoute)(
  ({ url }) => url.pathname === "/account/profile/changeDisplayName",
  (_a) =>
    __awaiter(void 0, [_a], void 0, function* ({ request }) {
      const response = yield fetch(request);
      if (!response.ok) {
        return response;
      }
      const cache = yield caches.open(sw_strategies_1.partialsCache);
      const allCacheKeys = yield cache.keys();
      const matchingResponses = allCacheKeys.filter((request) =>
        request.url.endsWith("commonHeader-Partial.html")
      );
      for (const req of matchingResponses) {
        yield cache.put(req, response.clone());
      }
      return response;
    }),
  "POST"
);
(0, workbox_routing_1.registerRoute)(
  ({ url }) => url.pathname === "/account/profile",
  sw_strategies_1.profilePageStream
);
// account/airtime
(0, workbox_routing_1.registerRoute)(
  ({ url }) => url.pathname === "/account/airtime",
  sw_strategies_1.airtimePageStream,
  "POST"
);
// account/beneficiaries
(0, workbox_routing_1.registerRoute)(
  ({ url }) => url.pathname === "/account/beneficiaries",
  sw_strategies_1.beneficiariesPageStream
);
// account/waste
(0, workbox_routing_1.registerRoute)(({ url }) => {
  const wastePathPattern = /.*\/account\/waste\/(paper|glass|plastics|metal)$/;
  return wastePathPattern.test(url.href);
}, sw_strategies_1.wastePageStream);
// account/tips
(0, workbox_routing_1.registerRoute)(({ url }) => {
  const tipsPathPattern = /.*\/account\/tips\/(paper|glass|plastics|metal)$/;
  return tipsPathPattern.test(url.href);
}, sw_strategies_1.tipsPageStream);
(0, workbox_precaching_1.precacheAndRoute)(self.__WB_MANIFEST || []);
//
