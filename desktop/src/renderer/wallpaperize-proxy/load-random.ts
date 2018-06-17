import { exec } from 'child_process';
import { BIN_NAME } from '@app/wallpaperize-proxy/init';

export const loadRandom = async (): Promise<boolean> => {
  return new Promise(r => {
    const cmd = `${BIN_NAME} random -l`;
    exec(cmd, () => r(true));
  }) as Promise<boolean>;
};
