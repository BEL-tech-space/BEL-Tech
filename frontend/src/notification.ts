import { IDBPDatabase, openDB, DBSchema } from "idb";

const INDEXDB_CURR_VERSION = 1;
const INDEX_PREV_VERSION = 0;

interface MyIndexDB extends DBSchema {
  "accounts-notifications": {
    key: string; // accountId (email)
    value: {
      accountId: string;
      notifications: Notification[];
    };
  };
  "accounts-recent": {
    key: string; // accountId (email)
    value: lastNotification;
  };
}

type Notification = {
  accountId: string;
  notificationId: string; // Unique per notification (but not globally)
  notificationHeading: string;
  notificationTitle: string;
  notificationBody: string;
  notificationTime: string;
};

type lastNotification = {
  accountId: string;
  time: string;
};

const dbPromise = function () {
  const db = openDB<MyIndexDB>("beltech", INDEXDB_CURR_VERSION, {
    upgrade(db) {
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

const addNewNotification = async function (
  accountId: string,
  newNotification: Notification
) {
  try {
    const db = await dbPromise();
    const accountData = await db.get("accounts-notifications", accountId);

    if (!accountData) {
      await db.put("accounts-notifications", {
        accountId,
        notifications: [newNotification],
      });
      return;
    }

    const notificationExists = accountData.notifications.some(
      (notification) =>
        notification.notificationId === newNotification.notificationId
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

const deleteNotification = async function (
  accountId: string,
  notificationId: string
) {
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
      notifications: updatedNotifications,
    });
  } catch (error) {
    console.error(
      `Error deleting notification for accountId: ${accountId}, notificationId: ${notificationId}`,
      error
    );
    throw error;
  }
};

const isNotificationRecent = async function (
  accountId: string,
  newNotification: Notification
) {
  try {
    const db = await dbPromise();
    const lastNotification = await db.get("accounts-recent", accountId);
    if (!lastNotification) {
      await db.put("accounts-recent", {
        accountId,
        time: newNotification.notificationTime,
      });
      return true; // Indicates an update was made
    }

    const newNotificationTime = new Date(newNotification.notificationTime);
    if (newNotificationTime > new Date(lastNotification.time)) {
      await db.put("accounts-recent", {
        accountId,
        time: newNotification.notificationTime,
      });
      return true; // Indicates an update was made
    }

    return false; // Indicates no update was necessary
  } catch (error) {
    console.error(
      `Error updating last notification for accountId: ${accountId}`,
      error
    );
    throw error;
  }
};
function getRandomBgColor(): string {
  const colors = ["bg-[#BB1D80]", "bg-[#1BCA71]", "bg-[#EBAB01]"];
  return colors[Math.floor(Math.random() * colors.length)];
}

declare global {
  interface Window {
    _hyperscript: any;
    getNotificationMarkup: (accountId: string) => void;
    deleteNotification: (accountId: string, notificationId: string) => void;
    updateNotificationCount: () => void;
  }
}

function updateNotifMarkup(markup: string) {
  const notificationBar = document.getElementById("notification-bar");
  if (notificationBar) {
    notificationBar.innerHTML = markup;
    window._hyperscript.processNode(notificationBar);
  }
}

window.deleteNotification = deleteNotification;
window.getNotificationMarkup = (accountId: string) => {
  (async () => {
    const db = await dbPromise();
    const accountData = await db.get("accounts-notifications", accountId);

    let markup = "";

    if (
      !accountData ||
      !accountData.notifications ||
      accountData.notifications.length === 0
    ) {
      markup = `<p class="no-new text-black font-poppins text-sm text-center m-auto">No new notifications</p>`;
    } else {
      markup = `
      <ul class="flex flex-col gap-[2.635rem] justify-center items-center">
      ${accountData.notifications
        .map((notif) => {
          return produceNotifLiMarkup(notif);
        })
        .join("")}
      <p class="hidden no-new text-black font-poppins text-sm text-center m-auto">No new notifications</p>
      </ul>
      `;
    }

    updateNotifMarkup(markup);
  })();
};

function produceNotifLiMarkup(notif: Notification) {
  return `
  <li id="${notif.notificationId}" class="flex justify-between w-full items-start gap-[0.98rem] [transition:transform_0.8s_ease-in,opacity_0.6s_ease-in]">
    <img id="closeBtn-img" src="/public/assets/svg/closeBtn.svg" class="close-btn w-[1.324rem] aspect-square"
      _="on click
        halt the event
        tell the closest <li/>
          add .slide-out then settle then remove yourself
          then set :notifId to your id
          then set :accountId to #notification-bar.dataset.accountid
          then call deleteNotification(:accountId, :notifId)
          then call updateNotificationCount()
          then if <li/> in #notification-bar is empty
            remove .hidden from .no-new
          end"
    />
    <div>
      <section class="w-full flex justify-between items-center mb-[0.441rem]">
        <p class="badge ${getRandomBgColor()} text-[#FFFFFF] font-bold text-[0.511rem] leading-[0.623rem] text-center px-[0.929rem] py-[0.078rem] rounded-[0.256rem]">${notif.notificationTitle}</p>
        <p class="text-[#3C3C4399] font-poppins font-extralight text-[0.477rem] leading-[0.715rem]">${new Date(notif.notificationTime).toLocaleString()}</p>
      </section>
      <div>
        <p class="text-[#000000] font-bold text-start text-[0.626rem] leading-[0.763rem] mb-[0.313rem]">${notif.notificationHeading}</p>
        <p class="text-[#000000] font-poppins font-extralight text-[0.626rem] text-start leading-[0.938rem] max-w-[50ch] break-words">${notif.notificationBody}</p>
      </div>
    </div>
  </li>
  `;
}

const allowedPaths = ["dashboard", "welcome"];
if (allowedPaths.some((path) => window.location.pathname.includes(path))) {
  (async () => {
    try {
      const res = await fetch("/account/notifications", {
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

      const data: Notification[] = await res.json();

      for (const notification of data) {
        const isNewNotification: boolean = await isNotificationRecent(
          accountId,
          notification
        );

        if (!isNewNotification) {
          console.log("Notification in the past, skipping", notification);
          continue;
        }

        console.log("Adding new notification", notification);

        await addNewNotification(accountId, notification);
        await updateNotificationCount();
      }
    } catch (error) {
      console.error("Error fetching notifications:", error);
    }
  })();
} else {
  (async () => {
    await updateNotificationCount();
  })();
}

window.updateNotificationCount = updateNotificationCount;
async function updateNotificationCount() {
  const notificationBar = document.getElementById("notification-bar");
  if (!notificationBar) {
    return;
  }
  const accountID = notificationBar.dataset.accountid;
  if (!accountID) {
    return;
  }
  const db = await dbPromise();
  const accountData = await db.get("accounts-notifications", accountID);
  if (!accountData) {
    return;
  }
  const count = accountData.notifications.length;
  const notificationCount = document.getElementById("notificationCount");
  if (!notificationCount) {
    return;
  }
  if (count === 0) {
    notificationCount.classList.add("hidden");
    notificationCount.classList.remove(
      "flex",
      "items-center",
      "justify-center",
      "w-4",
      "h-4"
    );
    return;
  }

  notificationCount.innerText = count.toString();
  notificationCount.classList.remove("hidden");
  notificationCount.classList.add(
    "flex",
    "items-center",
    "justify-center",
    "w-4",
    "h-4"
  );
}
