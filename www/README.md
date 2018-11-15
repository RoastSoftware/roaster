./www
=====
This folder holds the frontend code, `roasterc`, for the roast.sofware website.

The file structure follows the structure for MVC-pattern together with Mithril and webpack. TypeScript is used instead of JavaScript. The `dist` folder holds the minified JavaScript and
the `src` folder holds the TypeScript code using the Mithril framework.
An example file structure:

```
.
├── dist
│   ├── bundle.js
│   └── bundle.js.map
├── index.html
├── package.json
├── package-lock.json
├── src
├── tsconfig.json
└── webpack.config.js
```

## Setting up the development environment
Make sure you have *nodejs* and *npm* installed on your system.

Install *webpack* globally on your system:
```
npm install -g webpack
```

Then run *npm install* to install all the dependencies inside this directory
(./www):
```
npm install
```

Finally, run *npm start* to compile the source upon file changes to the `dist` folder inside this directory (./www):
```
npm start
```

To make a production ready compilation run:
```
npm run build
```

Done!
