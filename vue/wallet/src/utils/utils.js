// @ts-nocheck

export let localCache = {
  set (key, value) {
    localStorage.setItem(key, JSON.stringify(value));
  },
  get (key) {
    var val = localStorage.getItem(key);
    var dataobj = JSON.parse(val);
    return dataobj;
  },
  remove (key) {
    return localStorage.removeItem(key);
  },
};

/**
 * 保留四位小数
 * @param {*} num
 */
export function toFixed4 (num) {
  return Number(num.toString().match(/^\d+(?:\.\d{0,4})?/));
}
/**
 * 文件下载
 * @param {*} params 
 * @param {*} download 
 */
export function downLoadFile (path, download) {
  if (!path) {
    return;
  }
  let origin = window.location.origin;
  let elink = document.createElement("a");
  elink.download = download || new Date().getTime();
  elink.style.display = "none";
  // elink.href = `${origin}/api/claims/download/${loginUser.name}/${wallet.name}/${eviDetail.name}`;
  // elink.href = `${origin}/api/download/${params}`;
  elink.href = path;
  document.body.appendChild(elink);
  elink.click();
  document.body.removeChild(elink);
}

/**
 * 版本号比较
 * @param {*} v1 
 * @param {*} v2 
 * @returns 
 */
export function compareVersion (v1, v2) {
  if (v1 == v2) {
    return 0;
  }

  const vs1 = v1.split(".").map(a => parseInt(a));
  const vs2 = v2.split(".").map(a => parseInt(a));

  const length = Math.min(vs1.length, vs2.length);
  for (let i = 0; i < length; i++) {
    if (vs1[i] > vs2[i]) {
      return 1;
    } else if (vs1[i] < vs2[i]) {
      return -1;
    }
  }

  if (length == vs1.length) {
    return -1;
  } else {
    return 1;
  }
}