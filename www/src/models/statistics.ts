export default class StatisticsModel {
    static chartColors = {
      red: 'rgb(250, 20, 47)',
      orange: 'rgb(203, 75, 22)',
      yellow: 'rgb(181, 137, 0)',
      green: 'rgb(133, 153, 0)',
      blue: 'rgb(38, 139, 210)',
      magenta: 'rgb(211, 54, 130)',
      grey: 'rgb(201, 203, 207)',
      violet: 'rgb(108, 113, 196)',
      cyan: 'rgb(42, 161, 152)',
    };

    static months = ['January', 'February', 'March', 'April', 'May', 'June',
      'July', 'August', 'September', 'October', 'November', 'December'];
    static config = {
      type: 'line',
      data: {
        labels: ['January', 'February', 'March', 'April',
          'May', 'June', 'July'],
        datasets: [{
          label: 'My first dataset!',
          backgroundColor: StatisticsModel.chartColors.magenta,
          borderColor: StatisticsModel.chartColors.magenta,
          data: [
            100,
            10,
            70,
            -10,
            -90,
            30,
            5,
          ],
          fill: false,
        }, {
          label: 'My first dataset!',
          backgroundColor: StatisticsModel.chartColors.blue,
          borderColor: StatisticsModel.chartColors.blue,
          data: [
            100,
            50,
            50,
            10,
            -70,
            -100,
            45,
          ],
          fill: true,
        }],
      },
    };


    static options = {
      responsive: true,
      title: {
        display: true,
        text: 'Chart.js Line Chart',
      },
      tooltips: {
        mode: 'index',
        intersect: false,
      },
      hover: {
        mode: 'nearest',
        intersect: true,
      },
      scales: {
        xAxes: [{
          display: true,
          scaleLabel: {
            display: true,
            labelString: 'Month',
          },
        }],
        yAxes: [{
          display: true,
          scaleLabel: {
            display: true,
            labelString: 'Value',
          },
        }],
      },
    };
};
