import { exec } from 'child_process';

export const setWallpaper = async (path: string): Promise<void> => {
    const cmd = `wallpaperize set ${path}`;
    exec(cmd, console.log);
};
