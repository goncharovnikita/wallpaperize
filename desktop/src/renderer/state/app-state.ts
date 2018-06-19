export enum AppPath {
  Init = 'APP_PATH_INIT',
  Main = 'APP_PATH_MAIN',
  Saved = 'APP_PATH_SAVED'
}

export interface AppState {
  path: AppPath;
}
