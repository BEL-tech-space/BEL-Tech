"use strict";
(() => {
  // node_modules/.pnpm/idb@8.0.1/node_modules/idb/build/index.js
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
  var transactionDoneMap = /* @__PURE__ */ new WeakMap();
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
    if (getCursorAdvanceMethods().includes(func)) {
      return function(...args) {
        func.apply(unwrap(this), args);
        return wrap(this.request);
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
  var advanceMethodProps = ["continue", "continuePrimaryKey", "advance"];
  var methodMap = {};
  var advanceResults = /* @__PURE__ */ new WeakMap();
  var ittrProxiedCursorToOriginalProxy = /* @__PURE__ */ new WeakMap();
  var cursorIteratorTraps = {
    get(target, prop) {
      if (!advanceMethodProps.includes(prop))
        return target[prop];
      let cachedFunc = methodMap[prop];
      if (!cachedFunc) {
        cachedFunc = methodMap[prop] = function(...args) {
          advanceResults.set(this, ittrProxiedCursorToOriginalProxy.get(this)[prop](...args));
        };
      }
      return cachedFunc;
    }
  };
  async function* iterate(...args) {
    let cursor = this;
    if (!(cursor instanceof IDBCursor)) {
      cursor = await cursor.openCursor(...args);
    }
    if (!cursor)
      return;
    cursor = cursor;
    const proxiedCursor = new Proxy(cursor, cursorIteratorTraps);
    ittrProxiedCursorToOriginalProxy.set(proxiedCursor, cursor);
    reverseTransformCache.set(proxiedCursor, unwrap(cursor));
    while (cursor) {
      yield proxiedCursor;
      cursor = await (advanceResults.get(proxiedCursor) || cursor.continue());
      advanceResults.delete(proxiedCursor);
    }
  }
  function isIteratorProp(target, prop) {
    return prop === Symbol.asyncIterator && instanceOfAny(target, [IDBIndex, IDBObjectStore, IDBCursor]) || prop === "iterate" && instanceOfAny(target, [IDBIndex, IDBObjectStore]);
  }
  replaceTraps((oldTraps) => ({
    ...oldTraps,
    get(target, prop, receiver) {
      if (isIteratorProp(target, prop))
        return iterate;
      return oldTraps.get(target, prop, receiver);
    },
    has(target, prop) {
      return isIteratorProp(target, prop) || oldTraps.has(target, prop);
    }
  }));

  // src/notification.ts
  var INDEXDB_CURR_VERSION = 1;
  var dbPromise = function() {
    const db = openDB("beltech", INDEXDB_CURR_VERSION, {
      upgrade(db2) {
        console.log("Creating object store");
        if (!db2.objectStoreNames.contains("accounts-notifications")) {
          db2.createObjectStore("accounts-notifications", {
            keyPath: "accountId"
          });
        }
        if (!db2.objectStoreNames.contains("accounts-recent")) {
          db2.createObjectStore("accounts-recent", { keyPath: "accountId" });
        }
      }
    });
    return db;
  };
  var addNewNotification = async function(accountId, newNotification) {
    try {
      const db = await dbPromise();
      const accountData = await db.get("accounts-notifications", accountId);
      if (!accountData) {
        await db.put("accounts-notifications", {
          accountId,
          notifications: [newNotification]
        });
        return;
      }
      const notificationExists = accountData.notifications.some(
        (notification) => notification.notificationId === newNotification.notificationId
      );
      if (notificationExists) {
        return;
      }
      accountData.notifications.push(newNotification);
      await db.put("accounts-notifications", accountData);
    } catch (error) {
      console.error(
        `Error adding notification for accountId: ${accountId}`,
        error
      );
      throw error;
    }
  };
  var deleteNotification = async function(accountId, notificationId) {
    try {
      const db = await dbPromise();
      const accountData = await db.get("accounts-notifications", accountId);
      if (!accountData) {
        return;
      }
      const updatedNotifications = accountData.notifications.filter(
        (notification) => notification.notificationId !== notificationId
      );
      if (accountData.notifications.length === updatedNotifications.length) {
        return;
      }
      await db.put("accounts-notifications", {
        accountId,
        notifications: updatedNotifications
      });
    } catch (error) {
      console.error(
        `Error deleting notification for accountId: ${accountId}, notificationId: ${notificationId}`,
        error
      );
      throw error;
    }
  };
  var updateLastNotificationIfMoreRecent = async function(accountId, newNotification) {
    try {
      const db = await dbPromise();
      const lastNotification = await db.get("accounts-recent", accountId);
      if (!lastNotification) {
        await db.put("accounts-recent", {
          accountId,
          time: newNotification.notificationTime
        });
        return true;
      }
      const newNotificationTime = new Date(newNotification.notificationTime);
      if (newNotificationTime > new Date(lastNotification.time)) {
        await db.put("accounts-recent", {
          accountId,
          time: newNotification.notificationTime
        });
        return true;
      }
      return false;
    } catch (error) {
      console.error(
        `Error updating last notification for accountId: ${accountId}`,
        error
      );
      throw error;
    }
  };
  window.getNotificationMarkup = (accountId) => {
    (async () => {
      const db = await dbPromise();
      const accountData = await db.get("accounts-notifications", accountId);
      let markup = "";
      if (!accountData || !accountData.notifications || accountData.notifications.length === 0) {
        markup = `<p class="no-new text-black font-poppins text-sm text-center m-auto">No new notifications</p>`;
      } else {
        markup = `
        <ul class="flex flex-col gap-[2.635rem] justify-center items-center">
          ${accountData.notifications.map((notif) => {
          return `
              <li class="flex justify-between w-full items-start gap-[0.98rem] 
                [transition:transform_0.8s_ease-in,opacity_0.6s_ease-in]">
                <img
                  src="/public/assets/svg/closeBtn.svg"
                  class="close-btn w-[1.324rem] aspect-square"
                  _="on click
                    tell the closest <li/>
                      add .slide-out then settle then remove yourself
                      then if #notification-bar.querySelectorAll('li').length === 0
                        remove .hidden from .no-new then halt
                      end"
                />
                <div>
                  <section class="w-full flex justify-between items-center mb-[0.441rem]">
                    <p class="badge ${getRandomBgColor()} text-[#FFFFFF] font-bold text-[0.511rem] 
                      leading-[0.623rem] text-center px-[0.929rem] py-[0.078rem] rounded-[0.256rem]">
                      ${notif.notificationTitle}
                    </p>
                    <p class="text-[#3C3C4399] font-poppins font-normal text-[0.477rem] leading-[0.715rem]">
                      ${new Date(notif.notificationTime).toLocaleString()}
                    </p>
                  </section>
                  <div class="content">
                    <p class="text-[#000000] font-bold text-start text-[0.626rem] leading-[0.763rem] mb-[0.313rem]">
                      ${notif.notificationHeading}
                    </p>
                    <p class="text-[#000000] font-poppins font-normal text-[0.626rem] text-start leading-[0.938rem] max-w-[50ch] break-words">
                      ${notif.notificationBody}
                    </p>
                  </div>
                </div>
              </li>
            `;
        }).join("")}
        </ul>
      `;
      }
      const notificationBar = document.getElementById("notification-bar");
      if (notificationBar) {
        notificationBar.innerHTML = markup;
      }
    })();
  };
  function getRandomBgColor() {
    const colors = ["bg-[#BB1D80]", "bg-[#1BCA71]", "bg-[#EBAB01]"];
    return colors[Math.floor(Math.random() * colors.length)];
  }
  (async () => {
    try {
      const res = await fetch("/account/notifications", {
        headers: {
          Accept: "application/json"
        }
      });
      if (!res.ok) {
        console.warn("response not ok", res.status);
        return;
      }
      console.log("response ok", res.status);
      const accountId = res.headers.get("X-AccountId");
      if (!accountId) {
        console.warn("Missing X-AccountId header");
        return;
      }
      console.log("accountId", accountId);
      const data = await res.json();
      console.log("data", data);
      for (const notification of data) {
        const isNewNotification = await updateLastNotificationIfMoreRecent(accountId, notification);
        console.log("isNewNotification", isNewNotification);
        if (!isNewNotification) {
          console.log("Notification already exists, skipping", notification);
          continue;
        }
        console.log("Adding new notification", notification);
        await addNewNotification(accountId, notification);
      }
    } catch (error) {
      console.error("Error fetching notifications:", error);
    }
  })();
})();
//# sourceMappingURL=notification.js.map
