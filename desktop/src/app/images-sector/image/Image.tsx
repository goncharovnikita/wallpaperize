import React, { Component } from 'react';
import { setWallpaper } from '@app/wallpaperize-proxy/set-wallpaper';
import './Image.sass';

export interface ImageProps {
    image: string;
    cached: boolean;
    selected: boolean;
    getSrc: (s: string) => string;
}

export class Image extends Component<ImageProps, ImageProps> {
    constructor(props: ImageProps) {
        super(props);
        this.state = props;
    }

    async onClick() {
        await setWallpaper(this.state.getSrc(this.state.image));
        this.setState({...this.state, cached: true});
    }

    getCached = (): JSX.Element|undefined => {
        if (this.state.cached) {
            return <i className="position-absolute far fa-check-circle"></i>;
        }
    }

    getClassName(): string {
        if (!this.state.selected) {
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
                onClick={(_) => this.onClick()} src={this.state.getSrc(this.state.image)} alt=""/>
                {this.getCached()}
            </div>
        </div>
        );
    }
}
