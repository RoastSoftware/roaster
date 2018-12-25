import moment from 'moment';
import Network from '../services/network';
import {UserModel} from './user';

export enum StatisticsFilter {
  Global = 1, // eslint-disable-line no-unused-vars
  Friends, // eslint-disable-line no-unused-vars
  User, // eslint-disable-line no-unused-vars
};

const colors = {
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

class CountModel {
  public count: number;
  public filter: StatisticsFilter;
}

export class LineCountModel extends CountModel {
  public async update(): Promise {
    let uri = '/statistics/roast/lines';

    switch (this.filter) {
      case StatisticsFilter.Friends:
        uri += `?user=${UserModel.getUsername()}&friends=true`;
        break;
      case StatisticsFilter.User:
        uri += `?user=${UserModel.getUsername()}`;
        break;
    }

    return Network.request<Object>('GET', uri)
        .then((data: Object) => {
          this.count = data.lines;
          return;
        });
  };
}

export class RoastCountModel extends CountModel {
  public async update(): Promise {
    let uri = '/statistics/roast/count';

    switch (this.filter) {
      case StatisticsFilter.Friends:
        uri += `?user=${UserModel.getUsername()}&friends=true`;
        break;
      case StatisticsFilter.User:
        uri += `?user=${UserModel.getUsername()}`;
        break;
    }

    return Network.request<Object>('GET', uri)
        .then((data: Object) => {
          this.count = data.count;
          return;
        });
  };
}

class RoastMessageStatisticsModel {
  private data: Object = Object();

  public filter: StatisticsFilter = StatisticsFilter.Global;

  public update(): Promise {
    const interval = '30m';
    const end = moment();
    const start = moment().subtract(5, 'hours');

    let uri = `\
/statistics/roast/timeseries\
?start=${start.utc().format()}\
&end=${end.utc().format()}\
&interval=${interval}\
`;

    switch (this.filter) {
      case StatisticsFilter.Friends:
        uri += `&friends=true`;
      case StatisticsFilter.User:
        uri += `&user=${UserModel.getUsername()}`;
    }

    return Network.request<Object>('GET', uri)
        .then((stats: Object) => {
          this.data = stats.reverse();
        });
  };

  private getErrors(): Array {
    return this.data.map((p) => {
      return p.numberOfErrors;
    });
  };

  private getWarnings(): Array {
    return this.data.map((p) => {
      return p.numberOfWarnings;
    });
  };

  private getLabels(): Array {
    const labels = [];
    for (let i = 0; i < this.data.length; i++) {
      labels.push(moment(this.data[i].timestamp).format('HH:mm'));
    }

    return labels;
  };
}

export class RoastCountStatisticsModel extends RoastMessageStatisticsModel {
  private getCount(): Array {
    return this.data.map((p) => {
      return p.count;
    });
  };

  public getData(): Object {
    return {
      labels: this.getLabels(),
      datasets: [{
        label: 'Errors',
        yAxisID: 'messages',
        backgroundColor: colors.red,
        borderColor: colors.red,
        data: this.getErrors(),
        fill: false,
      }, {
        label: 'Warnings',
        yAxisID: 'messages',
        backgroundColor: colors.yellow,
        borderColor: colors.yellow,
        data: this.getWarnings(),
        fill: false,
      }, {
        label: 'Number of Roasts',
        yAxisID: 'count',
        backgroundColor: colors.darkblue,
        borderColor: colors.grey,
        data: this.getCount(),
        fill: true,
      }],
    };
  };

  public getConfig(): Object {
    return {
      type: 'line',
      options: {
        responsive: true,
        title: {
          display: true,
          text: 'Number of Roasts vs. Errors and Warnings',
        },
        scales: {
          yAxes: [{
            id: 'count',
            position: 'left',
            scaleLabel: {
              labelString: 'Number of Roasts',
              display: true,
            },
            gridLines: {
              color: colors.grey,
            },
            ticks: {
              beginAtZero: true,
              // Only allow integers for ticker.
              callback: (v) => {
                if (Number.isInteger(v)) {
                  return v;
                }
              },
            },
          }, {
            id: 'messages',
            position: 'right',
            scaleLabel: {
              labelString: 'Errors and Warnings',
              display: true,
            },
            gridLines: {
              color: colors.darkblue,
            },
            ticks: {
              beginAtZero: true,
              // Only allow integers for ticker.
              callback: (v) => {
                if (Number.isInteger(v)) {
                  return v;
                }
              },
            },
          }],
          xAxes: [{
            scaleLabel: {
              labelString: 'Time',
              display: true,
            },
            gridLines: {
              color: colors.darkblue,
            },
          }],
        },
      },
      data: this.getData(),
    };
  };
};


export class RoastLinesStatisticsModel extends RoastMessageStatisticsModel {
  private getLines(): Array {
    return this.data.map((p) => {
      return p.linesOfCode;
    });
  };

  public getData(): Object {
    return {
      labels: this.getLabels(),
      datasets: [{
        label: 'Errors',
        backgroundColor: colors.red,
        borderColor: colors.red,
        data: this.getErrors(),
        fill: false,
      }, {
        label: 'Warnings',
        backgroundColor: colors.yellow,
        borderColor: colors.yellow,
        data: this.getWarnings(),
        fill: false,
      }, {
        label: 'Lines Analyzed',
        backgroundColor: colors.darkblue,
        borderColor: colors.grey,
        data: this.getLines(),
        fill: true,
      }],
    };
  };

  public getConfig(): Object {
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
            scaleLabel: {
              labelString: 'n',
              display: true,
            },
            gridLines: {
              color: colors.darkblue,
            },
          }],
          xAxes: [{
            scaleLabel: {
              labelString: 'Time',
              display: true,
            },
            gridLines: {
              color: colors.darkblue,
            },
          }],
        },
      },
      data: this.getData(),
    };
  };
};

export class RoastDoughnutStatisticsModel {
    static dataDonut = {
      datasets: [{
        borderColor: 'rgba(0, 0, 0, 0.0)',
        backgroundColor: [
          colors.yellow,
          colors.cyan,
          colors.green,
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
