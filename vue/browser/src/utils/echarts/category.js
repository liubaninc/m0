// @ts-nocheck
import * as echarts from "echarts";

/**
 * @param {*} dom
 * @param {*} data
 * @param {*} opts
 */
export function createCategoryCharts(
  dom,
  data = {
    xAxis: [],
    series: [],
  },
  opts
) {
  var myChart = echarts.init(dom);
  let options = {
    title: {
      // text: "7天区块高度",
    },
    tooltip: {
      show: true,
      trigger: "axis",
      axisPointer: {
        type: "cross",
      },
      backgroundColor: "rgba(0,0,0,0.5)",
      textStyle: {
        color: "#ffffff",
        ellipsis: "...",
      },
      borderWidth: 0,
    },
    grid: {
      x: "12%", //x 偏移量
      y: "12%", // y 偏移量
      width: "80%", // 宽度
      height: "76%", // 高度
    },
    xAxis: {
      type: "category",
      boundaryGap: false,
      data: data.xAxis,
    },
    yAxis: {
      type: "value",
      axisLine: {
        show: true,
      },
      splitLine: {
        show: false,
      },
      axisTick: {
        show: true,
      },
    },
    series: [
      {
        data: data.series,
        type: "line",
        areaStyle: {
          color: new echarts.graphic.LinearGradient(0, 0, 0, 1, [
            {
              offset: 0,
              color: "rgba(58,77,233,0.8)",
            },
            {
              offset: 1,
              color: "rgba(58,77,233,0.3)",
            },
          ]),
        },
      },
    ],
  };
  myChart.setOption(options);
}
