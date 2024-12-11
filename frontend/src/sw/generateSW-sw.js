"use strict";
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
const workbox_build_1 = require("workbox-build");
function buildSW() {
    return __awaiter(this, void 0, void 0, function* () {
        const { count, size, warnings } = yield (0, workbox_build_1.generateSW)({
            globDirectory: 'public',
            globPatterns: ['**/*.{js,css,ttf}'],
            swDest: 'public/sw.js',
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
                    handler: 'CacheFirst',
                    options: {
                        cacheName: 'images-cache',
                        expiration: {
                            maxEntries: 300,
                            maxAgeSeconds: 60 * 60 * 24 * 30, // about a month
                        },
                    },
                }, {
                    urlPattern: /.*partial.*/,
                    handler: 'CacheFirst',
                    options: {
                        cacheName: 'partials-cache',
                        expiration: {
                            maxEntries: 100,
                            maxAgeSeconds: 60 * 60 * 24 * 30, // about a month
                        },
                    },
                },
            ],
        });
        console.log(`Generated a service worker, which will precache ${count} files, totaling ${size} bytes and \n warning: \n ${warnings}.`);
    });
}
buildSW();
