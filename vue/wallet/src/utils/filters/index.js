// @ts-nocheck
/**
 * @param {*} size
 */
export let formateSize = (size) => {
  return (size && (size / 1024).toFixed(2) + " K") || "0 K";
};
/**
 * @param {*} height
 * @returns
 */
export let trxStatus = (height) => {
  return height < 0 ? "未提交" : height == 0 ? "待确认" : "已确认";
};
/**
 * @param {*} fmt
 * @param {*} date
 * @returns
 */
export let dateFormat = (date, fmt) => {
  let ret;
  if (!date) return date;
  if (!(date instanceof Date)) {
    date = new Date(date);
  }
  const opt = {
    "Y+": date.getFullYear().toString(), // 年
    "m+": (date.getMonth() + 1).toString(), // 月
    "d+": date.getDate().toString(), // 日
    "H+": date.getHours().toString(), // 时
    "M+": date.getMinutes().toString(), // 分
    "S+": date.getSeconds().toString(), // 秒
  };
  for (let k in opt) {
    ret = new RegExp("(" + k + ")").exec(fmt);
    if (ret) {
      fmt = fmt.replace(
        ret[1],
        ret[1].length == 1 ? opt[k] : opt[k].padStart(ret[1].length, "0")
      );
    }
  }
  return fmt;
};

export let walType = (type) => {
  return type > 0 ? "多签" : "单签";
};

export let formateAmount = (amount) => {
  if (!amount) return amount;
  let amt = /^\d*/.exec(amount)[0];
  let names = amount.replace(/^\d*/, "$1").split("$1");
  return amt + (names.length ? " (" + names[1] + ")" : "");
};

export default {
  formateSize,
  trxStatus,
  dateFormat,
  walType,
  formateAmount,
};
