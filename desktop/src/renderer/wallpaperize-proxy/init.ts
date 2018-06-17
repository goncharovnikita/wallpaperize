// import { exec } from 'child_process';
import {
  stat,
  mkdirSync,
  writeFileSync,
  chmodSync,
  createWriteStream
} from 'fs';
import Axios from 'axios';
import { homedir } from 'os';
import * as os from 'os';
import * as fs from 'fs';
import { getInfo } from '@app/wallpaperize-proxy/get-info';
const APP_PLACEMENT = homedir() + '/.wallpaperize';
export const BIN_NAME = APP_PLACEMENT + '/wallpaperize';
const SITE_ROOT = 'https://wallpaperize.goncharovnikita.com/';
const GET_VERSION_PATH = SITE_ROOT + 'api/get/maxversion';
const BUILDS_PATH = SITE_ROOT + 'builds/';

export const init = async (): Promise<boolean> => {
  return new Promise(resolve => {
    stat(APP_PLACEMENT, async (err, stats) => {
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

const initAppdir = () => {
  mkdirSync(APP_PLACEMENT);
};

const initBin = async () => {
  const version = await getMaxVersion();
  const downloadLink = await getDownloadLink(version);
  const result = await Axios({
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
    } catch (e) {
      console.log(e);
      fs.unlinkSync(BIN_NAME);
    }
  };

  freader.readAsBinaryString(result.data);
};

const getDownloadLink = async (version: string): Promise<string> => {
  const platform = os.platform();
  switch (platform) {
    case 'darwin':
      return BUILDS_PATH + 'darwin-amd64-' + version;
    case 'linux':
      return BUILDS_PATH + 'linux-amd64-' + version;
    default:
      throw new Error('Unsupported platform');
  }
};

const getMaxVersion = async (): Promise<string> => {
  const result = await Axios.get(GET_VERSION_PATH);
  return result.data;
};
