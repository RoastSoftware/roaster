import ξ from 'mithril';
import base from './base';
import {UserModel} from '../models/user';

export default {
  view(vnode: CVnode) {
    return ξ(base,
        ξ('.ui.main.text.container[style=margin-top: 2em;]',
            ξ('img.ui.image.rounded.medium',
                {src: 'http://c0419384.cdn2.cloudfiles.rackspacecloud.com/adjnbivz-27337_l-avatar-main.jpg'}, 'User profile picture.'),
            ξ('h2',
                UserModel.getFullname()),
            ξ('p',
                ξ('i.user.icon'),
                UserModel.getUsername()),
            ξ('p',
                ξ('i.mail.icon'),
                UserModel.getEmail())
        )
    );
  },
}as ξ.Component;
