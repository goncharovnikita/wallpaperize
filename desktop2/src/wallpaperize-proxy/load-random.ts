import { exec } from 'child_process';

export const loadRandom = async (): Promise<boolean> => {
    return new Promise(r => {
        const cmd = `wallpaperize random -l`;
        exec(cmd, () => r(true));
    }) as Promise<boolean>;
};
