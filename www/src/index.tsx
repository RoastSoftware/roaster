import * as React from "react";
import * as ReactDOM from "react-dom";

import { Example } from "./components/Example";

ReactDOM.render(
    <Example compiler="TypeScript" framework="React" />,
    document.getElementById("example")
);
