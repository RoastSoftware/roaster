import ξ from 'mithril';
import base from './base';
import * as d3 from 'd3';

export default class Statistics implements ξ.ClassComponent {
  oncreate(vnode: ξ.CVnodeDOM) {
    d3.select('svg')
        .append('circle')
        .attr('r', 5)
        .attr('cx', '50%')
        .attr('cy', '50%')
        .attr('fill', 'red');
    // ξ.redraw();
  }

  view(vnode: ξ.CVnode) {
    return ξ(base,
        ξ('.ui.main.text.container[style=margin-top: 2em;]',
            ξ('p', 'Let there be GRAPHS! '),
            ξ('p', 'later...'),
            ξ('svg', {width: '100%', height: '100%'} )
        )
    );
  };
};
