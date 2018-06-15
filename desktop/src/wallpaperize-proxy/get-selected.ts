import { exec } from 'child_process';

export const getSelected = async (): Promise<string> => {
    return new Promise(r => {
        const cmd = `wallpaperize get selected`;
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
