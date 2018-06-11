import { Observable, ReplaySubject, from } from "rxjs";
import { Platform } from "../platform/platform";
import * as Axios from 'axios';

class App {
    private _baseURL = '';
    private _selectedPlatform: ReplaySubject<Platform>;
    private _selectedVersion: ReplaySubject<string>;

    constructor(env: string) {
        if (env === 'production') {
            this._baseURL = 'https://wallpaperize.goncharovnikita.com/api';
        } else {
            this._baseURL = 'http://localhost:8080/';
        }
        this._selectedPlatform = this._getSelectedPlatform();
        this._selectedVersion = this._getSelectedVersion();
    }

    getBaseURL(): string {
        return this._baseURL;
    }

    getSelectedPlatform(): Observable<Platform> {
        return this._selectedPlatform;
    }

    selectPlatform(p: Platform): void {
        this._selectedPlatform.next(p);
    }

    getSelectedVersion(): Observable<string> {
        return this._selectedVersion;
    }

    private  _getSelectedPlatform(): ReplaySubject<Platform> {
        const result = new ReplaySubject<Platform>();
        result.next(Platform.Mac);
        return result;
    }

    private _getSelectedVersion(): ReplaySubject<string> {
        const result = new ReplaySubject<string>();

        from(
            Axios.default.get(this._baseURL + 'get/maxversion')
        ).subscribe((r) => {
            result.next(r.data);
        });

        return result;
    }
}

const app = new App(process.env.NODE_ENV);
export const AppService = app;
