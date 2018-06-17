import { exec } from 'child_process';
import { BIN_NAME } from '@app/wallpaperize-proxy/init';

export const setWallpaper = async (path: string): Promise<void> => {
  const cmd = `${BIN_NAME} set ${path}`;
  exec(cmd, console.log);
};
