import ξ from 'mithril';
import base from './base';
import Chart from 'chart.js';
import Network from '../services/network';
import moment from 'moment';
import {RoastLinesStatisticsModel} from '../models/statistics';

export default class Statistics implements ξ.ClassComponent {
  linesStatistics: RoastLinesStatisticsModel = new RoastLinesStatisticsModel();

  updateStatistics() {
    const interval = '10m';
    const end = moment();
    const start = moment().subtract(1, 'days');
    const uri = `\
/statistics/roast/timeseries\
?start=${start.utc().format()}\
&end=${end.utc().format()}\
&interval=${interval}\
`;

    Network.request<RoastLinesStatisticsModel>('GET', uri)
        .then((stats: Object) => {
          this.linesStatistics.setData(stats);
        });
  };

  oncreate({dom}) {
    this.updateStatistics();
  };

  view(vnode: ξ.CVnode) {
    return ξ(base,
        ξ('.ui.main.text.container[style=margin-top: 1em;]',
            ξ('h1.ui.header',
                ξ('i.chart.bar.icon'),
                ξ('.content', 'STATISTICS',
                    ξ('.sub.header', 'All em\' statistics we\'ve collected.')),
            ),
            ξ('.ui.divider')),
        ξ('.ui.main.text.container',
            ξ('canvas#chart-area', {
              oncreate: ({dom}) => {
                const ctx = (
                  document.getElementById('chart-area') as HTMLCanvasElement)
                    .getContext('2d');

                new Chart(ctx, this.linesStatistics.getConfig());
              },
            }
            ),
        )
    );
  };
};
