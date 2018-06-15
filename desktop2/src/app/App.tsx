import * as React from 'react';
import { ImagesSector, ImagesSectorProps } from '@app/app/images-sector/ImagesSector';
import { getInfo } from '@app/wallpaperize-proxy/get-info';
import { loadRandom } from '@app/wallpaperize-proxy/load-random';

interface AppState {
    random: ImagesSectorProps;
    daily: ImagesSectorProps;
}

export class App extends React.Component {
    state: AppState;
    constructor(props: any) {
        super(props);
        this.state = {
            random: {
                title: 'Random Images',
                images: [],
                loadHandler: () => this._loadRandomHandler()
            },
            daily: {
                title: 'Daily Images',
                images: [],
                loadHandler: () => new Promise(r => r())
            }
        };
        this.setImages();
    }

    async setImages(): Promise<void> {
        const {random_images, daily_images} = await getInfo();
        console.log(random_images.length)
        this.setState({
            daily: {...this.state.daily, images: daily_images},
            random: {...this.state.random, images: random_images},
        });
    }

    render() {
        return (
            <div>
                <ImagesSector {...this.state.daily} />
                <ImagesSector {...this.state.random} />
            </div>
        );
    }

    private async _loadRandomHandler(): Promise<void> {
        await loadRandom();
        await this.setImages();
    }
}
