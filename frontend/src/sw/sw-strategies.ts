import {
  CacheFirst,
  NetworkFirst,
  StaleWhileRevalidate,
} from "workbox-strategies";
import { ExpirationPlugin } from "workbox-expiration";
import { strategy as composeStrategies } from "workbox-streams";

export const contentCache = "contents-cache";
export const partialsCache = "partials-cache";
export const imagesCache = "images-cache";

export const cacheFirstStrategy = new CacheFirst({
  cacheName: partialsCache,
  plugins: [
    new ExpirationPlugin({
      maxEntries: 30,
      maxAgeSeconds: 60 * 60 * 24 * 30 * 3, // about 3 months
    }),
  ],
});

export const networkFirstStrategy = new NetworkFirst({
  cacheName: contentCache,
  networkTimeoutSeconds: 5,
  plugins: [
    new ExpirationPlugin({
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
export const staleWhileRevalidateStrategy = new StaleWhileRevalidate({
  cacheName: contentCache,
  plugins: [
    new ExpirationPlugin({
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

export const dashboardPageStream = composeStrategies(
  [
    async ({ event }) =>
      staleWhileRevalidateStrategy.handle({
        request: new Request("/account/dashboard/openingHtmlTag-Partial.html"),
        event,
      }),
    async ({ event }) =>
      networkFirstStrategy.handle({
        request: new Request("/account/dashboard/commonHeader-Partial.html"),
        event,
      }),
    async ({ event }) =>
      staleWhileRevalidateStrategy.handle({
        request: new Request("/account/dashboard/mainOpeningTag-Partial.html"),
        event,
      }),
    async ({ event }) =>
      staleWhileRevalidateStrategy.handle({
        request: new Request("/account/dashboard/swiperOne"),
        event,
      }),
    async ({ event }) =>
      networkFirstStrategy.handle({
        request: new Request("/account/dashboard/echarts?month=0"),
        event,
      }),
    async ({ event }) =>
      staleWhileRevalidateStrategy.handle({
        request: new Request("/account/dashboard/swiperTwo"),
        event,
      }),
    async ({ event }) =>
      cacheFirstStrategy.handle({
        request: new Request("/account/dashboard/footer-Partial.html"),
        event,
      }),
    async ({ event }) =>
      staleWhileRevalidateStrategy.handle({
        request: new Request("/account/dashboard/closingHtmlTag-Partial.html"),
        event,
      }),
  ],
  {}
);

export const locationPageStream = composeStrategies(
  [
    async ({ event }) =>
      staleWhileRevalidateStrategy.handle({
        request: new Request("/account/location/openingHtmlTag-Partial.html"),
        event,
      }),
    async ({ event }) =>
      staleWhileRevalidateStrategy.handle({
        request: new Request("/account/location/header-Partial.html"),
        event,
      }),
    async ({ event }) =>
      networkFirstStrategy.handle({
        request: new Request("/account/location/main", {
          headers: [["X-Content-Mode", "partials"]],
        }),
        event,
      }),
    async ({ event }) =>
      cacheFirstStrategy.handle({
        request: new Request("/account/location/footer-Partial.html"),
        event,
      }),
    async ({ event }) =>
      staleWhileRevalidateStrategy.handle({
        request: new Request("/account/location/closingHtmlTag-Partial.html"),
        event,
      }),
  ],
  {}
);

export const walletPageStream = composeStrategies(
  [
    async ({ event }) =>
      staleWhileRevalidateStrategy.handle({
        request: new Request("/account/wallet/openingHtmlTag-Partial.html"),
        event,
      }),
    async ({ event }) =>
      networkFirstStrategy.handle({
        request: new Request("/account/wallet/commonHeader-Partial.html"),
        event,
      }),
    async ({ event }) =>
      cacheFirstStrategy.handle({
        request: new Request("/account/wallet/mainOpeningTag-Partial.html"),
        event,
      }),
    async ({ event }) =>
      networkFirstStrategy.handle({
        request: new Request("/account/wallet/balance"),
        event,
      }),
    async ({ event }) =>
      staleWhileRevalidateStrategy.handle({
        request: new Request("/account/wallet/options"),
        event,
      }),
    async ({ event }) =>
      networkFirstStrategy.handle({
        request: new Request("/account/wallet/coinData", {}),
        event,
      }),
    async ({ event }) =>
      staleWhileRevalidateStrategy.handle({
        request: new Request("/account/wallet/swiper"),
        event,
      }),
    async ({ event }) =>
      cacheFirstStrategy.handle({
        request: new Request("/account/wallet/footer-Partial.html"),
        event,
      }),
    async ({ event }) =>
      staleWhileRevalidateStrategy.handle({
        request: new Request("/account/wallet/closingHtmlTag-Partial.html"),
        event,
      }),
  ],
  {}
);

export const spendPageStream = composeStrategies(
  [
    async ({ event }) =>
      staleWhileRevalidateStrategy.handle({
        request: new Request("/account/spend/openingHtmlTag-Partial.html"),
        event,
      }),
    async ({ event }) =>
      networkFirstStrategy.handle({
        request: new Request("/account/spend/header-Partial.html"),
        event,
      }),
    async ({ event }) =>
      cacheFirstStrategy.handle({
        request: new Request("/account/spend/mainOpeningTag-Partial.html"),
        event,
      }),
    async ({ event }) =>
      networkFirstStrategy.handle({
        request: new Request("/account/spend/balance"),
        event,
      }),
    async ({ event }) =>
      staleWhileRevalidateStrategy.handle({
        request: new Request("/account/spend/otherhalf"),
        event,
      }),
    async ({ event }) =>
      cacheFirstStrategy.handle({
        request: new Request("/account/spend/footer-Partial.html"),
        event,
      }),
    async ({ event }) =>
      staleWhileRevalidateStrategy.handle({
        request: new Request("/account/spend/closingHtmlTag-Partial.html"),
        event,
      }),
  ],
  {}
);

export const profilePageStream = composeStrategies(
  [
    async ({ event }) =>
      staleWhileRevalidateStrategy.handle({
        request: new Request("/account/profile/openingHtmlTag-Partial.html"),
        event,
      }),
    async ({ event }) =>
      networkFirstStrategy.handle({
        request: new Request("/account/profile/commonHeader-Partial.html"),
        event,
      }),
    async ({ event }) =>
      cacheFirstStrategy.handle({
        request: new Request("/account/profile/mainOpeningTag-Partial.html"),
        event,
      }),
    async ({ event }) =>
      networkFirstStrategy.handle({
        request: new Request("/account/profile/displayIcon"),
        event,
      }),
    async ({ event }) =>
      networkFirstStrategy.handle({
        request: new Request("/account/profile/formPartial"),
        event,
      }),
    async ({ event }) =>
      cacheFirstStrategy.handle({
        request: new Request("/account/profile/footer-Partial.html"),
        event,
      }),
    async ({ event }) =>
      staleWhileRevalidateStrategy.handle({
        request: new Request("/account/profile/closingHtmlTag-Partial.html"),
        event,
      }),
  ],
  {}
);

export const airtimePageStream = composeStrategies(
  [
    async ({ event }) =>
      staleWhileRevalidateStrategy.handle({
        request: new Request("/account/airtime/openingHtmlTag-Partial.html"),
        event,
      }),
    async ({ event }) =>
      networkFirstStrategy.handle({
        request: new Request("/account/airtime/header-Partial.html"),
        event,
      }),
    async ({ event }) =>
      cacheFirstStrategy.handle({
        request: new Request("/account/airtime/mainOpeningTag-Partial.html"),
        event,
      }),
    async ({ event }) =>
      staleWhileRevalidateStrategy.handle({
        request: new Request("/account/airtime/inputNumber"),
        event,
      }),
    async ({ event }) =>
      networkFirstStrategy.handle({
        request: new Request("/account/airtime/beneficiaries"),
        event,
      }),
    async ({ event }) =>
      cacheFirstStrategy.handle({
        request: new Request("/account/airtime/footer-Partial.html"),
        event,
      }),
    async ({ event }) =>
      staleWhileRevalidateStrategy.handle({
        request: new Request("/account/airtime/closingHtmlTag-Partial.html"),
        event,
      }),
  ],
  {}
);

export const beneficiariesPageStream = composeStrategies(
  [
    async ({ event }) =>
      staleWhileRevalidateStrategy.handle({
        request: new Request(
          "/account/beneficiaries/openingHtmlTag-Partial.html"
        ),
        event,
      }),
    async ({ event, url }) =>
      networkFirstStrategy.handle({
        request: new Request("/account/beneficiaries/header-Partial.html"),
        event,
      }),
    async ({ event }) =>
      networkFirstStrategy.handle({
        request: new Request("/account/beneficiaries/main"),
        event,
      }),
    async ({ event }) =>
      cacheFirstStrategy.handle({
        request: new Request("/account/beneficiaries/footer-Partial.html"),
        event,
      }),
    async ({ event }) =>
      staleWhileRevalidateStrategy.handle({
        request: new Request(
          "/account/beneficiaries/closingHtmlTag-Partial.html"
        ),
        event,
      }),
  ],
  {}
);

export const wastePageStream = composeStrategies(
  [
    async ({ event }) =>
      staleWhileRevalidateStrategy.handle({
        request: new Request("/account/waste/openingHtmlTag-Partial.html"),
        event,
      }),
    async ({ event }) =>
      networkFirstStrategy.handle({
        request: new Request("/account/waste/commonHeader-Partial.html"),
        event,
      }),
    async ({ event }) =>
      cacheFirstStrategy.handle({
        request: new Request("/account/waste/mainOpeningTag-Partial.html"),
        event,
      }),
    async ({ event, url }) =>
      networkFirstStrategy.handle({
        request: matchingPathVariableForWaste(url, "quickInfo", "waste"),
        event,
      }),
    async ({ event }) =>
      staleWhileRevalidateStrategy.handle({
        request: new Request("/account/waste/filter"),
        event,
      }),
    async ({ event, url }) =>
      networkFirstStrategy.handle({
        request: matchingPathVariableForWaste(url, "history", "waste"),
        event,
      }),
    async ({ event }) =>
      cacheFirstStrategy.handle({
        request: new Request("/account/waste/footer-Partial.html"),
        event,
      }),
    async ({ event }) =>
      staleWhileRevalidateStrategy.handle({
        request: new Request("/account/waste/closingHtmlTag-Partial.html"),
        event,
      }),
  ],
  {}
);

export const tipsPageStream = composeStrategies(
  [
    async ({ event }) =>
      staleWhileRevalidateStrategy.handle({
        request: new Request("/account/tips/openingHtmlTag-Partial.html"),
        event,
      }),
    async ({ event }) =>
      staleWhileRevalidateStrategy.handle({
        request: new Request("/account/tips/tipsHeader-Partial.html"),
        event,
      }),
    async ({ event }) =>
      cacheFirstStrategy.handle({
        request: new Request("/account/tips/mainOpeningTag-Partial.html"),
        event,
      }),
    async ({ event, url }) =>
      networkFirstStrategy.handle({
        request: matchingPathVariableForWaste(url, "quickInfo", "tips"),
        event,
      }),
    async ({ event, url }) =>
      networkFirstStrategy.handle({
        request: matchingPathVariableForWaste(url, "article", "tips"),
        event,
      }),
    async ({ event }) =>
      cacheFirstStrategy.handle({
        request: new Request("/account/tips/footer-Partial.html"),
        event,
      }),
    async ({ event }) =>
      staleWhileRevalidateStrategy.handle({
        request: new Request("/account/tips/closingHtmlTag-Partial.html"),
        event,
      }),
  ],
  {}
);

function matchingPathVariableForWaste(
  url: URL,
  endpoint: string,
  page: string
) {
  const segments = url.pathname.split("/");
  const lastSegment = segments[segments.length - 1];
  return new Request(`/account/${page}/${endpoint}/${lastSegment}`);
}
