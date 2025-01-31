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
const idb_1 = require("idb");
const INDEXDB_CURR_VERSION = 1;
const INDEX_PREV_VERSION = 0;
const dbPromise = function () {
    const db = (0, idb_1.openDB)("beltech", INDEXDB_CURR_VERSION, {
        upgrade(db) {
            console.log("Creating object store");
            if (!db.objectStoreNames.contains("accounts-notifications")) {
                db.createObjectStore("accounts-notifications", {
                    keyPath: "accountId",
                });
            }
            if (!db.objectStoreNames.contains("accounts-recent")) {
                db.createObjectStore("accounts-recent", { keyPath: "accountId" });
            }
        },
    });
    return db;
};
const addNewNotification = function (accountId, newNotification) {
    return __awaiter(this, void 0, void 0, function* () {
        try {
            const db = yield dbPromise();
            const accountData = yield db.get("accounts-notifications", accountId);
            if (!accountData) {
                yield db.put("accounts-notifications", {
                    accountId,
                    notifications: [newNotification],
                });
                return;
            }
            const notificationExists = accountData.notifications.some((notification) => notification.notificationId === newNotification.notificationId);
            if (notificationExists) {
                return;
            }
            accountData.notifications.push(newNotification);
            yield db.put("accounts-notifications", accountData);
        }
        catch (error) {
            console.error(`Error adding notification for accountId: ${accountId}`, error);
            throw error;
        }
    });
};
const deleteNotification = function (accountId, notificationId) {
    return __awaiter(this, void 0, void 0, function* () {
        try {
            const db = yield dbPromise();
            const accountData = yield db.get("accounts-notifications", accountId);
            if (!accountData) {
                return;
            }
            const updatedNotifications = accountData.notifications.filter((notification) => notification.notificationId !== notificationId);
            if (accountData.notifications.length === updatedNotifications.length) {
                return;
            }
            yield db.put("accounts-notifications", {
                accountId,
                notifications: updatedNotifications,
            });
        }
        catch (error) {
            console.error(`Error deleting notification for accountId: ${accountId}, notificationId: ${notificationId}`, error);
            throw error;
        }
    });
};
const updateLastNotificationIfMoreRecent = function (accountId, newNotification) {
    return __awaiter(this, void 0, void 0, function* () {
        try {
            const db = yield dbPromise();
            const lastNotification = yield db.get("accounts-recent", accountId);
            if (!lastNotification) {
                yield db.put("accounts-recent", {
                    accountId,
                    time: newNotification.notificationTime,
                });
                return true; // Indicates an update was made
            }
            const newNotificationTime = new Date(newNotification.notificationTime);
            if (newNotificationTime > new Date(lastNotification.time)) {
                yield db.put("accounts-recent", {
                    accountId,
                    time: newNotification.notificationTime,
                });
                return true; // Indicates an update was made
            }
            return false; // Indicates no update was necessary
        }
        catch (error) {
            console.error(`Error updating last notification for accountId: ${accountId}`, error);
            throw error;
        }
    });
};
function getRandomBgColor() {
    const colors = ["bg-[#BB1D80]", "bg-[#1BCA71]", "bg-[#EBAB01]"];
    return colors[Math.floor(Math.random() * colors.length)];
}
function updateNotifMarkup(markup) {
    const notificationBar = document.getElementById("notification-bar");
    if (notificationBar) {
        notificationBar.innerHTML = markup;
        window._hyperscript.processNode(notificationBar);
    }
}
window.updateEmptyNotifMarkup = () => {
    const notificationBar = document.getElementById("notification-bar");
    if (notificationBar) {
        notificationBar.innerHTML =
            '<p class="no-new text-black font-poppins text-sm text-center m-auto">No new notification</p>';
    }
};
window.deleteNotification = deleteNotification;
window.getNotificationMarkup = (accountId) => {
    (() => __awaiter(void 0, void 0, void 0, function* () {
        const db = yield dbPromise();
        const accountData = yield db.get("accounts-notifications", accountId);
        let markup = "";
        if (!accountData ||
            !accountData.notifications ||
            accountData.notifications.length === 0) {
            markup = `<p class="no-new text-black font-poppins text-sm text-center m-auto">No new notifications</p>`;
        }
        else {
            markup = `
        <ul class="flex flex-col gap-[2.635rem] justify-center items-center">
          ${accountData.notifications
                .map((notif) => {
                return produceNotifLiMarkup(notif);
            })
                .join("")}
        </ul>
      `;
        }
        updateNotifMarkup(markup);
    }))();
};
function produceNotifLiMarkup(notif) {
    return `
  <li class="flex justify-between w-full items-start gap-[0.98rem] [transition:transform_0.8s_ease-in,opacity_0.6s_ease-in]">
    <img id="closeBtn-img"
      src="/public/assets/svg/closeBtn.svg"
      class="close-btn w-[1.324rem] aspect-square"
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
}
(() => __awaiter(void 0, void 0, void 0, function* () {
    try {
        const res = yield fetch("/account/notifications", {
            headers: {
                Accept: "application/json",
            },
        });
        if (!res.ok) {
            console.warn("response not ok", res.status);
            return;
        }
        const accountId = res.headers.get("X-AccountId");
        if (!accountId) {
            console.warn("Missing X-AccountId header");
            return;
        }
        // console.log("accountId", accountId);
        const data = yield res.json();
        // console.log("data", data);
        for (const notification of data) {
            const isNewNotification = yield updateLastNotificationIfMoreRecent(accountId, notification);
            // console.log("isNewNotification", isNewNotification);
            if (!isNewNotification) {
                console.log("Notification already exists, skipping", notification);
                continue;
            }
            console.log("Adding new notification", notification);
            yield addNewNotification(accountId, notification);
        }
    }
    catch (error) {
        console.error("Error fetching notifications:", error);
    }
}))();
// document.addEventListener("click", (event: Event) => {
//   const target = event.target as HTMLElement | null;
// if (target && target.classList.contains("close-btn")) {
//   const li = target.closest("li");
//   if (!li) {
//     return;
//   }
// li.classList.add("slide-out");
// setTimeout(() => {
//   li.remove();
//   deleteNotification("giwaoluwatobi@gmail.com", li.id);
//   const notificationBar = document.getElementById("notification-bar");
//   if (notificationBar) {
//     const notifications = notificationBar.querySelectorAll("li");
//     if (notifications.length === 0) {
//       updateNotifMarkup(
//         `<p class="no-new text-black font-poppins text-sm text-center m-auto">No new notifications</p>`
//       );
//     }
//   }
// }, 800);
// }
// });
