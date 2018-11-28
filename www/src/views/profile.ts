import ξ from 'mithril';
import base from './base';

export default {
  view(vnode: CVnode) {
    return ξ(base,
        ξ('.ui.main.text.container[style=margin-top: 2em;]',
            ξ('img.ui.image.rounded.medium',
                {src: 'http://c0419384.cdn2.cloudfiles.rackspacecloud.com/adjnbivz-27337_l-avatar-main.jpg'}, 'User profile picture.'),
            ξ('h2',
                'Mr. Bean-A-Tar'),
            ξ('h3',
                ξ('i.user.icon'),
                'mr-bean-a-tar'),
            ξ('p',
                ξ('i.mail.icon'),
                'mr@bean-a-tar.example.org')
        )
    );
  },
}as ξ.Component;
