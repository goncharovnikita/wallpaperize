export const getAppRoot = (): string => {
  if (process.env.NODE_ENV === 'production') {
    return `file://${__dirname}/index.html#`;
  }

  return `http://localhost:${process.env.ELECTRON_WEBPACK_WDS_PORT}#`;
};
