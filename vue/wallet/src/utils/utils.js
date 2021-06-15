// @ts-nocheck

export let localCache = {
  set(key, value) {
    localStorage.setItem(key, JSON.stringify(value));
  },
  get(key) {
    var val = localStorage.getItem(key);
    var dataobj = JSON.parse(val);
    return dataobj;
  },
  remove(key) {
    return localStorage.removeItem(key);
  },
};

/**
 * 保留四位小数
 * @param {*} num
 */
export function toFixed4(num) {
  return Number(num.toString().match(/^\d+(?:\.\d{0,4})?/));
}
