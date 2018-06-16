import React, { Component } from 'react';
import './Image.sass';

export interface ImageProps {
    image: string;
    cached: boolean;
    selected: boolean;
    onclick: () => void;
    getSrc: (s: string) => string;
}

export class Image extends Component<ImageProps, ImageProps> {
    constructor(props: ImageProps) {
        super(props);
        this.props = props;
    }

    getCached = (): JSX.Element|undefined => {
        if (this.props.cached) {
            return <div className="lead"><i className="position-absolute far fa-check-circle"></i></div>;
        }

        return undefined;
    }

    getClassName(): string {
        if (!this.props.selected) {
            return 'img-thumbnail pointer';
        } else {
            return 'img-thumbnail pointer bg-primary';
        }
    }

    render() {
        return (
        <div className="col-md-3 col-sm-6 col-6 mb-2">
            <div className="position-relative">
                <img className={this.getClassName()}
                onClick={this.props.onclick} src={this.props.getSrc(this.props.image)} alt=""/>
                {this.getCached()}
            </div>
        </div>
        );
    }
}
