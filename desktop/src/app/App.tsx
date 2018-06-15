import * as React from 'react';
import { ImagesSector, ImagesSectorProps } from '@app/app/images-sector/ImagesSector';
import { getInfo } from '@app/wallpaperize-proxy/get-info';
import { loadRandom } from '@app/wallpaperize-proxy/load-random';
import { getRandom } from '@app/wallpaperize-proxy/get-random';
import { getSelected } from '@app/wallpaperize-proxy/get-selected';

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
                selected: '',
                images: [],
                cachedImages: [],
                loadHandler: () => this._loadRandomHandler(),
                getSrc: this.getRemoteSrc
            },
            daily: {
                title: 'Daily Images',
                selected: '',
                images: [],
                cachedImages: [],
                loadHandler: () => new Promise(r => r()),
                getSrc: this.getSrc
            }
        };
        this.setImages();
    }

    async setImages(): Promise<void> {
        const { random_images, daily_images} = await getInfo();
        const random = await getRandom();
        const randomCached = random_images.map(i => i.split('/').pop());
        const selected = await getSelected();
        this.setState({
            daily: {
                ...this.state.daily,
                images: daily_images,
                cachedImages: daily_images
            },
            random: {
                ...this.state.random,
                images: random,
                cachedImages: randomCached,
                selected,
            },
        });
    }

    getSrc(path: string): string {
        return 'file://' + path;
    }

    getRemoteSrc(p: string): string {
        // return 'http://localhost:2015/' + p;
        return 'https://wallpaperize.goncharovnikita.com/i/random/' + p;
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
