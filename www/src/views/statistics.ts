import ξ from 'mithril';
import base from './base';
import Chart from 'chart.js';
import {
  RoastLinesStatisticsModel,
  RoastCountStatisticsModel,
  RoastCountModel,
  LineCountModel,
} from '../models/statistics';

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

class RoastCountChart implements ξ.ClassComponent {
  chart: Chart;
  ctx: CanvasRenderingContext2D;
  statistics: RoastCountStatisticsModel;

  oncreate({dom}) {
    this.statistics = new RoastCountStatisticsModel();

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

class RoastCount implements ξ.ClassComponent {
  roasts: RoastCountModel = new RoastCountModel();
  intervalID: number;

  oncreate({dom}) {
    this.roasts.update();

    this.intervalID = setInterval(() => {
      this.roasts.update();
    }, 5000);
  };

  onremove() {
    clearInterval(this.intervalID);
  };

  view(vnode: ξ.CVnode) {
    return ξ('h1.ui.icon.blue.header',
        ξ('.content[style=font-size: 3em;]', this.roasts.count,
            ξ('.sub.header[style=margin-top: 1em;]',
                'ROASTED SUBMISSIONS')));
  };
};

class LineCount implements ξ.ClassComponent {
  lines: LineCountModel = new LineCountModel();
  intervalID: number;

  oncreate({dom}) {
    this.lines.update();

    this.intervalID = setInterval(() => {
      this.lines.update();
    }, 5000);
  };

  onremove() {
    clearInterval(this.intervalID);
  };

  view(vnode: ξ.CVnode) {
    return ξ('h1.ui.icon.blue.header',
        ξ('.content[style=font-size: 3em;]',
            this.lines.count,
            ξ('.sub.header[style=margin-top: 1em;]',
                'TOTAL LINES OF CODE')));
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
        ξ('.ui.stackable.padded.grid',
            ξ('.row.center.aligned',
                ξ('.sixteen.wide.column',
                    ξ('.ui.compact.secondary.menu',
                        ξ('a.item.active',
                            ξ('i.globe.icon'), 'GLOBAL'),
                        ξ('a.item',
                            ξ('i.users.icon'), 'FRIENDS'),
                        ξ('a.item',
                            ξ('i.user.icon'), 'YOU'),
                    ),
                ),
            ),
            ξ('.row.center.aligned',
                ξ('.eight.wide.column',
                    ξ(LineCount),
                ),
                ξ('.eight.wide.column',
                    ξ(RoastCount),
                ),
            ),
            ξ('.row',
                ξ('.eight.wide.column',
                    ξ(RoastLinesChart),
                ),
                ξ('.eight.wide.column',
                    ξ(RoastCountChart),
                ),
            ),
        ),
    );
  };
};
