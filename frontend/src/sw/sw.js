"use strict";
var __createBinding = (this && this.__createBinding) || (Object.create ? (function(o, m, k, k2) {
    if (k2 === undefined) k2 = k;
    var desc = Object.getOwnPropertyDescriptor(m, k);
    if (!desc || ("get" in desc ? !m.__esModule : desc.writable || desc.configurable)) {
      desc = { enumerable: true, get: function() { return m[k]; } };
    }
    Object.defineProperty(o, k2, desc);
}) : (function(o, m, k, k2) {
    if (k2 === undefined) k2 = k;
    o[k2] = m[k];
}));
var __setModuleDefault = (this && this.__setModuleDefault) || (Object.create ? (function(o, v) {
    Object.defineProperty(o, "default", { enumerable: true, value: v });
}) : function(o, v) {
    o["default"] = v;
});
var __importStar = (this && this.__importStar) || (function () {
    var ownKeys = function(o) {
        ownKeys = Object.getOwnPropertyNames || function (o) {
            var ar = [];
            for (var k in o) if (Object.prototype.hasOwnProperty.call(o, k)) ar[ar.length] = k;
            return ar;
        };
        return ownKeys(o);
    };
    return function (mod) {
        if (mod && mod.__esModule) return mod;
        var result = {};
        if (mod != null) for (var k = ownKeys(mod), i = 0; i < k.length; i++) if (k[i] !== "default") __createBinding(result, mod, k[i]);
        __setModuleDefault(result, mod);
        return result;
    };
})();
var __awaiter = (this && this.__awaiter) || function (thisArg, _arguments, P, generator) {
    function adopt(value) { return value instanceof P ? value : new P(function (resolve) { resolve(value); }); }
    return new (P || (P = Promise))(function (resolve, reject) {
        function fulfilled(value) { try { step(generator.next(value)); } catch (e) { reject(e); } }
        function rejected(value) { try { step(generator["throw"](value)); } catch (e) { reject(e); } }
        function step(result) { result.done ? resolve(result.value) : adopt(result.value).then(fulfilled, rejected); }
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
const workbox_streams_1 = require("workbox-streams");
const workbox_core_1 = require("workbox-core");
const contentCache = 'contents-cache';
const partialsCache = 'partials-cache';
const imagesCache = 'images-cache';
(0, workbox_core_1.clientsClaim)();
// Enable navigation preload
navigationPreload.enable();
// RuntimeCaching
(0, workbox_routing_1.registerRoute)(RegExp(/.*\.(?:png|jpg|jpeg|svg)$/), new workbox_strategies_1.CacheFirst({
    cacheName: imagesCache,
    plugins: [
        new workbox_expiration_1.ExpirationPlugin({
            maxEntries: 300,
            maxAgeSeconds: 60 * 60 * 24 * 30, // about a month
        }),
    ],
}));
const cacheFirstPartialStrategy = new workbox_strategies_1.CacheFirst({
    cacheName: partialsCache,
    plugins: [
        new workbox_expiration_1.ExpirationPlugin({
            maxEntries: 30,
            maxAgeSeconds: 60 * 60 * 24 * 30, // about a month
        }),
        // {
        //   handlerDidError: async ({request}) => {
        //     return await matchPrecache('/offline-header.php');
        //   }
        // }
    ]
});
const networkFirstContentStrategy = new workbox_strategies_1.NetworkFirst({
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
const accountDashboardNavigation = (0, workbox_streams_1.strategy)([
    (_a) => __awaiter(void 0, [_a], void 0, function* ({ event }) {
        return cacheFirstPartialStrategy.handle({ request: new Request('/account/htmlAndHeaderTag-Partial.html'), event });
    }),
    (_a) => __awaiter(void 0, [_a], void 0, function* ({ event }) {
        return cacheFirstPartialStrategy.handle({ request: new Request('/account/dashboard/bodyTag-Partial.html'), event });
    }),
    (_a) => __awaiter(void 0, [_a], void 0, function* ({ event }) {
        return cacheFirstPartialStrategy.handle({ request: new Request('/account/dashboard/headerTypeOne-Partial.html'), event });
    }),
    (_a) => __awaiter(void 0, [_a], void 0, function* ({ event }) {
        return networkFirstContentStrategy.handle({ request: new Request('/account/dashboard', { headers: [['X-Content-Mode', 'partials']] }), event });
    }),
    (_a) => __awaiter(void 0, [_a], void 0, function* ({ event }) {
        return cacheFirstPartialStrategy.handle({ request: new Request('/account/dashboard/footer-Partial.html'), event });
    }),
    (_a) => __awaiter(void 0, [_a], void 0, function* ({ event }) {
        return cacheFirstPartialStrategy.handle({ request: new Request('/account/dashboard/scriptAndClosing-Partial.html'), event });
    }),
], {});
(0, workbox_routing_1.registerRoute)(({ url }) => url.pathname == '/account/dashboard', accountDashboardNavigation);
(0, workbox_routing_1.registerRoute)(({ url }) => url.pathname == '/account/scan', new workbox_strategies_1.CacheFirst({
    cacheName: contentCache,
    plugins: [
        new workbox_expiration_1.ExpirationPlugin({
            maxEntries: 30,
            maxAgeSeconds: 60 * 60 * 24 * 30, // about a month
        }),
    ]
}));
const locationPageStream = (0, workbox_streams_1.strategy)([
    (_a) => __awaiter(void 0, [_a], void 0, function* ({ event }) {
        return cacheFirstPartialStrategy.handle({ request: new Request('/account/htmlAndHeaderTag-Partial.html'), event });
    }),
    (_a) => __awaiter(void 0, [_a], void 0, function* ({ event }) {
        return cacheFirstPartialStrategy.handle({ request: new Request('/location/bodyTag-Partial.html'), event });
    }),
    (_a) => __awaiter(void 0, [_a], void 0, function* ({ event }) {
        return cacheFirstPartialStrategy.handle({ request: new Request('/location/headerTag-Partial.html'), event });
    }),
    (_a) => __awaiter(void 0, [_a], void 0, function* ({ event }) {
        return networkFirstContentStrategy.handle({ request: new Request('/account/dashboard', { headers: [['X-Content-Mode', 'partials']] }), event });
    }),
    (_a) => __awaiter(void 0, [_a], void 0, function* ({ event }) {
        return cacheFirstPartialStrategy.handle({ request: new Request('/location/footer-Partial.html'), event });
    }),
    (_a) => __awaiter(void 0, [_a], void 0, function* ({ event }) {
        return cacheFirstPartialStrategy.handle({ request: new Request('/location/scriptAndClosing-Partial.html'), event });
    }),
], {});
(0, workbox_precaching_1.precacheAndRoute)(self.__WB_MANIFEST || []);
