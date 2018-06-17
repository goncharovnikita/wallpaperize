import { Observable, ReplaySubject, from } from 'rxjs';
import { Platform } from '../platform/platform';
import Axios from 'axios';

export interface DownloadLinks {
  mac: string;
  linux: string;
  windows: string;
}

class App {
  private _baseURL = '';
  private _selectedPlatform: ReplaySubject<Platform>;
  private _selectedVersion: ReplaySubject<string>;
  private _downloadLinks: ReplaySubject<DownloadLinks>;

  constructor(env: string) {
    if (env === 'production') {
      this._baseURL = 'https://wallpaperize.goncharovnikita.com/api';
    } else {
      this._baseURL = 'http://localhost:3000/';
    }
    this._selectedPlatform = this._getSelectedPlatform();
    this._selectedVersion = this._getSelectedVersion();
    this._downloadLinks = this._getDownloadLinks();
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

  getDownloadLinks(): Observable<DownloadLinks> {
    return this._downloadLinks;
  }

  private _getSelectedPlatform(): ReplaySubject<Platform> {
    const result = new ReplaySubject<Platform>();
    switch (true) {
      case /mac(\w+)?/gim.test(navigator.platform):
        result.next(Platform.Mac);
        break;
      case /win(\w+)?/gim.test(navigator.platform):
        result.next(Platform.Windows);
        break;
      default:
        result.next(Platform.Linux);
    }
    return result;
  }

  private _getSelectedVersion(): ReplaySubject<string> {
    const result = new ReplaySubject<string>();

    from(Axios.get(this._baseURL + 'get/maxversion')).subscribe(r => {
      result.next(r.data);
    });

    return result;
  }

  private _getDownloadLinks(): ReplaySubject<DownloadLinks> {
    const result = new ReplaySubject<DownloadLinks>();

    from(Axios.get(this._baseURL + 'get/links')).subscribe(r =>
      result.next(r.data)
    );

    return result;
  }
}

const app = new App(process.env.NODE_ENV);
export const AppService = app;
