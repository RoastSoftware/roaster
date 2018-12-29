import ξ from 'mithril';
import base from './base';
import Chart from 'chart.js';
import {UserModel} from '../models/user';
import {
  RoastLinesStatisticsModel,
  RoastCountStatisticsModel,
  RoastCountModel,
  LineCountModel,
  StatisticsFilter,
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

  onupdate({attrs}) {
    if (this.statistics.filter != attrs.filter) {
      this.statistics.filter = attrs.filter;
      this.statistics.update().then(() => {
        this.chart.data = this.statistics.getData();
        this.chart.update();
      });
    }
  };

  view() {
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

  onupdate({attrs}) {
    if (this.statistics.filter != attrs.filter) {
      this.statistics.filter = attrs.filter;
      this.statistics.update().then(() => {
        this.chart.data = this.statistics.getData();
        this.chart.update();
      });
    }
  };

  view() {
    return ξ('canvas');
  };
};

class RoastCount implements ξ.ClassComponent {
  roasts: RoastCountModel;
  intervalID: number;

  oninit({attrs}) {
    this.roasts = new RoastCountModel(attrs.filter);
  };

  onupdate({attrs}) {
    if (this.roasts.filter != attrs.filter) {
      this.roasts.filter = attrs.filter;
      this.roasts.update();
    }
  };

  oncreate({dom}) {
    this.roasts.update();
    this.intervalID = setInterval(() => {
      this.roasts.update();
    }, 5000);
  };

  onremove() {
    clearInterval(this.intervalID);
  };

  view() {
    return ξ('h1.ui.icon.blue.header',
        ξ('.content[style=font-size: 3em;]', this.roasts.count,
            ξ('.sub.header[style=margin-top: 1em;]',
                'ROASTED SUBMISSIONS')));
  };
};

class LineCount implements ξ.ClassComponent {
  lines: LineCountModel;
  intervalID: number;

  oninit({attrs}) {
    this.lines = new LineCountModel(attrs.filter);
  };

  onupdate({attrs}) {
    if (this.lines.filter != attrs.filter) {
      this.lines.filter = attrs.filter;
      this.lines.update();
    }
  };

  oncreate({dom}) {
    this.lines.update();
    this.intervalID = setInterval(() => {
      this.lines.update();
    }, 5000);
  };

  onremove() {
    clearInterval(this.intervalID);
  };

  view() {
    return ξ('h1.ui.icon.blue.header',
        ξ('.content[style=font-size: 3em;]',
            this.lines.count,
            ξ('.sub.header[style=margin-top: 1em;]',
                'TOTAL LINES OF CODE')));
  };
};

export default class Statistics implements ξ.ClassComponent {
  currentFilter: StatisticsFilter = StatisticsFilter.Global;

  setFilter(filter: StatisticsFilter) {
    this.currentFilter = filter;
  };

  setActive(filter: StatisticsFilter) {
    return filter == this.currentFilter ? 'active' : '';
  }

  view() {
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
                        UserModel.isLoggedIn() ? [
                          ξ('a.item', {
                            class: this.setActive(StatisticsFilter.Global),
                            onclick: () => {
                              this.setFilter(StatisticsFilter.Global);
                            },
                          },
                          ξ('i.globe.icon'), 'GLOBAL'),
                          ξ('a.item', {
                            class: this.setActive(StatisticsFilter.Friends),
                            onclick: () => {
                              this.setFilter(StatisticsFilter.Friends);
                            },
                          },
                          ξ('i.users.icon'), 'FRIENDS'),
                          ξ('a.item', {
                            class: this.setActive(StatisticsFilter.User),
                            onclick: () => {
                              this.setFilter(StatisticsFilter.User);
                            },
                          },
                          ξ('i.user.icon'), 'YOU'),
                        ]: '',
                    ),
                ),
            ),
            ξ('.row.center.aligned',
                ξ('.eight.wide.column',
                    ξ(LineCount, {filter: this.currentFilter}),
                ),
                ξ('.eight.wide.column',
                    ξ(RoastCount, {filter: this.currentFilter}),
                ),
            ),
            ξ('.row',
                ξ('.eight.wide.column',
                    ξ(RoastLinesChart, {filter: this.currentFilter}),
                ),
                ξ('.eight.wide.column',
                    ξ(RoastCountChart, {filter: this.currentFilter}),
                ),
            ),
        ),
    );
  };
};
