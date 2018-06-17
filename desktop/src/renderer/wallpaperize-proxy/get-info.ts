import { exec } from 'child_process';
import { BIN_NAME } from '@app/wallpaperize-proxy/init';

export interface WallpaperizeInfo {
  app_version: string;
  arch: string;
  os: string;
  build: string;
  daily_images: string[];
  random_images: string[];
}

export const getInfo = async (): Promise<WallpaperizeInfo> => {
  return new Promise((resolve, reject) => {
    exec(BIN_NAME + ' info -o json', {encoding: 'latin1'}, (err, stdout, stderr) => {
      console.log(err)
      if (err) {
        reject(err);
        return;
      }

      if (stderr) {
        resolve(JSON.parse(stderr) as WallpaperizeInfo);
        return;
      }

      resolve(JSON.parse(stdout) as WallpaperizeInfo);
    });
  }) as Promise<WallpaperizeInfo>;
};
