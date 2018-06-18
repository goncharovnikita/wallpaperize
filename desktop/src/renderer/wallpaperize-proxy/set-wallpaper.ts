import { exec } from 'child_process';
import { BIN_NAME } from '@app/wallpaperize-proxy/init';

export const setWallpaper = async (path: string): Promise<void> => {
  path = path.replace('/min', '').replace('-min', '');
  const cmd = `${BIN_NAME} set ${path}`;
  exec(cmd, console.log);
};
