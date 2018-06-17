import * as path from 'path';
import { APP_PLACEMENT } from '@app/wallpaperize-proxy/init';
export const getBinname = (): string => {
  if (process.platform === 'win32') {
    return path.resolve(APP_PLACEMENT, 'wallpaperize.exe');
  } else {
    return path.resolve(APP_PLACEMENT, 'wallpaperize');
  }
}