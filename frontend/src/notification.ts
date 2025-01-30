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
  notificationId: string; // Unique per notification (but not globally)
  notificationHeading: string;
  notificationTitle: string;
  notificationBody: string;
  notificationTime: string; // Use ISO string for easy sorting/filtering
};

type lastNotification = {
  accountId: string;
  time: string;
};

export const dbPromise = function () {
  const db = openDB<MyIndexDB>("beltech", INDEXDB_CURR_VERSION, {
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

export const addNewNotification = async function (
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
export const deleteNotification = async function (
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

export const updateLastNotificationIfMoreRecent = async function (
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

// export const getNotificationMarkup = (accountId: string) => {
//   const db = await dbPromise();
//   const accountData = await db.get("accounts-notifications", accountId);

//   if (
//     !accountData ||
//     !accountData.notifications ||
//     accountData.notifications.length === 0
//   ) {
//     return `<p class="no-new text-black font-poppins text-sm text-center m-auto">No new notifications</p>`;
//   }

//   // Generate the notification items
//   const notificationsMarkup = accountData.notifications
//     .map((notif) => {
//       return `
//       <li class="flex justify-between items-start gap-[0.98rem] [transition:transform_0.8s_ease-in,opacity_0.6s_ease-in]">
//         <img
//           src="/public/assets/svg/closeBtn.svg"
//           class="close-btn w-[1.324rem] aspect-square"
//           _="on click
//             tell the closest <li/>
//               add .slide-out then settle then remove yourself
//               then if #notification-bar.querySelectorAll('li').length === 0
//                 remove .hidden from .no-new then halt
//               end"
//         />
//         <div>
//           <section class="w-full flex justify-between items-center mb-[0.441rem]">
//             <p class="badge bg-[#BB1D80] text-[#FFFFFF] font-bold text-[0.511rem]
//               leading-[0.623rem] text-center px-[0.929rem] py-[0.078rem] rounded-[0.256rem]">
//               ${notif.notificationTitle}
//             </p>
//             <p class="text-[#3C3C4399] font-poppins font-normal text-[0.477rem] leading-[0.715rem]">
//               ${new Date(notif.notificationTime).toLocaleString()}
//             </p>
//           </section>
//           <div class="content">
//             <p class="text-[#000000] font-bold text-start text-[0.626rem] leading-[0.763rem] mb-[0.313rem]">
//               ${notif.notificationHeading}
//             </p>
//             <p class="text-[#000000] font-poppins font-normal text-[0.626rem] text-start leading-[0.938rem]">
//               ${notif.notificationBody}
//             </p>
//           </div>
//         </div>
//       </li>
//     `;
//     })
//     .join(""); // Join the array into a single string

//   return `
//     <ul class="flex flex-col gap-[2.635rem] justify-center items-center">
//       ${notificationsMarkup}
//     </ul>
//   `;
// };

declare global {
  interface Window {
    getNotificationMarkup: (accountId: string) => void;
  }
}

window.getNotificationMarkup = (accountId: string) => {
  // Return a placeholder while we fetch data asynchronously
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
            })
            .join("")}
        </ul>
      `;
    }

    // Update the notification bar content dynamically
    const notificationBar = document.getElementById("notification-bar");
    if (notificationBar) {
      notificationBar.innerHTML = markup;
    }
  })();
  // return `<p class="text-gray-500 text-sm text-center">Loading notifications...</p>`; // Placeholder while loading
};
function getRandomBgColor() {
  const colors = ["bg-[#BB1D80]", "bg-[#1BCA71]", "bg-[#EBAB01]"];
  return colors[Math.floor(Math.random() * colors.length)];
}

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

    console.log("response ok", res.status);

    const accountId = res.headers.get("X-AccountId");
    if (!accountId) {
      console.warn("Missing X-AccountId header");
      return;
    }

    console.log("accountId", accountId);

    const data: Notification[] = await res.json();

    console.log("data", data);

    for (const notification of data) {
      const isNewNotification: boolean =
        await updateLastNotificationIfMoreRecent(accountId, notification);

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
