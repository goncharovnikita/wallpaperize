import { Dispatch as ReduxDispatch, Store as ReduxStore, Action } from 'redux';

export type AppState = {
  path: AppPath;
};

export enum AppPath {
  Init = 'APP_PATH_INIT',
  Main = 'APP_PATH_MAIN',
  Saved = 'APP_PATH_SAVED'
}

export type GetState = () => AppState;

export type Dispatch = ReduxDispatch<Action<string>>;

export type Store = ReduxStore<AppState, Action<string>>;
