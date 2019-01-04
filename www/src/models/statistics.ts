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
  public count: number = 0;
  public filter: StatisticsFilter = StatisticsFilter.Global;
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
    const interval = '1h';
    const end = moment();
    const start = moment().subtract(24, 'hours');

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

class RoastRatioModel {
  public linesOfCode: number = 0;
  public numberOfErrors: number = 0;
  public numberOfWarnings: number = 0;
  public filter: StatisticsFilter = StatisticsFilter.Global;
  public username: string = '';

  protected constructor(filter: number, username?: string) {
    this.filter = filter;
    if (username) {
      this.username = username;
    }
  };

  public async update(): Promise {
    let uri = '/statistics/roast/ratio';

    switch (this.filter) {
      case StatisticsFilter.Friends:
        uri += `?user=${this.username}&friends=true`;
        break;
      case StatisticsFilter.User:
        uri += `?user=${this.username}`;
        break;
    }

    return Network.request<RoastRatioModel>('GET', uri)
        .then((ratio: RoastRatioModel) => {
          Object.assign(this, ratio);
          return;
        });
  };
}

export class RoastDoughnutStatisticsModel extends RoastRatioModel {
  constructor(filter: number, username?: string) {
    super(filter, username);
  };

  private getSuccessLines(): number {
    return this.linesOfCode - this.numberOfErrors - this.numberOfWarnings;
  };

  private getSuccessRatio(): number {
    return Math.round((1 - ((this.numberOfErrors + this.numberOfWarnings) /
      this.getSuccessLines())) * 100);
  }

  public getData(): Object {
    return {
      datasets: [{
        borderColor: 'rgba(0, 0, 0, 0.3)',
        backgroundColor: [
          colors.yellow,
          colors.red,
          colors.green,
        ],
        data: [
          this.numberOfWarnings,
          this.numberOfErrors,
          this.getSuccessLines(),
        ],
      }],

      labels: [
        'Warnings',
        'Errors',
        'Success',
      ],
    };
  };

  public getConfig(): Object {
    const obj = this;

    return {
      type: 'doughnut',
      plugins: [{
        beforeDraw: function(chart) {
          const width = chart.chart.width;
          const height = chart.chart.height;
          const ctx = chart.chart.ctx;

          ctx.restore();

          const fontSize = (height / 114).toFixed(2);

          ctx.font = fontSize + 'em sans-serif';
          ctx.textBaseline = 'middle';
          ctx.fillStyle = 'white';

          const text = obj.getSuccessRatio() + '%';
          const textX = Math.round((width - ctx.measureText(text).width) / 2);
          const textY = height / 2;

          ctx.fillText(text, textX, textY);
          ctx.save();
        },
      }],
      options: {
        aspectRatio: 1,
        title: {
          display: true,
          text: 'ROAST® RATIO™',
          position: 'bottom',
          fontStyle: 'bold',
          fontSize: 16,
        },
      },
      data: this.getData(),
    };
  };
};
