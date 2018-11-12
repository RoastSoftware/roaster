./www
=====
This folder holds the frontend code, `roasterc`, for the roast.sofware website.

The file structure follows the standard React + Webpack structure. TypeScript is
used instead of JavaScript. The `dist` folder holds the minified JavaScript and
the `src` folder holds the TypeScript code using the React framework.
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
│   ├── components
│   │   └── Example.tsx
│   └── index.tsx
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

Finally, run *webpack* to compile the source to the `dist` folder inside this
directory (./www):
```
webpack
```

Done!
