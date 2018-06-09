import * as React from 'react';
import * as express from 'express';
import {renderToString} from 'react-dom/server';
import { AppComponent } from './app/App';
import * as path from 'path';
import * as fs from 'fs';

const port = process.env.PORT || 4200;
const app = express();

app.get('/', (req, res) => {
    const filePath = path.resolve(process.cwd(), 'dist', 'index.html');

    fs.readFile(filePath, 'utf8', (err, htmlData) => {
        if (err) {
            console.error('err', err);
            return res.status(404).end();
        }

        // render the app as a string
        const html = renderToString(<AppComponent />);

        // inject the rendered app into our html and send it
        return res.send(
            htmlData.replace(
                '<div id="app-root"></div>',
                `<div id="app-root">${html}</div>`
            )
        );
    });
});

app.use(express.static(process.cwd() + '/dist'));

app.listen(port, () => {
    console.log('app listen on ::', port);
});
