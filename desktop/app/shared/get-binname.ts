import * as path from 'path';
import { homedir } from 'os';

export const getAppPlacement = () => path.resolve(homedir(), '.wallpaperize');

export const getBinname = (): string => {
  const appPlacement = getAppPlacement();

  if (process.platform === 'win32') {
    return path.resolve(appPlacement, 'wallpaperize.exe');
  }

  return path.resolve(appPlacement, 'wallpaperize');
};

