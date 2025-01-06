"use strict";
(() => {
  // node_modules/.pnpm/workbox-core@7.3.0/node_modules/workbox-core/_version.js
  try {
    self["workbox:core:7.2.0"] && _();
  } catch (e) {
  }

  // node_modules/.pnpm/workbox-core@7.3.0/node_modules/workbox-core/_private/logger.js
  var logger = false ? null : (() => {
    if (!("__WB_DISABLE_DEV_LOGS" in globalThis)) {
      self.__WB_DISABLE_DEV_LOGS = false;
    }
    let inGroup = false;
    const methodToColorMap = {
      debug: `#7f8c8d`,
      log: `#2ecc71`,
      warn: `#f39c12`,
      error: `#c0392b`,
      groupCollapsed: `#3498db`,
      groupEnd: null
      // No colored prefix on groupEnd
    };
    const print = function(method, args) {
      if (self.__WB_DISABLE_DEV_LOGS) {
        return;
      }
      if (method === "groupCollapsed") {
        if (/^((?!chrome|android).)*safari/i.test(navigator.userAgent)) {
          console[method](...args);
          return;
        }
      }
      const styles = [
        `background: ${methodToColorMap[method]}`,
        `border-radius: 0.5em`,
        `color: white`,
        `font-weight: bold`,
        `padding: 2px 0.5em`
      ];
      const logPrefix = inGroup ? [] : ["%cworkbox", styles.join(";")];
      console[method](...logPrefix, ...args);
      if (method === "groupCollapsed") {
        inGroup = true;
      }
      if (method === "groupEnd") {
        inGroup = false;
      }
    };
    const api = {};
    const loggerMethods = Object.keys(methodToColorMap);
    for (const key of loggerMethods) {
      const method = key;
      api[method] = (...args) => {
        print(method, args);
      };
    }
    return api;
  })();

  // node_modules/.pnpm/workbox-navigation-preload@7.3.0/node_modules/workbox-navigation-preload/_version.js
  try {
    self["workbox:navigation-preload:7.2.0"] && _();
  } catch (e) {
  }

  // node_modules/.pnpm/workbox-navigation-preload@7.3.0/node_modules/workbox-navigation-preload/isSupported.js
  function isSupported() {
    return Boolean(self.registration && self.registration.navigationPreload);
  }

  // node_modules/.pnpm/workbox-navigation-preload@7.3.0/node_modules/workbox-navigation-preload/enable.js
  function enable(headerValue) {
    if (isSupported()) {
      self.addEventListener("activate", (event) => {
        event.waitUntil(self.registration.navigationPreload.enable().then(() => {
          if (headerValue) {
            void self.registration.navigationPreload.setHeaderValue(headerValue);
          }
          if (true) {
            logger.log(`Navigation preload is enabled.`);
          }
        }));
      });
    } else {
      if (true) {
        logger.log(`Navigation preload is not supported in this browser.`);
      }
    }
  }

  // node_modules/.pnpm/workbox-core@7.3.0/node_modules/workbox-core/models/messages/messages.js
  var messages = {
    "invalid-value": ({ paramName, validValueDescription, value }) => {
      if (!paramName || !validValueDescription) {
        throw new Error(`Unexpected input to 'invalid-value' error.`);
      }
      return `The '${paramName}' parameter was given a value with an unexpected value. ${validValueDescription} Received a value of ${JSON.stringify(value)}.`;
    },
    "not-an-array": ({ moduleName, className, funcName, paramName }) => {
      if (!moduleName || !className || !funcName || !paramName) {
        throw new Error(`Unexpected input to 'not-an-array' error.`);
      }
      return `The parameter '${paramName}' passed into '${moduleName}.${className}.${funcName}()' must be an array.`;
    },
    "incorrect-type": ({ expectedType, paramName, moduleName, className, funcName }) => {
      if (!expectedType || !paramName || !moduleName || !funcName) {
        throw new Error(`Unexpected input to 'incorrect-type' error.`);
      }
      const classNameStr = className ? `${className}.` : "";
      return `The parameter '${paramName}' passed into '${moduleName}.${classNameStr}${funcName}()' must be of type ${expectedType}.`;
    },
    "incorrect-class": ({ expectedClassName, paramName, moduleName, className, funcName, isReturnValueProblem }) => {
      if (!expectedClassName || !moduleName || !funcName) {
        throw new Error(`Unexpected input to 'incorrect-class' error.`);
      }
      const classNameStr = className ? `${className}.` : "";
      if (isReturnValueProblem) {
        return `The return value from '${moduleName}.${classNameStr}${funcName}()' must be an instance of class ${expectedClassName}.`;
      }
      return `The parameter '${paramName}' passed into '${moduleName}.${classNameStr}${funcName}()' must be an instance of class ${expectedClassName}.`;
    },
    "missing-a-method": ({ expectedMethod, paramName, moduleName, className, funcName }) => {
      if (!expectedMethod || !paramName || !moduleName || !className || !funcName) {
        throw new Error(`Unexpected input to 'missing-a-method' error.`);
      }
      return `${moduleName}.${className}.${funcName}() expected the '${paramName}' parameter to expose a '${expectedMethod}' method.`;
    },
    "add-to-cache-list-unexpected-type": ({ entry }) => {
      return `An unexpected entry was passed to 'workbox-precaching.PrecacheController.addToCacheList()' The entry '${JSON.stringify(entry)}' isn't supported. You must supply an array of strings with one or more characters, objects with a url property or Request objects.`;
    },
    "add-to-cache-list-conflicting-entries": ({ firstEntry, secondEntry }) => {
      if (!firstEntry || !secondEntry) {
        throw new Error(`Unexpected input to 'add-to-cache-list-duplicate-entries' error.`);
      }
      return `Two of the entries passed to 'workbox-precaching.PrecacheController.addToCacheList()' had the URL ${firstEntry} but different revision details. Workbox is unable to cache and version the asset correctly. Please remove one of the entries.`;
    },
    "plugin-error-request-will-fetch": ({ thrownErrorMessage }) => {
      if (!thrownErrorMessage) {
        throw new Error(`Unexpected input to 'plugin-error-request-will-fetch', error.`);
      }
      return `An error was thrown by a plugins 'requestWillFetch()' method. The thrown error message was: '${thrownErrorMessage}'.`;
    },
    "invalid-cache-name": ({ cacheNameId, value }) => {
      if (!cacheNameId) {
        throw new Error(`Expected a 'cacheNameId' for error 'invalid-cache-name'`);
      }
      return `You must provide a name containing at least one character for setCacheDetails({${cacheNameId}: '...'}). Received a value of '${JSON.stringify(value)}'`;
    },
    "unregister-route-but-not-found-with-method": ({ method }) => {
      if (!method) {
        throw new Error(`Unexpected input to 'unregister-route-but-not-found-with-method' error.`);
      }
      return `The route you're trying to unregister was not  previously registered for the method type '${method}'.`;
    },
    "unregister-route-route-not-registered": () => {
      return `The route you're trying to unregister was not previously registered.`;
    },
    "queue-replay-failed": ({ name }) => {
      return `Replaying the background sync queue '${name}' failed.`;
    },
    "duplicate-queue-name": ({ name }) => {
      return `The Queue name '${name}' is already being used. All instances of backgroundSync.Queue must be given unique names.`;
    },
    "expired-test-without-max-age": ({ methodName, paramName }) => {
      return `The '${methodName}()' method can only be used when the '${paramName}' is used in the constructor.`;
    },
    "unsupported-route-type": ({ moduleName, className, funcName, paramName }) => {
      return `The supplied '${paramName}' parameter was an unsupported type. Please check the docs for ${moduleName}.${className}.${funcName} for valid input types.`;
    },
    "not-array-of-class": ({ value, expectedClass, moduleName, className, funcName, paramName }) => {
      return `The supplied '${paramName}' parameter must be an array of '${expectedClass}' objects. Received '${JSON.stringify(value)},'. Please check the call to ${moduleName}.${className}.${funcName}() to fix the issue.`;
    },
    "max-entries-or-age-required": ({ moduleName, className, funcName }) => {
      return `You must define either config.maxEntries or config.maxAgeSecondsin ${moduleName}.${className}.${funcName}`;
    },
    "statuses-or-headers-required": ({ moduleName, className, funcName }) => {
      return `You must define either config.statuses or config.headersin ${moduleName}.${className}.${funcName}`;
    },
    "invalid-string": ({ moduleName, funcName, paramName }) => {
      if (!paramName || !moduleName || !funcName) {
        throw new Error(`Unexpected input to 'invalid-string' error.`);
      }
      return `When using strings, the '${paramName}' parameter must start with 'http' (for cross-origin matches) or '/' (for same-origin matches). Please see the docs for ${moduleName}.${funcName}() for more info.`;
    },
    "channel-name-required": () => {
      return `You must provide a channelName to construct a BroadcastCacheUpdate instance.`;
    },
    "invalid-responses-are-same-args": () => {
      return `The arguments passed into responsesAreSame() appear to be invalid. Please ensure valid Responses are used.`;
    },
    "expire-custom-caches-only": () => {
      return `You must provide a 'cacheName' property when using the expiration plugin with a runtime caching strategy.`;
    },
    "unit-must-be-bytes": ({ normalizedRangeHeader }) => {
      if (!normalizedRangeHeader) {
        throw new Error(`Unexpected input to 'unit-must-be-bytes' error.`);
      }
      return `The 'unit' portion of the Range header must be set to 'bytes'. The Range header provided was "${normalizedRangeHeader}"`;
    },
    "single-range-only": ({ normalizedRangeHeader }) => {
      if (!normalizedRangeHeader) {
        throw new Error(`Unexpected input to 'single-range-only' error.`);
      }
      return `Multiple ranges are not supported. Please use a  single start value, and optional end value. The Range header provided was "${normalizedRangeHeader}"`;
    },
    "invalid-range-values": ({ normalizedRangeHeader }) => {
      if (!normalizedRangeHeader) {
        throw new Error(`Unexpected input to 'invalid-range-values' error.`);
      }
      return `The Range header is missing both start and end values. At least one of those values is needed. The Range header provided was "${normalizedRangeHeader}"`;
    },
    "no-range-header": () => {
      return `No Range header was found in the Request provided.`;
    },
    "range-not-satisfiable": ({ size, start, end }) => {
      return `The start (${start}) and end (${end}) values in the Range are not satisfiable by the cached response, which is ${size} bytes.`;
    },
    "attempt-to-cache-non-get-request": ({ url, method }) => {
      return `Unable to cache '${url}' because it is a '${method}' request and only 'GET' requests can be cached.`;
    },
    "cache-put-with-no-response": ({ url }) => {
      return `There was an attempt to cache '${url}' but the response was not defined.`;
    },
    "no-response": ({ url, error }) => {
      let message = `The strategy could not generate a response for '${url}'.`;
      if (error) {
        message += ` The underlying error is ${error}.`;
      }
      return message;
    },
    "bad-precaching-response": ({ url, status }) => {
      return `The precaching request for '${url}' failed` + (status ? ` with an HTTP status of ${status}.` : `.`);
    },
    "non-precached-url": ({ url }) => {
      return `createHandlerBoundToURL('${url}') was called, but that URL is not precached. Please pass in a URL that is precached instead.`;
    },
    "add-to-cache-list-conflicting-integrities": ({ url }) => {
      return `Two of the entries passed to 'workbox-precaching.PrecacheController.addToCacheList()' had the URL ${url} with different integrity values. Please remove one of them.`;
    },
    "missing-precache-entry": ({ cacheName, url }) => {
      return `Unable to find a precached response in ${cacheName} for ${url}.`;
    },
    "cross-origin-copy-response": ({ origin }) => {
      return `workbox-core.copyResponse() can only be used with same-origin responses. It was passed a response with origin ${origin}.`;
    },
    "opaque-streams-source": ({ type }) => {
      const message = `One of the workbox-streams sources resulted in an '${type}' response.`;
      if (type === "opaqueredirect") {
        return `${message} Please do not use a navigation request that results in a redirect as a source.`;
      }
      return `${message} Please ensure your sources are CORS-enabled.`;
    }
  };

  // node_modules/.pnpm/workbox-core@7.3.0/node_modules/workbox-core/models/messages/messageGenerator.js
  var generatorFunction = (code, details = {}) => {
    const message = messages[code];
    if (!message) {
      throw new Error(`Unable to find message for code '${code}'.`);
    }
    return message(details);
  };
  var messageGenerator = false ? fallback : generatorFunction;

  // node_modules/.pnpm/workbox-core@7.3.0/node_modules/workbox-core/_private/WorkboxError.js
  var WorkboxError = class extends Error {
    /**
     *
     * @param {string} errorCode The error code that
     * identifies this particular error.
     * @param {Object=} details Any relevant arguments
     * that will help developers identify issues should
     * be added as a key on the context object.
     */
    constructor(errorCode, details) {
      const message = messageGenerator(errorCode, details);
      super(message);
      this.name = errorCode;
      this.details = details;
    }
  };

  // node_modules/.pnpm/workbox-core@7.3.0/node_modules/workbox-core/_private/assert.js
  var isArray = (value, details) => {
    if (!Array.isArray(value)) {
      throw new WorkboxError("not-an-array", details);
    }
  };
  var hasMethod = (object, expectedMethod, details) => {
    const type = typeof object[expectedMethod];
    if (type !== "function") {
      details["expectedMethod"] = expectedMethod;
      throw new WorkboxError("missing-a-method", details);
    }
  };
  var isType = (object, expectedType, details) => {
    if (typeof object !== expectedType) {
      details["expectedType"] = expectedType;
      throw new WorkboxError("incorrect-type", details);
    }
  };
  var isInstance = (object, expectedClass, details) => {
    if (!(object instanceof expectedClass)) {
      details["expectedClassName"] = expectedClass.name;
      throw new WorkboxError("incorrect-class", details);
    }
  };
  var isOneOf = (value, validValues, details) => {
    if (!validValues.includes(value)) {
      details["validValueDescription"] = `Valid values are ${JSON.stringify(validValues)}.`;
      throw new WorkboxError("invalid-value", details);
    }
  };
  var isArrayOfClass = (value, expectedClass, details) => {
    const error = new WorkboxError("not-array-of-class", details);
    if (!Array.isArray(value)) {
      throw error;
    }
    for (const item of value) {
      if (!(item instanceof expectedClass)) {
        throw error;
      }
    }
  };
  var finalAssertExports = false ? null : {
    hasMethod,
    isArray,
    isInstance,
    isOneOf,
    isType,
    isArrayOfClass
  };

  // node_modules/.pnpm/workbox-core@7.3.0/node_modules/workbox-core/_private/cacheNames.js
  var _cacheNameDetails = {
    googleAnalytics: "googleAnalytics",
    precache: "precache-v2",
    prefix: "workbox",
    runtime: "runtime",
    suffix: typeof registration !== "undefined" ? registration.scope : ""
  };
  var _createCacheName = (cacheName) => {
    return [_cacheNameDetails.prefix, cacheName, _cacheNameDetails.suffix].filter((value) => value && value.length > 0).join("-");
  };
  var eachCacheNameDetail = (fn) => {
    for (const key of Object.keys(_cacheNameDetails)) {
      fn(key);
    }
  };
  var cacheNames = {
    updateDetails: (details) => {
      eachCacheNameDetail((key) => {
        if (typeof details[key] === "string") {
          _cacheNameDetails[key] = details[key];
        }
      });
    },
    getGoogleAnalyticsName: (userCacheName) => {
      return userCacheName || _createCacheName(_cacheNameDetails.googleAnalytics);
    },
    getPrecacheName: (userCacheName) => {
      return userCacheName || _createCacheName(_cacheNameDetails.precache);
    },
    getPrefix: () => {
      return _cacheNameDetails.prefix;
    },
    getRuntimeName: (userCacheName) => {
      return userCacheName || _createCacheName(_cacheNameDetails.runtime);
    },
    getSuffix: () => {
      return _cacheNameDetails.suffix;
    }
  };

  // node_modules/.pnpm/workbox-core@7.3.0/node_modules/workbox-core/_private/getFriendlyURL.js
  var getFriendlyURL = (url) => {
    const urlObj = new URL(String(url), location.href);
    return urlObj.href.replace(new RegExp(`^${location.origin}`), "");
  };

  // node_modules/.pnpm/workbox-core@7.3.0/node_modules/workbox-core/_private/cacheMatchIgnoreParams.js
  function stripParams(fullURL, ignoreParams) {
    const strippedURL = new URL(fullURL);
    for (const param of ignoreParams) {
      strippedURL.searchParams.delete(param);
    }
    return strippedURL.href;
  }
  async function cacheMatchIgnoreParams(cache, request, ignoreParams, matchOptions) {
    const strippedRequestURL = stripParams(request.url, ignoreParams);
    if (request.url === strippedRequestURL) {
      return cache.match(request, matchOptions);
    }
    const keysOptions = Object.assign(Object.assign({}, matchOptions), { ignoreSearch: true });
    const cacheKeys = await cache.keys(request, keysOptions);
    for (const cacheKey of cacheKeys) {
      const strippedCacheKeyURL = stripParams(cacheKey.url, ignoreParams);
      if (strippedRequestURL === strippedCacheKeyURL) {
        return cache.match(cacheKey, matchOptions);
      }
    }
    return;
  }

  // node_modules/.pnpm/workbox-core@7.3.0/node_modules/workbox-core/_private/Deferred.js
  var Deferred = class {
    /**
     * Creates a promise and exposes its resolve and reject functions as methods.
     */
    constructor() {
      this.promise = new Promise((resolve, reject) => {
        this.resolve = resolve;
        this.reject = reject;
      });
    }
  };

  // node_modules/.pnpm/workbox-core@7.3.0/node_modules/workbox-core/models/quotaErrorCallbacks.js
  var quotaErrorCallbacks = /* @__PURE__ */ new Set();

  // node_modules/.pnpm/workbox-core@7.3.0/node_modules/workbox-core/_private/executeQuotaErrorCallbacks.js
  async function executeQuotaErrorCallbacks() {
    if (true) {
      logger.log(`About to run ${quotaErrorCallbacks.size} callbacks to clean up caches.`);
    }
    for (const callback of quotaErrorCallbacks) {
      await callback();
      if (true) {
        logger.log(callback, "is complete.");
      }
    }
    if (true) {
      logger.log("Finished running callbacks.");
    }
  }

  // node_modules/.pnpm/workbox-core@7.3.0/node_modules/workbox-core/_private/timeout.js
  function timeout(ms) {
    return new Promise((resolve) => setTimeout(resolve, ms));
  }

  // node_modules/.pnpm/workbox-strategies@7.3.0/node_modules/workbox-strategies/_version.js
  try {
    self["workbox:strategies:7.2.0"] && _();
  } catch (e) {
  }

  // node_modules/.pnpm/workbox-strategies@7.3.0/node_modules/workbox-strategies/StrategyHandler.js
  function toRequest(input) {
    return typeof input === "string" ? new Request(input) : input;
  }
  var StrategyHandler = class {
    /**
     * Creates a new instance associated with the passed strategy and event
     * that's handling the request.
     *
     * The constructor also initializes the state that will be passed to each of
     * the plugins handling this request.
     *
     * @param {workbox-strategies.Strategy} strategy
     * @param {Object} options
     * @param {Request|string} options.request A request to run this strategy for.
     * @param {ExtendableEvent} options.event The event associated with the
     *     request.
     * @param {URL} [options.url]
     * @param {*} [options.params] The return value from the
     *     {@link workbox-routing~matchCallback} (if applicable).
     */
    constructor(strategy2, options) {
      this._cacheKeys = {};
      if (true) {
        finalAssertExports.isInstance(options.event, ExtendableEvent, {
          moduleName: "workbox-strategies",
          className: "StrategyHandler",
          funcName: "constructor",
          paramName: "options.event"
        });
      }
      Object.assign(this, options);
      this.event = options.event;
      this._strategy = strategy2;
      this._handlerDeferred = new Deferred();
      this._extendLifetimePromises = [];
      this._plugins = [...strategy2.plugins];
      this._pluginStateMap = /* @__PURE__ */ new Map();
      for (const plugin of this._plugins) {
        this._pluginStateMap.set(plugin, {});
      }
      this.event.waitUntil(this._handlerDeferred.promise);
    }
    /**
     * Fetches a given request (and invokes any applicable plugin callback
     * methods) using the `fetchOptions` (for non-navigation requests) and
     * `plugins` defined on the `Strategy` object.
     *
     * The following plugin lifecycle methods are invoked when using this method:
     * - `requestWillFetch()`
     * - `fetchDidSucceed()`
     * - `fetchDidFail()`
     *
     * @param {Request|string} input The URL or request to fetch.
     * @return {Promise<Response>}
     */
    async fetch(input) {
      const { event } = this;
      let request = toRequest(input);
      if (request.mode === "navigate" && event instanceof FetchEvent && event.preloadResponse) {
        const possiblePreloadResponse = await event.preloadResponse;
        if (possiblePreloadResponse) {
          if (true) {
            logger.log(`Using a preloaded navigation response for '${getFriendlyURL(request.url)}'`);
          }
          return possiblePreloadResponse;
        }
      }
      const originalRequest = this.hasCallback("fetchDidFail") ? request.clone() : null;
      try {
        for (const cb of this.iterateCallbacks("requestWillFetch")) {
          request = await cb({ request: request.clone(), event });
        }
      } catch (err) {
        if (err instanceof Error) {
          throw new WorkboxError("plugin-error-request-will-fetch", {
            thrownErrorMessage: err.message
          });
        }
      }
      const pluginFilteredRequest = request.clone();
      try {
        let fetchResponse;
        fetchResponse = await fetch(request, request.mode === "navigate" ? void 0 : this._strategy.fetchOptions);
        if (true) {
          logger.debug(`Network request for '${getFriendlyURL(request.url)}' returned a response with status '${fetchResponse.status}'.`);
        }
        for (const callback of this.iterateCallbacks("fetchDidSucceed")) {
          fetchResponse = await callback({
            event,
            request: pluginFilteredRequest,
            response: fetchResponse
          });
        }
        return fetchResponse;
      } catch (error) {
        if (true) {
          logger.log(`Network request for '${getFriendlyURL(request.url)}' threw an error.`, error);
        }
        if (originalRequest) {
          await this.runCallbacks("fetchDidFail", {
            error,
            event,
            originalRequest: originalRequest.clone(),
            request: pluginFilteredRequest.clone()
          });
        }
        throw error;
      }
    }
    /**
     * Calls `this.fetch()` and (in the background) runs `this.cachePut()` on
     * the response generated by `this.fetch()`.
     *
     * The call to `this.cachePut()` automatically invokes `this.waitUntil()`,
     * so you do not have to manually call `waitUntil()` on the event.
     *
     * @param {Request|string} input The request or URL to fetch and cache.
     * @return {Promise<Response>}
     */
    async fetchAndCachePut(input) {
      const response = await this.fetch(input);
      const responseClone = response.clone();
      void this.waitUntil(this.cachePut(input, responseClone));
      return response;
    }
    /**
     * Matches a request from the cache (and invokes any applicable plugin
     * callback methods) using the `cacheName`, `matchOptions`, and `plugins`
     * defined on the strategy object.
     *
     * The following plugin lifecycle methods are invoked when using this method:
     * - cacheKeyWillBeUsed()
     * - cachedResponseWillBeUsed()
     *
     * @param {Request|string} key The Request or URL to use as the cache key.
     * @return {Promise<Response|undefined>} A matching response, if found.
     */
    async cacheMatch(key) {
      const request = toRequest(key);
      let cachedResponse;
      const { cacheName, matchOptions } = this._strategy;
      const effectiveRequest = await this.getCacheKey(request, "read");
      const multiMatchOptions = Object.assign(Object.assign({}, matchOptions), { cacheName });
      cachedResponse = await caches.match(effectiveRequest, multiMatchOptions);
      if (true) {
        if (cachedResponse) {
          logger.debug(`Found a cached response in '${cacheName}'.`);
        } else {
          logger.debug(`No cached response found in '${cacheName}'.`);
        }
      }
      for (const callback of this.iterateCallbacks("cachedResponseWillBeUsed")) {
        cachedResponse = await callback({
          cacheName,
          matchOptions,
          cachedResponse,
          request: effectiveRequest,
          event: this.event
        }) || void 0;
      }
      return cachedResponse;
    }
    /**
     * Puts a request/response pair in the cache (and invokes any applicable
     * plugin callback methods) using the `cacheName` and `plugins` defined on
     * the strategy object.
     *
     * The following plugin lifecycle methods are invoked when using this method:
     * - cacheKeyWillBeUsed()
     * - cacheWillUpdate()
     * - cacheDidUpdate()
     *
     * @param {Request|string} key The request or URL to use as the cache key.
     * @param {Response} response The response to cache.
     * @return {Promise<boolean>} `false` if a cacheWillUpdate caused the response
     * not be cached, and `true` otherwise.
     */
    async cachePut(key, response) {
      const request = toRequest(key);
      await timeout(0);
      const effectiveRequest = await this.getCacheKey(request, "write");
      if (true) {
        if (effectiveRequest.method && effectiveRequest.method !== "GET") {
          throw new WorkboxError("attempt-to-cache-non-get-request", {
            url: getFriendlyURL(effectiveRequest.url),
            method: effectiveRequest.method
          });
        }
        const vary = response.headers.get("Vary");
        if (vary) {
          logger.debug(`The response for ${getFriendlyURL(effectiveRequest.url)} has a 'Vary: ${vary}' header. Consider setting the {ignoreVary: true} option on your strategy to ensure cache matching and deletion works as expected.`);
        }
      }
      if (!response) {
        if (true) {
          logger.error(`Cannot cache non-existent response for '${getFriendlyURL(effectiveRequest.url)}'.`);
        }
        throw new WorkboxError("cache-put-with-no-response", {
          url: getFriendlyURL(effectiveRequest.url)
        });
      }
      const responseToCache = await this._ensureResponseSafeToCache(response);
      if (!responseToCache) {
        if (true) {
          logger.debug(`Response '${getFriendlyURL(effectiveRequest.url)}' will not be cached.`, responseToCache);
        }
        return false;
      }
      const { cacheName, matchOptions } = this._strategy;
      const cache = await self.caches.open(cacheName);
      const hasCacheUpdateCallback = this.hasCallback("cacheDidUpdate");
      const oldResponse = hasCacheUpdateCallback ? await cacheMatchIgnoreParams(
        // TODO(philipwalton): the `__WB_REVISION__` param is a precaching
        // feature. Consider into ways to only add this behavior if using
        // precaching.
        cache,
        effectiveRequest.clone(),
        ["__WB_REVISION__"],
        matchOptions
      ) : null;
      if (true) {
        logger.debug(`Updating the '${cacheName}' cache with a new Response for ${getFriendlyURL(effectiveRequest.url)}.`);
      }
      try {
        await cache.put(effectiveRequest, hasCacheUpdateCallback ? responseToCache.clone() : responseToCache);
      } catch (error) {
        if (error instanceof Error) {
          if (error.name === "QuotaExceededError") {
            await executeQuotaErrorCallbacks();
          }
          throw error;
        }
      }
      for (const callback of this.iterateCallbacks("cacheDidUpdate")) {
        await callback({
          cacheName,
          oldResponse,
          newResponse: responseToCache.clone(),
          request: effectiveRequest,
          event: this.event
        });
      }
      return true;
    }
    /**
     * Checks the list of plugins for the `cacheKeyWillBeUsed` callback, and
     * executes any of those callbacks found in sequence. The final `Request`
     * object returned by the last plugin is treated as the cache key for cache
     * reads and/or writes. If no `cacheKeyWillBeUsed` plugin callbacks have
     * been registered, the passed request is returned unmodified
     *
     * @param {Request} request
     * @param {string} mode
     * @return {Promise<Request>}
     */
    async getCacheKey(request, mode) {
      const key = `${request.url} | ${mode}`;
      if (!this._cacheKeys[key]) {
        let effectiveRequest = request;
        for (const callback of this.iterateCallbacks("cacheKeyWillBeUsed")) {
          effectiveRequest = toRequest(await callback({
            mode,
            request: effectiveRequest,
            event: this.event,
            // params has a type any can't change right now.
            params: this.params
            // eslint-disable-line
          }));
        }
        this._cacheKeys[key] = effectiveRequest;
      }
      return this._cacheKeys[key];
    }
    /**
     * Returns true if the strategy has at least one plugin with the given
     * callback.
     *
     * @param {string} name The name of the callback to check for.
     * @return {boolean}
     */
    hasCallback(name) {
      for (const plugin of this._strategy.plugins) {
        if (name in plugin) {
          return true;
        }
      }
      return false;
    }
    /**
     * Runs all plugin callbacks matching the given name, in order, passing the
     * given param object (merged ith the current plugin state) as the only
     * argument.
     *
     * Note: since this method runs all plugins, it's not suitable for cases
     * where the return value of a callback needs to be applied prior to calling
     * the next callback. See
     * {@link workbox-strategies.StrategyHandler#iterateCallbacks}
     * below for how to handle that case.
     *
     * @param {string} name The name of the callback to run within each plugin.
     * @param {Object} param The object to pass as the first (and only) param
     *     when executing each callback. This object will be merged with the
     *     current plugin state prior to callback execution.
     */
    async runCallbacks(name, param) {
      for (const callback of this.iterateCallbacks(name)) {
        await callback(param);
      }
    }
    /**
     * Accepts a callback and returns an iterable of matching plugin callbacks,
     * where each callback is wrapped with the current handler state (i.e. when
     * you call each callback, whatever object parameter you pass it will
     * be merged with the plugin's current state).
     *
     * @param {string} name The name fo the callback to run
     * @return {Array<Function>}
     */
    *iterateCallbacks(name) {
      for (const plugin of this._strategy.plugins) {
        if (typeof plugin[name] === "function") {
          const state = this._pluginStateMap.get(plugin);
          const statefulCallback = (param) => {
            const statefulParam = Object.assign(Object.assign({}, param), { state });
            return plugin[name](statefulParam);
          };
          yield statefulCallback;
        }
      }
    }
    /**
     * Adds a promise to the
     * [extend lifetime promises]{@link https://w3c.github.io/ServiceWorker/#extendableevent-extend-lifetime-promises}
     * of the event event associated with the request being handled (usually a
     * `FetchEvent`).
     *
     * Note: you can await
     * {@link workbox-strategies.StrategyHandler~doneWaiting}
     * to know when all added promises have settled.
     *
     * @param {Promise} promise A promise to add to the extend lifetime promises
     *     of the event that triggered the request.
     */
    waitUntil(promise) {
      this._extendLifetimePromises.push(promise);
      return promise;
    }
    /**
     * Returns a promise that resolves once all promises passed to
     * {@link workbox-strategies.StrategyHandler~waitUntil}
     * have settled.
     *
     * Note: any work done after `doneWaiting()` settles should be manually
     * passed to an event's `waitUntil()` method (not this handler's
     * `waitUntil()` method), otherwise the service worker thread my be killed
     * prior to your work completing.
     */
    async doneWaiting() {
      let promise;
      while (promise = this._extendLifetimePromises.shift()) {
        await promise;
      }
    }
    /**
     * Stops running the strategy and immediately resolves any pending
     * `waitUntil()` promises.
     */
    destroy() {
      this._handlerDeferred.resolve(null);
    }
    /**
     * This method will call cacheWillUpdate on the available plugins (or use
     * status === 200) to determine if the Response is safe and valid to cache.
     *
     * @param {Request} options.request
     * @param {Response} options.response
     * @return {Promise<Response|undefined>}
     *
     * @private
     */
    async _ensureResponseSafeToCache(response) {
      let responseToCache = response;
      let pluginsUsed = false;
      for (const callback of this.iterateCallbacks("cacheWillUpdate")) {
        responseToCache = await callback({
          request: this.request,
          response: responseToCache,
          event: this.event
        }) || void 0;
        pluginsUsed = true;
        if (!responseToCache) {
          break;
        }
      }
      if (!pluginsUsed) {
        if (responseToCache && responseToCache.status !== 200) {
          responseToCache = void 0;
        }
        if (true) {
          if (responseToCache) {
            if (responseToCache.status !== 200) {
              if (responseToCache.status === 0) {
                logger.warn(`The response for '${this.request.url}' is an opaque response. The caching strategy that you're using will not cache opaque responses by default.`);
              } else {
                logger.debug(`The response for '${this.request.url}' returned a status code of '${response.status}' and won't be cached as a result.`);
              }
            }
          }
        }
      }
      return responseToCache;
    }
  };

  // node_modules/.pnpm/workbox-strategies@7.3.0/node_modules/workbox-strategies/Strategy.js
  var Strategy = class {
    /**
     * Creates a new instance of the strategy and sets all documented option
     * properties as public instance properties.
     *
     * Note: if a custom strategy class extends the base Strategy class and does
     * not need more than these properties, it does not need to define its own
     * constructor.
     *
     * @param {Object} [options]
     * @param {string} [options.cacheName] Cache name to store and retrieve
     * requests. Defaults to the cache names provided by
     * {@link workbox-core.cacheNames}.
     * @param {Array<Object>} [options.plugins] [Plugins]{@link https://developers.google.com/web/tools/workbox/guides/using-plugins}
     * to use in conjunction with this caching strategy.
     * @param {Object} [options.fetchOptions] Values passed along to the
     * [`init`](https://developer.mozilla.org/en-US/docs/Web/API/WindowOrWorkerGlobalScope/fetch#Parameters)
     * of [non-navigation](https://github.com/GoogleChrome/workbox/issues/1796)
     * `fetch()` requests made by this strategy.
     * @param {Object} [options.matchOptions] The
     * [`CacheQueryOptions`]{@link https://w3c.github.io/ServiceWorker/#dictdef-cachequeryoptions}
     * for any `cache.match()` or `cache.put()` calls made by this strategy.
     */
    constructor(options = {}) {
      this.cacheName = cacheNames.getRuntimeName(options.cacheName);
      this.plugins = options.plugins || [];
      this.fetchOptions = options.fetchOptions;
      this.matchOptions = options.matchOptions;
    }
    /**
     * Perform a request strategy and returns a `Promise` that will resolve with
     * a `Response`, invoking all relevant plugin callbacks.
     *
     * When a strategy instance is registered with a Workbox
     * {@link workbox-routing.Route}, this method is automatically
     * called when the route matches.
     *
     * Alternatively, this method can be used in a standalone `FetchEvent`
     * listener by passing it to `event.respondWith()`.
     *
     * @param {FetchEvent|Object} options A `FetchEvent` or an object with the
     *     properties listed below.
     * @param {Request|string} options.request A request to run this strategy for.
     * @param {ExtendableEvent} options.event The event associated with the
     *     request.
     * @param {URL} [options.url]
     * @param {*} [options.params]
     */
    handle(options) {
      const [responseDone] = this.handleAll(options);
      return responseDone;
    }
    /**
     * Similar to {@link workbox-strategies.Strategy~handle}, but
     * instead of just returning a `Promise` that resolves to a `Response` it
     * it will return an tuple of `[response, done]` promises, where the former
     * (`response`) is equivalent to what `handle()` returns, and the latter is a
     * Promise that will resolve once any promises that were added to
     * `event.waitUntil()` as part of performing the strategy have completed.
     *
     * You can await the `done` promise to ensure any extra work performed by
     * the strategy (usually caching responses) completes successfully.
     *
     * @param {FetchEvent|Object} options A `FetchEvent` or an object with the
     *     properties listed below.
     * @param {Request|string} options.request A request to run this strategy for.
     * @param {ExtendableEvent} options.event The event associated with the
     *     request.
     * @param {URL} [options.url]
     * @param {*} [options.params]
     * @return {Array<Promise>} A tuple of [response, done]
     *     promises that can be used to determine when the response resolves as
     *     well as when the handler has completed all its work.
     */
    handleAll(options) {
      if (options instanceof FetchEvent) {
        options = {
          event: options,
          request: options.request
        };
      }
      const event = options.event;
      const request = typeof options.request === "string" ? new Request(options.request) : options.request;
      const params = "params" in options ? options.params : void 0;
      const handler = new StrategyHandler(this, { event, request, params });
      const responseDone = this._getResponse(handler, request, event);
      const handlerDone = this._awaitComplete(responseDone, handler, request, event);
      return [responseDone, handlerDone];
    }
    async _getResponse(handler, request, event) {
      await handler.runCallbacks("handlerWillStart", { event, request });
      let response = void 0;
      try {
        response = await this._handle(request, handler);
        if (!response || response.type === "error") {
          throw new WorkboxError("no-response", { url: request.url });
        }
      } catch (error) {
        if (error instanceof Error) {
          for (const callback of handler.iterateCallbacks("handlerDidError")) {
            response = await callback({ error, event, request });
            if (response) {
              break;
            }
          }
        }
        if (!response) {
          throw error;
        } else if (true) {
          logger.log(`While responding to '${getFriendlyURL(request.url)}', an ${error instanceof Error ? error.toString() : ""} error occurred. Using a fallback response provided by a handlerDidError plugin.`);
        }
      }
      for (const callback of handler.iterateCallbacks("handlerWillRespond")) {
        response = await callback({ event, request, response });
      }
      return response;
    }
    async _awaitComplete(responseDone, handler, request, event) {
      let response;
      let error;
      try {
        response = await responseDone;
      } catch (error2) {
      }
      try {
        await handler.runCallbacks("handlerDidRespond", {
          event,
          request,
          response
        });
        await handler.doneWaiting();
      } catch (waitUntilError) {
        if (waitUntilError instanceof Error) {
          error = waitUntilError;
        }
      }
      await handler.runCallbacks("handlerDidComplete", {
        event,
        request,
        response,
        error
      });
      handler.destroy();
      if (error) {
        throw error;
      }
    }
  };

  // node_modules/.pnpm/workbox-strategies@7.3.0/node_modules/workbox-strategies/utils/messages.js
  var messages2 = {
    strategyStart: (strategyName, request) => `Using ${strategyName} to respond to '${getFriendlyURL(request.url)}'`,
    printFinalResponse: (response) => {
      if (response) {
        logger.groupCollapsed(`View the final response here.`);
        logger.log(response || "[No response returned]");
        logger.groupEnd();
      }
    }
  };

  // node_modules/.pnpm/workbox-strategies@7.3.0/node_modules/workbox-strategies/CacheFirst.js
  var CacheFirst = class extends Strategy {
    /**
     * @private
     * @param {Request|string} request A request to run this strategy for.
     * @param {workbox-strategies.StrategyHandler} handler The event that
     *     triggered the request.
     * @return {Promise<Response>}
     */
    async _handle(request, handler) {
      const logs = [];
      if (true) {
        finalAssertExports.isInstance(request, Request, {
          moduleName: "workbox-strategies",
          className: this.constructor.name,
          funcName: "makeRequest",
          paramName: "request"
        });
      }
      let response = await handler.cacheMatch(request);
      let error = void 0;
      if (!response) {
        if (true) {
          logs.push(`No response found in the '${this.cacheName}' cache. Will respond with a network request.`);
        }
        try {
          response = await handler.fetchAndCachePut(request);
        } catch (err) {
          if (err instanceof Error) {
            error = err;
          }
        }
        if (true) {
          if (response) {
            logs.push(`Got response from network.`);
          } else {
            logs.push(`Unable to get a response from the network.`);
          }
        }
      } else {
        if (true) {
          logs.push(`Found a cached response in the '${this.cacheName}' cache.`);
        }
      }
      if (true) {
        logger.groupCollapsed(messages2.strategyStart(this.constructor.name, request));
        for (const log of logs) {
          logger.log(log);
        }
        messages2.printFinalResponse(response);
        logger.groupEnd();
      }
      if (!response) {
        throw new WorkboxError("no-response", { url: request.url, error });
      }
      return response;
    }
  };

  // node_modules/.pnpm/workbox-strategies@7.3.0/node_modules/workbox-strategies/plugins/cacheOkAndOpaquePlugin.js
  var cacheOkAndOpaquePlugin = {
    /**
     * Returns a valid response (to allow caching) if the status is 200 (OK) or
     * 0 (opaque).
     *
     * @param {Object} options
     * @param {Response} options.response
     * @return {Response|null}
     *
     * @private
     */
    cacheWillUpdate: async ({ response }) => {
      if (response.status === 200 || response.status === 0) {
        return response;
      }
      return null;
    }
  };

  // node_modules/.pnpm/workbox-strategies@7.3.0/node_modules/workbox-strategies/NetworkFirst.js
  var NetworkFirst = class extends Strategy {
    /**
     * @param {Object} [options]
     * @param {string} [options.cacheName] Cache name to store and retrieve
     * requests. Defaults to cache names provided by
     * {@link workbox-core.cacheNames}.
     * @param {Array<Object>} [options.plugins] [Plugins]{@link https://developers.google.com/web/tools/workbox/guides/using-plugins}
     * to use in conjunction with this caching strategy.
     * @param {Object} [options.fetchOptions] Values passed along to the
     * [`init`](https://developer.mozilla.org/en-US/docs/Web/API/WindowOrWorkerGlobalScope/fetch#Parameters)
     * of [non-navigation](https://github.com/GoogleChrome/workbox/issues/1796)
     * `fetch()` requests made by this strategy.
     * @param {Object} [options.matchOptions] [`CacheQueryOptions`](https://w3c.github.io/ServiceWorker/#dictdef-cachequeryoptions)
     * @param {number} [options.networkTimeoutSeconds] If set, any network requests
     * that fail to respond within the timeout will fallback to the cache.
     *
     * This option can be used to combat
     * "[lie-fi]{@link https://developers.google.com/web/fundamentals/performance/poor-connectivity/#lie-fi}"
     * scenarios.
     */
    constructor(options = {}) {
      super(options);
      if (!this.plugins.some((p) => "cacheWillUpdate" in p)) {
        this.plugins.unshift(cacheOkAndOpaquePlugin);
      }
      this._networkTimeoutSeconds = options.networkTimeoutSeconds || 0;
      if (true) {
        if (this._networkTimeoutSeconds) {
          finalAssertExports.isType(this._networkTimeoutSeconds, "number", {
            moduleName: "workbox-strategies",
            className: this.constructor.name,
            funcName: "constructor",
            paramName: "networkTimeoutSeconds"
          });
        }
      }
    }
    /**
     * @private
     * @param {Request|string} request A request to run this strategy for.
     * @param {workbox-strategies.StrategyHandler} handler The event that
     *     triggered the request.
     * @return {Promise<Response>}
     */
    async _handle(request, handler) {
      const logs = [];
      if (true) {
        finalAssertExports.isInstance(request, Request, {
          moduleName: "workbox-strategies",
          className: this.constructor.name,
          funcName: "handle",
          paramName: "makeRequest"
        });
      }
      const promises = [];
      let timeoutId;
      if (this._networkTimeoutSeconds) {
        const { id, promise } = this._getTimeoutPromise({ request, logs, handler });
        timeoutId = id;
        promises.push(promise);
      }
      const networkPromise = this._getNetworkPromise({
        timeoutId,
        request,
        logs,
        handler
      });
      promises.push(networkPromise);
      const response = await handler.waitUntil((async () => {
        return await handler.waitUntil(Promise.race(promises)) || // If Promise.race() resolved with null, it might be due to a network
        // timeout + a cache miss. If that were to happen, we'd rather wait until
        // the networkPromise resolves instead of returning null.
        // Note that it's fine to await an already-resolved promise, so we don't
        // have to check to see if it's still "in flight".
        await networkPromise;
      })());
      if (true) {
        logger.groupCollapsed(messages2.strategyStart(this.constructor.name, request));
        for (const log of logs) {
          logger.log(log);
        }
        messages2.printFinalResponse(response);
        logger.groupEnd();
      }
      if (!response) {
        throw new WorkboxError("no-response", { url: request.url });
      }
      return response;
    }
    /**
     * @param {Object} options
     * @param {Request} options.request
     * @param {Array} options.logs A reference to the logs array
     * @param {Event} options.event
     * @return {Promise<Response>}
     *
     * @private
     */
    _getTimeoutPromise({ request, logs, handler }) {
      let timeoutId;
      const timeoutPromise = new Promise((resolve) => {
        const onNetworkTimeout = async () => {
          if (true) {
            logs.push(`Timing out the network response at ${this._networkTimeoutSeconds} seconds.`);
          }
          resolve(await handler.cacheMatch(request));
        };
        timeoutId = setTimeout(onNetworkTimeout, this._networkTimeoutSeconds * 1e3);
      });
      return {
        promise: timeoutPromise,
        id: timeoutId
      };
    }
    /**
     * @param {Object} options
     * @param {number|undefined} options.timeoutId
     * @param {Request} options.request
     * @param {Array} options.logs A reference to the logs Array.
     * @param {Event} options.event
     * @return {Promise<Response>}
     *
     * @private
     */
    async _getNetworkPromise({ timeoutId, request, logs, handler }) {
      let error;
      let response;
      try {
        response = await handler.fetchAndCachePut(request);
      } catch (fetchError) {
        if (fetchError instanceof Error) {
          error = fetchError;
        }
      }
      if (timeoutId) {
        clearTimeout(timeoutId);
      }
      if (true) {
        if (response) {
          logs.push(`Got response from network.`);
        } else {
          logs.push(`Unable to get a response from the network. Will respond with a cached response.`);
        }
      }
      if (error || !response) {
        response = await handler.cacheMatch(request);
        if (true) {
          if (response) {
            logs.push(`Found a cached response in the '${this.cacheName}' cache.`);
          } else {
            logs.push(`No response found in the '${this.cacheName}' cache.`);
          }
        }
      }
      return response;
    }
  };

  // node_modules/.pnpm/workbox-strategies@7.3.0/node_modules/workbox-strategies/StaleWhileRevalidate.js
  var StaleWhileRevalidate = class extends Strategy {
    /**
     * @param {Object} [options]
     * @param {string} [options.cacheName] Cache name to store and retrieve
     * requests. Defaults to cache names provided by
     * {@link workbox-core.cacheNames}.
     * @param {Array<Object>} [options.plugins] [Plugins]{@link https://developers.google.com/web/tools/workbox/guides/using-plugins}
     * to use in conjunction with this caching strategy.
     * @param {Object} [options.fetchOptions] Values passed along to the
     * [`init`](https://developer.mozilla.org/en-US/docs/Web/API/WindowOrWorkerGlobalScope/fetch#Parameters)
     * of [non-navigation](https://github.com/GoogleChrome/workbox/issues/1796)
     * `fetch()` requests made by this strategy.
     * @param {Object} [options.matchOptions] [`CacheQueryOptions`](https://w3c.github.io/ServiceWorker/#dictdef-cachequeryoptions)
     */
    constructor(options = {}) {
      super(options);
      if (!this.plugins.some((p) => "cacheWillUpdate" in p)) {
        this.plugins.unshift(cacheOkAndOpaquePlugin);
      }
    }
    /**
     * @private
     * @param {Request|string} request A request to run this strategy for.
     * @param {workbox-strategies.StrategyHandler} handler The event that
     *     triggered the request.
     * @return {Promise<Response>}
     */
    async _handle(request, handler) {
      const logs = [];
      if (true) {
        finalAssertExports.isInstance(request, Request, {
          moduleName: "workbox-strategies",
          className: this.constructor.name,
          funcName: "handle",
          paramName: "request"
        });
      }
      const fetchAndCachePromise = handler.fetchAndCachePut(request).catch(() => {
      });
      void handler.waitUntil(fetchAndCachePromise);
      let response = await handler.cacheMatch(request);
      let error;
      if (response) {
        if (true) {
          logs.push(`Found a cached response in the '${this.cacheName}' cache. Will update with the network response in the background.`);
        }
      } else {
        if (true) {
          logs.push(`No response found in the '${this.cacheName}' cache. Will wait for the network response.`);
        }
        try {
          response = await fetchAndCachePromise;
        } catch (err) {
          if (err instanceof Error) {
            error = err;
          }
        }
      }
      if (true) {
        logger.groupCollapsed(messages2.strategyStart(this.constructor.name, request));
        for (const log of logs) {
          logger.log(log);
        }
        messages2.printFinalResponse(response);
        logger.groupEnd();
      }
      if (!response) {
        throw new WorkboxError("no-response", { url: request.url, error });
      }
      return response;
    }
  };

  // node_modules/.pnpm/workbox-core@7.1.0/node_modules/workbox-core/_version.js
  try {
    self["workbox:core:7.0.0"] && _();
  } catch (e) {
  }

  // node_modules/.pnpm/workbox-core@7.1.0/node_modules/workbox-core/models/messages/messages.js
  var messages3 = {
    "invalid-value": ({ paramName, validValueDescription, value }) => {
      if (!paramName || !validValueDescription) {
        throw new Error(`Unexpected input to 'invalid-value' error.`);
      }
      return `The '${paramName}' parameter was given a value with an unexpected value. ${validValueDescription} Received a value of ${JSON.stringify(value)}.`;
    },
    "not-an-array": ({ moduleName, className, funcName, paramName }) => {
      if (!moduleName || !className || !funcName || !paramName) {
        throw new Error(`Unexpected input to 'not-an-array' error.`);
      }
      return `The parameter '${paramName}' passed into '${moduleName}.${className}.${funcName}()' must be an array.`;
    },
    "incorrect-type": ({ expectedType, paramName, moduleName, className, funcName }) => {
      if (!expectedType || !paramName || !moduleName || !funcName) {
        throw new Error(`Unexpected input to 'incorrect-type' error.`);
      }
      const classNameStr = className ? `${className}.` : "";
      return `The parameter '${paramName}' passed into '${moduleName}.${classNameStr}${funcName}()' must be of type ${expectedType}.`;
    },
    "incorrect-class": ({ expectedClassName, paramName, moduleName, className, funcName, isReturnValueProblem }) => {
      if (!expectedClassName || !moduleName || !funcName) {
        throw new Error(`Unexpected input to 'incorrect-class' error.`);
      }
      const classNameStr = className ? `${className}.` : "";
      if (isReturnValueProblem) {
        return `The return value from '${moduleName}.${classNameStr}${funcName}()' must be an instance of class ${expectedClassName}.`;
      }
      return `The parameter '${paramName}' passed into '${moduleName}.${classNameStr}${funcName}()' must be an instance of class ${expectedClassName}.`;
    },
    "missing-a-method": ({ expectedMethod, paramName, moduleName, className, funcName }) => {
      if (!expectedMethod || !paramName || !moduleName || !className || !funcName) {
        throw new Error(`Unexpected input to 'missing-a-method' error.`);
      }
      return `${moduleName}.${className}.${funcName}() expected the '${paramName}' parameter to expose a '${expectedMethod}' method.`;
    },
    "add-to-cache-list-unexpected-type": ({ entry }) => {
      return `An unexpected entry was passed to 'workbox-precaching.PrecacheController.addToCacheList()' The entry '${JSON.stringify(entry)}' isn't supported. You must supply an array of strings with one or more characters, objects with a url property or Request objects.`;
    },
    "add-to-cache-list-conflicting-entries": ({ firstEntry, secondEntry }) => {
      if (!firstEntry || !secondEntry) {
        throw new Error(`Unexpected input to 'add-to-cache-list-duplicate-entries' error.`);
      }
      return `Two of the entries passed to 'workbox-precaching.PrecacheController.addToCacheList()' had the URL ${firstEntry} but different revision details. Workbox is unable to cache and version the asset correctly. Please remove one of the entries.`;
    },
    "plugin-error-request-will-fetch": ({ thrownErrorMessage }) => {
      if (!thrownErrorMessage) {
        throw new Error(`Unexpected input to 'plugin-error-request-will-fetch', error.`);
      }
      return `An error was thrown by a plugins 'requestWillFetch()' method. The thrown error message was: '${thrownErrorMessage}'.`;
    },
    "invalid-cache-name": ({ cacheNameId, value }) => {
      if (!cacheNameId) {
        throw new Error(`Expected a 'cacheNameId' for error 'invalid-cache-name'`);
      }
      return `You must provide a name containing at least one character for setCacheDetails({${cacheNameId}: '...'}). Received a value of '${JSON.stringify(value)}'`;
    },
    "unregister-route-but-not-found-with-method": ({ method }) => {
      if (!method) {
        throw new Error(`Unexpected input to 'unregister-route-but-not-found-with-method' error.`);
      }
      return `The route you're trying to unregister was not  previously registered for the method type '${method}'.`;
    },
    "unregister-route-route-not-registered": () => {
      return `The route you're trying to unregister was not previously registered.`;
    },
    "queue-replay-failed": ({ name }) => {
      return `Replaying the background sync queue '${name}' failed.`;
    },
    "duplicate-queue-name": ({ name }) => {
      return `The Queue name '${name}' is already being used. All instances of backgroundSync.Queue must be given unique names.`;
    },
    "expired-test-without-max-age": ({ methodName, paramName }) => {
      return `The '${methodName}()' method can only be used when the '${paramName}' is used in the constructor.`;
    },
    "unsupported-route-type": ({ moduleName, className, funcName, paramName }) => {
      return `The supplied '${paramName}' parameter was an unsupported type. Please check the docs for ${moduleName}.${className}.${funcName} for valid input types.`;
    },
    "not-array-of-class": ({ value, expectedClass, moduleName, className, funcName, paramName }) => {
      return `The supplied '${paramName}' parameter must be an array of '${expectedClass}' objects. Received '${JSON.stringify(value)},'. Please check the call to ${moduleName}.${className}.${funcName}() to fix the issue.`;
    },
    "max-entries-or-age-required": ({ moduleName, className, funcName }) => {
      return `You must define either config.maxEntries or config.maxAgeSecondsin ${moduleName}.${className}.${funcName}`;
    },
    "statuses-or-headers-required": ({ moduleName, className, funcName }) => {
      return `You must define either config.statuses or config.headersin ${moduleName}.${className}.${funcName}`;
    },
    "invalid-string": ({ moduleName, funcName, paramName }) => {
      if (!paramName || !moduleName || !funcName) {
        throw new Error(`Unexpected input to 'invalid-string' error.`);
      }
      return `When using strings, the '${paramName}' parameter must start with 'http' (for cross-origin matches) or '/' (for same-origin matches). Please see the docs for ${moduleName}.${funcName}() for more info.`;
    },
    "channel-name-required": () => {
      return `You must provide a channelName to construct a BroadcastCacheUpdate instance.`;
    },
    "invalid-responses-are-same-args": () => {
      return `The arguments passed into responsesAreSame() appear to be invalid. Please ensure valid Responses are used.`;
    },
    "expire-custom-caches-only": () => {
      return `You must provide a 'cacheName' property when using the expiration plugin with a runtime caching strategy.`;
    },
    "unit-must-be-bytes": ({ normalizedRangeHeader }) => {
      if (!normalizedRangeHeader) {
        throw new Error(`Unexpected input to 'unit-must-be-bytes' error.`);
      }
      return `The 'unit' portion of the Range header must be set to 'bytes'. The Range header provided was "${normalizedRangeHeader}"`;
    },
    "single-range-only": ({ normalizedRangeHeader }) => {
      if (!normalizedRangeHeader) {
        throw new Error(`Unexpected input to 'single-range-only' error.`);
      }
      return `Multiple ranges are not supported. Please use a  single start value, and optional end value. The Range header provided was "${normalizedRangeHeader}"`;
    },
    "invalid-range-values": ({ normalizedRangeHeader }) => {
      if (!normalizedRangeHeader) {
        throw new Error(`Unexpected input to 'invalid-range-values' error.`);
      }
      return `The Range header is missing both start and end values. At least one of those values is needed. The Range header provided was "${normalizedRangeHeader}"`;
    },
    "no-range-header": () => {
      return `No Range header was found in the Request provided.`;
    },
    "range-not-satisfiable": ({ size, start, end }) => {
      return `The start (${start}) and end (${end}) values in the Range are not satisfiable by the cached response, which is ${size} bytes.`;
    },
    "attempt-to-cache-non-get-request": ({ url, method }) => {
      return `Unable to cache '${url}' because it is a '${method}' request and only 'GET' requests can be cached.`;
    },
    "cache-put-with-no-response": ({ url }) => {
      return `There was an attempt to cache '${url}' but the response was not defined.`;
    },
    "no-response": ({ url, error }) => {
      let message = `The strategy could not generate a response for '${url}'.`;
      if (error) {
        message += ` The underlying error is ${error}.`;
      }
      return message;
    },
    "bad-precaching-response": ({ url, status }) => {
      return `The precaching request for '${url}' failed` + (status ? ` with an HTTP status of ${status}.` : `.`);
    },
    "non-precached-url": ({ url }) => {
      return `createHandlerBoundToURL('${url}') was called, but that URL is not precached. Please pass in a URL that is precached instead.`;
    },
    "add-to-cache-list-conflicting-integrities": ({ url }) => {
      return `Two of the entries passed to 'workbox-precaching.PrecacheController.addToCacheList()' had the URL ${url} with different integrity values. Please remove one of them.`;
    },
    "missing-precache-entry": ({ cacheName, url }) => {
      return `Unable to find a precached response in ${cacheName} for ${url}.`;
    },
    "cross-origin-copy-response": ({ origin }) => {
      return `workbox-core.copyResponse() can only be used with same-origin responses. It was passed a response with origin ${origin}.`;
    },
    "opaque-streams-source": ({ type }) => {
      const message = `One of the workbox-streams sources resulted in an '${type}' response.`;
      if (type === "opaqueredirect") {
        return `${message} Please do not use a navigation request that results in a redirect as a source.`;
      }
      return `${message} Please ensure your sources are CORS-enabled.`;
    }
  };

  // node_modules/.pnpm/workbox-core@7.1.0/node_modules/workbox-core/models/messages/messageGenerator.js
  var generatorFunction2 = (code, details = {}) => {
    const message = messages3[code];
    if (!message) {
      throw new Error(`Unable to find message for code '${code}'.`);
    }
    return message(details);
  };
  var messageGenerator2 = false ? fallback : generatorFunction2;

  // node_modules/.pnpm/workbox-core@7.1.0/node_modules/workbox-core/_private/WorkboxError.js
  var WorkboxError2 = class extends Error {
    /**
     *
     * @param {string} errorCode The error code that
     * identifies this particular error.
     * @param {Object=} details Any relevant arguments
     * that will help developers identify issues should
     * be added as a key on the context object.
     */
    constructor(errorCode, details) {
      const message = messageGenerator2(errorCode, details);
      super(message);
      this.name = errorCode;
      this.details = details;
    }
  };

  // node_modules/.pnpm/workbox-core@7.1.0/node_modules/workbox-core/_private/assert.js
  var isArray2 = (value, details) => {
    if (!Array.isArray(value)) {
      throw new WorkboxError2("not-an-array", details);
    }
  };
  var hasMethod2 = (object, expectedMethod, details) => {
    const type = typeof object[expectedMethod];
    if (type !== "function") {
      details["expectedMethod"] = expectedMethod;
      throw new WorkboxError2("missing-a-method", details);
    }
  };
  var isType2 = (object, expectedType, details) => {
    if (typeof object !== expectedType) {
      details["expectedType"] = expectedType;
      throw new WorkboxError2("incorrect-type", details);
    }
  };
  var isInstance2 = (object, expectedClass, details) => {
    if (!(object instanceof expectedClass)) {
      details["expectedClassName"] = expectedClass.name;
      throw new WorkboxError2("incorrect-class", details);
    }
  };
  var isOneOf2 = (value, validValues, details) => {
    if (!validValues.includes(value)) {
      details["validValueDescription"] = `Valid values are ${JSON.stringify(validValues)}.`;
      throw new WorkboxError2("invalid-value", details);
    }
  };
  var isArrayOfClass2 = (value, expectedClass, details) => {
    const error = new WorkboxError2("not-array-of-class", details);
    if (!Array.isArray(value)) {
      throw error;
    }
    for (const item of value) {
      if (!(item instanceof expectedClass)) {
        throw error;
      }
    }
  };
  var finalAssertExports2 = false ? null : {
    hasMethod: hasMethod2,
    isArray: isArray2,
    isInstance: isInstance2,
    isOneOf: isOneOf2,
    isType: isType2,
    isArrayOfClass: isArrayOfClass2
  };

  // node_modules/.pnpm/workbox-core@7.1.0/node_modules/workbox-core/_private/logger.js
  var logger2 = false ? null : (() => {
    if (!("__WB_DISABLE_DEV_LOGS" in globalThis)) {
      self.__WB_DISABLE_DEV_LOGS = false;
    }
    let inGroup = false;
    const methodToColorMap = {
      debug: `#7f8c8d`,
      log: `#2ecc71`,
      warn: `#f39c12`,
      error: `#c0392b`,
      groupCollapsed: `#3498db`,
      groupEnd: null
      // No colored prefix on groupEnd
    };
    const print = function(method, args) {
      if (self.__WB_DISABLE_DEV_LOGS) {
        return;
      }
      if (method === "groupCollapsed") {
        if (/^((?!chrome|android).)*safari/i.test(navigator.userAgent)) {
          console[method](...args);
          return;
        }
      }
      const styles = [
        `background: ${methodToColorMap[method]}`,
        `border-radius: 0.5em`,
        `color: white`,
        `font-weight: bold`,
        `padding: 2px 0.5em`
      ];
      const logPrefix = inGroup ? [] : ["%cworkbox", styles.join(";")];
      console[method](...logPrefix, ...args);
      if (method === "groupCollapsed") {
        inGroup = true;
      }
      if (method === "groupEnd") {
        inGroup = false;
      }
    };
    const api = {};
    const loggerMethods = Object.keys(methodToColorMap);
    for (const key of loggerMethods) {
      const method = key;
      api[method] = (...args) => {
        print(method, args);
      };
    }
    return api;
  })();

  // node_modules/.pnpm/workbox-routing@7.1.0/node_modules/workbox-routing/_version.js
  try {
    self["workbox:routing:7.0.0"] && _();
  } catch (e) {
  }

  // node_modules/.pnpm/workbox-routing@7.1.0/node_modules/workbox-routing/utils/constants.js
  var defaultMethod = "GET";
  var validMethods = [
    "DELETE",
    "GET",
    "HEAD",
    "PATCH",
    "POST",
    "PUT"
  ];

  // node_modules/.pnpm/workbox-routing@7.1.0/node_modules/workbox-routing/utils/normalizeHandler.js
  var normalizeHandler = (handler) => {
    if (handler && typeof handler === "object") {
      if (true) {
        finalAssertExports2.hasMethod(handler, "handle", {
          moduleName: "workbox-routing",
          className: "Route",
          funcName: "constructor",
          paramName: "handler"
        });
      }
      return handler;
    } else {
      if (true) {
        finalAssertExports2.isType(handler, "function", {
          moduleName: "workbox-routing",
          className: "Route",
          funcName: "constructor",
          paramName: "handler"
        });
      }
      return { handle: handler };
    }
  };

  // node_modules/.pnpm/workbox-routing@7.1.0/node_modules/workbox-routing/Route.js
  var Route = class {
    /**
     * Constructor for Route class.
     *
     * @param {workbox-routing~matchCallback} match
     * A callback function that determines whether the route matches a given
     * `fetch` event by returning a non-falsy value.
     * @param {workbox-routing~handlerCallback} handler A callback
     * function that returns a Promise resolving to a Response.
     * @param {string} [method='GET'] The HTTP method to match the Route
     * against.
     */
    constructor(match, handler, method = defaultMethod) {
      if (true) {
        finalAssertExports2.isType(match, "function", {
          moduleName: "workbox-routing",
          className: "Route",
          funcName: "constructor",
          paramName: "match"
        });
        if (method) {
          finalAssertExports2.isOneOf(method, validMethods, { paramName: "method" });
        }
      }
      this.handler = normalizeHandler(handler);
      this.match = match;
      this.method = method;
    }
    /**
     *
     * @param {workbox-routing-handlerCallback} handler A callback
     * function that returns a Promise resolving to a Response
     */
    setCatchHandler(handler) {
      this.catchHandler = normalizeHandler(handler);
    }
  };

  // node_modules/.pnpm/workbox-routing@7.1.0/node_modules/workbox-routing/RegExpRoute.js
  var RegExpRoute = class extends Route {
    /**
     * If the regular expression contains
     * [capture groups]{@link https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/RegExp#grouping-back-references},
     * the captured values will be passed to the
     * {@link workbox-routing~handlerCallback} `params`
     * argument.
     *
     * @param {RegExp} regExp The regular expression to match against URLs.
     * @param {workbox-routing~handlerCallback} handler A callback
     * function that returns a Promise resulting in a Response.
     * @param {string} [method='GET'] The HTTP method to match the Route
     * against.
     */
    constructor(regExp, handler, method) {
      if (true) {
        finalAssertExports2.isInstance(regExp, RegExp, {
          moduleName: "workbox-routing",
          className: "RegExpRoute",
          funcName: "constructor",
          paramName: "pattern"
        });
      }
      const match = ({ url }) => {
        const result = regExp.exec(url.href);
        if (!result) {
          return;
        }
        if (url.origin !== location.origin && result.index !== 0) {
          if (true) {
            logger2.debug(`The regular expression '${regExp.toString()}' only partially matched against the cross-origin URL '${url.toString()}'. RegExpRoute's will only handle cross-origin requests if they match the entire URL.`);
          }
          return;
        }
        return result.slice(1);
      };
      super(match, handler, method);
    }
  };

  // node_modules/.pnpm/workbox-core@7.1.0/node_modules/workbox-core/_private/getFriendlyURL.js
  var getFriendlyURL2 = (url) => {
    const urlObj = new URL(String(url), location.href);
    return urlObj.href.replace(new RegExp(`^${location.origin}`), "");
  };

  // node_modules/.pnpm/workbox-routing@7.1.0/node_modules/workbox-routing/Router.js
  var Router = class {
    /**
     * Initializes a new Router.
     */
    constructor() {
      this._routes = /* @__PURE__ */ new Map();
      this._defaultHandlerMap = /* @__PURE__ */ new Map();
    }
    /**
     * @return {Map<string, Array<workbox-routing.Route>>} routes A `Map` of HTTP
     * method name ('GET', etc.) to an array of all the corresponding `Route`
     * instances that are registered.
     */
    get routes() {
      return this._routes;
    }
    /**
     * Adds a fetch event listener to respond to events when a route matches
     * the event's request.
     */
    addFetchListener() {
      self.addEventListener("fetch", (event) => {
        const { request } = event;
        const responsePromise = this.handleRequest({ request, event });
        if (responsePromise) {
          event.respondWith(responsePromise);
        }
      });
    }
    /**
     * Adds a message event listener for URLs to cache from the window.
     * This is useful to cache resources loaded on the page prior to when the
     * service worker started controlling it.
     *
     * The format of the message data sent from the window should be as follows.
     * Where the `urlsToCache` array may consist of URL strings or an array of
     * URL string + `requestInit` object (the same as you'd pass to `fetch()`).
     *
     * ```
     * {
     *   type: 'CACHE_URLS',
     *   payload: {
     *     urlsToCache: [
     *       './script1.js',
     *       './script2.js',
     *       ['./script3.js', {mode: 'no-cors'}],
     *     ],
     *   },
     * }
     * ```
     */
    addCacheListener() {
      self.addEventListener("message", (event) => {
        if (event.data && event.data.type === "CACHE_URLS") {
          const { payload } = event.data;
          if (true) {
            logger2.debug(`Caching URLs from the window`, payload.urlsToCache);
          }
          const requestPromises = Promise.all(payload.urlsToCache.map((entry) => {
            if (typeof entry === "string") {
              entry = [entry];
            }
            const request = new Request(...entry);
            return this.handleRequest({ request, event });
          }));
          event.waitUntil(requestPromises);
          if (event.ports && event.ports[0]) {
            void requestPromises.then(() => event.ports[0].postMessage(true));
          }
        }
      });
    }
    /**
     * Apply the routing rules to a FetchEvent object to get a Response from an
     * appropriate Route's handler.
     *
     * @param {Object} options
     * @param {Request} options.request The request to handle.
     * @param {ExtendableEvent} options.event The event that triggered the
     *     request.
     * @return {Promise<Response>|undefined} A promise is returned if a
     *     registered route can handle the request. If there is no matching
     *     route and there's no `defaultHandler`, `undefined` is returned.
     */
    handleRequest({ request, event }) {
      if (true) {
        finalAssertExports2.isInstance(request, Request, {
          moduleName: "workbox-routing",
          className: "Router",
          funcName: "handleRequest",
          paramName: "options.request"
        });
      }
      const url = new URL(request.url, location.href);
      if (!url.protocol.startsWith("http")) {
        if (true) {
          logger2.debug(`Workbox Router only supports URLs that start with 'http'.`);
        }
        return;
      }
      const sameOrigin = url.origin === location.origin;
      const { params, route } = this.findMatchingRoute({
        event,
        request,
        sameOrigin,
        url
      });
      let handler = route && route.handler;
      const debugMessages = [];
      if (true) {
        if (handler) {
          debugMessages.push([`Found a route to handle this request:`, route]);
          if (params) {
            debugMessages.push([
              `Passing the following params to the route's handler:`,
              params
            ]);
          }
        }
      }
      const method = request.method;
      if (!handler && this._defaultHandlerMap.has(method)) {
        if (true) {
          debugMessages.push(`Failed to find a matching route. Falling back to the default handler for ${method}.`);
        }
        handler = this._defaultHandlerMap.get(method);
      }
      if (!handler) {
        if (true) {
          logger2.debug(`No route found for: ${getFriendlyURL2(url)}`);
        }
        return;
      }
      if (true) {
        logger2.groupCollapsed(`Router is responding to: ${getFriendlyURL2(url)}`);
        debugMessages.forEach((msg) => {
          if (Array.isArray(msg)) {
            logger2.log(...msg);
          } else {
            logger2.log(msg);
          }
        });
        logger2.groupEnd();
      }
      let responsePromise;
      try {
        responsePromise = handler.handle({ url, request, event, params });
      } catch (err) {
        responsePromise = Promise.reject(err);
      }
      const catchHandler = route && route.catchHandler;
      if (responsePromise instanceof Promise && (this._catchHandler || catchHandler)) {
        responsePromise = responsePromise.catch(async (err) => {
          if (catchHandler) {
            if (true) {
              logger2.groupCollapsed(`Error thrown when responding to:  ${getFriendlyURL2(url)}. Falling back to route's Catch Handler.`);
              logger2.error(`Error thrown by:`, route);
              logger2.error(err);
              logger2.groupEnd();
            }
            try {
              return await catchHandler.handle({ url, request, event, params });
            } catch (catchErr) {
              if (catchErr instanceof Error) {
                err = catchErr;
              }
            }
          }
          if (this._catchHandler) {
            if (true) {
              logger2.groupCollapsed(`Error thrown when responding to:  ${getFriendlyURL2(url)}. Falling back to global Catch Handler.`);
              logger2.error(`Error thrown by:`, route);
              logger2.error(err);
              logger2.groupEnd();
            }
            return this._catchHandler.handle({ url, request, event });
          }
          throw err;
        });
      }
      return responsePromise;
    }
    /**
     * Checks a request and URL (and optionally an event) against the list of
     * registered routes, and if there's a match, returns the corresponding
     * route along with any params generated by the match.
     *
     * @param {Object} options
     * @param {URL} options.url
     * @param {boolean} options.sameOrigin The result of comparing `url.origin`
     *     against the current origin.
     * @param {Request} options.request The request to match.
     * @param {Event} options.event The corresponding event.
     * @return {Object} An object with `route` and `params` properties.
     *     They are populated if a matching route was found or `undefined`
     *     otherwise.
     */
    findMatchingRoute({ url, sameOrigin, request, event }) {
      const routes = this._routes.get(request.method) || [];
      for (const route of routes) {
        let params;
        const matchResult = route.match({ url, sameOrigin, request, event });
        if (matchResult) {
          if (true) {
            if (matchResult instanceof Promise) {
              logger2.warn(`While routing ${getFriendlyURL2(url)}, an async matchCallback function was used. Please convert the following route to use a synchronous matchCallback function:`, route);
            }
          }
          params = matchResult;
          if (Array.isArray(params) && params.length === 0) {
            params = void 0;
          } else if (matchResult.constructor === Object && // eslint-disable-line
          Object.keys(matchResult).length === 0) {
            params = void 0;
          } else if (typeof matchResult === "boolean") {
            params = void 0;
          }
          return { route, params };
        }
      }
      return {};
    }
    /**
     * Define a default `handler` that's called when no routes explicitly
     * match the incoming request.
     *
     * Each HTTP method ('GET', 'POST', etc.) gets its own default handler.
     *
     * Without a default handler, unmatched requests will go against the
     * network as if there were no service worker present.
     *
     * @param {workbox-routing~handlerCallback} handler A callback
     * function that returns a Promise resulting in a Response.
     * @param {string} [method='GET'] The HTTP method to associate with this
     * default handler. Each method has its own default.
     */
    setDefaultHandler(handler, method = defaultMethod) {
      this._defaultHandlerMap.set(method, normalizeHandler(handler));
    }
    /**
     * If a Route throws an error while handling a request, this `handler`
     * will be called and given a chance to provide a response.
     *
     * @param {workbox-routing~handlerCallback} handler A callback
     * function that returns a Promise resulting in a Response.
     */
    setCatchHandler(handler) {
      this._catchHandler = normalizeHandler(handler);
    }
    /**
     * Registers a route with the router.
     *
     * @param {workbox-routing.Route} route The route to register.
     */
    registerRoute(route) {
      if (true) {
        finalAssertExports2.isType(route, "object", {
          moduleName: "workbox-routing",
          className: "Router",
          funcName: "registerRoute",
          paramName: "route"
        });
        finalAssertExports2.hasMethod(route, "match", {
          moduleName: "workbox-routing",
          className: "Router",
          funcName: "registerRoute",
          paramName: "route"
        });
        finalAssertExports2.isType(route.handler, "object", {
          moduleName: "workbox-routing",
          className: "Router",
          funcName: "registerRoute",
          paramName: "route"
        });
        finalAssertExports2.hasMethod(route.handler, "handle", {
          moduleName: "workbox-routing",
          className: "Router",
          funcName: "registerRoute",
          paramName: "route.handler"
        });
        finalAssertExports2.isType(route.method, "string", {
          moduleName: "workbox-routing",
          className: "Router",
          funcName: "registerRoute",
          paramName: "route.method"
        });
      }
      if (!this._routes.has(route.method)) {
        this._routes.set(route.method, []);
      }
      this._routes.get(route.method).push(route);
    }
    /**
     * Unregisters a route with the router.
     *
     * @param {workbox-routing.Route} route The route to unregister.
     */
    unregisterRoute(route) {
      if (!this._routes.has(route.method)) {
        throw new WorkboxError2("unregister-route-but-not-found-with-method", {
          method: route.method
        });
      }
      const routeIndex = this._routes.get(route.method).indexOf(route);
      if (routeIndex > -1) {
        this._routes.get(route.method).splice(routeIndex, 1);
      } else {
        throw new WorkboxError2("unregister-route-route-not-registered");
      }
    }
  };

  // node_modules/.pnpm/workbox-routing@7.1.0/node_modules/workbox-routing/utils/getOrCreateDefaultRouter.js
  var defaultRouter;
  var getOrCreateDefaultRouter = () => {
    if (!defaultRouter) {
      defaultRouter = new Router();
      defaultRouter.addFetchListener();
      defaultRouter.addCacheListener();
    }
    return defaultRouter;
  };

  // node_modules/.pnpm/workbox-routing@7.1.0/node_modules/workbox-routing/registerRoute.js
  function registerRoute(capture, handler, method) {
    let route;
    if (typeof capture === "string") {
      const captureUrl = new URL(capture, location.href);
      if (true) {
        if (!(capture.startsWith("/") || capture.startsWith("http"))) {
          throw new WorkboxError2("invalid-string", {
            moduleName: "workbox-routing",
            funcName: "registerRoute",
            paramName: "capture"
          });
        }
        const valueToCheck = capture.startsWith("http") ? captureUrl.pathname : capture;
        const wildcards = "[*:?+]";
        if (new RegExp(`${wildcards}`).exec(valueToCheck)) {
          logger2.debug(`The '$capture' parameter contains an Express-style wildcard character (${wildcards}). Strings are now always interpreted as exact matches; use a RegExp for partial or wildcard matches.`);
        }
      }
      const matchCallback = ({ url }) => {
        if (true) {
          if (url.pathname === captureUrl.pathname && url.origin !== captureUrl.origin) {
            logger2.debug(`${capture} only partially matches the cross-origin URL ${url.toString()}. This route will only handle cross-origin requests if they match the entire URL.`);
          }
        }
        return url.href === captureUrl.href;
      };
      route = new Route(matchCallback, handler, method);
    } else if (capture instanceof RegExp) {
      route = new RegExpRoute(capture, handler, method);
    } else if (typeof capture === "function") {
      route = new Route(capture, handler, method);
    } else if (capture instanceof Route) {
      route = capture;
    } else {
      throw new WorkboxError2("unsupported-route-type", {
        moduleName: "workbox-routing",
        funcName: "registerRoute",
        paramName: "capture"
      });
    }
    const defaultRouter3 = getOrCreateDefaultRouter();
    defaultRouter3.registerRoute(route);
    return route;
  }

  // node_modules/.pnpm/workbox-core@7.3.0/node_modules/workbox-core/_private/waitUntil.js
  function waitUntil(event, asyncFn) {
    const returnPromise = asyncFn();
    event.waitUntil(returnPromise);
    return returnPromise;
  }

  // node_modules/.pnpm/workbox-precaching@7.3.0/node_modules/workbox-precaching/_version.js
  try {
    self["workbox:precaching:7.2.0"] && _();
  } catch (e) {
  }

  // node_modules/.pnpm/workbox-precaching@7.3.0/node_modules/workbox-precaching/utils/createCacheKey.js
  var REVISION_SEARCH_PARAM = "__WB_REVISION__";
  function createCacheKey(entry) {
    if (!entry) {
      throw new WorkboxError("add-to-cache-list-unexpected-type", { entry });
    }
    if (typeof entry === "string") {
      const urlObject = new URL(entry, location.href);
      return {
        cacheKey: urlObject.href,
        url: urlObject.href
      };
    }
    const { revision, url } = entry;
    if (!url) {
      throw new WorkboxError("add-to-cache-list-unexpected-type", { entry });
    }
    if (!revision) {
      const urlObject = new URL(url, location.href);
      return {
        cacheKey: urlObject.href,
        url: urlObject.href
      };
    }
    const cacheKeyURL = new URL(url, location.href);
    const originalURL = new URL(url, location.href);
    cacheKeyURL.searchParams.set(REVISION_SEARCH_PARAM, revision);
    return {
      cacheKey: cacheKeyURL.href,
      url: originalURL.href
    };
  }

  // node_modules/.pnpm/workbox-precaching@7.3.0/node_modules/workbox-precaching/utils/PrecacheInstallReportPlugin.js
  var PrecacheInstallReportPlugin = class {
    constructor() {
      this.updatedURLs = [];
      this.notUpdatedURLs = [];
      this.handlerWillStart = async ({ request, state }) => {
        if (state) {
          state.originalRequest = request;
        }
      };
      this.cachedResponseWillBeUsed = async ({ event, state, cachedResponse }) => {
        if (event.type === "install") {
          if (state && state.originalRequest && state.originalRequest instanceof Request) {
            const url = state.originalRequest.url;
            if (cachedResponse) {
              this.notUpdatedURLs.push(url);
            } else {
              this.updatedURLs.push(url);
            }
          }
        }
        return cachedResponse;
      };
    }
  };

  // node_modules/.pnpm/workbox-precaching@7.3.0/node_modules/workbox-precaching/utils/PrecacheCacheKeyPlugin.js
  var PrecacheCacheKeyPlugin = class {
    constructor({ precacheController: precacheController2 }) {
      this.cacheKeyWillBeUsed = async ({ request, params }) => {
        const cacheKey = (params === null || params === void 0 ? void 0 : params.cacheKey) || this._precacheController.getCacheKeyForURL(request.url);
        return cacheKey ? new Request(cacheKey, { headers: request.headers }) : request;
      };
      this._precacheController = precacheController2;
    }
  };

  // node_modules/.pnpm/workbox-precaching@7.3.0/node_modules/workbox-precaching/utils/printCleanupDetails.js
  var logGroup = (groupTitle, deletedURLs) => {
    logger.groupCollapsed(groupTitle);
    for (const url of deletedURLs) {
      logger.log(url);
    }
    logger.groupEnd();
  };
  function printCleanupDetails(deletedURLs) {
    const deletionCount = deletedURLs.length;
    if (deletionCount > 0) {
      logger.groupCollapsed(`During precaching cleanup, ${deletionCount} cached request${deletionCount === 1 ? " was" : "s were"} deleted.`);
      logGroup("Deleted Cache Requests", deletedURLs);
      logger.groupEnd();
    }
  }

  // node_modules/.pnpm/workbox-precaching@7.3.0/node_modules/workbox-precaching/utils/printInstallDetails.js
  function _nestedGroup(groupTitle, urls) {
    if (urls.length === 0) {
      return;
    }
    logger.groupCollapsed(groupTitle);
    for (const url of urls) {
      logger.log(url);
    }
    logger.groupEnd();
  }
  function printInstallDetails(urlsToPrecache, urlsAlreadyPrecached) {
    const precachedCount = urlsToPrecache.length;
    const alreadyPrecachedCount = urlsAlreadyPrecached.length;
    if (precachedCount || alreadyPrecachedCount) {
      let message = `Precaching ${precachedCount} file${precachedCount === 1 ? "" : "s"}.`;
      if (alreadyPrecachedCount > 0) {
        message += ` ${alreadyPrecachedCount} file${alreadyPrecachedCount === 1 ? " is" : "s are"} already cached.`;
      }
      logger.groupCollapsed(message);
      _nestedGroup(`View newly precached URLs.`, urlsToPrecache);
      _nestedGroup(`View previously precached URLs.`, urlsAlreadyPrecached);
      logger.groupEnd();
    }
  }

  // node_modules/.pnpm/workbox-core@7.3.0/node_modules/workbox-core/_private/canConstructResponseFromBodyStream.js
  var supportStatus;
  function canConstructResponseFromBodyStream() {
    if (supportStatus === void 0) {
      const testResponse = new Response("");
      if ("body" in testResponse) {
        try {
          new Response(testResponse.body);
          supportStatus = true;
        } catch (error) {
          supportStatus = false;
        }
      }
      supportStatus = false;
    }
    return supportStatus;
  }

  // node_modules/.pnpm/workbox-core@7.3.0/node_modules/workbox-core/copyResponse.js
  async function copyResponse(response, modifier) {
    let origin = null;
    if (response.url) {
      const responseURL = new URL(response.url);
      origin = responseURL.origin;
    }
    if (origin !== self.location.origin) {
      throw new WorkboxError("cross-origin-copy-response", { origin });
    }
    const clonedResponse = response.clone();
    const responseInit = {
      headers: new Headers(clonedResponse.headers),
      status: clonedResponse.status,
      statusText: clonedResponse.statusText
    };
    const modifiedResponseInit = modifier ? modifier(responseInit) : responseInit;
    const body = canConstructResponseFromBodyStream() ? clonedResponse.body : await clonedResponse.blob();
    return new Response(body, modifiedResponseInit);
  }

  // node_modules/.pnpm/workbox-precaching@7.3.0/node_modules/workbox-precaching/PrecacheStrategy.js
  var PrecacheStrategy = class _PrecacheStrategy extends Strategy {
    /**
     *
     * @param {Object} [options]
     * @param {string} [options.cacheName] Cache name to store and retrieve
     * requests. Defaults to the cache names provided by
     * {@link workbox-core.cacheNames}.
     * @param {Array<Object>} [options.plugins] {@link https://developers.google.com/web/tools/workbox/guides/using-plugins|Plugins}
     * to use in conjunction with this caching strategy.
     * @param {Object} [options.fetchOptions] Values passed along to the
     * {@link https://developer.mozilla.org/en-US/docs/Web/API/WindowOrWorkerGlobalScope/fetch#Parameters|init}
     * of all fetch() requests made by this strategy.
     * @param {Object} [options.matchOptions] The
     * {@link https://w3c.github.io/ServiceWorker/#dictdef-cachequeryoptions|CacheQueryOptions}
     * for any `cache.match()` or `cache.put()` calls made by this strategy.
     * @param {boolean} [options.fallbackToNetwork=true] Whether to attempt to
     * get the response from the network if there's a precache miss.
     */
    constructor(options = {}) {
      options.cacheName = cacheNames.getPrecacheName(options.cacheName);
      super(options);
      this._fallbackToNetwork = options.fallbackToNetwork === false ? false : true;
      this.plugins.push(_PrecacheStrategy.copyRedirectedCacheableResponsesPlugin);
    }
    /**
     * @private
     * @param {Request|string} request A request to run this strategy for.
     * @param {workbox-strategies.StrategyHandler} handler The event that
     *     triggered the request.
     * @return {Promise<Response>}
     */
    async _handle(request, handler) {
      const response = await handler.cacheMatch(request);
      if (response) {
        return response;
      }
      if (handler.event && handler.event.type === "install") {
        return await this._handleInstall(request, handler);
      }
      return await this._handleFetch(request, handler);
    }
    async _handleFetch(request, handler) {
      let response;
      const params = handler.params || {};
      if (this._fallbackToNetwork) {
        if (true) {
          logger.warn(`The precached response for ${getFriendlyURL(request.url)} in ${this.cacheName} was not found. Falling back to the network.`);
        }
        const integrityInManifest = params.integrity;
        const integrityInRequest = request.integrity;
        const noIntegrityConflict = !integrityInRequest || integrityInRequest === integrityInManifest;
        response = await handler.fetch(new Request(request, {
          integrity: request.mode !== "no-cors" ? integrityInRequest || integrityInManifest : void 0
        }));
        if (integrityInManifest && noIntegrityConflict && request.mode !== "no-cors") {
          this._useDefaultCacheabilityPluginIfNeeded();
          const wasCached = await handler.cachePut(request, response.clone());
          if (true) {
            if (wasCached) {
              logger.log(`A response for ${getFriendlyURL(request.url)} was used to "repair" the precache.`);
            }
          }
        }
      } else {
        throw new WorkboxError("missing-precache-entry", {
          cacheName: this.cacheName,
          url: request.url
        });
      }
      if (true) {
        const cacheKey = params.cacheKey || await handler.getCacheKey(request, "read");
        logger.groupCollapsed(`Precaching is responding to: ` + getFriendlyURL(request.url));
        logger.log(`Serving the precached url: ${getFriendlyURL(cacheKey instanceof Request ? cacheKey.url : cacheKey)}`);
        logger.groupCollapsed(`View request details here.`);
        logger.log(request);
        logger.groupEnd();
        logger.groupCollapsed(`View response details here.`);
        logger.log(response);
        logger.groupEnd();
        logger.groupEnd();
      }
      return response;
    }
    async _handleInstall(request, handler) {
      this._useDefaultCacheabilityPluginIfNeeded();
      const response = await handler.fetch(request);
      const wasCached = await handler.cachePut(request, response.clone());
      if (!wasCached) {
        throw new WorkboxError("bad-precaching-response", {
          url: request.url,
          status: response.status
        });
      }
      return response;
    }
    /**
     * This method is complex, as there a number of things to account for:
     *
     * The `plugins` array can be set at construction, and/or it might be added to
     * to at any time before the strategy is used.
     *
     * At the time the strategy is used (i.e. during an `install` event), there
     * needs to be at least one plugin that implements `cacheWillUpdate` in the
     * array, other than `copyRedirectedCacheableResponsesPlugin`.
     *
     * - If this method is called and there are no suitable `cacheWillUpdate`
     * plugins, we need to add `defaultPrecacheCacheabilityPlugin`.
     *
     * - If this method is called and there is exactly one `cacheWillUpdate`, then
     * we don't have to do anything (this might be a previously added
     * `defaultPrecacheCacheabilityPlugin`, or it might be a custom plugin).
     *
     * - If this method is called and there is more than one `cacheWillUpdate`,
     * then we need to check if one is `defaultPrecacheCacheabilityPlugin`. If so,
     * we need to remove it. (This situation is unlikely, but it could happen if
     * the strategy is used multiple times, the first without a `cacheWillUpdate`,
     * and then later on after manually adding a custom `cacheWillUpdate`.)
     *
     * See https://github.com/GoogleChrome/workbox/issues/2737 for more context.
     *
     * @private
     */
    _useDefaultCacheabilityPluginIfNeeded() {
      let defaultPluginIndex = null;
      let cacheWillUpdatePluginCount = 0;
      for (const [index, plugin] of this.plugins.entries()) {
        if (plugin === _PrecacheStrategy.copyRedirectedCacheableResponsesPlugin) {
          continue;
        }
        if (plugin === _PrecacheStrategy.defaultPrecacheCacheabilityPlugin) {
          defaultPluginIndex = index;
        }
        if (plugin.cacheWillUpdate) {
          cacheWillUpdatePluginCount++;
        }
      }
      if (cacheWillUpdatePluginCount === 0) {
        this.plugins.push(_PrecacheStrategy.defaultPrecacheCacheabilityPlugin);
      } else if (cacheWillUpdatePluginCount > 1 && defaultPluginIndex !== null) {
        this.plugins.splice(defaultPluginIndex, 1);
      }
    }
  };
  PrecacheStrategy.defaultPrecacheCacheabilityPlugin = {
    async cacheWillUpdate({ response }) {
      if (!response || response.status >= 400) {
        return null;
      }
      return response;
    }
  };
  PrecacheStrategy.copyRedirectedCacheableResponsesPlugin = {
    async cacheWillUpdate({ response }) {
      return response.redirected ? await copyResponse(response) : response;
    }
  };

  // node_modules/.pnpm/workbox-precaching@7.3.0/node_modules/workbox-precaching/PrecacheController.js
  var PrecacheController = class {
    /**
     * Create a new PrecacheController.
     *
     * @param {Object} [options]
     * @param {string} [options.cacheName] The cache to use for precaching.
     * @param {string} [options.plugins] Plugins to use when precaching as well
     * as responding to fetch events for precached assets.
     * @param {boolean} [options.fallbackToNetwork=true] Whether to attempt to
     * get the response from the network if there's a precache miss.
     */
    constructor({ cacheName, plugins = [], fallbackToNetwork = true } = {}) {
      this._urlsToCacheKeys = /* @__PURE__ */ new Map();
      this._urlsToCacheModes = /* @__PURE__ */ new Map();
      this._cacheKeysToIntegrities = /* @__PURE__ */ new Map();
      this._strategy = new PrecacheStrategy({
        cacheName: cacheNames.getPrecacheName(cacheName),
        plugins: [
          ...plugins,
          new PrecacheCacheKeyPlugin({ precacheController: this })
        ],
        fallbackToNetwork
      });
      this.install = this.install.bind(this);
      this.activate = this.activate.bind(this);
    }
    /**
     * @type {workbox-precaching.PrecacheStrategy} The strategy created by this controller and
     * used to cache assets and respond to fetch events.
     */
    get strategy() {
      return this._strategy;
    }
    /**
     * Adds items to the precache list, removing any duplicates and
     * stores the files in the
     * {@link workbox-core.cacheNames|"precache cache"} when the service
     * worker installs.
     *
     * This method can be called multiple times.
     *
     * @param {Array<Object|string>} [entries=[]] Array of entries to precache.
     */
    precache(entries) {
      this.addToCacheList(entries);
      if (!this._installAndActiveListenersAdded) {
        self.addEventListener("install", this.install);
        self.addEventListener("activate", this.activate);
        this._installAndActiveListenersAdded = true;
      }
    }
    /**
     * This method will add items to the precache list, removing duplicates
     * and ensuring the information is valid.
     *
     * @param {Array<workbox-precaching.PrecacheController.PrecacheEntry|string>} entries
     *     Array of entries to precache.
     */
    addToCacheList(entries) {
      if (true) {
        finalAssertExports.isArray(entries, {
          moduleName: "workbox-precaching",
          className: "PrecacheController",
          funcName: "addToCacheList",
          paramName: "entries"
        });
      }
      const urlsToWarnAbout = [];
      for (const entry of entries) {
        if (typeof entry === "string") {
          urlsToWarnAbout.push(entry);
        } else if (entry && entry.revision === void 0) {
          urlsToWarnAbout.push(entry.url);
        }
        const { cacheKey, url } = createCacheKey(entry);
        const cacheMode = typeof entry !== "string" && entry.revision ? "reload" : "default";
        if (this._urlsToCacheKeys.has(url) && this._urlsToCacheKeys.get(url) !== cacheKey) {
          throw new WorkboxError("add-to-cache-list-conflicting-entries", {
            firstEntry: this._urlsToCacheKeys.get(url),
            secondEntry: cacheKey
          });
        }
        if (typeof entry !== "string" && entry.integrity) {
          if (this._cacheKeysToIntegrities.has(cacheKey) && this._cacheKeysToIntegrities.get(cacheKey) !== entry.integrity) {
            throw new WorkboxError("add-to-cache-list-conflicting-integrities", {
              url
            });
          }
          this._cacheKeysToIntegrities.set(cacheKey, entry.integrity);
        }
        this._urlsToCacheKeys.set(url, cacheKey);
        this._urlsToCacheModes.set(url, cacheMode);
        if (urlsToWarnAbout.length > 0) {
          const warningMessage = `Workbox is precaching URLs without revision info: ${urlsToWarnAbout.join(", ")}
This is generally NOT safe. Learn more at https://bit.ly/wb-precache`;
          if (false) {
            console.warn(warningMessage);
          } else {
            logger.warn(warningMessage);
          }
        }
      }
    }
    /**
     * Precaches new and updated assets. Call this method from the service worker
     * install event.
     *
     * Note: this method calls `event.waitUntil()` for you, so you do not need
     * to call it yourself in your event handlers.
     *
     * @param {ExtendableEvent} event
     * @return {Promise<workbox-precaching.InstallResult>}
     */
    install(event) {
      return waitUntil(event, async () => {
        const installReportPlugin = new PrecacheInstallReportPlugin();
        this.strategy.plugins.push(installReportPlugin);
        for (const [url, cacheKey] of this._urlsToCacheKeys) {
          const integrity = this._cacheKeysToIntegrities.get(cacheKey);
          const cacheMode = this._urlsToCacheModes.get(url);
          const request = new Request(url, {
            integrity,
            cache: cacheMode,
            credentials: "same-origin"
          });
          await Promise.all(this.strategy.handleAll({
            params: { cacheKey },
            request,
            event
          }));
        }
        const { updatedURLs, notUpdatedURLs } = installReportPlugin;
        if (true) {
          printInstallDetails(updatedURLs, notUpdatedURLs);
        }
        return { updatedURLs, notUpdatedURLs };
      });
    }
    /**
     * Deletes assets that are no longer present in the current precache manifest.
     * Call this method from the service worker activate event.
     *
     * Note: this method calls `event.waitUntil()` for you, so you do not need
     * to call it yourself in your event handlers.
     *
     * @param {ExtendableEvent} event
     * @return {Promise<workbox-precaching.CleanupResult>}
     */
    activate(event) {
      return waitUntil(event, async () => {
        const cache = await self.caches.open(this.strategy.cacheName);
        const currentlyCachedRequests = await cache.keys();
        const expectedCacheKeys = new Set(this._urlsToCacheKeys.values());
        const deletedURLs = [];
        for (const request of currentlyCachedRequests) {
          if (!expectedCacheKeys.has(request.url)) {
            await cache.delete(request);
            deletedURLs.push(request.url);
          }
        }
        if (true) {
          printCleanupDetails(deletedURLs);
        }
        return { deletedURLs };
      });
    }
    /**
     * Returns a mapping of a precached URL to the corresponding cache key, taking
     * into account the revision information for the URL.
     *
     * @return {Map<string, string>} A URL to cache key mapping.
     */
    getURLsToCacheKeys() {
      return this._urlsToCacheKeys;
    }
    /**
     * Returns a list of all the URLs that have been precached by the current
     * service worker.
     *
     * @return {Array<string>} The precached URLs.
     */
    getCachedURLs() {
      return [...this._urlsToCacheKeys.keys()];
    }
    /**
     * Returns the cache key used for storing a given URL. If that URL is
     * unversioned, like `/index.html', then the cache key will be the original
     * URL with a search parameter appended to it.
     *
     * @param {string} url A URL whose cache key you want to look up.
     * @return {string} The versioned URL that corresponds to a cache key
     * for the original URL, or undefined if that URL isn't precached.
     */
    getCacheKeyForURL(url) {
      const urlObject = new URL(url, location.href);
      return this._urlsToCacheKeys.get(urlObject.href);
    }
    /**
     * @param {string} url A cache key whose SRI you want to look up.
     * @return {string} The subresource integrity associated with the cache key,
     * or undefined if it's not set.
     */
    getIntegrityForCacheKey(cacheKey) {
      return this._cacheKeysToIntegrities.get(cacheKey);
    }
    /**
     * This acts as a drop-in replacement for
     * [`cache.match()`](https://developer.mozilla.org/en-US/docs/Web/API/Cache/match)
     * with the following differences:
     *
     * - It knows what the name of the precache is, and only checks in that cache.
     * - It allows you to pass in an "original" URL without versioning parameters,
     * and it will automatically look up the correct cache key for the currently
     * active revision of that URL.
     *
     * E.g., `matchPrecache('index.html')` will find the correct precached
     * response for the currently active service worker, even if the actual cache
     * key is `'/index.html?__WB_REVISION__=1234abcd'`.
     *
     * @param {string|Request} request The key (without revisioning parameters)
     * to look up in the precache.
     * @return {Promise<Response|undefined>}
     */
    async matchPrecache(request) {
      const url = request instanceof Request ? request.url : request;
      const cacheKey = this.getCacheKeyForURL(url);
      if (cacheKey) {
        const cache = await self.caches.open(this.strategy.cacheName);
        return cache.match(cacheKey);
      }
      return void 0;
    }
    /**
     * Returns a function that looks up `url` in the precache (taking into
     * account revision information), and returns the corresponding `Response`.
     *
     * @param {string} url The precached URL which will be used to lookup the
     * `Response`.
     * @return {workbox-routing~handlerCallback}
     */
    createHandlerBoundToURL(url) {
      const cacheKey = this.getCacheKeyForURL(url);
      if (!cacheKey) {
        throw new WorkboxError("non-precached-url", { url });
      }
      return (options) => {
        options.request = new Request(url);
        options.params = Object.assign({ cacheKey }, options.params);
        return this.strategy.handle(options);
      };
    }
  };

  // node_modules/.pnpm/workbox-precaching@7.3.0/node_modules/workbox-precaching/utils/getOrCreatePrecacheController.js
  var precacheController;
  var getOrCreatePrecacheController = () => {
    if (!precacheController) {
      precacheController = new PrecacheController();
    }
    return precacheController;
  };

  // node_modules/.pnpm/workbox-routing@7.3.0/node_modules/workbox-routing/_version.js
  try {
    self["workbox:routing:7.2.0"] && _();
  } catch (e) {
  }

  // node_modules/.pnpm/workbox-routing@7.3.0/node_modules/workbox-routing/utils/constants.js
  var defaultMethod2 = "GET";
  var validMethods2 = [
    "DELETE",
    "GET",
    "HEAD",
    "PATCH",
    "POST",
    "PUT"
  ];

  // node_modules/.pnpm/workbox-routing@7.3.0/node_modules/workbox-routing/utils/normalizeHandler.js
  var normalizeHandler2 = (handler) => {
    if (handler && typeof handler === "object") {
      if (true) {
        finalAssertExports.hasMethod(handler, "handle", {
          moduleName: "workbox-routing",
          className: "Route",
          funcName: "constructor",
          paramName: "handler"
        });
      }
      return handler;
    } else {
      if (true) {
        finalAssertExports.isType(handler, "function", {
          moduleName: "workbox-routing",
          className: "Route",
          funcName: "constructor",
          paramName: "handler"
        });
      }
      return { handle: handler };
    }
  };

  // node_modules/.pnpm/workbox-routing@7.3.0/node_modules/workbox-routing/Route.js
  var Route2 = class {
    /**
     * Constructor for Route class.
     *
     * @param {workbox-routing~matchCallback} match
     * A callback function that determines whether the route matches a given
     * `fetch` event by returning a non-falsy value.
     * @param {workbox-routing~handlerCallback} handler A callback
     * function that returns a Promise resolving to a Response.
     * @param {string} [method='GET'] The HTTP method to match the Route
     * against.
     */
    constructor(match, handler, method = defaultMethod2) {
      if (true) {
        finalAssertExports.isType(match, "function", {
          moduleName: "workbox-routing",
          className: "Route",
          funcName: "constructor",
          paramName: "match"
        });
        if (method) {
          finalAssertExports.isOneOf(method, validMethods2, { paramName: "method" });
        }
      }
      this.handler = normalizeHandler2(handler);
      this.match = match;
      this.method = method;
    }
    /**
     *
     * @param {workbox-routing-handlerCallback} handler A callback
     * function that returns a Promise resolving to a Response
     */
    setCatchHandler(handler) {
      this.catchHandler = normalizeHandler2(handler);
    }
  };

  // node_modules/.pnpm/workbox-routing@7.3.0/node_modules/workbox-routing/RegExpRoute.js
  var RegExpRoute2 = class extends Route2 {
    /**
     * If the regular expression contains
     * [capture groups]{@link https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/RegExp#grouping-back-references},
     * the captured values will be passed to the
     * {@link workbox-routing~handlerCallback} `params`
     * argument.
     *
     * @param {RegExp} regExp The regular expression to match against URLs.
     * @param {workbox-routing~handlerCallback} handler A callback
     * function that returns a Promise resulting in a Response.
     * @param {string} [method='GET'] The HTTP method to match the Route
     * against.
     */
    constructor(regExp, handler, method) {
      if (true) {
        finalAssertExports.isInstance(regExp, RegExp, {
          moduleName: "workbox-routing",
          className: "RegExpRoute",
          funcName: "constructor",
          paramName: "pattern"
        });
      }
      const match = ({ url }) => {
        const result = regExp.exec(url.href);
        if (!result) {
          return;
        }
        if (url.origin !== location.origin && result.index !== 0) {
          if (true) {
            logger.debug(`The regular expression '${regExp.toString()}' only partially matched against the cross-origin URL '${url.toString()}'. RegExpRoute's will only handle cross-origin requests if they match the entire URL.`);
          }
          return;
        }
        return result.slice(1);
      };
      super(match, handler, method);
    }
  };

  // node_modules/.pnpm/workbox-routing@7.3.0/node_modules/workbox-routing/Router.js
  var Router2 = class {
    /**
     * Initializes a new Router.
     */
    constructor() {
      this._routes = /* @__PURE__ */ new Map();
      this._defaultHandlerMap = /* @__PURE__ */ new Map();
    }
    /**
     * @return {Map<string, Array<workbox-routing.Route>>} routes A `Map` of HTTP
     * method name ('GET', etc.) to an array of all the corresponding `Route`
     * instances that are registered.
     */
    get routes() {
      return this._routes;
    }
    /**
     * Adds a fetch event listener to respond to events when a route matches
     * the event's request.
     */
    addFetchListener() {
      self.addEventListener("fetch", (event) => {
        const { request } = event;
        const responsePromise = this.handleRequest({ request, event });
        if (responsePromise) {
          event.respondWith(responsePromise);
        }
      });
    }
    /**
     * Adds a message event listener for URLs to cache from the window.
     * This is useful to cache resources loaded on the page prior to when the
     * service worker started controlling it.
     *
     * The format of the message data sent from the window should be as follows.
     * Where the `urlsToCache` array may consist of URL strings or an array of
     * URL string + `requestInit` object (the same as you'd pass to `fetch()`).
     *
     * ```
     * {
     *   type: 'CACHE_URLS',
     *   payload: {
     *     urlsToCache: [
     *       './script1.js',
     *       './script2.js',
     *       ['./script3.js', {mode: 'no-cors'}],
     *     ],
     *   },
     * }
     * ```
     */
    addCacheListener() {
      self.addEventListener("message", (event) => {
        if (event.data && event.data.type === "CACHE_URLS") {
          const { payload } = event.data;
          if (true) {
            logger.debug(`Caching URLs from the window`, payload.urlsToCache);
          }
          const requestPromises = Promise.all(payload.urlsToCache.map((entry) => {
            if (typeof entry === "string") {
              entry = [entry];
            }
            const request = new Request(...entry);
            return this.handleRequest({ request, event });
          }));
          event.waitUntil(requestPromises);
          if (event.ports && event.ports[0]) {
            void requestPromises.then(() => event.ports[0].postMessage(true));
          }
        }
      });
    }
    /**
     * Apply the routing rules to a FetchEvent object to get a Response from an
     * appropriate Route's handler.
     *
     * @param {Object} options
     * @param {Request} options.request The request to handle.
     * @param {ExtendableEvent} options.event The event that triggered the
     *     request.
     * @return {Promise<Response>|undefined} A promise is returned if a
     *     registered route can handle the request. If there is no matching
     *     route and there's no `defaultHandler`, `undefined` is returned.
     */
    handleRequest({ request, event }) {
      if (true) {
        finalAssertExports.isInstance(request, Request, {
          moduleName: "workbox-routing",
          className: "Router",
          funcName: "handleRequest",
          paramName: "options.request"
        });
      }
      const url = new URL(request.url, location.href);
      if (!url.protocol.startsWith("http")) {
        if (true) {
          logger.debug(`Workbox Router only supports URLs that start with 'http'.`);
        }
        return;
      }
      const sameOrigin = url.origin === location.origin;
      const { params, route } = this.findMatchingRoute({
        event,
        request,
        sameOrigin,
        url
      });
      let handler = route && route.handler;
      const debugMessages = [];
      if (true) {
        if (handler) {
          debugMessages.push([`Found a route to handle this request:`, route]);
          if (params) {
            debugMessages.push([
              `Passing the following params to the route's handler:`,
              params
            ]);
          }
        }
      }
      const method = request.method;
      if (!handler && this._defaultHandlerMap.has(method)) {
        if (true) {
          debugMessages.push(`Failed to find a matching route. Falling back to the default handler for ${method}.`);
        }
        handler = this._defaultHandlerMap.get(method);
      }
      if (!handler) {
        if (true) {
          logger.debug(`No route found for: ${getFriendlyURL(url)}`);
        }
        return;
      }
      if (true) {
        logger.groupCollapsed(`Router is responding to: ${getFriendlyURL(url)}`);
        debugMessages.forEach((msg) => {
          if (Array.isArray(msg)) {
            logger.log(...msg);
          } else {
            logger.log(msg);
          }
        });
        logger.groupEnd();
      }
      let responsePromise;
      try {
        responsePromise = handler.handle({ url, request, event, params });
      } catch (err) {
        responsePromise = Promise.reject(err);
      }
      const catchHandler = route && route.catchHandler;
      if (responsePromise instanceof Promise && (this._catchHandler || catchHandler)) {
        responsePromise = responsePromise.catch(async (err) => {
          if (catchHandler) {
            if (true) {
              logger.groupCollapsed(`Error thrown when responding to:  ${getFriendlyURL(url)}. Falling back to route's Catch Handler.`);
              logger.error(`Error thrown by:`, route);
              logger.error(err);
              logger.groupEnd();
            }
            try {
              return await catchHandler.handle({ url, request, event, params });
            } catch (catchErr) {
              if (catchErr instanceof Error) {
                err = catchErr;
              }
            }
          }
          if (this._catchHandler) {
            if (true) {
              logger.groupCollapsed(`Error thrown when responding to:  ${getFriendlyURL(url)}. Falling back to global Catch Handler.`);
              logger.error(`Error thrown by:`, route);
              logger.error(err);
              logger.groupEnd();
            }
            return this._catchHandler.handle({ url, request, event });
          }
          throw err;
        });
      }
      return responsePromise;
    }
    /**
     * Checks a request and URL (and optionally an event) against the list of
     * registered routes, and if there's a match, returns the corresponding
     * route along with any params generated by the match.
     *
     * @param {Object} options
     * @param {URL} options.url
     * @param {boolean} options.sameOrigin The result of comparing `url.origin`
     *     against the current origin.
     * @param {Request} options.request The request to match.
     * @param {Event} options.event The corresponding event.
     * @return {Object} An object with `route` and `params` properties.
     *     They are populated if a matching route was found or `undefined`
     *     otherwise.
     */
    findMatchingRoute({ url, sameOrigin, request, event }) {
      const routes = this._routes.get(request.method) || [];
      for (const route of routes) {
        let params;
        const matchResult = route.match({ url, sameOrigin, request, event });
        if (matchResult) {
          if (true) {
            if (matchResult instanceof Promise) {
              logger.warn(`While routing ${getFriendlyURL(url)}, an async matchCallback function was used. Please convert the following route to use a synchronous matchCallback function:`, route);
            }
          }
          params = matchResult;
          if (Array.isArray(params) && params.length === 0) {
            params = void 0;
          } else if (matchResult.constructor === Object && // eslint-disable-line
          Object.keys(matchResult).length === 0) {
            params = void 0;
          } else if (typeof matchResult === "boolean") {
            params = void 0;
          }
          return { route, params };
        }
      }
      return {};
    }
    /**
     * Define a default `handler` that's called when no routes explicitly
     * match the incoming request.
     *
     * Each HTTP method ('GET', 'POST', etc.) gets its own default handler.
     *
     * Without a default handler, unmatched requests will go against the
     * network as if there were no service worker present.
     *
     * @param {workbox-routing~handlerCallback} handler A callback
     * function that returns a Promise resulting in a Response.
     * @param {string} [method='GET'] The HTTP method to associate with this
     * default handler. Each method has its own default.
     */
    setDefaultHandler(handler, method = defaultMethod2) {
      this._defaultHandlerMap.set(method, normalizeHandler2(handler));
    }
    /**
     * If a Route throws an error while handling a request, this `handler`
     * will be called and given a chance to provide a response.
     *
     * @param {workbox-routing~handlerCallback} handler A callback
     * function that returns a Promise resulting in a Response.
     */
    setCatchHandler(handler) {
      this._catchHandler = normalizeHandler2(handler);
    }
    /**
     * Registers a route with the router.
     *
     * @param {workbox-routing.Route} route The route to register.
     */
    registerRoute(route) {
      if (true) {
        finalAssertExports.isType(route, "object", {
          moduleName: "workbox-routing",
          className: "Router",
          funcName: "registerRoute",
          paramName: "route"
        });
        finalAssertExports.hasMethod(route, "match", {
          moduleName: "workbox-routing",
          className: "Router",
          funcName: "registerRoute",
          paramName: "route"
        });
        finalAssertExports.isType(route.handler, "object", {
          moduleName: "workbox-routing",
          className: "Router",
          funcName: "registerRoute",
          paramName: "route"
        });
        finalAssertExports.hasMethod(route.handler, "handle", {
          moduleName: "workbox-routing",
          className: "Router",
          funcName: "registerRoute",
          paramName: "route.handler"
        });
        finalAssertExports.isType(route.method, "string", {
          moduleName: "workbox-routing",
          className: "Router",
          funcName: "registerRoute",
          paramName: "route.method"
        });
      }
      if (!this._routes.has(route.method)) {
        this._routes.set(route.method, []);
      }
      this._routes.get(route.method).push(route);
    }
    /**
     * Unregisters a route with the router.
     *
     * @param {workbox-routing.Route} route The route to unregister.
     */
    unregisterRoute(route) {
      if (!this._routes.has(route.method)) {
        throw new WorkboxError("unregister-route-but-not-found-with-method", {
          method: route.method
        });
      }
      const routeIndex = this._routes.get(route.method).indexOf(route);
      if (routeIndex > -1) {
        this._routes.get(route.method).splice(routeIndex, 1);
      } else {
        throw new WorkboxError("unregister-route-route-not-registered");
      }
    }
  };

  // node_modules/.pnpm/workbox-routing@7.3.0/node_modules/workbox-routing/utils/getOrCreateDefaultRouter.js
  var defaultRouter2;
  var getOrCreateDefaultRouter2 = () => {
    if (!defaultRouter2) {
      defaultRouter2 = new Router2();
      defaultRouter2.addFetchListener();
      defaultRouter2.addCacheListener();
    }
    return defaultRouter2;
  };

  // node_modules/.pnpm/workbox-routing@7.3.0/node_modules/workbox-routing/registerRoute.js
  function registerRoute2(capture, handler, method) {
    let route;
    if (typeof capture === "string") {
      const captureUrl = new URL(capture, location.href);
      if (true) {
        if (!(capture.startsWith("/") || capture.startsWith("http"))) {
          throw new WorkboxError("invalid-string", {
            moduleName: "workbox-routing",
            funcName: "registerRoute",
            paramName: "capture"
          });
        }
        const valueToCheck = capture.startsWith("http") ? captureUrl.pathname : capture;
        const wildcards = "[*:?+]";
        if (new RegExp(`${wildcards}`).exec(valueToCheck)) {
          logger.debug(`The '$capture' parameter contains an Express-style wildcard character (${wildcards}). Strings are now always interpreted as exact matches; use a RegExp for partial or wildcard matches.`);
        }
      }
      const matchCallback = ({ url }) => {
        if (true) {
          if (url.pathname === captureUrl.pathname && url.origin !== captureUrl.origin) {
            logger.debug(`${capture} only partially matches the cross-origin URL ${url.toString()}. This route will only handle cross-origin requests if they match the entire URL.`);
          }
        }
        return url.href === captureUrl.href;
      };
      route = new Route2(matchCallback, handler, method);
    } else if (capture instanceof RegExp) {
      route = new RegExpRoute2(capture, handler, method);
    } else if (typeof capture === "function") {
      route = new Route2(capture, handler, method);
    } else if (capture instanceof Route2) {
      route = capture;
    } else {
      throw new WorkboxError("unsupported-route-type", {
        moduleName: "workbox-routing",
        funcName: "registerRoute",
        paramName: "capture"
      });
    }
    const defaultRouter3 = getOrCreateDefaultRouter2();
    defaultRouter3.registerRoute(route);
    return route;
  }

  // node_modules/.pnpm/workbox-precaching@7.3.0/node_modules/workbox-precaching/utils/removeIgnoredSearchParams.js
  function removeIgnoredSearchParams(urlObject, ignoreURLParametersMatching = []) {
    for (const paramName of [...urlObject.searchParams.keys()]) {
      if (ignoreURLParametersMatching.some((regExp) => regExp.test(paramName))) {
        urlObject.searchParams.delete(paramName);
      }
    }
    return urlObject;
  }

  // node_modules/.pnpm/workbox-precaching@7.3.0/node_modules/workbox-precaching/utils/generateURLVariations.js
  function* generateURLVariations(url, { ignoreURLParametersMatching = [/^utm_/, /^fbclid$/], directoryIndex = "index.html", cleanURLs = true, urlManipulation } = {}) {
    const urlObject = new URL(url, location.href);
    urlObject.hash = "";
    yield urlObject.href;
    const urlWithoutIgnoredParams = removeIgnoredSearchParams(urlObject, ignoreURLParametersMatching);
    yield urlWithoutIgnoredParams.href;
    if (directoryIndex && urlWithoutIgnoredParams.pathname.endsWith("/")) {
      const directoryURL = new URL(urlWithoutIgnoredParams.href);
      directoryURL.pathname += directoryIndex;
      yield directoryURL.href;
    }
    if (cleanURLs) {
      const cleanURL = new URL(urlWithoutIgnoredParams.href);
      cleanURL.pathname += ".html";
      yield cleanURL.href;
    }
    if (urlManipulation) {
      const additionalURLs = urlManipulation({ url: urlObject });
      for (const urlToAttempt of additionalURLs) {
        yield urlToAttempt.href;
      }
    }
  }

  // node_modules/.pnpm/workbox-precaching@7.3.0/node_modules/workbox-precaching/PrecacheRoute.js
  var PrecacheRoute = class extends Route2 {
    /**
     * @param {PrecacheController} precacheController A `PrecacheController`
     * instance used to both match requests and respond to fetch events.
     * @param {Object} [options] Options to control how requests are matched
     * against the list of precached URLs.
     * @param {string} [options.directoryIndex=index.html] The `directoryIndex` will
     * check cache entries for a URLs ending with '/' to see if there is a hit when
     * appending the `directoryIndex` value.
     * @param {Array<RegExp>} [options.ignoreURLParametersMatching=[/^utm_/, /^fbclid$/]] An
     * array of regex's to remove search params when looking for a cache match.
     * @param {boolean} [options.cleanURLs=true] The `cleanURLs` option will
     * check the cache for the URL with a `.html` added to the end of the end.
     * @param {workbox-precaching~urlManipulation} [options.urlManipulation]
     * This is a function that should take a URL and return an array of
     * alternative URLs that should be checked for precache matches.
     */
    constructor(precacheController2, options) {
      const match = ({ request }) => {
        const urlsToCacheKeys = precacheController2.getURLsToCacheKeys();
        for (const possibleURL of generateURLVariations(request.url, options)) {
          const cacheKey = urlsToCacheKeys.get(possibleURL);
          if (cacheKey) {
            const integrity = precacheController2.getIntegrityForCacheKey(cacheKey);
            return { cacheKey, integrity };
          }
        }
        if (true) {
          logger.debug(`Precaching did not find a match for ` + getFriendlyURL(request.url));
        }
        return;
      };
      super(match, precacheController2.strategy);
    }
  };

  // node_modules/.pnpm/workbox-precaching@7.3.0/node_modules/workbox-precaching/addRoute.js
  function addRoute(options) {
    const precacheController2 = getOrCreatePrecacheController();
    const precacheRoute = new PrecacheRoute(precacheController2, options);
    registerRoute2(precacheRoute);
  }

  // node_modules/.pnpm/workbox-precaching@7.3.0/node_modules/workbox-precaching/precache.js
  function precache(entries) {
    const precacheController2 = getOrCreatePrecacheController();
    precacheController2.precache(entries);
  }

  // node_modules/.pnpm/workbox-precaching@7.3.0/node_modules/workbox-precaching/precacheAndRoute.js
  function precacheAndRoute(entries, options) {
    precache(entries);
    addRoute(options);
  }

  // node_modules/.pnpm/workbox-core@7.3.0/node_modules/workbox-core/_private/dontWaitFor.js
  function dontWaitFor(promise) {
    void promise.then(() => {
    });
  }

  // node_modules/.pnpm/idb@7.1.1/node_modules/idb/build/wrap-idb-value.js
  var instanceOfAny = (object, constructors) => constructors.some((c) => object instanceof c);
  var idbProxyableTypes;
  var cursorAdvanceMethods;
  function getIdbProxyableTypes() {
    return idbProxyableTypes || (idbProxyableTypes = [
      IDBDatabase,
      IDBObjectStore,
      IDBIndex,
      IDBCursor,
      IDBTransaction
    ]);
  }
  function getCursorAdvanceMethods() {
    return cursorAdvanceMethods || (cursorAdvanceMethods = [
      IDBCursor.prototype.advance,
      IDBCursor.prototype.continue,
      IDBCursor.prototype.continuePrimaryKey
    ]);
  }
  var cursorRequestMap = /* @__PURE__ */ new WeakMap();
  var transactionDoneMap = /* @__PURE__ */ new WeakMap();
  var transactionStoreNamesMap = /* @__PURE__ */ new WeakMap();
  var transformCache = /* @__PURE__ */ new WeakMap();
  var reverseTransformCache = /* @__PURE__ */ new WeakMap();
  function promisifyRequest(request) {
    const promise = new Promise((resolve, reject) => {
      const unlisten = () => {
        request.removeEventListener("success", success);
        request.removeEventListener("error", error);
      };
      const success = () => {
        resolve(wrap(request.result));
        unlisten();
      };
      const error = () => {
        reject(request.error);
        unlisten();
      };
      request.addEventListener("success", success);
      request.addEventListener("error", error);
    });
    promise.then((value) => {
      if (value instanceof IDBCursor) {
        cursorRequestMap.set(value, request);
      }
    }).catch(() => {
    });
    reverseTransformCache.set(promise, request);
    return promise;
  }
  function cacheDonePromiseForTransaction(tx) {
    if (transactionDoneMap.has(tx))
      return;
    const done = new Promise((resolve, reject) => {
      const unlisten = () => {
        tx.removeEventListener("complete", complete);
        tx.removeEventListener("error", error);
        tx.removeEventListener("abort", error);
      };
      const complete = () => {
        resolve();
        unlisten();
      };
      const error = () => {
        reject(tx.error || new DOMException("AbortError", "AbortError"));
        unlisten();
      };
      tx.addEventListener("complete", complete);
      tx.addEventListener("error", error);
      tx.addEventListener("abort", error);
    });
    transactionDoneMap.set(tx, done);
  }
  var idbProxyTraps = {
    get(target, prop, receiver) {
      if (target instanceof IDBTransaction) {
        if (prop === "done")
          return transactionDoneMap.get(target);
        if (prop === "objectStoreNames") {
          return target.objectStoreNames || transactionStoreNamesMap.get(target);
        }
        if (prop === "store") {
          return receiver.objectStoreNames[1] ? void 0 : receiver.objectStore(receiver.objectStoreNames[0]);
        }
      }
      return wrap(target[prop]);
    },
    set(target, prop, value) {
      target[prop] = value;
      return true;
    },
    has(target, prop) {
      if (target instanceof IDBTransaction && (prop === "done" || prop === "store")) {
        return true;
      }
      return prop in target;
    }
  };
  function replaceTraps(callback) {
    idbProxyTraps = callback(idbProxyTraps);
  }
  function wrapFunction(func) {
    if (func === IDBDatabase.prototype.transaction && !("objectStoreNames" in IDBTransaction.prototype)) {
      return function(storeNames, ...args) {
        const tx = func.call(unwrap(this), storeNames, ...args);
        transactionStoreNamesMap.set(tx, storeNames.sort ? storeNames.sort() : [storeNames]);
        return wrap(tx);
      };
    }
    if (getCursorAdvanceMethods().includes(func)) {
      return function(...args) {
        func.apply(unwrap(this), args);
        return wrap(cursorRequestMap.get(this));
      };
    }
    return function(...args) {
      return wrap(func.apply(unwrap(this), args));
    };
  }
  function transformCachableValue(value) {
    if (typeof value === "function")
      return wrapFunction(value);
    if (value instanceof IDBTransaction)
      cacheDonePromiseForTransaction(value);
    if (instanceOfAny(value, getIdbProxyableTypes()))
      return new Proxy(value, idbProxyTraps);
    return value;
  }
  function wrap(value) {
    if (value instanceof IDBRequest)
      return promisifyRequest(value);
    if (transformCache.has(value))
      return transformCache.get(value);
    const newValue = transformCachableValue(value);
    if (newValue !== value) {
      transformCache.set(value, newValue);
      reverseTransformCache.set(newValue, value);
    }
    return newValue;
  }
  var unwrap = (value) => reverseTransformCache.get(value);

  // node_modules/.pnpm/idb@7.1.1/node_modules/idb/build/index.js
  function openDB(name, version, { blocked, upgrade, blocking, terminated } = {}) {
    const request = indexedDB.open(name, version);
    const openPromise = wrap(request);
    if (upgrade) {
      request.addEventListener("upgradeneeded", (event) => {
        upgrade(wrap(request.result), event.oldVersion, event.newVersion, wrap(request.transaction), event);
      });
    }
    if (blocked) {
      request.addEventListener("blocked", (event) => blocked(
        // Casting due to https://github.com/microsoft/TypeScript-DOM-lib-generator/pull/1405
        event.oldVersion,
        event.newVersion,
        event
      ));
    }
    openPromise.then((db) => {
      if (terminated)
        db.addEventListener("close", () => terminated());
      if (blocking) {
        db.addEventListener("versionchange", (event) => blocking(event.oldVersion, event.newVersion, event));
      }
    }).catch(() => {
    });
    return openPromise;
  }
  function deleteDB(name, { blocked } = {}) {
    const request = indexedDB.deleteDatabase(name);
    if (blocked) {
      request.addEventListener("blocked", (event) => blocked(
        // Casting due to https://github.com/microsoft/TypeScript-DOM-lib-generator/pull/1405
        event.oldVersion,
        event
      ));
    }
    return wrap(request).then(() => void 0);
  }
  var readMethods = ["get", "getKey", "getAll", "getAllKeys", "count"];
  var writeMethods = ["put", "add", "delete", "clear"];
  var cachedMethods = /* @__PURE__ */ new Map();
  function getMethod(target, prop) {
    if (!(target instanceof IDBDatabase && !(prop in target) && typeof prop === "string")) {
      return;
    }
    if (cachedMethods.get(prop))
      return cachedMethods.get(prop);
    const targetFuncName = prop.replace(/FromIndex$/, "");
    const useIndex = prop !== targetFuncName;
    const isWrite = writeMethods.includes(targetFuncName);
    if (
      // Bail if the target doesn't exist on the target. Eg, getAll isn't in Edge.
      !(targetFuncName in (useIndex ? IDBIndex : IDBObjectStore).prototype) || !(isWrite || readMethods.includes(targetFuncName))
    ) {
      return;
    }
    const method = async function(storeName, ...args) {
      const tx = this.transaction(storeName, isWrite ? "readwrite" : "readonly");
      let target2 = tx.store;
      if (useIndex)
        target2 = target2.index(args.shift());
      return (await Promise.all([
        target2[targetFuncName](...args),
        isWrite && tx.done
      ]))[0];
    };
    cachedMethods.set(prop, method);
    return method;
  }
  replaceTraps((oldTraps) => ({
    ...oldTraps,
    get: (target, prop, receiver) => getMethod(target, prop) || oldTraps.get(target, prop, receiver),
    has: (target, prop) => !!getMethod(target, prop) || oldTraps.has(target, prop)
  }));

  // node_modules/.pnpm/workbox-expiration@7.3.0/node_modules/workbox-expiration/_version.js
  try {
    self["workbox:expiration:7.2.0"] && _();
  } catch (e) {
  }

  // node_modules/.pnpm/workbox-expiration@7.3.0/node_modules/workbox-expiration/models/CacheTimestampsModel.js
  var DB_NAME = "workbox-expiration";
  var CACHE_OBJECT_STORE = "cache-entries";
  var normalizeURL = (unNormalizedUrl) => {
    const url = new URL(unNormalizedUrl, location.href);
    url.hash = "";
    return url.href;
  };
  var CacheTimestampsModel = class {
    /**
     *
     * @param {string} cacheName
     *
     * @private
     */
    constructor(cacheName) {
      this._db = null;
      this._cacheName = cacheName;
    }
    /**
     * Performs an upgrade of indexedDB.
     *
     * @param {IDBPDatabase<CacheDbSchema>} db
     *
     * @private
     */
    _upgradeDb(db) {
      const objStore = db.createObjectStore(CACHE_OBJECT_STORE, { keyPath: "id" });
      objStore.createIndex("cacheName", "cacheName", { unique: false });
      objStore.createIndex("timestamp", "timestamp", { unique: false });
    }
    /**
     * Performs an upgrade of indexedDB and deletes deprecated DBs.
     *
     * @param {IDBPDatabase<CacheDbSchema>} db
     *
     * @private
     */
    _upgradeDbAndDeleteOldDbs(db) {
      this._upgradeDb(db);
      if (this._cacheName) {
        void deleteDB(this._cacheName);
      }
    }
    /**
     * @param {string} url
     * @param {number} timestamp
     *
     * @private
     */
    async setTimestamp(url, timestamp) {
      url = normalizeURL(url);
      const entry = {
        url,
        timestamp,
        cacheName: this._cacheName,
        // Creating an ID from the URL and cache name won't be necessary once
        // Edge switches to Chromium and all browsers we support work with
        // array keyPaths.
        id: this._getId(url)
      };
      const db = await this.getDb();
      const tx = db.transaction(CACHE_OBJECT_STORE, "readwrite", {
        durability: "relaxed"
      });
      await tx.store.put(entry);
      await tx.done;
    }
    /**
     * Returns the timestamp stored for a given URL.
     *
     * @param {string} url
     * @return {number | undefined}
     *
     * @private
     */
    async getTimestamp(url) {
      const db = await this.getDb();
      const entry = await db.get(CACHE_OBJECT_STORE, this._getId(url));
      return entry === null || entry === void 0 ? void 0 : entry.timestamp;
    }
    /**
     * Iterates through all the entries in the object store (from newest to
     * oldest) and removes entries once either `maxCount` is reached or the
     * entry's timestamp is less than `minTimestamp`.
     *
     * @param {number} minTimestamp
     * @param {number} maxCount
     * @return {Array<string>}
     *
     * @private
     */
    async expireEntries(minTimestamp, maxCount) {
      const db = await this.getDb();
      let cursor = await db.transaction(CACHE_OBJECT_STORE).store.index("timestamp").openCursor(null, "prev");
      const entriesToDelete = [];
      let entriesNotDeletedCount = 0;
      while (cursor) {
        const result = cursor.value;
        if (result.cacheName === this._cacheName) {
          if (minTimestamp && result.timestamp < minTimestamp || maxCount && entriesNotDeletedCount >= maxCount) {
            entriesToDelete.push(cursor.value);
          } else {
            entriesNotDeletedCount++;
          }
        }
        cursor = await cursor.continue();
      }
      const urlsDeleted = [];
      for (const entry of entriesToDelete) {
        await db.delete(CACHE_OBJECT_STORE, entry.id);
        urlsDeleted.push(entry.url);
      }
      return urlsDeleted;
    }
    /**
     * Takes a URL and returns an ID that will be unique in the object store.
     *
     * @param {string} url
     * @return {string}
     *
     * @private
     */
    _getId(url) {
      return this._cacheName + "|" + normalizeURL(url);
    }
    /**
     * Returns an open connection to the database.
     *
     * @private
     */
    async getDb() {
      if (!this._db) {
        this._db = await openDB(DB_NAME, 1, {
          upgrade: this._upgradeDbAndDeleteOldDbs.bind(this)
        });
      }
      return this._db;
    }
  };

  // node_modules/.pnpm/workbox-expiration@7.3.0/node_modules/workbox-expiration/CacheExpiration.js
  var CacheExpiration = class {
    /**
     * To construct a new CacheExpiration instance you must provide at least
     * one of the `config` properties.
     *
     * @param {string} cacheName Name of the cache to apply restrictions to.
     * @param {Object} config
     * @param {number} [config.maxEntries] The maximum number of entries to cache.
     * Entries used the least will be removed as the maximum is reached.
     * @param {number} [config.maxAgeSeconds] The maximum age of an entry before
     * it's treated as stale and removed.
     * @param {Object} [config.matchOptions] The [`CacheQueryOptions`](https://developer.mozilla.org/en-US/docs/Web/API/Cache/delete#Parameters)
     * that will be used when calling `delete()` on the cache.
     */
    constructor(cacheName, config = {}) {
      this._isRunning = false;
      this._rerunRequested = false;
      if (true) {
        finalAssertExports.isType(cacheName, "string", {
          moduleName: "workbox-expiration",
          className: "CacheExpiration",
          funcName: "constructor",
          paramName: "cacheName"
        });
        if (!(config.maxEntries || config.maxAgeSeconds)) {
          throw new WorkboxError("max-entries-or-age-required", {
            moduleName: "workbox-expiration",
            className: "CacheExpiration",
            funcName: "constructor"
          });
        }
        if (config.maxEntries) {
          finalAssertExports.isType(config.maxEntries, "number", {
            moduleName: "workbox-expiration",
            className: "CacheExpiration",
            funcName: "constructor",
            paramName: "config.maxEntries"
          });
        }
        if (config.maxAgeSeconds) {
          finalAssertExports.isType(config.maxAgeSeconds, "number", {
            moduleName: "workbox-expiration",
            className: "CacheExpiration",
            funcName: "constructor",
            paramName: "config.maxAgeSeconds"
          });
        }
      }
      this._maxEntries = config.maxEntries;
      this._maxAgeSeconds = config.maxAgeSeconds;
      this._matchOptions = config.matchOptions;
      this._cacheName = cacheName;
      this._timestampModel = new CacheTimestampsModel(cacheName);
    }
    /**
     * Expires entries for the given cache and given criteria.
     */
    async expireEntries() {
      if (this._isRunning) {
        this._rerunRequested = true;
        return;
      }
      this._isRunning = true;
      const minTimestamp = this._maxAgeSeconds ? Date.now() - this._maxAgeSeconds * 1e3 : 0;
      const urlsExpired = await this._timestampModel.expireEntries(minTimestamp, this._maxEntries);
      const cache = await self.caches.open(this._cacheName);
      for (const url of urlsExpired) {
        await cache.delete(url, this._matchOptions);
      }
      if (true) {
        if (urlsExpired.length > 0) {
          logger.groupCollapsed(`Expired ${urlsExpired.length} ${urlsExpired.length === 1 ? "entry" : "entries"} and removed ${urlsExpired.length === 1 ? "it" : "them"} from the '${this._cacheName}' cache.`);
          logger.log(`Expired the following ${urlsExpired.length === 1 ? "URL" : "URLs"}:`);
          urlsExpired.forEach((url) => logger.log(`    ${url}`));
          logger.groupEnd();
        } else {
          logger.debug(`Cache expiration ran and found no entries to remove.`);
        }
      }
      this._isRunning = false;
      if (this._rerunRequested) {
        this._rerunRequested = false;
        dontWaitFor(this.expireEntries());
      }
    }
    /**
     * Update the timestamp for the given URL. This ensures the when
     * removing entries based on maximum entries, most recently used
     * is accurate or when expiring, the timestamp is up-to-date.
     *
     * @param {string} url
     */
    async updateTimestamp(url) {
      if (true) {
        finalAssertExports.isType(url, "string", {
          moduleName: "workbox-expiration",
          className: "CacheExpiration",
          funcName: "updateTimestamp",
          paramName: "url"
        });
      }
      await this._timestampModel.setTimestamp(url, Date.now());
    }
    /**
     * Can be used to check if a URL has expired or not before it's used.
     *
     * This requires a look up from IndexedDB, so can be slow.
     *
     * Note: This method will not remove the cached entry, call
     * `expireEntries()` to remove indexedDB and Cache entries.
     *
     * @param {string} url
     * @return {boolean}
     */
    async isURLExpired(url) {
      if (!this._maxAgeSeconds) {
        if (true) {
          throw new WorkboxError(`expired-test-without-max-age`, {
            methodName: "isURLExpired",
            paramName: "maxAgeSeconds"
          });
        }
        return false;
      } else {
        const timestamp = await this._timestampModel.getTimestamp(url);
        const expireOlderThan = Date.now() - this._maxAgeSeconds * 1e3;
        return timestamp !== void 0 ? timestamp < expireOlderThan : true;
      }
    }
    /**
     * Removes the IndexedDB object store used to keep track of cache expiration
     * metadata.
     */
    async delete() {
      this._rerunRequested = false;
      await this._timestampModel.expireEntries(Infinity);
    }
  };

  // node_modules/.pnpm/workbox-core@7.3.0/node_modules/workbox-core/registerQuotaErrorCallback.js
  function registerQuotaErrorCallback(callback) {
    if (true) {
      finalAssertExports.isType(callback, "function", {
        moduleName: "workbox-core",
        funcName: "register",
        paramName: "callback"
      });
    }
    quotaErrorCallbacks.add(callback);
    if (true) {
      logger.log("Registered a callback to respond to quota errors.", callback);
    }
  }

  // node_modules/.pnpm/workbox-expiration@7.3.0/node_modules/workbox-expiration/ExpirationPlugin.js
  var ExpirationPlugin = class {
    /**
     * @param {ExpirationPluginOptions} config
     * @param {number} [config.maxEntries] The maximum number of entries to cache.
     * Entries used the least will be removed as the maximum is reached.
     * @param {number} [config.maxAgeSeconds] The maximum age of an entry before
     * it's treated as stale and removed.
     * @param {Object} [config.matchOptions] The [`CacheQueryOptions`](https://developer.mozilla.org/en-US/docs/Web/API/Cache/delete#Parameters)
     * that will be used when calling `delete()` on the cache.
     * @param {boolean} [config.purgeOnQuotaError] Whether to opt this cache in to
     * automatic deletion if the available storage quota has been exceeded.
     */
    constructor(config = {}) {
      this.cachedResponseWillBeUsed = async ({ event, request, cacheName, cachedResponse }) => {
        if (!cachedResponse) {
          return null;
        }
        const isFresh = this._isResponseDateFresh(cachedResponse);
        const cacheExpiration = this._getCacheExpiration(cacheName);
        dontWaitFor(cacheExpiration.expireEntries());
        const updateTimestampDone = cacheExpiration.updateTimestamp(request.url);
        if (event) {
          try {
            event.waitUntil(updateTimestampDone);
          } catch (error) {
            if (true) {
              if ("request" in event) {
                logger.warn(`Unable to ensure service worker stays alive when updating cache entry for '${getFriendlyURL(event.request.url)}'.`);
              }
            }
          }
        }
        return isFresh ? cachedResponse : null;
      };
      this.cacheDidUpdate = async ({ cacheName, request }) => {
        if (true) {
          finalAssertExports.isType(cacheName, "string", {
            moduleName: "workbox-expiration",
            className: "Plugin",
            funcName: "cacheDidUpdate",
            paramName: "cacheName"
          });
          finalAssertExports.isInstance(request, Request, {
            moduleName: "workbox-expiration",
            className: "Plugin",
            funcName: "cacheDidUpdate",
            paramName: "request"
          });
        }
        const cacheExpiration = this._getCacheExpiration(cacheName);
        await cacheExpiration.updateTimestamp(request.url);
        await cacheExpiration.expireEntries();
      };
      if (true) {
        if (!(config.maxEntries || config.maxAgeSeconds)) {
          throw new WorkboxError("max-entries-or-age-required", {
            moduleName: "workbox-expiration",
            className: "Plugin",
            funcName: "constructor"
          });
        }
        if (config.maxEntries) {
          finalAssertExports.isType(config.maxEntries, "number", {
            moduleName: "workbox-expiration",
            className: "Plugin",
            funcName: "constructor",
            paramName: "config.maxEntries"
          });
        }
        if (config.maxAgeSeconds) {
          finalAssertExports.isType(config.maxAgeSeconds, "number", {
            moduleName: "workbox-expiration",
            className: "Plugin",
            funcName: "constructor",
            paramName: "config.maxAgeSeconds"
          });
        }
      }
      this._config = config;
      this._maxAgeSeconds = config.maxAgeSeconds;
      this._cacheExpirations = /* @__PURE__ */ new Map();
      if (config.purgeOnQuotaError) {
        registerQuotaErrorCallback(() => this.deleteCacheAndMetadata());
      }
    }
    /**
     * A simple helper method to return a CacheExpiration instance for a given
     * cache name.
     *
     * @param {string} cacheName
     * @return {CacheExpiration}
     *
     * @private
     */
    _getCacheExpiration(cacheName) {
      if (cacheName === cacheNames.getRuntimeName()) {
        throw new WorkboxError("expire-custom-caches-only");
      }
      let cacheExpiration = this._cacheExpirations.get(cacheName);
      if (!cacheExpiration) {
        cacheExpiration = new CacheExpiration(cacheName, this._config);
        this._cacheExpirations.set(cacheName, cacheExpiration);
      }
      return cacheExpiration;
    }
    /**
     * @param {Response} cachedResponse
     * @return {boolean}
     *
     * @private
     */
    _isResponseDateFresh(cachedResponse) {
      if (!this._maxAgeSeconds) {
        return true;
      }
      const dateHeaderTimestamp = this._getDateHeaderTimestamp(cachedResponse);
      if (dateHeaderTimestamp === null) {
        return true;
      }
      const now = Date.now();
      return dateHeaderTimestamp >= now - this._maxAgeSeconds * 1e3;
    }
    /**
     * This method will extract the data header and parse it into a useful
     * value.
     *
     * @param {Response} cachedResponse
     * @return {number|null}
     *
     * @private
     */
    _getDateHeaderTimestamp(cachedResponse) {
      if (!cachedResponse.headers.has("date")) {
        return null;
      }
      const dateHeader = cachedResponse.headers.get("date");
      const parsedDate = new Date(dateHeader);
      const headerTime = parsedDate.getTime();
      if (isNaN(headerTime)) {
        return null;
      }
      return headerTime;
    }
    /**
     * This is a helper method that performs two operations:
     *
     * - Deletes *all* the underlying Cache instances associated with this plugin
     * instance, by calling caches.delete() on your behalf.
     * - Deletes the metadata from IndexedDB used to keep track of expiration
     * details for each Cache instance.
     *
     * When using cache expiration, calling this method is preferable to calling
     * `caches.delete()` directly, since this will ensure that the IndexedDB
     * metadata is also cleanly removed and open IndexedDB instances are deleted.
     *
     * Note that if you're *not* using cache expiration for a given cache, calling
     * `caches.delete()` and passing in the cache's name should be sufficient.
     * There is no Workbox-specific method needed for cleanup in that case.
     */
    async deleteCacheAndMetadata() {
      for (const [cacheName, cacheExpiration] of this._cacheExpirations) {
        await self.caches.delete(cacheName);
        await cacheExpiration.delete();
      }
      this._cacheExpirations = /* @__PURE__ */ new Map();
    }
  };

  // node_modules/.pnpm/workbox-core@7.3.0/node_modules/workbox-core/_private/canConstructReadableStream.js
  var supportStatus2;
  function canConstructReadableStream() {
    if (supportStatus2 === void 0) {
      try {
        new ReadableStream({ start() {
        } });
        supportStatus2 = true;
      } catch (error) {
        supportStatus2 = false;
      }
    }
    return supportStatus2;
  }

  // node_modules/.pnpm/workbox-core@7.3.0/node_modules/workbox-core/clientsClaim.js
  function clientsClaim() {
    self.addEventListener("activate", () => self.clients.claim());
  }

  // node_modules/.pnpm/workbox-streams@7.3.0/node_modules/workbox-streams/_version.js
  try {
    self["workbox:streams:7.2.0"] && _();
  } catch (e) {
  }

  // node_modules/.pnpm/workbox-streams@7.3.0/node_modules/workbox-streams/concatenate.js
  function _getReaderFromSource(source) {
    if (source instanceof Response) {
      if (source.body) {
        return source.body.getReader();
      }
      throw new WorkboxError("opaque-streams-source", { type: source.type });
    }
    if (source instanceof ReadableStream) {
      return source.getReader();
    }
    return new Response(source).body.getReader();
  }
  function concatenate(sourcePromises) {
    if (true) {
      finalAssertExports.isArray(sourcePromises, {
        moduleName: "workbox-streams",
        funcName: "concatenate",
        paramName: "sourcePromises"
      });
    }
    const readerPromises = sourcePromises.map((sourcePromise) => {
      return Promise.resolve(sourcePromise).then((source) => {
        return _getReaderFromSource(source);
      });
    });
    const streamDeferred = new Deferred();
    let i = 0;
    const logMessages = [];
    const stream = new ReadableStream({
      pull(controller) {
        return readerPromises[i].then((reader) => {
          if (reader instanceof ReadableStreamDefaultReader) {
            return reader.read();
          } else {
            return;
          }
        }).then((result) => {
          if (result === null || result === void 0 ? void 0 : result.done) {
            if (true) {
              logMessages.push([
                "Reached the end of source:",
                sourcePromises[i]
              ]);
            }
            i++;
            if (i >= readerPromises.length) {
              if (true) {
                logger.groupCollapsed(`Concatenating ${readerPromises.length} sources.`);
                for (const message of logMessages) {
                  if (Array.isArray(message)) {
                    logger.log(...message);
                  } else {
                    logger.log(message);
                  }
                }
                logger.log("Finished reading all sources.");
                logger.groupEnd();
              }
              controller.close();
              streamDeferred.resolve();
              return;
            }
            return this.pull(controller);
          } else {
            controller.enqueue(result === null || result === void 0 ? void 0 : result.value);
          }
        }).catch((error) => {
          if (true) {
            logger.error("An error occurred:", error);
          }
          streamDeferred.reject(error);
          throw error;
        });
      },
      cancel() {
        if (true) {
          logger.warn("The ReadableStream was cancelled.");
        }
        streamDeferred.resolve();
      }
    });
    return { done: streamDeferred.promise, stream };
  }

  // node_modules/.pnpm/workbox-streams@7.3.0/node_modules/workbox-streams/utils/createHeaders.js
  function createHeaders(headersInit = {}) {
    const headers = new Headers(headersInit);
    if (!headers.has("content-type")) {
      headers.set("content-type", "text/html");
    }
    return headers;
  }

  // node_modules/.pnpm/workbox-streams@7.3.0/node_modules/workbox-streams/concatenateToResponse.js
  function concatenateToResponse(sourcePromises, headersInit) {
    const { done, stream } = concatenate(sourcePromises);
    const headers = createHeaders(headersInit);
    const response = new Response(stream, { headers });
    return { done, response };
  }

  // node_modules/.pnpm/workbox-streams@7.3.0/node_modules/workbox-streams/isSupported.js
  function isSupported2() {
    return canConstructReadableStream();
  }

  // node_modules/.pnpm/workbox-streams@7.3.0/node_modules/workbox-streams/strategy.js
  function strategy(sourceFunctions, headersInit) {
    return async ({ event, request, url, params }) => {
      const sourcePromises = sourceFunctions.map((fn) => {
        return Promise.resolve(fn({ event, request, url, params }));
      });
      if (isSupported2()) {
        const { done, response } = concatenateToResponse(sourcePromises, headersInit);
        if (event) {
          event.waitUntil(done);
        }
        return response;
      }
      if (true) {
        logger.log(`The current browser doesn't support creating response streams. Falling back to non-streaming response instead.`);
      }
      const blobPartsPromises = sourcePromises.map(async (sourcePromise) => {
        const source = await sourcePromise;
        if (source instanceof Response) {
          return source.blob();
        } else {
          return new Response(source).blob();
        }
      });
      const blobParts = await Promise.all(blobPartsPromises);
      const headers = createHeaders(headersInit);
      return new Response(new Blob(blobParts), { headers });
    };
  }

  // src/sw/sw-strategies.ts
  var contentCache = "contents-cache";
  var partialsCache = "partials-cache";
  var cacheFirstStrategy = new CacheFirst({
    cacheName: partialsCache,
    plugins: [
      new ExpirationPlugin({
        maxEntries: 30,
        maxAgeSeconds: 60 * 60 * 24 * 30
        // about a month
      })
    ]
  });
  var networkFirstStrategy = new NetworkFirst({
    cacheName: contentCache,
    networkTimeoutSeconds: 5
  });
  var staleWhileRevalidateStrategy = new StaleWhileRevalidate({
    cacheName: contentCache
  });
  var dashboardPageStream = strategy(
    [
      async ({ event }) => cacheFirstStrategy.handle({
        request: new Request("/account/dashboard/openingHtmlTag-Partial.html"),
        event
      }),
      async ({ event }) => cacheFirstStrategy.handle({
        request: new Request("/account/dashboard/commonHeader-Partial.html"),
        event
      }),
      async ({ event }) => cacheFirstStrategy.handle({
        request: new Request("/account/dashboard/mainOpeningTag-Partial.html"),
        event
      }),
      async ({ event }) => staleWhileRevalidateStrategy.handle({
        request: new Request("/account/dashboard/swiperOne"),
        event
      }),
      async ({ event }) => staleWhileRevalidateStrategy.handle({
        request: new Request("/account/dashboard/echarts?month=0"),
        event
      }),
      async ({ event }) => staleWhileRevalidateStrategy.handle({
        request: new Request("/account/dashboard/swiperTwo"),
        event
      }),
      async ({ event }) => cacheFirstStrategy.handle({
        request: new Request("/account/dashboard/footer-Partial.html"),
        event
      }),
      async ({ event }) => cacheFirstStrategy.handle({
        request: new Request("/account/dashboard/closingHtmlTag-Partial.html"),
        event
      })
    ],
    {}
  );
  var locationPageStream = strategy(
    [
      async ({ event }) => cacheFirstStrategy.handle({
        request: new Request("/account/location/openingHtmlTag-Partial.html"),
        event
      }),
      async ({ event }) => cacheFirstStrategy.handle({
        request: new Request("/account/location/header-Partial.html"),
        event
      }),
      async ({ event }) => networkFirstStrategy.handle({
        request: new Request("/account/location/main", {
          headers: [["X-Content-Mode", "partials"]]
        }),
        event
      }),
      async ({ event }) => cacheFirstStrategy.handle({
        request: new Request("/account/location/footer-Partial.html"),
        event
      }),
      async ({ event }) => cacheFirstStrategy.handle({
        request: new Request("/account/location/closingHtmlTag-Partial.html"),
        event
      })
    ],
    {}
  );
  var walletPageStream = strategy(
    [
      async ({ event }) => cacheFirstStrategy.handle({
        request: new Request("/account/wallet/openingHtmlTag-Partial.html"),
        event
      }),
      async ({ event }) => cacheFirstStrategy.handle({
        request: new Request("/account/wallet/commonHeader-Partial.html"),
        event
      }),
      async ({ event }) => cacheFirstStrategy.handle({
        request: new Request("/account/wallet/mainOpeningTag-Partial.html"),
        event
      }),
      async ({ event }) => networkFirstStrategy.handle({
        request: new Request("/account/wallet/balance"),
        event
      }),
      async ({ event }) => cacheFirstStrategy.handle({
        request: new Request("/account/wallet/options"),
        event
      }),
      async ({ event }) => networkFirstStrategy.handle({
        request: new Request("/account/wallet/coinData", {}),
        event
      }),
      async ({ event }) => staleWhileRevalidateStrategy.handle({
        request: new Request("/account/wallet/swiper"),
        event
      }),
      async ({ event }) => cacheFirstStrategy.handle({
        request: new Request("/account/wallet/footer-Partial.html"),
        event
      }),
      async ({ event }) => cacheFirstStrategy.handle({
        request: new Request("/account/wallet/closingHtmlTag-Partial.html"),
        event
      })
    ],
    {}
  );
  var spendPageStream = strategy(
    [
      async ({ event }) => cacheFirstStrategy.handle({
        request: new Request("/account/spend/openingHtmlTag-Partial.html"),
        event
      }),
      async ({ event }) => cacheFirstStrategy.handle({
        request: new Request("/account/spend/header-Partial.html"),
        event
      }),
      async ({ event }) => cacheFirstStrategy.handle({
        request: new Request("/account/spend/mainOpeningTag-Partial.html"),
        event
      }),
      async ({ event }) => networkFirstStrategy.handle({
        request: new Request("/account/spend/balance"),
        event
      }),
      async ({ event }) => cacheFirstStrategy.handle({
        request: new Request("/account/spend/otherhalf"),
        event
      }),
      async ({ event }) => cacheFirstStrategy.handle({
        request: new Request("/account/spend/footer-Partial.html"),
        event
      }),
      async ({ event }) => cacheFirstStrategy.handle({
        request: new Request("/account/spend/closingHtmlTag-Partial.html"),
        event
      })
    ],
    {}
  );
  var profilePageStream = strategy(
    [
      async ({ event }) => cacheFirstStrategy.handle({
        request: new Request("/account/profile/openingHtmlTag-Partial.html"),
        event
      }),
      async ({ event }) => cacheFirstStrategy.handle({
        request: new Request("/account/profile/commonHeader-Partial.html"),
        event
      }),
      async ({ event }) => cacheFirstStrategy.handle({
        request: new Request("/account/profile/mainOpeningTag-Partial.html"),
        event
      }),
      async ({ event }) => staleWhileRevalidateStrategy.handle({
        request: new Request("/account/profile/displayIcon"),
        event
      }),
      async ({ event }) => staleWhileRevalidateStrategy.handle({
        request: new Request("/account/profile/formPartial"),
        event
      }),
      async ({ event }) => cacheFirstStrategy.handle({
        request: new Request("/account/profile/footer-Partial.html"),
        event
      }),
      async ({ event }) => cacheFirstStrategy.handle({
        request: new Request("/account/profile/closingHtmlTag-Partial.html"),
        event
      })
    ],
    {}
  );
  var airtimePageStream = strategy(
    [
      async ({ event }) => cacheFirstStrategy.handle({
        request: new Request("/account/airtime/openingHtmlTag-Partial.html"),
        event
      }),
      async ({ event }) => cacheFirstStrategy.handle({
        request: new Request("/account/airtime/header-Partial.html"),
        event
      }),
      async ({ event }) => cacheFirstStrategy.handle({
        request: new Request("/account/airtime/mainOpeningTag-Partial.html"),
        event
      }),
      async ({ event }) => cacheFirstStrategy.handle({
        request: new Request("/account/airtime/inputNumber"),
        event
      }),
      async ({ event }) => networkFirstStrategy.handle({
        request: new Request("/account/airtime/beneficiaries"),
        event
      }),
      async ({ event }) => cacheFirstStrategy.handle({
        request: new Request("/account/airtime/footer-Partial.html"),
        event
      }),
      async ({ event }) => cacheFirstStrategy.handle({
        request: new Request("/account/airtime/closingHtmlTag-Partial.html"),
        event
      })
    ],
    {}
  );
  var beneficiariesPageStream = strategy(
    [
      async ({ event }) => cacheFirstStrategy.handle({
        request: new Request(
          "/account/beneficiaries/openingHtmlTag-Partial.html"
        ),
        event
      }),
      async ({ event, url }) => cacheFirstStrategy.handle({
        request: new Request("/account/beneficiaries/header-Partial.html"),
        event
      }),
      async ({ event }) => networkFirstStrategy.handle({
        request: new Request("/account/beneficiaries/main"),
        event
      }),
      async ({ event }) => cacheFirstStrategy.handle({
        request: new Request("/account/beneficiaries/footer-Partial.html"),
        event
      }),
      async ({ event }) => cacheFirstStrategy.handle({
        request: new Request(
          "/account/beneficiaries/closingHtmlTag-Partial.html"
        ),
        event
      })
    ],
    {}
  );
  var wastePageStream = strategy(
    [
      async ({ event }) => cacheFirstStrategy.handle({
        request: new Request("/account/waste/openingHtmlTag-Partial.html"),
        event
      }),
      async ({ event }) => cacheFirstStrategy.handle({
        request: new Request("/account/waste/commonHeader-Partial.html"),
        event
      }),
      async ({ event }) => cacheFirstStrategy.handle({
        request: new Request("/account/waste/mainOpeningTag-Partial.html"),
        event
      }),
      async ({ event, url }) => networkFirstStrategy.handle({
        request: matchingPathVariableForWaste(url, "quickInfo"),
        event
      }),
      async ({ event }) => cacheFirstStrategy.handle({
        request: new Request("/account/waste/filter"),
        event
      }),
      async ({ event, url }) => networkFirstStrategy.handle({
        request: matchingPathVariableForWaste(url, "history"),
        event
      }),
      async ({ event }) => cacheFirstStrategy.handle({
        request: new Request("/account/waste/footer-Partial.html"),
        event
      }),
      async ({ event }) => cacheFirstStrategy.handle({
        request: new Request("/account/waste/closingHtmlTag-Partial.html"),
        event
      })
    ],
    {}
  );
  function matchingPathVariableForWaste(url, endpoint) {
    const segments = url.pathname.split("/");
    const lastSegment = segments[segments.length - 1];
    return new Request(`/account/waste/${endpoint}/${lastSegment}`);
  }

  // src/sw/sw.ts
  console.log("Service Worker running in", "development", "mode");
  clientsClaim();
  enable();
  registerRoute(
    ({ url }) => url.pathname === "/account/signin-Partial.html" || url.pathname === "/account/signup-Partial.html",
    new CacheFirst({
      cacheName: partialsCache,
      plugins: [
        new ExpirationPlugin({
          maxEntries: 30,
          maxAgeSeconds: 60 * 60 * 24 * 30
          // about a month
        })
      ]
    })
  );
  registerRoute(
    new RegExp("^.*/account/dashboard/chart\\?month=[1-4]$"),
    new StaleWhileRevalidate({
      cacheName: contentCache,
      plugins: [
        new ExpirationPlugin({
          maxEntries: 30,
          maxAgeSeconds: 60 * 60 * 24 * 30
          // about a month
        })
      ]
    })
  );
  registerRoute(
    new RegExp("^.*/account/dashboard/chart\\?month=0$"),
    new NetworkFirst({
      networkTimeoutSeconds: 3,
      cacheName: contentCache,
      plugins: [
        new ExpirationPlugin({
          maxEntries: 30,
          maxAgeSeconds: 60 * 60 * 24 * 30
          // about a month
        })
      ]
    })
  );
  registerRoute(
    ({ url }) => url.pathname === "/account/dashboard",
    dashboardPageStream
  );
  registerRoute(
    ({ url }) => url.pathname === "/account/scan",
    new CacheFirst({
      cacheName: contentCache,
      plugins: [
        new ExpirationPlugin({
          maxEntries: 30,
          maxAgeSeconds: 60 * 60 * 24 * 30
          // about a month
        })
      ]
    })
  );
  registerRoute(
    ({ url }) => url.pathname === "/account/location",
    locationPageStream
  );
  registerRoute(
    ({ url }) => url.pathname === "/account/wallet",
    walletPageStream
  );
  registerRoute(({ url }) => url.pathname === "/account/spend", spendPageStream);
  registerRoute(
    ({ url }) => url.pathname === "/account/profile/changeDisplayName",
    async ({ request }) => {
      const response = await fetch(request);
      if (!response.ok) {
        return response;
      }
      const cache = await caches.open(partialsCache);
      const allCacheKeys = await cache.keys();
      const matchingResponses = allCacheKeys.filter(
        (request2) => request2.url.endsWith("commonHeader-Partial.html")
      );
      for (const req of matchingResponses) {
        await cache.put(req, response.clone());
      }
      return response;
    },
    "POST"
  );
  registerRoute(
    ({ url }) => url.pathname === "/account/profile",
    profilePageStream
  );
  registerRoute(
    ({ url }) => url.pathname === "/account/airtime",
    airtimePageStream,
    "POST"
  );
  registerRoute(
    ({ url }) => url.pathname === "/account/beneficiaries",
    beneficiariesPageStream
  );
  registerRoute(({ url }) => {
    const wastePathPattern = /.*\/account\/waste\/(paper|glass|plastics|metal)$/;
    return wastePathPattern.test(url.href);
  }, wastePageStream);
  registerRoute(
    ({ url }) => {
      const tipsPathPattern = /.*\/account\/tips\/(paper|glass|plastics|metal)$/;
      return tipsPathPattern.test(url.href);
    },
    new CacheFirst({
      cacheName: contentCache,
      plugins: [
        new ExpirationPlugin({
          maxEntries: 30,
          maxAgeSeconds: 60 * 60 * 24 * 30
          // about a month
        })
      ]
    })
  );
  precacheAndRoute([{"revision":"9d2eb3bea7b25860ef40be12d98a6cdc","url":"assets/app_shortcut/bin/100.png"},{"revision":"8db02fcf4e9706bf2a3840aaa027b6db","url":"assets/app_shortcut/bin/114.png"},{"revision":"293bea63d7e7c2cee8023b83a1331944","url":"assets/app_shortcut/bin/120.png"},{"revision":"183c7b5833a2912154af9368d91eb2dd","url":"assets/app_shortcut/bin/128.png"},{"revision":"4c0edd073a1ce08671c395055d992623","url":"assets/app_shortcut/bin/144.png"},{"revision":"cd21adfd5447fde96ceeda136dff56a1","url":"assets/app_shortcut/bin/152.png"},{"revision":"f41c155a908810ba6277c533b1900567","url":"assets/app_shortcut/bin/167.png"},{"revision":"463a7d07ef4cbe6b1c9f7438fed4f692","url":"assets/app_shortcut/bin/180.png"},{"revision":"427c5d9f0a29d8c1dc09f7ddacc0cf3b","url":"assets/app_shortcut/bin/192.png"},{"revision":"f44465f3c83125525278f09b5e21514f","url":"assets/app_shortcut/bin/32.png"},{"revision":"693d2292119a664a662c82f122752e8e","url":"assets/app_shortcut/bin/40.png"},{"revision":"3e1533db9c1037fae6f3f1d9661fae52","url":"assets/app_shortcut/bin/50.png"},{"revision":"9a3f82a6c274c812f22dfa0940e0c439","url":"assets/app_shortcut/bin/57.png"},{"revision":"70d2a2ee8d339d164c1ef54cc10bc747","url":"assets/app_shortcut/bin/58.png"},{"revision":"fec43542c62ff2f47a68f70176b5ea60","url":"assets/app_shortcut/bin/60.png"},{"revision":"5983158087bddeca7e587d4f83bb1ec1","url":"assets/app_shortcut/bin/64.png"},{"revision":"8103a0b63cea1c7e86d79086e8249d7a","url":"assets/app_shortcut/bin/72.png"},{"revision":"635573025353d632cb77c10f0dbf4141","url":"assets/app_shortcut/bin/76.png"},{"revision":"90864c33123622d822b670b3eb4e1974","url":"assets/app_shortcut/bin/80.png"},{"revision":"dd5a12579248b9b8f9a2a5d972923ae2","url":"assets/app_shortcut/bin/87.png"},{"revision":"1f9952c588a74e1ee5ae58c42d3c0d55","url":"assets/app_shortcut/collector/100.png"},{"revision":"bd6f5034a2ba7d7f1e186087eb58704b","url":"assets/app_shortcut/collector/114.png"},{"revision":"9678ce36babc863de591f17e9624009a","url":"assets/app_shortcut/collector/120.png"},{"revision":"dbf424a9d401c12ad794929f477c4b98","url":"assets/app_shortcut/collector/128.png"},{"revision":"f0be4f87c2cdc709d6bd37e4c248cf0c","url":"assets/app_shortcut/collector/144.png"},{"revision":"52123fa6a3dd9fa41c4a7ba731fbdac2","url":"assets/app_shortcut/collector/152.png"},{"revision":"0be657b5e164596b8fe328a6a4b2649c","url":"assets/app_shortcut/collector/167.png"},{"revision":"7cad924a3e38be895fb6edae06b7d8ca","url":"assets/app_shortcut/collector/180.png"},{"revision":"a6b94d8d98e67c1ccfad2b13af515d10","url":"assets/app_shortcut/collector/192.png"},{"revision":"3d11b1261159d4f686bd02d2f014c765","url":"assets/app_shortcut/collector/32.png"},{"revision":"8969e2cc112b6a04900ca6fbb5b3ab5a","url":"assets/app_shortcut/collector/40.png"},{"revision":"d2dbf8dddadb3c8f3a8f98bb32f01a31","url":"assets/app_shortcut/collector/50.png"},{"revision":"fd372ee29d3c74e5f702b3f41ee227d3","url":"assets/app_shortcut/collector/57.png"},{"revision":"736e6403016634e84b4807f21eaa2c08","url":"assets/app_shortcut/collector/58.png"},{"revision":"16cc2b3fec91bb3ae2b1360839800cf5","url":"assets/app_shortcut/collector/60.png"},{"revision":"690996cad354f8ebaee34e65ea35640a","url":"assets/app_shortcut/collector/64.png"},{"revision":"c545b68fe58687b941d52149a9eba045","url":"assets/app_shortcut/collector/72.png"},{"revision":"91ba66113149fe1b96dd68fac7c6e747","url":"assets/app_shortcut/collector/76.png"},{"revision":"8928ca4f0ef9a4cf5ae99535b54810c8","url":"assets/app_shortcut/collector/80.png"},{"revision":"88dbafc5d0a33b8cf77cbf4e1a41b15f","url":"assets/app_shortcut/collector/87.png"},{"revision":"198315bf7752ee0313ce0a626e1b8fd8","url":"assets/app_shortcut/scan/100.png"},{"revision":"47f6d165f5937a9dbed88dd6e2d26293","url":"assets/app_shortcut/scan/114.png"},{"revision":"c85b5ccf6d9e482a8a5d9467810c9a3c","url":"assets/app_shortcut/scan/120.png"},{"revision":"2702d476d13f143bb2440f3b399aef6c","url":"assets/app_shortcut/scan/128.png"},{"revision":"8e782c36848aa92393ffc1946860b85d","url":"assets/app_shortcut/scan/144.png"},{"revision":"52014bdcc434d7b50eb80c4f827129f3","url":"assets/app_shortcut/scan/152.png"},{"revision":"525189339a10791c8fe04707729ca236","url":"assets/app_shortcut/scan/167.png"},{"revision":"9f1c2d4ce7a9354441b5eae07ea47287","url":"assets/app_shortcut/scan/180.png"},{"revision":"601cfe6496b1fc80196937be0e2f4958","url":"assets/app_shortcut/scan/192.png"},{"revision":"4a171bed338b8486daee7eee473eb8ad","url":"assets/app_shortcut/scan/32.png"},{"revision":"7b71dcf45e38f6d26d8bb486a740e538","url":"assets/app_shortcut/scan/40.png"},{"revision":"ff5c6b21db5b69af0c0ccab2c7313fde","url":"assets/app_shortcut/scan/50.png"},{"revision":"8d550d10d90bfbafb39bd3dde03e9ea8","url":"assets/app_shortcut/scan/57.png"},{"revision":"09e16b9e79f201633bdf8b2fe927367f","url":"assets/app_shortcut/scan/58.png"},{"revision":"f1f337885e7f80f747d385bf2cc627f7","url":"assets/app_shortcut/scan/60.png"},{"revision":"5d732a029ad26c50a968cce7191d3d9e","url":"assets/app_shortcut/scan/64.png"},{"revision":"9e71f2bfb63002a24360a90c4a000b59","url":"assets/app_shortcut/scan/72.png"},{"revision":"d2c62376296f0b5606435782e418a966","url":"assets/app_shortcut/scan/76.png"},{"revision":"23cb3d3304dabf14dba14a269811143d","url":"assets/app_shortcut/scan/80.png"},{"revision":"772acf40bb3dc7152eed8411ee0aaa4a","url":"assets/app_shortcut/scan/87.png"},{"revision":"5e077c15f6e1d334dd4e9be62b28ac75","url":"assets/fonts/montserrat/static/Montserrat-Regular.ttf"},{"revision":"cc10461cb5e0a6f2621c7179f4d6de17","url":"assets/fonts/montserrat/static/Montserrat-SemiBold.ttf"},{"revision":"093ee89be9ede30383f39a899c485a82","url":"assets/fonts/poppins/Poppins-Regular.ttf"},{"revision":"3cc21e10177d0fa7cb2a3d7f73e14942","url":"assets/icons/android/android-launchericon-144-144.png"},{"revision":"132b96cc1f95d173d97d2d84e7f8adef","url":"assets/icons/android/android-launchericon-192-192.png"},{"revision":"5747207c94a1be96e7a75dd770b697be","url":"assets/icons/android/android-launchericon-48-48.png"},{"revision":"e02074974e62abf5cd36d4d434f1635f","url":"assets/icons/android/android-launchericon-512-512.png"},{"revision":"2ed6b53f5b6a045ac52b3653deeb66a6","url":"assets/icons/android/android-launchericon-72-72.png"},{"revision":"7bb2a233b5bd7699777631b173326cf6","url":"assets/icons/android/android-launchericon-96-96.png"},{"revision":"7786b8e652c290cd1bcbeac8ec639df5","url":"assets/icons/ios/100.png"},{"revision":"4a22bd5bc8d2d006ca42fa51671970d4","url":"assets/icons/ios/1024.png"},{"revision":"78a9b13b4ec1897756f500605f22bf70","url":"assets/icons/ios/114.png"},{"revision":"05dacfb129e61cc8dfedd75a875d8b3d","url":"assets/icons/ios/120.png"},{"revision":"8d2a39523bbfa547409d6700156dfc50","url":"assets/icons/ios/128.png"},{"revision":"3cc21e10177d0fa7cb2a3d7f73e14942","url":"assets/icons/ios/144.png"},{"revision":"13d6ba668019351ba66f7df70f3ea598","url":"assets/icons/ios/152.png"},{"revision":"9708d4942095f2202580f275fee8f9a4","url":"assets/icons/ios/16.png"},{"revision":"9458a5c4e9e2cc0d8a20c4dfb17f7b6d","url":"assets/icons/ios/167.png"},{"revision":"21d6115d586383088eb7fa0a5b53616f","url":"assets/icons/ios/180.png"},{"revision":"132b96cc1f95d173d97d2d84e7f8adef","url":"assets/icons/ios/192.png"},{"revision":"51fdf7cc9aae314cef16b7826226d65d","url":"assets/icons/ios/20.png"},{"revision":"f6d4f909a4bcadd6445e03926b729df3","url":"assets/icons/ios/256.png"},{"revision":"e93631f97dc9727ff3a2b62ad256fbed","url":"assets/icons/ios/29.png"},{"revision":"283a011d5cd81c47ce4731dd4a09913e","url":"assets/icons/ios/32.png"},{"revision":"9f440a6f6d4f9d38bed7801572772235","url":"assets/icons/ios/40.png"},{"revision":"f621936cea96675f4cef8512b1bea6ab","url":"assets/icons/ios/50.png"},{"revision":"e02074974e62abf5cd36d4d434f1635f","url":"assets/icons/ios/512.png"},{"revision":"6e8c0d35fabdc1e7d7210bc7c045579f","url":"assets/icons/ios/57.png"},{"revision":"890a780b6a11bd34dcc02baa00f1b289","url":"assets/icons/ios/58.png"},{"revision":"0c79fbc353d73ed19564e32a0d043e85","url":"assets/icons/ios/60.png"},{"revision":"aa7ac2b0aab5ec1443e0009f564b7db3","url":"assets/icons/ios/64.png"},{"revision":"2ed6b53f5b6a045ac52b3653deeb66a6","url":"assets/icons/ios/72.png"},{"revision":"9cee7f6d7540bb2574d7a1d47589fde8","url":"assets/icons/ios/76.png"},{"revision":"83fd4935fd6695eedcda54f5466ed617","url":"assets/icons/ios/80.png"},{"revision":"15ae8bb4f07a30646ddefcea106183b3","url":"assets/icons/ios/87.png"},{"revision":"4349a65089ea734359da8d82e78124d1","url":"assets/images/airtel.png"},{"revision":"77b2e12d3673d49a71e9d992823e70c1","url":"assets/images/airtel.webp"},{"revision":"98e2a6eb4e6cb630795feb47d206c724","url":"assets/images/arrow-vert.png"},{"revision":"595fd420311cb113d081b32d56fe9376","url":"assets/images/arrow-vert.webp"},{"revision":"6842a2b5dfc4b44e3cd3562037205c79","url":"assets/images/arrow.png"},{"revision":"318ac3310fda6eeba87af0e836429491","url":"assets/images/arrow.webp"},{"revision":"68daaf12037e3dc66404a42ec74baa10","url":"assets/images/background.jpg"},{"revision":"f63b5f10a623ff16749ca7b5ae66359d","url":"assets/images/background.webp"},{"revision":"6301c49ced161e3431577cd05b75f826","url":"assets/images/beltech_background.png"},{"revision":"e6cc44748b694e5fbabd9ac357b8e446","url":"assets/images/beltech_background.webp"},{"revision":"373a702aa7121ec51cab20af3ce93ccf","url":"assets/images/beltech_logo_.png"},{"revision":"273cdc2da99bffa1e2750efc31afbd4f","url":"assets/images/beltech_logo_.webp"},{"revision":"1b7f3a520e597f36e5a1cae28ca3c4ae","url":"assets/images/beltech_logo_104x107.jpg"},{"revision":"592dbd9f931bfc2c5fa9dbb65efb7d1e","url":"assets/images/beltech_logo_104x107.webp"},{"revision":"7167710b4396424e143e0796b9f00397","url":"assets/images/beltech_logo_16x16.png"},{"revision":"57b32cdb11504567dc795b5c48126cd8","url":"assets/images/beltech_logo_16x16.webp"},{"revision":"1c5a7f56c46bab3265dac472b509de08","url":"assets/images/beltech_logo_2.jpg"},{"revision":"7a9f5f9ffb969002b4e0a25c7eb527ce","url":"assets/images/beltech_logo_2.webp"},{"revision":"8a860281292d3a0dffeb4d16186e078a","url":"assets/images/beltech_logo_32x32.png"},{"revision":"658cc7443076d46cbd7498a696ffc339","url":"assets/images/beltech_logo_32x32.webp"},{"revision":"c2f6fdae826f550ade4d493f021291a3","url":"assets/images/beltech_logo_35x36.jpg"},{"revision":"2ebe4be626f1fc5f7c7ee0dc87d954fb","url":"assets/images/beltech_logo_35x36.webp"},{"revision":"4b9037cbf05f11e9770dc03cf48af9a3","url":"assets/images/beltech_logo_square.png"},{"revision":"a7d80b8021d741935d9a12d0c669266f","url":"assets/images/beltech_logo_square.webp"},{"revision":"26b98af9daef39bed019ebeedf0ab91e","url":"assets/images/beltech_logo.png"},{"revision":"bf933e423dcf2c1b7f78628ed7fd65dd","url":"assets/images/beltech_logo.webp"},{"revision":"0ada98e903c1a8892fc980f91ac60857","url":"assets/images/beltech_wide.png"},{"revision":"31db00948f881c4c60e776e19d88e31b","url":"assets/images/beltech_wide.webp"},{"revision":"f91aee73250ab84e45fb4f6b205a80a2","url":"assets/images/beltech.png"},{"revision":"bd68b90b6686b06c21031c04428c1f56","url":"assets/images/beltech.webp"},{"revision":"64b5c2ccd74050583e5a96c44ca5e3d7","url":"assets/images/bitcoin.png"},{"revision":"38bf14f38eb4dcbf73e2a71946302de8","url":"assets/images/bitcoin.webp"},{"revision":"093249306e5d95c9a091419a6d77015d","url":"assets/images/bonus.jpg"},{"revision":"63428698a20e1bb6cc0ea52bb55b838c","url":"assets/images/bonus.webp"},{"revision":"94429bba13f67578a02e98ab4a9a637d","url":"assets/images/book.jpg"},{"revision":"b1be73eba1a8b1d6dbd82532332da8fe","url":"assets/images/book.webp"},{"revision":"5b6e8a628df167d162d2a426612c0e33","url":"assets/images/call-to-action-img.png"},{"revision":"470a0c075bd12099e2a4550dddd1124c","url":"assets/images/call-to-action-img.webp"},{"revision":"4bbdebbbb7fa1e907875c443853d1401","url":"assets/images/downTrend.png"},{"revision":"bd2ffcd6a4ca1670dd203bb1537c1c29","url":"assets/images/downTrend.webp"},{"revision":"d46afcbf4c1911b45dad8aa4106f1d8b","url":"assets/images/ethereum.png"},{"revision":"f176b9411fcf5498fcf96d0efed44662","url":"assets/images/ethereum.webp"},{"revision":"27def2649a087fe99cef55d48a5d8cbc","url":"assets/images/financial-growth.png"},{"revision":"1ec0fc36d535f50c75920c27ce28e29f","url":"assets/images/financial-growth.webp"},{"revision":"57e93d8a2aa0125ed353c88e0a34f4fc","url":"assets/images/footer-logo.png"},{"revision":"ae70f383bf32122669487917140d7908","url":"assets/images/footer-logo.webp"},{"revision":"241e6a8814c6dde87bf5e5e65522d8d5","url":"assets/images/friends.jpg"},{"revision":"63245aa80f671f87ff81d2de81044c37","url":"assets/images/friends.webp"},{"revision":"cadc5c03571a19a183539d8b5cb1477f","url":"assets/images/glass.jpg"},{"revision":"d532ed16b3530190b176c27d5e0fee89","url":"assets/images/glass.webp"},{"revision":"668b9df9f517b864140289a3b113fa0e","url":"assets/images/glo.png"},{"revision":"d84af5b1607de13addc253ad5f69869c","url":"assets/images/glo.webp"},{"revision":"9c82316281dbfff4ca13b0a645020d57","url":"assets/images/grid.png"},{"revision":"536c75c83afc2e7a706ccc1606484dcc","url":"assets/images/grid.webp"},{"revision":"1c5a7f56c46bab3265dac472b509de08","url":"assets/images/maskAny.jpeg"},{"revision":"7a9f5f9ffb969002b4e0a25c7eb527ce","url":"assets/images/maskAny.webp"},{"revision":"63e6bd0b0e749946d1fd71d20dab322a","url":"assets/images/metal.jpeg"},{"revision":"5d4e25805b6f33427e650dac4bf589b9","url":"assets/images/metal.webp"},{"revision":"96fb21df65238da74b971ccecc61350c","url":"assets/images/metal1.jpeg"},{"revision":"9469c44b9255ec327ca83a48e6453ab5","url":"assets/images/metal1.webp"},{"revision":"3b467a598c4379a858e28e4173ddc76e","url":"assets/images/metal2.jpeg"},{"revision":"d5337b9a3ca8747af79f0461e6a079c1","url":"assets/images/metal2.jpg"},{"revision":"6a63b98d1e46155c4e31e0f325936fe1","url":"assets/images/metal2.webp"},{"revision":"b44f889a3d65391518eecfc822d87bf0","url":"assets/images/mtn.png"},{"revision":"8877ad5c89d43a0bff4e76aa62aeb222","url":"assets/images/mtn.webp"},{"revision":"2a323f3fb0935c775506f7b99b3a62ef","url":"assets/images/name.jpg"},{"revision":"ade3c3fc3f04c5902b23b7c276a52ee5","url":"assets/images/name.webp"},{"revision":"3fa7afe62e9f2cf3ec3e6a2f20906edf","url":"assets/images/pancakeswap.png"},{"revision":"3c19f272692c9a8aa21832f5382d4a83","url":"assets/images/pancakeswap.webp"},{"revision":"ac35aabb05fb5f1ba6b1d727897d8448","url":"assets/images/paper_bag.jpg"},{"revision":"8cb816de8ad1dc4070ca5c8a5f02bd16","url":"assets/images/paper_bag.webp"},{"revision":"3f3ba3ac0feb1e0eefaeb938c1ae4a64","url":"assets/images/plastic_2.jpg"},{"revision":"63e14dc597c5e10f3cad281e34e2d761","url":"assets/images/plastic_2.webp"},{"revision":"90995cd2f97d2c417ef1f86525d25bda","url":"assets/images/qr_code.jpg"},{"revision":"102ecb001faf4657dd84c9d6275af1a4","url":"assets/images/qr_code.webp"},{"revision":"cfef005b69b6d115f97d137a889aee37","url":"assets/images/qr-scan.png"},{"revision":"fbfee90d756a458c1a138d102b0c6c7e","url":"assets/images/qr-scan.webp"},{"revision":"489243e171f1d6f388de05398cabb7ca","url":"assets/images/Recycle_small.png"},{"revision":"aeb77d7de9394eee2602a3ff7a5c7e34","url":"assets/images/Recycle_small.webp"},{"revision":"df0868faaa680bc02b605328b4550e19","url":"assets/images/scan.jpeg"},{"revision":"e2bd2b8a3788cb2ff3fae4e79a04b2a3","url":"assets/images/scan.webp"},{"revision":"fb9fdc74eb0d938442abd74276cf6b66","url":"assets/images/scan2.png"},{"revision":"4e9914eb6a4dc8ef4b63462af4ba8024","url":"assets/images/scan2.webp"},{"revision":"42e37300a6b12037fbeba4b04c2e7586","url":"assets/images/sign-into-user-account.png"},{"revision":"981f59b2db3b176854e53c337cc95e59","url":"assets/images/sign-into-user-account.webp"},{"revision":"3da5f7bbc76d299241505a6aa6a9ae52","url":"assets/images/smiley.png"},{"revision":"8b269b789316b5a0bf2513d919410cdf","url":"assets/images/smiley.webp"},{"revision":"e3aafb66de4f0539c9db524fbc0173c2","url":"assets/images/solana.png"},{"revision":"75ece93f3e029ba8236de22668b199aa","url":"assets/images/solana.webp"},{"revision":"c97794c3ddf7528f7a7591ddf22df0eb","url":"assets/images/sports_bottle.jpg"},{"revision":"3429f74c4d4fef3258478194ae109287","url":"assets/images/sports_bottle.webp"},{"revision":"2611b94379165a2c762c0499b6e7303d","url":"assets/images/upTrend.png"},{"revision":"902108a049f4dffefe1aa52d7ba0d82e","url":"assets/images/upTrend.webp"},{"revision":"25bcb059251caa734d596e764838b1aa","url":"assets/images/usdc.png"},{"revision":"161f4c1cdf4a3b22474bc4d72d135140","url":"assets/images/usdc.webp"},{"revision":"32ab02dd09c37e6f137c6d597b0b1502","url":"assets/images/world.png"},{"revision":"3b2a7ad7da2b7916c75c8efaa0c9d7dc","url":"assets/images/world.webp"},{"revision":"f5c2c3f0c21cdf3bddde52ca635629b9","url":"assets/screens/10.2__iPad_landscape.png"},{"revision":"e82c519408df28321bfdf34fe47bc916","url":"assets/screens/10.2__iPad_portrait.png"},{"revision":"a7a4ece008cedac07e46a8a1b45d5ccb","url":"assets/screens/10.5__iPad_Air_landscape.png"},{"revision":"c6e1fa42d974c002c4d90079020cb17c","url":"assets/screens/10.5__iPad_Air_portrait.png"},{"revision":"da3a937b9cfe59413ad437cb90478588","url":"assets/screens/10.9__iPad_Air_landscape.png"},{"revision":"2271cec1dd1cf1fbcccaaea3ca54208e","url":"assets/screens/10.9__iPad_Air_portrait.png"},{"revision":"da0aad3ffd743dea1aa7056905b9e9e8","url":"assets/screens/11__iPad_Pro__10.5__iPad_Pro_landscape.png"},{"revision":"6c4c54c452a32d786577ed20ed19c9fa","url":"assets/screens/11__iPad_Pro__10.5__iPad_Pro_portrait.png"},{"revision":"33572db520a06256d3bcfe3ec2e351e6","url":"assets/screens/11__iPad_Pro_M4_landscape.png"},{"revision":"33bd6edbecc898b1ea58444063174192","url":"assets/screens/11__iPad_Pro_M4_portrait.png"},{"revision":"f60c5788e58f524b4ecec0b598f1cedd","url":"assets/screens/12.9__iPad_Pro_landscape.png"},{"revision":"707b5a3ba085c60ac1e315d059cfd9d4","url":"assets/screens/12.9__iPad_Pro_portrait.png"},{"revision":"26eefcc89959a12785c6f2fb7c776f9e","url":"assets/screens/13__iPad_Pro_M4_landscape.png"},{"revision":"418edcf6c8c3efcfef71186fe7274bd1","url":"assets/screens/13__iPad_Pro_M4_portrait.png"},{"revision":"8af445ffba5179709e0e04db59a02502","url":"assets/screens/4__iPhone_SE__iPod_touch_5th_generation_and_later_landscape.png"},{"revision":"f9495d085c9c608f71459ad20711eee0","url":"assets/screens/4__iPhone_SE__iPod_touch_5th_generation_and_later_portrait.png"},{"revision":"6ff2487cf8adfcc7a25ee6a73a7b3366","url":"assets/screens/8.3__iPad_Mini_landscape.png"},{"revision":"e829fd63915a8da2117ca5a7c3720c32","url":"assets/screens/8.3__iPad_Mini_portrait.png"},{"revision":"081f0a832a27a3b3c68f3549623a09e6","url":"assets/screens/9.7__iPad_Pro__7.9__iPad_mini__9.7__iPad_Air__9.7__iPad_landscape.png"},{"revision":"6a4a4b5670c40a526711230187beb69f","url":"assets/screens/9.7__iPad_Pro__7.9__iPad_mini__9.7__iPad_Air__9.7__iPad_portrait.png"},{"revision":"de0c104549071fbf59d9d185c0744bee","url":"assets/screens/icon.png"},{"revision":"31fd25ac6251dcc6a1e90b04ff429292","url":"assets/screens/iPhone_11__iPhone_XR_landscape.png"},{"revision":"16053045f8c782f8bf37abac59abb544","url":"assets/screens/iPhone_11__iPhone_XR_portrait.png"},{"revision":"dd003237fc0d056bba9313b0588c9bbf","url":"assets/screens/iPhone_11_Pro_Max__iPhone_XS_Max_landscape.png"},{"revision":"eccf261e736c0e8e149c66bc951365e8","url":"assets/screens/iPhone_11_Pro_Max__iPhone_XS_Max_portrait.png"},{"revision":"082aa6299562bbdd20e3d6112dc25415","url":"assets/screens/iPhone_13_mini__iPhone_12_mini__iPhone_11_Pro__iPhone_XS__iPhone_X_landscape.png"},{"revision":"6af710b8a556236725f22e83620b4e56","url":"assets/screens/iPhone_13_mini__iPhone_12_mini__iPhone_11_Pro__iPhone_XS__iPhone_X_portrait.png"},{"revision":"bc178ab2b5a3aa480d7f7f13a2485a82","url":"assets/screens/iPhone_14__iPhone_13_Pro__iPhone_13__iPhone_12_Pro__iPhone_12_landscape.png"},{"revision":"e2e67ebe0eee3813a9ba677b6fb93ee9","url":"assets/screens/iPhone_14__iPhone_13_Pro__iPhone_13__iPhone_12_Pro__iPhone_12_portrait.png"},{"revision":"806f4f9e495fc9a9d4e5427c4588b1bd","url":"assets/screens/iPhone_14_Plus__iPhone_13_Pro_Max__iPhone_12_Pro_Max_landscape.png"},{"revision":"cd7d3ce83cd66578398dc134aa4ee54a","url":"assets/screens/iPhone_14_Plus__iPhone_13_Pro_Max__iPhone_12_Pro_Max_portrait.png"},{"revision":"ae22e2646a8e1fe973d5862fd4132338","url":"assets/screens/iPhone_16__iPhone_15_Pro__iPhone_15__iPhone_14_Pro_landscape.png"},{"revision":"a3402e670625c09c8ee48f7b2557dca1","url":"assets/screens/iPhone_16__iPhone_15_Pro__iPhone_15__iPhone_14_Pro_portrait.png"},{"revision":"3ac85a8eff8208fd5580cc511000f8b7","url":"assets/screens/iPhone_16_Plus__iPhone_15_Pro_Max__iPhone_15_Plus__iPhone_14_Pro_Max_landscape.png"},{"revision":"3c823439dfcadb1bd440efdb18874bb9","url":"assets/screens/iPhone_16_Plus__iPhone_15_Pro_Max__iPhone_15_Plus__iPhone_14_Pro_Max_portrait.png"},{"revision":"50a33ccbba112671fb0805ddb82bb051","url":"assets/screens/iPhone_16_Pro_landscape.png"},{"revision":"b8e9d871993f89c492758fb203cbcaa5","url":"assets/screens/iPhone_16_Pro_Max_landscape.png"},{"revision":"dac6ed01e699ec18f7baefeefcc51ed2","url":"assets/screens/iPhone_16_Pro_Max_portrait.png"},{"revision":"f846f45bb77bb3afafe51ccf6e710629","url":"assets/screens/iPhone_16_Pro_portrait.png"},{"revision":"d5d16de89702e5b44512900979491925","url":"assets/screens/iPhone_8__iPhone_7__iPhone_6s__iPhone_6__4.7__iPhone_SE_landscape.png"},{"revision":"ab09c7fa3ee69770fbedbdd31cf52963","url":"assets/screens/iPhone_8__iPhone_7__iPhone_6s__iPhone_6__4.7__iPhone_SE_portrait.png"},{"revision":"7e1884bfc1775b764ca9f7c9137849a7","url":"assets/screens/iPhone_8_Plus__iPhone_7_Plus__iPhone_6s_Plus__iPhone_6_Plus_landscape.png"},{"revision":"2a01862fbdc247d76720a0d3a7164309","url":"assets/screens/iPhone_8_Plus__iPhone_7_Plus__iPhone_6s_Plus__iPhone_6_Plus_portrait.png"},{"revision":"adb55f7b6bd155aee484821913573808","url":"assets/scripts/beltech.js"},{"revision":"710c7b9eee37884740dcd5a97f5b3a7b","url":"assets/scripts/echarts.min.js"},{"revision":"d88d3200cb547c9e86ee972975561dcd","url":"assets/scripts/html5-qrcode.min.js"},{"revision":"93063fe04ea37e3b3441590f9abf141a","url":"assets/scripts/htmx.min.js"},{"revision":"35467f89118bd098ed1030b77d78a032","url":"assets/scripts/hyperscript.min.js"},{"revision":"8fb8fee4fcc3cc86ff6c724154c49c42","url":"assets/scripts/jquery.min.js"},{"revision":"96760f90b96ce00be8e4a204dd3d4b98","url":"assets/scripts/swiper-bundle.min.js"},{"revision":"b36f28de584845317de40a7219c82b1c","url":"assets/scripts/toastr.min.js"},{"revision":"5151a944c0a94fe1b1c29ec6caef868c","url":"assets/styles/input.css"},{"revision":"fabdbe3066f3328e4eb37b59cda7413f","url":"assets/styles/output.css"},{"revision":"766238d2d89efd82440dd641f2c569f5","url":"assets/styles/swiper-bundle.min.css"},{"revision":"bc96861d9899e4e68fb2e59c363d8c60","url":"assets/styles/toastr.min.css"},{"revision":"8352c33464e1a27cee5ef6a641d9582a","url":"assets/svg/-.svg"},{"revision":"717be459fb1e21f71818ebe878ea3303","url":"assets/svg/+.svg"},{"revision":"2c6a0cd05942b4e1124dc5a1c636e1f6","url":"assets/svg/20240321_173552.svg"},{"revision":"7718e9bff2d7c8b4241f80f6212f414b","url":"assets/svg/20240322_081249.svg"},{"revision":"8a1473ec409dd1f18559c3cbd120a90c","url":"assets/svg/3d_avatar_3.svg"},{"revision":"1ecde7b4d5c33ea6e4436d464d7a18bb","url":"assets/svg/9mobile.svg"},{"revision":"444682cfc71b2dccd7039c9edd013e57","url":"assets/svg/agentSort.svg"},{"revision":"a356d6a834fa6c6ae5f4c42d71c70f05","url":"assets/svg/airtel.svg"},{"revision":"90a4f314ed0de641eacb1c6fe29c069d","url":"assets/svg/android.svg"},{"revision":"760aafd779e4b7d16e1efb10641340bc","url":"assets/svg/apple.svg"},{"revision":"b4793db50c8d4f9cdbab54d08427aca6","url":"assets/svg/arrow_right.svg"},{"revision":"953de7e0d42e777d13f51950a16f682c","url":"assets/svg/back_arrow.svg"},{"revision":"24f3069c826705cbfaf5fde783743fbe","url":"assets/svg/beltechClearLogo.svg"},{"revision":"d5e548cd2f6758d01e90fae19a154fc7","url":"assets/svg/beltechLogoIndexPage.svg"},{"revision":"d78684acee54cac567b84ddd60017935","url":"assets/svg/bi_qr-code.svg"},{"revision":"a20cf484562d9a4076aa79a3daa330e8","url":"assets/svg/binAgent.svg"},{"revision":"62eceb2756ba45eafc57b207f9eec723","url":"assets/svg/cash out.svg"},{"revision":"4450d58d422ccdfcdc7a58f19533cc72","url":"assets/svg/collect_icon_bold.svg"},{"revision":"b0c40cae6bec0a56011ef0394edd23e5","url":"assets/svg/collect_icon_linear.svg"},{"revision":"e46f0a75d6146347e4c08d5b31e75788","url":"assets/svg/collectorAgent.svg"},{"revision":"f42cecba3f63d0115c63288b5c821e53","url":"assets/svg/confetti.svg"},{"revision":"eae78e7034bd2c0c01920d2971595ba8","url":"assets/svg/connect wallet.svg"},{"revision":"612c88e868b1744a78f4f227c1c70a2b","url":"assets/svg/connectWallet.svg"},{"revision":"efb42e9e2d24e6939724d7644e836c3f","url":"assets/svg/cryptoNews.svg"},{"revision":"76e1f653a725f61513ccedc63d305108","url":"assets/svg/editPFP.svg"},{"revision":"62088362b66b7a1faaabaa3f544038e6","url":"assets/svg/exchange linear.svg"},{"revision":"7343730a9d318b268b88ed975d14b8ec","url":"assets/svg/exchange.svg"},{"revision":"511696215153edc29a63709332195957","url":"assets/svg/eye-off.svg"},{"revision":"6f14ef4ece6dcc17ee077be8e7f95736","url":"assets/svg/eye.svg"},{"revision":"6b446a380e308ad3d84cf955d088f3be","url":"assets/svg/facebookLogo.svg"},{"revision":"320d76587828d5e6dee1a3435d5dc114","url":"assets/svg/filter_small.svg"},{"revision":"6a4f5cf5b739324d72ae3624f359a032","url":"assets/svg/flying_world.svg"},{"revision":"9d888438a313f3081349092fee4302ad","url":"assets/svg/glassSwipper.svg"},{"revision":"2a92922d15808c1cb9af738f77989283","url":"assets/svg/glo.svg"},{"revision":"3b4bae837a4c84427e8c2eb19b0bf91f","url":"assets/svg/global.svg"},{"revision":"a32befb389d5900111492e9fcc615ecf","url":"assets/svg/googleLogo.svg"},{"revision":"bef39d33a6cbd7ec44e3134aa15a0814","url":"assets/svg/greenEdit.svg"},{"revision":"390cd22af96671d30b6c26daae1db511","url":"assets/svg/greenEye.svg"},{"revision":"f343df11f224b123418a15f2cb2e7a7f","url":"assets/svg/grid.svg"},{"revision":"3779f6f9d272aa83945e1de6dd0f21de","url":"assets/svg/history_icon_bold.svg"},{"revision":"ba1d8b617ed29defed9bce5121068714","url":"assets/svg/history_icon_linear.svg"},{"revision":"92e593051d7629cffa1afcd1b5edcf14","url":"assets/svg/history.svg"},{"revision":"6324d5829cc8cb9d0126b501b5c7d46e","url":"assets/svg/home bold.svg"},{"revision":"903124e15c3955d2c8fb47f3a74bfac3","url":"assets/svg/home linear.svg"},{"revision":"fd6fb1d9c5f8aaba084cfe39e30e01fb","url":"assets/svg/learn.svg"},{"revision":"957c64539fd5ca72a8d872ebbbcb4fa5","url":"assets/svg/linker_arrow.svg"},{"revision":"22ae8e4508549a29205953376ee9ee39","url":"assets/svg/location bold.svg"},{"revision":"cbf8d208455706e7a786be39eb6ec617","url":"assets/svg/location linear.svg"},{"revision":"7f5c3648f80056e6e278f5c559290b1a","url":"assets/svg/login.svg"},{"revision":"cf619d9beff3f608f43cb25975e9276b","url":"assets/svg/logout_bold.svg"},{"revision":"7e7add644feb84fff262fff6ed17d2c6","url":"assets/svg/logout_linear.svg"},{"revision":"b2da15e55ef3aad718ea45aa999d20c6","url":"assets/svg/manWithPlastics.svg"},{"revision":"3287442a4d85d1990470a87fd5f7fc77","url":"assets/svg/metalSwipper.svg"},{"revision":"285c218ec5f38a2c9af87b1cbac0679f","url":"assets/svg/mtn.svg"},{"revision":"679db0306325bd33762430d801bd00d2","url":"assets/svg/new_exchange_bold.svg"},{"revision":"2669111a5b94d862a07ff0ae3b414d8c","url":"assets/svg/new_exchange_linear.svg"},{"revision":"e3346a22e7bfeb94f3a2a3913f9783c0","url":"assets/svg/nft_box.svg"},{"revision":"5ed9415df609f6d2ec8544adc33ee1e7","url":"assets/svg/notification.svg"},{"revision":"4c50576ce5a6a97c9d4cc5850d120783","url":"assets/svg/octicon_history-16.svg"},{"revision":"bccd3d01c3d7e30af0509da3124d73ab","url":"assets/svg/paperSwipper.svg"},{"revision":"e78b0ec6287411c7a490c02b46a6b3d7","url":"assets/svg/profile-add.svg"},{"revision":"f80cc4d0f6c3747b32471dcd526938cb","url":"assets/svg/profile-home.svg"},{"revision":"32446bff012b9e8fc92d6ba2af93b24e","url":"assets/svg/profile.svg"},{"revision":"e600e463b453adae9e855db18cd1afe9","url":"assets/svg/recycle_3d.svg"},{"revision":"e09d2c3be2f79eb1260c5e055e0990ca","url":"assets/svg/recycling_people.svg"},{"revision":"e9b55faa4c931c0210bb6b6e56f70dff","url":"assets/svg/redBin.svg"},{"revision":"a799da63db230cce37b7ab50d98e10d4","url":"assets/svg/scan bold.svg"},{"revision":"574d4dcf15531871475642a73588d91d","url":"assets/svg/scan linear.svg"},{"revision":"59f8ff971a21cbeeae15b63ebba263ba","url":"assets/svg/search-svgrepo-com.svg"},{"revision":"59f8ff971a21cbeeae15b63ebba263ba","url":"assets/svg/search.svg"},{"revision":"7a45e022716a9078d75b35b5ec353db5","url":"assets/svg/spendVoucher.svg"},{"revision":"9c043057b815cb4e34acc3fab39ae18c","url":"assets/svg/subscribe.svg"},{"revision":"6c4d43e5955f447bf997809f0d002e22","url":"assets/svg/unauthorized.svg"},{"revision":"935bf9107ae84cf782aa274e906925de","url":"assets/svg/user_rewards.svg"},{"revision":"9a5200976d1be705553a679fc7fff5c9","url":"assets/svg/user_scan.svg"},{"revision":"2ce1c6095ac389567091a5726a5b87b6","url":"assets/svg/user_signin.svg"},{"revision":"8991f69e9e6ce5a3fada072372672930","url":"assets/svg/vivid-ticket.svg"},{"revision":"108a6ecea300a80ea8c70f1ab0fb34bb","url":"assets/svg/voucher.svg"},{"revision":"114c18d03d9681aa93b5a9280facd12f","url":"assets/svg/wallet bold.svg"},{"revision":"09550603ddaf720aa1e1521ffee6b3c3","url":"assets/svg/wallet linear.svg"},{"revision":"408d6d8f64b302e6caf34f732abf1141","url":"assets/svg/waste_bin.svg"},{"revision":"156d977bd57cbcceac44060f6c479cad","url":"manifest.json"}] || []);
})();
//# sourceMappingURL=sw.js.map
