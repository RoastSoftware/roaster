import ξ from 'mithril';
import base from './base';
import Chart from 'chart.js';
import Network from '../services/network';
import moment from 'moment';
import {RoastLinesStatisticsModel} from '../models/statistics';

class RoastLinesChart implements ξ.ClassComponent {
  chart: Chart;
  ctx: CanvasRenderingContext2D;
  statistics: RoastLinesStatisticsModel;

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

    this.ctx = (dom as HTMLCanvasElement).getContext('2d');
    this.chart = new Chart(this.ctx,
        this.statistics.getConfig());
  };

  view(vnode: ξ.CVnode) {
    return ξ('canvas');
  };
};

export default class Statistics implements ξ.ClassComponent {
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
            ξ(RoastLinesChart),
        ),
    );
  };
};
