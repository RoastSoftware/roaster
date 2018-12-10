import ξ from 'mithril';
import base from './base';

export default {
  view(vnode: CVnode) {
    return [
      ξ(base,
          ξ('.ui.main.text.container[style=margin-top: 2em;]',
              ξ('h1', `What is Roaster?`),
              ξ('p', `Roaster is a site that will run static code analysis on\
                  code you write or paste into the IDE on the home page.\
                  What this means is that it will check your code for \
                  syntactic errors, formatting that does not conform with the\
                  set code style etc. This information is then used to show\
                  you messages about where your code can improve and is also\
                  the foundation for calculating your Roast® Score™.`),
              ξ('h2', `Roast® Score™`),
              ξ('p', `The Roast® Score™ is your way of showing your coding \
                  skill and progress to yourself and your friends!
                  It's a quality metric based on a formula, considering your\
                  warnings (low impact on score), \
                  errors (high impact on score) and the amount of lines you've\
                  analysed. `),
              ξ('h2', `Getting started`),
              ξ('p', `You can of course analyze your code without registering\
              an account but then you will loose the benefit of progress\
              tracking. To analyze your first line, you will want to check \
                  that the correct version of python is selected to the bottom\
                  right of the right column in the editor view, otherwise you \
                  will get incorrect results.\
                  Next, click the "ROAST ME" button! Now, if you ran our\
                  prefilled example snippet you should see one warning and\
                  one error in the result section of the editor view.\
                 `),
              ξ('h3', ` Congratulations, you've made your first tiny steps to\
                  writing better code!\ We hope that you will find joy\
                  in using Roast® Inc.`),
              ξ('p', `// Roast® Inc. Team.`)

          ),
      ),
    ];
  },
} as ξ.Component;
