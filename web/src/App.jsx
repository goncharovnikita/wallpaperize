import React from "react";
import axios from "axios";

import styles from "./app.module.scss";

let initialPlatform;

switch (true) {
    case /mac(\w+)?/gim.test(navigator.platform):
        initialPlatform = "mac";
        break;
    case /win(\w+)?/gim.test(navigator.platform):
        initialPlatform = "windows";
        break;
    default:
        initialPlatform = "linux";
}

const BASE_URL = "https://api.wallpaperize.goncharovnikita.com";

function Button({ os, activeOs, onClick, children }) {
    const handleClick = React.useCallback(() => onClick(os), [os, onClick]);

    return (
        <button
            className={`${styles.button} ${
                os === activeOs ? styles.active : ""
            }`}
            onClick={handleClick}
        >
            {children}
        </button>
    );
}

function App() {
    const [activeOs, setActiveOs] = React.useState(initialPlatform);
    const [downloadLinks, setDownloadLinks] = React.useState([]);
    const [maxVersion, setMaxVersion] = React.useState();

    React.useEffect(() => {
        (async () => {
            const linksRes = await axios.get(`${BASE_URL}/get/links`);
            const maxversionRes = await axios.get(`${BASE_URL}/get/maxversion`);

            setDownloadLinks(linksRes.data);
            setMaxVersion(maxversionRes.data);
        })();
    }, []);

    const downloadLinkHref = React.useMemo(() => {
        if (downloadLinks) {
            return downloadLinks[activeOs];
        }

        return null;
    }, [activeOs, downloadLinks]);

    return (
        <div className={styles.container}>
            <div className={styles.box}>
                <h1 className={styles.title}>Wallpaperize</h1>
                <p className={styles.subtitle}>
                    Set the image in high quality as your desktop wallpaper
                </p>
                <hr />
                <div className={styles.buttons}>
                    <Button os="mac" activeOs={activeOs} onClick={setActiveOs}>
                        Mac OS
                    </Button>
                    <Button
                        os="linux"
                        activeOs={activeOs}
                        onClick={setActiveOs}
                    >
                        Linux
                    </Button>
                    <Button
                        os="windows"
                        activeOs={activeOs}
                        onClick={setActiveOs}
                    >
                        Windows
                    </Button>
                </div>
                <hr />
                <div className={styles.buttons}>
                    <a
                        href={downloadLinkHref}
                        rel="noopener noreferrer"
                        target="_blank"
                    >
                        <button className={styles.button}>Download</button>
                    </a>
                </div>
            </div>
        </div>
    );
}

export default App;
