import * as React from 'react';
import { ImagesSector, ImagesSectorProps } from '@app/app/images-sector/ImagesSector';
import { getInfo } from '@app/wallpaperize-proxy/get-info';
import { loadRandom } from '@app/wallpaperize-proxy/load-random';
import { getRandom } from '@app/wallpaperize-proxy/get-random';
import { getSelected } from '@app/wallpaperize-proxy/get-selected';
import { AppService } from '@app/app/app.service';
import * as ospath from 'path';
import { Menu } from '@app/menu/Menu';
import './App.sass';

interface AppState {
    random: ImagesSectorProps;
    daily: ImagesSectorProps;
}

export class App extends React.Component {
    randomPath = 'https://wallpaperize.goncharovnikita.com/i/random/';
    state: AppState;
    service: AppService;
    constructor(props: any) {
        super(props);
        this.state = this.initState();
        this.setImages();
        this.service = new AppService();
    }

    setImages = async (): Promise<void> => {
        const { random_images, daily_images} = await getInfo();
        const random = await getRandom();
        const randomCached = random_images.map(i => i.split('/').pop());
        const selected = await getSelected();
        this.setState({
            daily: {
                images: daily_images,
                cachedImages: daily_images
            },
            random: {
                images: random,
                cachedImages: randomCached,
                selected,
            },
        });
    }

    getSrc = (path: string): string => {
        console.log('file://' + ospath.resolve(path));
        return 'file://' + ospath.resolve(path);
    }

    getRemoteSrc = (p: string): string => {
        // return 'http://localhost:2015/' + p;
        return this.randomPath + 'min/' + p.replace('.jpg', '-min.jpg');
    }

    render() {
        return (
            <div className="d-flex h-100">
                <Menu />
                <div className="main-sector container">
                    <ImagesSector {...this.state.daily} />
                    <ImagesSector {...this.state.random} />
                </div>
            </div>
        );
    }

    afterTitle = <div className="lead d-flex align-items-end">
        {/* <i onClick={this.setImages} className="fas fa-xs ml-2 fa-sync"></i> */}
        <small onClick={this.setImages} className="text-muted pointer">refresh</small>
    </div>;

    private async _loadRandomHandler(): Promise<void> {
        await loadRandom();
        await this.setImages();
    }

    private initState(): AppState {
        return {
            random: {
                title: 'Random Images',
                selected: '',
                images: [],
                cachedImages: [],
                afterTitle: this.afterTitle,
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
    }
}
