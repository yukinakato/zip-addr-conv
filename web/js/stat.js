"use strict";

import $ from "jquery";
import Chart from "chart.js";

$(function () {
  $.get({
    url: "/getstat",
    dataType: "json",
    success: function (data) {
      $(".since").text(`${data.since} total ${data.total}`);
      let labels = [];
      for (let i = 0; i < 101; i++) {
        labels.push(i * 10);
      }
      let ctx = document.getElementById("stat").getContext("2d");
      new Chart(ctx, {
        type: "bar",
        data: {
          labels: labels,
          datasets: [
            {
              label: "Required time (ms)",
              data: data.time,
              backgroundColor: "rgba(31, 191, 255, 0.9)",
            },
          ],
        },
        options: {
          aspectRatio: 1.25,
          scales: {
            xAxes: [
              {
                ticks: {
                  maxTicksLimit: 20,
                },
                gridLines: {
                  offsetGridLines: false,
                },
              },
            ],
            yAxes: [
              {
                ticks: {
                  stepSize: 10,
                },
                gridLines: {
                  drawBorder: false,
                },
              },
            ],
          },
        },
      });
    },
    error: function () {
      console.log("getting stat failed");
    },
  });
});
