import * as React from 'react';
import { Image } from '@app/images-sector/image/Image';
import { setWallpaper } from '@app/wallpaperize-proxy/set-wallpaper';

export interface ImagesSectorProps {
    title: string;
    selected: string;
    images: string[];
    cachedImages: string[];
    afterTitle?: JSX.Element;
    loadHandler: () => Promise<void>;
    getSrc: (s: string) => string;
}

export class ImagesSector extends React.Component<ImagesSectorProps, ImagesSectorProps> {
    constructor(public props: ImagesSectorProps) {
        super(props);
        this.state = props;
    }

    shouldComponentUpdate(nextProps: ImagesSectorProps, nextState: ImagesSectorProps): boolean {
        const shouldUpdate = nextProps.images.some(i => this.state.images.indexOf(i) === -1);
        const selectedChanged = this.state.selected !== nextState.selected;
        if (shouldUpdate) {
            this.setState(nextProps);
        }
        const result = shouldUpdate || selectedChanged;
        return result;
    }

    mapImages(imgs: string[]): JSX.Element[] {
        return imgs.map((img) => {
            const cached = this.state.cachedImages.some(v => v === img);
            const selected = img === this.state.selected;
            return (
                <Image key={img} {
                    ...{image: img, getSrc: this.state.getSrc, cached, selected, onclick: this.onImgClick(img)}} />
            );
        });
    }

    async onBtnClick(): Promise<void> {
        await this.state.loadHandler();
    }

    render() {
        return (
            <div>
                <h3 className="display-3">
                {this.state.title}{this.state.afterTitle}
                </h3>
                <hr/>
                <div className="row">
                    {this.mapImages(this.state.images)}
                </div>
            </div>
        );
    }

    onImgClick = (img: string): () => void => {
        return () => {
            this.setState((prevState: ImagesSectorProps) => {
                return {
                    selected: img,
                    cachedImages: [...prevState.cachedImages, img]
                };
            });
            setWallpaper(this.state.getSrc(img));
        };
    }
}
