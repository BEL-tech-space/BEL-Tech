"use strict";(()=>{var p=(t,n)=>n.some(e=>t instanceof e),b,D;function P(){return b||(b=[IDBDatabase,IDBObjectStore,IDBIndex,IDBCursor,IDBTransaction])}function L(){return D||(D=[IDBCursor.prototype.advance,IDBCursor.prototype.continue,IDBCursor.prototype.continuePrimaryKey])}var h=new WeakMap,m=new WeakMap,d=new WeakMap;function v(t){let n=new Promise((e,i)=>{let o=()=>{t.removeEventListener("success",r),t.removeEventListener("error",c)},r=()=>{e(f(t.result)),o()},c=()=>{i(t.error),o()};t.addEventListener("success",r),t.addEventListener("error",c)});return d.set(n,t),n}function C(t){if(h.has(t))return;let n=new Promise((e,i)=>{let o=()=>{t.removeEventListener("complete",r),t.removeEventListener("error",c),t.removeEventListener("abort",c)},r=()=>{e(),o()},c=()=>{i(t.error||new DOMException("AbortError","AbortError")),o()};t.addEventListener("complete",r),t.addEventListener("error",c),t.addEventListener("abort",c)});h.set(t,n)}var w={get(t,n,e){if(t instanceof IDBTransaction){if(n==="done")return h.get(t);if(n==="store")return e.objectStoreNames[1]?void 0:e.objectStore(e.objectStoreNames[0])}return f(t[n])},set(t,n,e){return t[n]=e,!0},has(t,n){return t instanceof IDBTransaction&&(n==="done"||n==="store")?!0:n in t}};function E(t){w=t(w)}function T(t){return L().includes(t)?function(...n){return t.apply(y(this),n),f(this.request)}:function(...n){return f(t.apply(y(this),n))}}function j(t){return typeof t=="function"?T(t):(t instanceof IDBTransaction&&C(t),p(t,P())?new Proxy(t,w):t)}function f(t){if(t instanceof IDBRequest)return v(t);if(m.has(t))return m.get(t);let n=j(t);return n!==t&&(m.set(t,n),d.set(n,t)),n}var y=t=>d.get(t);function M(t,n,{blocked:e,upgrade:i,blocking:o,terminated:r}={}){let c=indexedDB.open(t,n),u=f(c);return i&&c.addEventListener("upgradeneeded",a=>{i(f(c.result),a.oldVersion,a.newVersion,f(c.transaction),a)}),e&&c.addEventListener("blocked",a=>e(a.oldVersion,a.newVersion,a)),u.then(a=>{r&&a.addEventListener("close",()=>r()),o&&a.addEventListener("versionchange",s=>o(s.oldVersion,s.newVersion,s))}).catch(()=>{}),u}var S=["get","getKey","getAll","getAllKeys","count"],A=["put","add","delete","clear"],g=new Map;function B(t,n){if(!(t instanceof IDBDatabase&&!(n in t)&&typeof n=="string"))return;if(g.get(n))return g.get(n);let e=n.replace(/FromIndex$/,""),i=n!==e,o=A.includes(e);if(!(e in(i?IDBIndex:IDBObjectStore).prototype)||!(o||S.includes(e)))return;let r=async function(c,...u){let a=this.transaction(c,o?"readwrite":"readonly"),s=a.store;return i&&(s=s.index(u.shift())),(await Promise.all([s[e](...u),o&&a.done]))[0]};return g.set(n,r),r}E(t=>({...t,get:(n,e,i)=>B(n,e)||t.get(n,e,i),has:(n,e)=>!!B(n,e)||t.has(n,e)}));var V=["continue","continuePrimaryKey","advance"],x={},I=new WeakMap,k=new WeakMap,$={get(t,n){if(!V.includes(n))return t[n];let e=x[n];return e||(e=x[n]=function(...i){I.set(this,k.get(this)[n](...i))}),e}};async function*F(...t){let n=this;if(n instanceof IDBCursor||(n=await n.openCursor(...t)),!n)return;n=n;let e=new Proxy(n,$);for(k.set(e,n),d.set(e,y(n));n;)yield e,n=await(I.get(e)||n.continue()),I.delete(e)}function N(t,n){return n===Symbol.asyncIterator&&p(t,[IDBIndex,IDBObjectStore,IDBCursor])||n==="iterate"&&p(t,[IDBIndex,IDBObjectStore])}E(t=>({...t,get(n,e,i){return N(n,e)?F:t.get(n,e,i)},has(n,e){return N(n,e)||t.has(n,e)}}));var O=1;var l=function(){return M("beltech",O,{upgrade(n){console.log("Creating object store"),n.objectStoreNames.contains("accounts-notifications")||n.createObjectStore("accounts-notifications",{keyPath:"accountId"}),n.objectStoreNames.contains("accounts-recent")||n.createObjectStore("accounts-recent",{keyPath:"accountId"})}})},_=async function(t,n){try{let e=await l(),i=await e.get("accounts-notifications",t);if(!i){await e.put("accounts-notifications",{accountId:t,notifications:[n]});return}if(i.notifications.some(r=>r.notificationId===n.notificationId))return;i.notifications.push(n),await e.put("accounts-notifications",i)}catch(e){throw console.error(`Error adding notification for accountId: ${t}`,e),e}},R=async function(t,n){try{let e=await l(),i=await e.get("accounts-notifications",t);if(!i)return;let o=i.notifications.filter(r=>r.notificationId!==n);if(i.notifications.length===o.length)return;await e.put("accounts-notifications",{accountId:t,notifications:o})}catch(e){throw console.error(`Error deleting notification for accountId: ${t}, notificationId: ${n}`,e),e}},W=async function(t,n){try{let e=await l(),i=await e.get("accounts-recent",t);return i?new Date(n.notificationTime)>new Date(i.time)?(await e.put("accounts-recent",{accountId:t,time:n.notificationTime}),!0):!1:(await e.put("accounts-recent",{accountId:t,time:n.notificationTime}),!0)}catch(e){throw console.error(`Error updating last notification for accountId: ${t}`,e),e}};function H(){let t=["bg-[#BB1D80]","bg-[#1BCA71]","bg-[#EBAB01]"];return t[Math.floor(Math.random()*t.length)]}function K(t){let n=document.getElementById("notification-bar");n&&(n.innerHTML=t,window._hyperscript.processNode(n))}window.updateEmptyNotifMarkup=()=>{let t=document.getElementById("notification-bar");t&&(t.innerHTML='<p class="no-new text-black font-poppins text-sm text-center m-auto">No new notification</p>')};window.deleteNotification=R;window.getNotificationMarkup=t=>{(async()=>{let e=await(await l()).get("accounts-notifications",t),i="";!e||!e.notifications||e.notifications.length===0?i='<p class="no-new text-black font-poppins text-sm text-center m-auto">No new notifications</p>':i=`
        <ul class="flex flex-col gap-[2.635rem] justify-center items-center">
          ${e.notifications.map(o=>X(o)).join("")}
        </ul>
      `,K(i)})()};function X(t){return`
  <li id="${t.notificationId}" class="flex justify-between w-full items-start gap-[0.98rem] [transition:transform_0.8s_ease-in,opacity_0.6s_ease-in]">
    <img id="closeBtn-img"
      src="/public/assets/svg/closeBtn.svg"
      class="close-btn w-[1.324rem] aspect-square"
      _="on click
        halt the event
        tell the closest <li/>
          add .slide-out then settle then remove yourself
          then set :notifId to your id then log :notifId
          then set :accountId to #notification-bar.dataset.accountid then log :accountId
          then call deleteNotification(:accountId, :notifId)
          then if <li/> in #notification-bar is empty
            call updateEmptyNotifMarkup()
          end"
    />
    <div>
      <section class="w-full flex justify-between items-center mb-[0.441rem]">
        <p class="badge ${H()} text-[#FFFFFF] font-bold text-[0.511rem] leading-[0.623rem] text-center px-[0.929rem] py-[0.078rem] rounded-[0.256rem]">${t.notificationTitle}</p>
        <p class="text-[#3C3C4399] font-poppins font-extralight text-[0.477rem] leading-[0.715rem]">
          ${new Date(t.notificationTime).toLocaleString()}
        </p>
      </section>
      <div>
        <p class="text-[#000000] font-bold text-start text-[0.626rem] leading-[0.763rem] mb-[0.313rem]">
          ${t.notificationHeading}
        </p>
        <p class="text-[#000000] font-poppins font-extralight text-[0.626rem] text-start leading-[0.938rem] max-w-[50ch] break-words">
          ${t.notificationBody}
        </p>
      </div>
    </div>
  </li>
  `}(async()=>{try{let t=await fetch("/account/notifications",{headers:{Accept:"application/json"}});if(!t.ok){console.warn("response not ok",t.status);return}let n=t.headers.get("X-AccountId");if(!n){console.warn("Missing X-AccountId header");return}let e=await t.json();for(let i of e){if(!await W(n,i)){console.log("Notification already exists, skipping",i);continue}console.log("Adding new notification",i),await _(n,i)}}catch(t){console.error("Error fetching notifications:",t)}})();})();
//# sourceMappingURL=notification.js.map
