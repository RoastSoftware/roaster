import ξ from 'mithril';
import base from './base';
import Chart from 'chart.js';
import Model from '../models/statistics';

export default class Statistics implements ξ.ClassComponent {
  view(vnode: ξ.CVnode) {
    return ξ(base,
        ξ('.ui.main.text.container[style=margin-top: 2em;]',
            ξ('p', 'Let there be GRAPHS! '),
            ξ('p', 'later...'),
            ξ('canvas#chart-area', {
              oncreate: ({dom}) => {
                const ctx = (document.getElementById(
                    'chart-area')as HTMLCanvasElement)
                    .getContext('2d');
                new Chart(ctx, Model.config);
              },
            }
            ),
        )
    );
  };
};
