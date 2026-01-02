import { dirname, resolve } from 'path';
import { fileURLToPath } from 'url';

const __dirname = dirname(fileURLToPath(import.meta.url));

export default {
  useTabs: false,
  singleQuote: true,
  trailingComma: 'none',
  printWidth: 100,
  plugins: [
    '@trivago/prettier-plugin-sort-imports',
    'prettier-plugin-svelte',
    'prettier-plugin-tailwindcss'
  ],
  overrides: [
    {
      files: '*.svelte',
      options: {
        parser: 'svelte'
      }
    }
  ],
  // in a monorepo, the svelte vscode extension will resolve this from the root dir
  // the prettier config is nested in /web, so it needs to resolve from the root dir as well
  tailwindStylesheet: resolve(__dirname, './src/routes/layout.css'),
  bracketSameLine: true,
  endOfLine: 'lf',
  importOrder: [
    '<BUILTIN_MODULES>',
    '<THIRD_PARTY_MODULES>',
    '^$app/(.*)$',
    '^lib/components/(.*)$',
    '^lib/assets/(.*)$',
    '^lib/(.*)$',
    '^[./]',
    '<THIRD_PARTY_TS_TYPES>',
    '<TS_TYPES>'
  ],
  importOrderSeparation: true
};
