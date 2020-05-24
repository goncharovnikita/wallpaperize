import axios from 'axios';
import os from 'os';
import { exec } from 'child_process';
import { stat, mkdirSync, writeFileSync, chmodSync } from 'fs';

import { getBinname, getAppPlacement } from '../shared/get-binname';

export const APP_PLACEMENT = getAppPlacement();
export const BIN_NAME = getBinname();

const SITE_ROOT = 'https://wallpaperize.goncharovnikita.com/';
const GET_VERSION_PATH = `${SITE_ROOT}api/get/maxversion`;
const BUILDS_PATH = `${SITE_ROOT}builds/`;

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
    exec(
      `${BIN_NAME} info -o json`,
      { encoding: 'latin1' },
      (err, stdout, stderr) => {
        if (err) {
          console.error(err);
          reject(err);
          return;
        }

        if (stderr) {
          resolve(JSON.parse(stderr) as WallpaperizeInfo);
          return;
        }

        resolve(JSON.parse(stdout) as WallpaperizeInfo);
      }
    );
  }) as Promise<WallpaperizeInfo>;
};

export const getRandom = async (): Promise<string[]> => {
  const url = 'https://api.wallpaperize.goncharovnikita.com/get/random';
  const resp = await axios.get(url);
  const result = resp.data as string[];

  return result;
};

export const getSelected = async (): Promise<string> => {
  return new Promise(resolve => {
    const cmd = `${BIN_NAME} get selected`;
    exec(cmd, (_, s, e) => {
      if (s) {
        const result = s.split('/').pop() as string;
        resolve(result.trim());
        return;
      }

      if (e) {
        const result = e.split('/').pop() as string;
        resolve(result.trim());
      }
    });
  }) as Promise<string>;
};

const initAppdir = () => {
  mkdirSync(APP_PLACEMENT);
};

const getMaxVersion = async (): Promise<string> => {
  const result = await axios.get(GET_VERSION_PATH);
  return result.data;
};

const getDownloadLink = async (version: string): Promise<string> => {
  const platform = os.platform();
  switch (platform) {
    case 'darwin':
      return `${BUILDS_PATH}darwin-amd64-${version}`;
    case 'linux':
      return `${BUILDS_PATH}linux-amd64-${version}`;
    case 'win32':
      return `${BUILDS_PATH}windows-amd64-${version}`;
    default:
      throw new Error('Unsupported platform');
  }
};

const initBin = async () => {
  return new Promise(async resolve => {
    const version = await getMaxVersion();
    const downloadLink = await getDownloadLink(version);
    const result = await axios({
      method: 'GET',
      url: downloadLink,
      responseType: 'blob'
    });
    const freader = new FileReader();
    freader.onload = async () => {
      writeFileSync(BIN_NAME, freader.result, 'binary');

      chmodSync(BIN_NAME, '0777');

      try {
        const info = await getInfo();
        if (info.app_version === version) {
          console.log('app successfully initialized');
        } else {
          console.log(info.app_version);
        }
        resolve();
      } catch (e) {
        console.log(e);
        // fs.unlinkSync(BIN_NAME);
      }
    };

    freader.readAsBinaryString(result.data);
  });
};

export const init = async (): Promise<boolean> => {
  return new Promise(resolve => {
    stat(APP_PLACEMENT, async err => {
      if (err) {
        console.log(err);
        initAppdir();
        await initBin();
      }

      stat(BIN_NAME, async (e, _) => {
        if (e) {
          await initBin();
        }

        console.log('App initialized');
        resolve(true);
      });
    });
  }) as Promise<boolean>;
};

export const loadRandom = async (): Promise<boolean> => {
  return new Promise(resolve => {
    const cmd = `${BIN_NAME} random -l`;
    exec(cmd, () => resolve(true));
  }) as Promise<boolean>;
};

export const setWallpaper = async (path: string): Promise<void> => {
  const resolvedPath = path
    .replace('/random_images_min', '/random_images')
    .replace('-min', '');
  const cmd = `${BIN_NAME} set ${resolvedPath}`;

  exec(cmd, console.log);
};
