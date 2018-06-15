import React, { Component } from 'react';
import { setWallpaper } from '@app/wallpaperize-proxy/set-wallpaper';

export class Image extends Component {
    state: {image: string};
    constructor(props: {image: string}) {
        super(props);
        this.state = {image: props.image};
    }

    onClick() {
        setWallpaper(this.state.image);
    }

    getSrc(path: string): string {
        return 'file://' + path;
    }

    render() {
        return (
        <div className="col-md-3 col-sm-6">
            <img onClick={(_) => this.onClick()} className="img-thumbnail" src={this.getSrc(this.state.image)} alt=""/>
        </div>
        );
    }
}
