import { ipcRenderer } from 'electron';

export class AppService {
  constructor() {
    this.listenForUpdate();
  }

  listenForUpdate(): void {
    ipcRenderer.on('updater-message', (_: any, msg: any) => {
      console.log(msg);
    });
  }
}
