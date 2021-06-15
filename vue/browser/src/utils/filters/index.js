/**
 * @param {*} size
 */
export let formateSize = (size) => {
  return (size && (size / 1024).toFixed(2) + " K") || "0 K";
};
/**
 * @param {*} status
 */
export let trxStatus = (status) => {
  return status ? "交易成功" : "交易失败";
};
export default {
  formateSize,
  trxStatus,
};
