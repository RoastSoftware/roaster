import * as React from "react";

export interface ExampleProps { compiler: string; framework: string; }

// 'ExampleProps' describes the shape of props.
// State is never set so we use the '{}' type.
export class Example extends React.Component<ExampleProps, {}> {
    render() {
        return <h1>Example file compiled by {this.props.compiler} using the {this.props.framework} framework!</h1>;
    }
}
