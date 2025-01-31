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
exports.tipsPageStream = exports.wastePageStream = exports.beneficiariesPageStream = exports.airtimePageStream = exports.profilePageStream = exports.spendPageStream = exports.walletPageStream = exports.locationPageStream = exports.dashboardPageStream = exports.staleWhileRevalidateStrategy = exports.networkFirstStrategy = exports.cacheFirstStrategy = exports.imagesCache = exports.partialsCache = exports.contentCache = void 0;
const workbox_strategies_1 = require("workbox-strategies");
const workbox_expiration_1 = require("workbox-expiration");
const workbox_streams_1 = require("workbox-streams");
exports.contentCache = "contents-cache";
exports.partialsCache = "partials-cache";
exports.imagesCache = "images-cache";
exports.cacheFirstStrategy = new workbox_strategies_1.CacheFirst({
    cacheName: exports.partialsCache,
    plugins: [
        new workbox_expiration_1.ExpirationPlugin({
            maxEntries: 30,
            maxAgeSeconds: 60 * 60 * 24 * 30 * 3, // about 3 months
        }),
    ],
});
exports.networkFirstStrategy = new workbox_strategies_1.NetworkFirst({
    cacheName: exports.contentCache,
    networkTimeoutSeconds: 5,
    plugins: [
        new workbox_expiration_1.ExpirationPlugin({
            maxEntries: 30,
            maxAgeSeconds: 60 * 60 * 24 * 30 * 3, // about 3 months
        }),
    ],
});
exports.staleWhileRevalidateStrategy = new workbox_strategies_1.StaleWhileRevalidate({
    cacheName: exports.contentCache,
    plugins: [
        new workbox_expiration_1.ExpirationPlugin({
            maxEntries: 30,
            maxAgeSeconds: 60 * 60 * 24 * 30 * 3, // about 3 months
        }),
        // {
        //   fetchDidSucceed: async ({ request, response, event, state }) => {
        //     const clonedResponse = response.clone();
        //     if (clonedResponse.headers.get("X-Session-Expired")) {
        //       console.warn("Session expired. Redirecting to signup page.");
        //       return Response.redirect("/account/login.html", 302);
        //     }
        //     return response;
        //   },
        // },
    ],
});
exports.dashboardPageStream = (0, workbox_streams_1.strategy)([
    (_a) => __awaiter(void 0, [_a], void 0, function* ({ event }) {
        return exports.staleWhileRevalidateStrategy.handle({
            request: new Request("/account/dashboard/openingHtmlTag-Partial.html"),
            event,
        });
    }),
    (_a) => __awaiter(void 0, [_a], void 0, function* ({ event }) {
        return exports.networkFirstStrategy.handle({
            request: new Request("/account/dashboard/commonHeader-Partial.html"),
            event,
        });
    }),
    (_a) => __awaiter(void 0, [_a], void 0, function* ({ event }) {
        return exports.staleWhileRevalidateStrategy.handle({
            request: new Request("/account/dashboard/mainOpeningTag-Partial.html"),
            event,
        });
    }),
    (_a) => __awaiter(void 0, [_a], void 0, function* ({ event }) {
        return exports.staleWhileRevalidateStrategy.handle({
            request: new Request("/account/dashboard/swiperOne"),
            event,
        });
    }),
    (_a) => __awaiter(void 0, [_a], void 0, function* ({ event }) {
        return exports.networkFirstStrategy.handle({
            request: new Request("/account/dashboard/echarts?month=0"),
            event,
        });
    }),
    (_a) => __awaiter(void 0, [_a], void 0, function* ({ event }) {
        return exports.staleWhileRevalidateStrategy.handle({
            request: new Request("/account/dashboard/swiperTwo"),
            event,
        });
    }),
    (_a) => __awaiter(void 0, [_a], void 0, function* ({ event }) {
        return exports.cacheFirstStrategy.handle({
            request: new Request("/account/dashboard/footer-Partial.html"),
            event,
        });
    }),
    (_a) => __awaiter(void 0, [_a], void 0, function* ({ event }) {
        return exports.staleWhileRevalidateStrategy.handle({
            request: new Request("/account/dashboard/closingHtmlTag-Partial.html"),
            event,
        });
    }),
], {});
exports.locationPageStream = (0, workbox_streams_1.strategy)([
    (_a) => __awaiter(void 0, [_a], void 0, function* ({ event }) {
        return exports.staleWhileRevalidateStrategy.handle({
            request: new Request("/account/location/openingHtmlTag-Partial.html"),
            event,
        });
    }),
    (_a) => __awaiter(void 0, [_a], void 0, function* ({ event }) {
        return exports.staleWhileRevalidateStrategy.handle({
            request: new Request("/account/location/header-Partial.html"),
            event,
        });
    }),
    (_a) => __awaiter(void 0, [_a], void 0, function* ({ event }) {
        return exports.networkFirstStrategy.handle({
            request: new Request("/account/location/main", {
                headers: [["X-Content-Mode", "partials"]],
            }),
            event,
        });
    }),
    (_a) => __awaiter(void 0, [_a], void 0, function* ({ event }) {
        return exports.cacheFirstStrategy.handle({
            request: new Request("/account/location/footer-Partial.html"),
            event,
        });
    }),
    (_a) => __awaiter(void 0, [_a], void 0, function* ({ event }) {
        return exports.staleWhileRevalidateStrategy.handle({
            request: new Request("/account/location/closingHtmlTag-Partial.html"),
            event,
        });
    }),
], {});
exports.walletPageStream = (0, workbox_streams_1.strategy)([
    (_a) => __awaiter(void 0, [_a], void 0, function* ({ event }) {
        return exports.staleWhileRevalidateStrategy.handle({
            request: new Request("/account/wallet/openingHtmlTag-Partial.html"),
            event,
        });
    }),
    (_a) => __awaiter(void 0, [_a], void 0, function* ({ event }) {
        return exports.networkFirstStrategy.handle({
            request: new Request("/account/wallet/commonHeader-Partial.html"),
            event,
        });
    }),
    (_a) => __awaiter(void 0, [_a], void 0, function* ({ event }) {
        return exports.cacheFirstStrategy.handle({
            request: new Request("/account/wallet/mainOpeningTag-Partial.html"),
            event,
        });
    }),
    (_a) => __awaiter(void 0, [_a], void 0, function* ({ event }) {
        return exports.networkFirstStrategy.handle({
            request: new Request("/account/wallet/balance"),
            event,
        });
    }),
    (_a) => __awaiter(void 0, [_a], void 0, function* ({ event }) {
        return exports.staleWhileRevalidateStrategy.handle({
            request: new Request("/account/wallet/options"),
            event,
        });
    }),
    (_a) => __awaiter(void 0, [_a], void 0, function* ({ event }) {
        return exports.networkFirstStrategy.handle({
            request: new Request("/account/wallet/coinData", {}),
            event,
        });
    }),
    (_a) => __awaiter(void 0, [_a], void 0, function* ({ event }) {
        return exports.staleWhileRevalidateStrategy.handle({
            request: new Request("/account/wallet/swiper"),
            event,
        });
    }),
    (_a) => __awaiter(void 0, [_a], void 0, function* ({ event }) {
        return exports.cacheFirstStrategy.handle({
            request: new Request("/account/wallet/footer-Partial.html"),
            event,
        });
    }),
    (_a) => __awaiter(void 0, [_a], void 0, function* ({ event }) {
        return exports.staleWhileRevalidateStrategy.handle({
            request: new Request("/account/wallet/closingHtmlTag-Partial.html"),
            event,
        });
    }),
], {});
exports.spendPageStream = (0, workbox_streams_1.strategy)([
    (_a) => __awaiter(void 0, [_a], void 0, function* ({ event }) {
        return exports.staleWhileRevalidateStrategy.handle({
            request: new Request("/account/spend/openingHtmlTag-Partial.html"),
            event,
        });
    }),
    (_a) => __awaiter(void 0, [_a], void 0, function* ({ event }) {
        return exports.networkFirstStrategy.handle({
            request: new Request("/account/spend/header-Partial.html"),
            event,
        });
    }),
    (_a) => __awaiter(void 0, [_a], void 0, function* ({ event }) {
        return exports.cacheFirstStrategy.handle({
            request: new Request("/account/spend/mainOpeningTag-Partial.html"),
            event,
        });
    }),
    (_a) => __awaiter(void 0, [_a], void 0, function* ({ event }) {
        return exports.networkFirstStrategy.handle({
            request: new Request("/account/spend/balance"),
            event,
        });
    }),
    (_a) => __awaiter(void 0, [_a], void 0, function* ({ event }) {
        return exports.staleWhileRevalidateStrategy.handle({
            request: new Request("/account/spend/otherhalf"),
            event,
        });
    }),
    (_a) => __awaiter(void 0, [_a], void 0, function* ({ event }) {
        return exports.cacheFirstStrategy.handle({
            request: new Request("/account/spend/footer-Partial.html"),
            event,
        });
    }),
    (_a) => __awaiter(void 0, [_a], void 0, function* ({ event }) {
        return exports.staleWhileRevalidateStrategy.handle({
            request: new Request("/account/spend/closingHtmlTag-Partial.html"),
            event,
        });
    }),
], {});
exports.profilePageStream = (0, workbox_streams_1.strategy)([
    (_a) => __awaiter(void 0, [_a], void 0, function* ({ event }) {
        return exports.staleWhileRevalidateStrategy.handle({
            request: new Request("/account/profile/openingHtmlTag-Partial.html"),
            event,
        });
    }),
    (_a) => __awaiter(void 0, [_a], void 0, function* ({ event }) {
        return exports.networkFirstStrategy.handle({
            request: new Request("/account/profile/commonHeader-Partial.html"),
            event,
        });
    }),
    (_a) => __awaiter(void 0, [_a], void 0, function* ({ event }) {
        return exports.cacheFirstStrategy.handle({
            request: new Request("/account/profile/mainOpeningTag-Partial.html"),
            event,
        });
    }),
    (_a) => __awaiter(void 0, [_a], void 0, function* ({ event }) {
        return exports.networkFirstStrategy.handle({
            request: new Request("/account/profile/displayIcon"),
            event,
        });
    }),
    (_a) => __awaiter(void 0, [_a], void 0, function* ({ event }) {
        return exports.networkFirstStrategy.handle({
            request: new Request("/account/profile/formPartial"),
            event,
        });
    }),
    (_a) => __awaiter(void 0, [_a], void 0, function* ({ event }) {
        return exports.cacheFirstStrategy.handle({
            request: new Request("/account/profile/footer-Partial.html"),
            event,
        });
    }),
    (_a) => __awaiter(void 0, [_a], void 0, function* ({ event }) {
        return exports.staleWhileRevalidateStrategy.handle({
            request: new Request("/account/profile/closingHtmlTag-Partial.html"),
            event,
        });
    }),
], {});
exports.airtimePageStream = (0, workbox_streams_1.strategy)([
    (_a) => __awaiter(void 0, [_a], void 0, function* ({ event }) {
        return exports.staleWhileRevalidateStrategy.handle({
            request: new Request("/account/airtime/openingHtmlTag-Partial.html"),
            event,
        });
    }),
    (_a) => __awaiter(void 0, [_a], void 0, function* ({ event }) {
        return exports.networkFirstStrategy.handle({
            request: new Request("/account/airtime/header-Partial.html"),
            event,
        });
    }),
    (_a) => __awaiter(void 0, [_a], void 0, function* ({ event }) {
        return exports.cacheFirstStrategy.handle({
            request: new Request("/account/airtime/mainOpeningTag-Partial.html"),
            event,
        });
    }),
    (_a) => __awaiter(void 0, [_a], void 0, function* ({ event }) {
        return exports.staleWhileRevalidateStrategy.handle({
            request: new Request("/account/airtime/inputNumber"),
            event,
        });
    }),
    (_a) => __awaiter(void 0, [_a], void 0, function* ({ event }) {
        return exports.networkFirstStrategy.handle({
            request: new Request("/account/airtime/beneficiaries"),
            event,
        });
    }),
    (_a) => __awaiter(void 0, [_a], void 0, function* ({ event }) {
        return exports.cacheFirstStrategy.handle({
            request: new Request("/account/airtime/footer-Partial.html"),
            event,
        });
    }),
    (_a) => __awaiter(void 0, [_a], void 0, function* ({ event }) {
        return exports.staleWhileRevalidateStrategy.handle({
            request: new Request("/account/airtime/closingHtmlTag-Partial.html"),
            event,
        });
    }),
], {});
exports.beneficiariesPageStream = (0, workbox_streams_1.strategy)([
    (_a) => __awaiter(void 0, [_a], void 0, function* ({ event }) {
        return exports.staleWhileRevalidateStrategy.handle({
            request: new Request("/account/beneficiaries/openingHtmlTag-Partial.html"),
            event,
        });
    }),
    (_a) => __awaiter(void 0, [_a], void 0, function* ({ event, url }) {
        return exports.networkFirstStrategy.handle({
            request: new Request("/account/beneficiaries/header-Partial.html"),
            event,
        });
    }),
    (_a) => __awaiter(void 0, [_a], void 0, function* ({ event }) {
        return exports.networkFirstStrategy.handle({
            request: new Request("/account/beneficiaries/main"),
            event,
        });
    }),
    (_a) => __awaiter(void 0, [_a], void 0, function* ({ event }) {
        return exports.cacheFirstStrategy.handle({
            request: new Request("/account/beneficiaries/footer-Partial.html"),
            event,
        });
    }),
    (_a) => __awaiter(void 0, [_a], void 0, function* ({ event }) {
        return exports.staleWhileRevalidateStrategy.handle({
            request: new Request("/account/beneficiaries/closingHtmlTag-Partial.html"),
            event,
        });
    }),
], {});
exports.wastePageStream = (0, workbox_streams_1.strategy)([
    (_a) => __awaiter(void 0, [_a], void 0, function* ({ event }) {
        return exports.staleWhileRevalidateStrategy.handle({
            request: new Request("/account/waste/openingHtmlTag-Partial.html"),
            event,
        });
    }),
    (_a) => __awaiter(void 0, [_a], void 0, function* ({ event }) {
        return exports.networkFirstStrategy.handle({
            request: new Request("/account/waste/commonHeader-Partial.html"),
            event,
        });
    }),
    (_a) => __awaiter(void 0, [_a], void 0, function* ({ event }) {
        return exports.cacheFirstStrategy.handle({
            request: new Request("/account/waste/mainOpeningTag-Partial.html"),
            event,
        });
    }),
    (_a) => __awaiter(void 0, [_a], void 0, function* ({ event, url }) {
        return exports.networkFirstStrategy.handle({
            request: matchingPathVariableForWaste(url, "quickInfo", "waste"),
            event,
        });
    }),
    (_a) => __awaiter(void 0, [_a], void 0, function* ({ event }) {
        return exports.staleWhileRevalidateStrategy.handle({
            request: new Request("/account/waste/filter"),
            event,
        });
    }),
    (_a) => __awaiter(void 0, [_a], void 0, function* ({ event, url }) {
        return exports.networkFirstStrategy.handle({
            request: matchingPathVariableForWaste(url, "history", "waste"),
            event,
        });
    }),
    (_a) => __awaiter(void 0, [_a], void 0, function* ({ event }) {
        return exports.cacheFirstStrategy.handle({
            request: new Request("/account/waste/footer-Partial.html"),
            event,
        });
    }),
    (_a) => __awaiter(void 0, [_a], void 0, function* ({ event }) {
        return exports.staleWhileRevalidateStrategy.handle({
            request: new Request("/account/waste/closingHtmlTag-Partial.html"),
            event,
        });
    }),
], {});
exports.tipsPageStream = (0, workbox_streams_1.strategy)([
    (_a) => __awaiter(void 0, [_a], void 0, function* ({ event }) {
        return exports.staleWhileRevalidateStrategy.handle({
            request: new Request("/account/tips/openingHtmlTag-Partial.html"),
            event,
        });
    }),
    (_a) => __awaiter(void 0, [_a], void 0, function* ({ event }) {
        return exports.staleWhileRevalidateStrategy.handle({
            request: new Request("/account/tips/tipsHeader-Partial.html"),
            event,
        });
    }),
    (_a) => __awaiter(void 0, [_a], void 0, function* ({ event }) {
        return exports.cacheFirstStrategy.handle({
            request: new Request("/account/tips/mainOpeningTag-Partial.html"),
            event,
        });
    }),
    (_a) => __awaiter(void 0, [_a], void 0, function* ({ event, url }) {
        return exports.networkFirstStrategy.handle({
            request: matchingPathVariableForWaste(url, "quickInfo", "tips"),
            event,
        });
    }),
    (_a) => __awaiter(void 0, [_a], void 0, function* ({ event, url }) {
        return exports.networkFirstStrategy.handle({
            request: matchingPathVariableForWaste(url, "article", "tips"),
            event,
        });
    }),
    (_a) => __awaiter(void 0, [_a], void 0, function* ({ event }) {
        return exports.cacheFirstStrategy.handle({
            request: new Request("/account/tips/footer-Partial.html"),
            event,
        });
    }),
    (_a) => __awaiter(void 0, [_a], void 0, function* ({ event }) {
        return exports.staleWhileRevalidateStrategy.handle({
            request: new Request("/account/tips/closingHtmlTag-Partial.html"),
            event,
        });
    }),
], {});
function matchingPathVariableForWaste(url, endpoint, page) {
    const segments = url.pathname.split("/");
    const lastSegment = segments[segments.length - 1];
    return new Request(`/account/${page}/${endpoint}/${lastSegment}`);
}
