import * as React from 'react';

export class Slide extends React.Component {
    render() {
        return (
            <div className="h-100 w-100 d-flex flex-column justify-content-around">
            {this.props.children}
            </div>
        );
    }
}
