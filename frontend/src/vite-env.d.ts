/// <reference types="vite/client" />

declare module '*.css' {
    const content: { [className: string]: string };
    export default content;
  }

interface ImportMetaEnv {
    VITE_API_URL: string;
    VITE_APP_NAME: string;
  }

interface ImportMeta {
    readonly env: ImportMetaEnv;
  }
