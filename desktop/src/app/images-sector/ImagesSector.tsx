import * as React from 'react';
import { Image } from '@app/app/images-sector/image/Image';

export interface ImagesSectorProps {
    title: string;
    selected: string;
    images: string[];
    cachedImages: string[];
    loadHandler: () => Promise<void>;
    getSrc: (s: string) => string;
}

export class ImagesSector extends React.Component<ImagesSectorProps, ImagesSectorProps> {
    constructor(props: ImagesSectorProps) {
        super(props);
        this.state = props;
    }

    shouldComponentUpdate(nextProps: ImagesSectorProps, _: any): boolean {
        const shouldUpdate = nextProps.images !== this.state.images;
        if (shouldUpdate) {
            this.setState(nextProps);
        }
        return shouldUpdate;
    }

    getImages(): JSX.Element[] {
        return this.state.images.map((img, index) => {
            const cached = this.state.cachedImages.some(v => v === img);
            const selected = img === this.state.selected;
            return <Image key={index} {...{image: img, getSrc: this.state.getSrc, cached, selected}} />;
        });
    }

    async onBtnClick(): Promise<void> {
        await this.state.loadHandler();
    }

    render() {
        return (
            <div>
                <h3 className="display-3">
                {this.state.title}
                </h3>
                <hr/>
                <div className="row">
                    {this.getImages()}
                </div>
            </div>
        );
    }
}
