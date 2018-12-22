import moment from 'moment';
import Network from '../services/network';

class StandardColorModel {
    static chartColors = {
      red: 'rgb(250, 50, 47)',
      orange: 'rgb(203, 75, 22)',
      yellow: 'rgb(181, 137, 0)',
      green: 'rgb(133, 153, 0)',
      blue: 'rgb(38, 139, 210)',
      magenta: 'rgb(211, 54, 130)',
      grey: 'rgb(131, 148, 150)',
      violet: 'rgb(108, 113, 196)',
      cyan: 'rgb(0, 181, 173)',
      darkblue: 'rgb(7, 54, 66)',
    };
};

export class RoastLinesStatisticsModel {
  data: Object = Object();

  setData(data: Object) {
    this.data = data;
  }

  update(): Promise {
    const interval = '30m';
    const end = moment();
    const start = moment().subtract(5, 'hours');
    const uri = `\
/statistics/roast/timeseries\
?start=${start.utc().format()}\
&end=${end.utc().format()}\
&interval=${interval}\
`;

    return Network.request<RoastLinesStatisticsModel>('GET', uri)
        .then((stats: Object) => {
          this.setData(stats.reverse());
        });
  };

  getErrors(): Array {
    return this.data.map((p) => {
      return p.numberOfErrors;
    });
  };

  getWarnings(): Array {
    return this.data.map((p) => {
      return p.numberOfWarnings;
    });
  };

  getLines(): Array {
    return this.data.map((p) => {
      return p.linesOfCode;
    });
  };

  getLabels(): Array {
    const labels = [];
    for (let i = 0; i < this.data.length; i++) {
      console.log(this.data[i]);
      labels.push(moment(this.data[i].timestamp).format('HH:mm'));
    }

    return labels;
  };

  getConfig(): Object {
    return {
      type: 'line',
      options: {
        responsive: true,
        title: {
          display: true,
          text: 'Number of Lines Analyzed vs. Errors and Warnings',
        },
        scales: {
          yAxes: [{
            gridLines: {
              color: StandardColorModel.chartColors.darkblue,
            },
          }],
          xAxes: [{
            gridLines: {
              color: StandardColorModel.chartColors.darkblue,
            },
          }],
        },
      },
      data: {
        labels: this.getLabels(),
        datasets: [{
          label: 'Errors',
          backgroundColor: StandardColorModel.chartColors.red,
          borderColor: StandardColorModel.chartColors.red,
          data: this.getErrors(),
          fill: false,
        }, {
          label: 'Warnings',
          backgroundColor: StandardColorModel.chartColors.yellow,
          borderColor: StandardColorModel.chartColors.yellow,
          data: this.getWarnings(),
          fill: false,
        }, {
          label: 'Lines Analyzed',
          backgroundColor: StandardColorModel.chartColors.darkblue,
          borderColor: StandardColorModel.chartColors.grey,
          data: this.getLines(),
          fill: true,
        }],
      },
    };
  };
};

export class RoastDoughnutStatisticsModel {
    static dataDonut = {
      datasets: [{
        borderColor: 'rgba(0, 0, 0, 0.0)',
        backgroundColor: [
          StandardColorModel.chartColors.yellow,
          StandardColorModel.chartColors.cyan,
          StandardColorModel.chartColors.green,
        ],
        data: [10, 20, 30],
      }],

      // These labels appear in the legend and in the
      // tooltips when hovering different arcs
      /* labels: [
            'ERR',
            'WARN',
            'Blue'
        ] */
    };
    static optionsDonut = {
      aspectRatio: 1,
      title: {
        display: true,
        text: 'ROAST SCORE',
        position: 'bottom',
        fontStyle: 'bold',
        fontSize: 16,
      },
    }

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
