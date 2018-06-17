import { exec } from 'child_process';
import { BIN_NAME } from '@app/wallpaperize-proxy/init';

export const getSelected = async (): Promise<string> => {
  return new Promise(r => {
    const cmd = `${BIN_NAME} get selected`;
    exec(cmd, (_, s, e) => {
      if (s) {
        const result = s.split('/').pop() as string;
        return r(result.trim());
      }

      if (e) {
        const result = e.split('/').pop() as string;
        return r(result.trim());
      }
    });
  }) as Promise<string>;
};
