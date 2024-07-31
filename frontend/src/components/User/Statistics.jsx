import { useEffect } from "react";
import "../../css/User.css";
import { Chart } from "chart.js/auto";

function Statistics() {
  useEffect(() => {
    const ctx = document.getElementById("myChart").getContext("2d");
    ctx.canvas.width = 400;
    ctx.canvas.height = 400;
    const myChart = new Chart(ctx, {
      type: "doughnut",
      data: {
        labels: ["Completed", "Not Completed"],
        datasets: [
          {
            label: "% Completion",
            data: [90, 10],
            backgroundColor: ["rgb(54, 162, 235)", "rgb(255, 99, 132)"],
            hoverOffset: 4,
          },
        ],
      },
      options: {
        plugins: {
          legend: {
            position: "right",
            align: "center",
          },
        },
        responsive: false,
        maintainAspectRatio: false,
      },
    });

    return () => {
      myChart.destroy();
    };
  }, []);

  return (
    <div className="stats-main">
      <div className="stats-upper-level">
        <div className="completion-chart-container">
          <p className="stats-header">Game completion percentage</p>
          <canvas id="myChart" className="completion-pie"></canvas>
        </div>
        <div className="completion-chart-container">
          <p className="stats-header">Game completion percentage</p>
          <canvas id="myChart" className="completion-pie"></canvas>
        </div>
      </div>
    </div>
  );
}

export default Statistics;
