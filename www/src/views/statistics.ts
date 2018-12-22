import ξ from 'mithril';
import base from './base';
import Chart from 'chart.js';
import {RoastLinesStatisticsModel} from '../models/statistics';

class RoastLinesChart implements ξ.ClassComponent {
  chart: Chart;
  ctx: CanvasRenderingContext2D;
  statistics: RoastLinesStatisticsModel;

  oncreate({dom}) {
    this.statistics = new RoastLinesStatisticsModel();

    this.statistics.update().then(() => {
      this.ctx = (dom as HTMLCanvasElement).getContext('2d');
      this.chart = new Chart(this.ctx,
          this.statistics.getConfig());
    });
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
