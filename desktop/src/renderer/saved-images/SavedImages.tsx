import * as React from 'react';
import { ImagesSectorProps, ImagesSector } from '@app/images-sector/ImagesSector';
import * as ospath from 'path';
import { getInfo } from '@app/wallpaperize-proxy/get-info';

interface State {
    images: ImagesSectorProps;
}

export class SavedImages extends React.Component<{}, State> {

    constructor(props: any) {
        super(props);
        this.state = this.initState();
        this.loadImages();
    }

    render(): JSX.Element {
        return (
            <div className="container main-sector">
                <ImagesSector {...this.state.images} />
            </div>
        );
    }

    getSrc = (path: string): string => {
        return 'file://' + ospath.resolve(path);
    }

    loadImages = async () => {
        const { random_images } = await getInfo();
        this.setState((prevState) => {
            return {
                images: {
                    ...prevState.images,
                    images: random_images
                }
            };
        });
    }

    private initState(): State {
        return {
            images: {
                title: 'Saved Images',
                selected: '',
                cachedImages: [],
                getSrc: this.getSrc,
                images: [],
                loadHandler: () => Promise.resolve()
            }
        };
    }
}
